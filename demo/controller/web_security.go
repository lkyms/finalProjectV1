package controller

import (
	"demo/dao"
	"demo/middleware"
	"demo/model"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type MessageInterface struct {
	PhoneNumber string `form:"phone" json:"phone" binding:"required"`
	Type        int    `form:"type" json:"type"`
}

// 短信验证码发送
func SignUpSendMessage(c *gin.Context) {
	var s MessageInterface
	if err := c.Bind(&s); err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	var strKey string
	cap := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	if s.Type == 0 {
		// 注册的验证码
		strKey = fmt.Sprintf("RegCaptcha:%s", s.PhoneNumber)
	} else {
		// 忘记密码的验证码
		strKey = fmt.Sprintf("FinCaptcha:%s", s.PhoneNumber)
	}
	// res := util.Send(cap, s.PhoneNumber) // 就剩八条了啊，省着点用
	// 存到redis
	// 验证码两分钟有效

	fmt.Println(strKey)
	if err := dao.Rdb.Set(strKey, cap, 120*time.Second).Err(); err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(200, gin.H{
		"status":  0,
		"message": "发送成功",
		//"More":    string(res),
		"cap": cap, // 测试用的

	})

}

type SignInterface struct {
	PhoneNumber string `form:"phone" json:"phone" binding:"required"`
	Username    string `form:"username" json:"username" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	Cap         string `form:"cap" json:"cap" binding:"required"`
}

func SignUp(c *gin.Context) {
	// 注册
	var s SignInterface
	if err := c.Bind(&s); err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	// 验证码审核
	var val string
	var err error
	if val, err = dao.Rdb.Get("RegCaptcha:" + s.PhoneNumber).Result(); err != nil {
		log.Println("验证码获取失败:", err)
		c.JSON(200, gin.H{
			"status":  1,
			"message": "验证码过期",
		})
		return
	}
	if val != s.Cap {
		c.JSON(200, gin.H{
			"status":  1,
			"message": "验证码错误",
		})
		return
	}

	//
	var u model.User
	u.Phone = s.PhoneNumber
	u.Password = s.Password
	u.Username = s.Username

	// 创建用户
	if err := u.Create(); err != nil {

		if strings.Contains(err.Error(), "Error 1062") {
			c.JSON(200, gin.H{
				"status":  1,
				"message": "用户名或手机号已注册",
			})
		} else {
			log.Printf("create in reg err:%v\n", err)
			c.JSON(200, gin.H{
				"status":  1,
				"message": "创建用户错误",
			})
		}
		return
	} else {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "success!!",
		})
	}
}

type SignInInterface struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	// 登录
	var s SignInInterface
	if err := c.Bind(&s); err != nil {
		c.JSON(200, gin.H{
			"Status":  1,
			"Message": err.Error(),
		})
		return
	}
	var u model.User

	u.Username = s.Username
	var getUser model.User
	var err error
	getUser, err = u.Get()
	if err != nil {
		log.Printf("err:%v\n", err)
	} else {
		// 加密匹配
		byteHash := []byte(getUser.Password)
		err2 := bcrypt.CompareHashAndPassword(byteHash, []byte(s.Password))
		if err2 != nil {
			// 登录失败 密码不符
			c.JSON(200, gin.H{
				"Status":  1,
				"Message": "密码错误",
			})
			return
		}
		//登录成功 密码符合

		// 下面设置token 等
		var token string
		token, err = middleware.GenerateToken(getUser.Username, strconv.Itoa(int(getUser.ID)))
		if err != nil {
			log.Printf("Token Create err:%v\n", err)
			return
		}
		log.Println(token)
		//存到redis

		keyStr := fmt.Sprintf("Token:%d", getUser.ID)
		if err := dao.Rdb.Set(keyStr, token, 24*time.Hour).Err(); err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(200, gin.H{
			"Status":  0,
			"Message": "密码正确",
			"token":   token,
		})
	}
}

func FindPassword(c *gin.Context) {
	// 注册
	var s SignInterface
	if err := c.Bind(&s); err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	// 验证码审核
	var val string
	var err error
	if val, err = dao.Rdb.Get("FinCaptcha:" + s.PhoneNumber).Result(); err != nil {
		log.Println("验证码获取失败:", err)
		c.JSON(200, gin.H{
			"status":  1,
			"message": "验证码过期",
		})
		return
	}
	if val != s.Cap {
		c.JSON(200, gin.H{
			"status":  1,
			"message": "验证码错误",
		})
		return
	}

	//
	var u model.User
	u.Phone = s.PhoneNumber
	u.Username = s.Username
	getUser, err := u.Get()
	if err != nil {
		c.JSON(200, gin.H{
			"Status":  1,
			"Message": "此手机号未注册或用户名不匹配",
		})
		return
	} else {
		// 加密匹配
		byteHash := []byte(getUser.Password)
		err2 := bcrypt.CompareHashAndPassword(byteHash, []byte(s.Password))
		if err2 == nil {
			// 密码就是旧密码
			c.JSON(200, gin.H{
				"status":  1,
				"message": "你tm不知道密码吗，找回个锤子",
			})
			return
		}
		u.ID = getUser.ID
		if err := u.Update(s.Password); err != nil {
			c.JSON(200, gin.H{
				"status":  1,
				"message": "更新密码失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"status":  0,
			"message": "重置成功！",
		})

	}
}

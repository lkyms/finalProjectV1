package controller

import (
	"demo/model"
	"demo/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 设置头像，返回头像url(json)
func SetAvatar(c *gin.Context) {
	// 从jwt中获取是哪个用户，若jwt不支持就在定一个key
	file, head, _ := c.Request.FormFile("file")
	size := head.Size
	url, err := util.UploadAvator(file, size)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	var u model.User

	id, _ := c.Get("uid")
	t, _ := strconv.Atoi(id.(string))
	u.ID = uint(t)
	u.AvatarUrl = url
	if err = u.Update(u); err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  0,
		"message": "success",
		"url":     url,
	})

}

// 获取指定username的头像的url
func GetAvatar(c *gin.Context) {
	username := c.Query("username")
	var u, g model.User
	u.Username = username
	var err error
	if g, err = u.Get(); err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
		"url":    g.AvatarUrl,
	})

}

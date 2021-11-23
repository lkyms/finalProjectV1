// 此接口支持发送验证码短信、订单通知短信；
// 调试期间，请使用测试专用短信模板：您的验证码是：1234。请不要把验证码泄露给其他人；
// 请求参数中的account和password分别为 APIID、APIKEY，请在本页面上方处获取。

package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func Send(captcha, mobileNumber string) (data []byte) {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := GetConfig("message.APIID")   //查看用户名 登录用户中心->验证码通知短信>产品总览->API接口信息->APIID
	_password := GetConfig("message.APIKEY") //查看密码 登录用户中心->验证码通知短信>产品总览->API接口信息->APIKEY
	_mobile := mobileNumber
	_content := fmt.Sprintf("您的验证码是：%s。请不要把验证码泄露给其他人。", captcha)
	v.Set("account", _account)
	v.Set("password", getMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	if err != nil {
		log.Printf("Message sending error:%v\n", err)
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ = ioutil.ReadAll(resp.Body)
	//fmt.Println(string(data), err)
	return
}

package util

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	accessKey string = GetConfig("qiniu.accessKey")
	secretKey string = GetConfig("qiniu.secretKey")
	bucket    string = GetConfig("qiniu.bucket")
	qiurl     string = GetConfig("qiniu.url")
)

func UploadAvator(file multipart.File, size int64) (retUrl string, err error) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{} // 返回的玩意
	// 可选配置
	putExtra := storage.PutExtra{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, size, &putExtra)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(ret.Key, ret.Hash)
	retUrl = qiurl + ret.Key
	return
}

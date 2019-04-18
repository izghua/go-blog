/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-13
 * Time: 00:51
 */
package service

import (
	"context"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/utils/alarm"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)


// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

// Upload file to Qiniu
// LocalFile is the local file, such as "./static/images/uploads/2.jpeg"
// FileName is the name what  qiniu name as
// The storage Zone is default
func Qiniu(localFile string,fileName string) {
	accessKey := conf.QiNiuAccessKey
	secretKey := conf.QiNiuSecretKey
	//localFile := "./static/images/uploads/2.jpeg"
	bucket := conf.QiNiuBucket
	key := fileName
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	switch conf.QiNiuZone {
	case "HUABEI":
		cfg.Zone = &storage.ZoneHuabei
	case "HUADONG":
		cfg.Zone = &storage.ZoneHuadong
	case "BEIMEI":
		cfg.Zone = &storage.ZoneBeimei
	case "HUANAN":
		cfg.Zone = &storage.ZoneHuanan
	case "XINJIAPO":
		cfg.Zone = &storage.ZoneXinjiapo
	default:
		cfg.Zone = &storage.ZoneHuabei
	}
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			//"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		zgh.ZLog().Error("message","service.QiNiu upload file","err",err.Error())
		alarm.Alarm("文件上传七牛失败了,文件名是"+fileName)
		return
	}
	//fmt.Println("234222",ret,"最后是谁啥")
	//fmt.Println(ret.Key,"234222", ret.Hash,"最后是谁啥")
	zgh.ZLog().Info("message","service.QiNiu upload file","end","succeed")
	return
}



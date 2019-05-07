/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:50
 */
package conf

import (
	"time"
)

const (

	AppUrl = "http://localhost:8081/static/images/uploads"

	DefaultLimit = "20"

	DbUser = "root"
	DbPassword = "Passw0rd"
	DbPort = "3306"
	DbDataBase = "go-blog"
	DbHost = "127.0.0.1"

	AlarmType = "mail,wechat"

	MailUser = "test@g9zz.com"
	MailPwd = "1234abcd#"
	MailHost = "smtp.mxhichina.com:25"


	HashIdSalt = "i must add a salt what is only for me"
	HashIdLength = 8


	RedisAddr = "localhost:6379"
	RedisPwd = ""
	RedisDb = 0


	QCaptchaAid = "2040723710"
	QCaptchaSecreptKey = "0hG5RLcAjsrktVjvV5YRRQQ**"

	BackUpFilePath = "./backup/"
	BackUpDuration = time.Hour * 24
	BackUpSentTo = "xzghua@gmail.com"

	DataCacheTimeDuration = 720
	ImgUploadUrl = "http://localhost:8081/console/post/imgUpload"
	ImgUploadDst = "./static/images/uploads/"

	//qiniu
	QiNiuUploadImg = true
	QiNiuHostName = "http://pl89sz86l.bkt.clouddn.com/"
	QiNiuAccessKey = "Mk80G9bd_VcsLvMamVXhqxrWiChc76Vz3UGlbJqA"
	QiNiuSecretKey = "us0URcelzGY-mcoSY1Lw3mkZrTP1QCsegRxArTvZ"
	QiNiuBucket = "go-blog"
	QiNiuZone = "HUABEI" //you can use "HUADONG","HUABEI","BEIMEI","HUANAN","XINJIAPO"


	CateListKey =  "all:cate:sort"
	TagListKey =  "all:tag"

	Theme = 0
	Title = "默认Title"
	Keywords = "默认关键词,叶落山城秋"
	Description = "个人网站,https://github.com/izghua/go-blog"
	RecordNumber = "000-0000"
)


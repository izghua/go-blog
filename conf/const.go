/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:50
 */
package conf

import "time"

const (
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
)


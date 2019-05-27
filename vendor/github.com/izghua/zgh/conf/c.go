package conf

import "time"

const (
	DBHOST = "127.0.0.1"
	DBPORT = "3306"
	DBPASSWORD = "Passw0rd"
	DBUSERNAME = "root"
	DBDATABASE = "izghua"

	ALARMCRITICAL = "critical"
	ALARMWARNING  = "warning"
	ALARMALERT    = "alert"

	MAIlTYPE = "html"

	HASHIDSALT = "salt"
	HASHIDMINLENGTH = 8


	REDISADDR = ""
	REDISPWD = ""
	REDISDB = 0

	JWTISS = "izghua"
	JWTAUDIENCE = "zgh"
	JWTJTI = "izghua"
	JWTSECRETKEY = "izghua"
	JWTTOKENKEY = "login:token:"
	JWTTOKENLIFE = time.Hour * time.Duration(72)


	QCapUrl = "https://ssl.captcha.qq.com/ticket/verify"



)

// Log
const  (
	LOGFILEPATH = "./log"
	LOGFILENAME = "zog"
	LOGFILESUFFIX = "log"
	LOGFILEMAXSIZE = 0
	LOGFILEMAXNSIZE = 1
	LOGTIMEZONE = "Asia/Chongqing"
)

const (
	BackUpDest = "./backup"
	BackUpDuration = "0 0 0 * * *"
	BackUpSqlFileName = "-sql-backup.sql"
	BackUpFilePath = "./backup/"
)
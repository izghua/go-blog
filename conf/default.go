/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:53
 */
package conf

import (
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conf"
	"github.com/izghua/zgh/conn"
	"github.com/izghua/zgh/jwt"
	"github.com/izghua/zgh/utils/alarm"
	"github.com/izghua/zgh/utils/backup"
	"github.com/izghua/zgh/utils/hashid"
	"github.com/izghua/zgh/utils/mail"
	"github.com/izghua/zgh/utils/qq_captcha"
	string2 "github.com/izghua/zgh/utils/string"
	"github.com/speps/go-hashids"
	"time"
)

var (
	SqlServer *xorm.Engine
	ZHashId *hashids.HashID
	CacheClient *redis.Client
	MailClient *mail.EmailParam
)


func DefaultInit() {
	ZLogInit()
	DbInit()
	AlarmInit()
	MailInit()
	ZHashIdInit()
	RedisInit()
	JwtInit()
	QCaptchaInit()
	// the customer error code init
	conf.SetMsg(Msg)
	//BackUpInit()
}

func ZLogInit() {
	zog := new(zgh.ZLogParam)
	fileName := zog.SetFileName("go-blog")
	err := zog.ZLogInit(fileName)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
}

func DbInit () {
	sp := new(conn.Sp)
	dbUser := sp.SetDbUserName(DbUser)
	dbPwd := sp.SetDbPassword(DbPassword)
	dbPort := sp.SetDbPort(DbPort)
	dbHost := sp.SetDbHost(DbHost)
	dbdb := sp.SetDbDataBase(DbDataBase)
	sqlServer,err := conn.InitMysql(dbUser,dbPwd,dbPort,dbHost,dbdb)
	SqlServer = sqlServer
	if err != nil {
		zgh.ZLog().Error("有错误",err.Error())
	}
}

func BackUpInit() {
	bp := new(backup.BackUpParam)
	dest := "./zip/"+time.Now().Format("2006-01-02")+".zip"
	backu := bp.SetFilePath(BackUpFilePath).
		SetFiles("./backup").
		SetDest(dest).SetCronSpec(BackUpDuration)
	data := make(map[string]string)
	data[time.Now().Format("2006-01-02")+".zip"] = dest
	bp.Ep = MailClient
	subject := time.Now().Format("2006-01-02") + "备份邮件"
	bp.Ep.SetSubject(mail.EmailType(subject)).SetAttaches(data).SetBody(mail.EmailType(
		`<html><body>
		<p><img src="https://golang.org/doc/gopher/doc.png"></p><br/>
		<h1>天黑了,有点饿了.</h1>
		`+ string2.RandString(10) +`
		</body></html>`)).SetTo(BackUpSentTo)
	err := backu.Backup()
	if err != nil {
		zgh.ZLog().Error("message","backup has error","error",err.Error())
	} else {
		zgh.ZLog().Info("message","Congratulations for backup")
	}
}



func AlarmInit() {
	a := new(alarm.AlarmParam)
	alarmT := a.SetType(AlarmType)
	mailTo := a.SetMailTo("xzghua@gmail.com")
	err := a.AlarmInit(alarmT,mailTo)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
}

func MailInit() {
	m := new(mail.EmailParam)
	mailUser := m.SetMailUser(MailUser)
	mailPwd := m.SetMailPwd(MailPwd)
	mailHost :=  m.SetMailHost(MailHost)
	mails,err := m.MailInit(mailPwd,mailHost,mailUser)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
	MailClient = mails
}



func ZHashIdInit() {
	hd := new(hashid.HashIdParams)
	salt := hd.SetHashIdSalt(HashIdSalt)
	hdLength := hd.SetHashIdLength(HashIdLength)
	zHashId,err := hd.HashIdInit(hdLength,salt)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
	ZHashId = zHashId

}

func RedisInit() {
	rc := new(conn.RedisClient)
	addr := rc.SetRedisAddr(RedisAddr)
	pwd := rc.SetRedisPwd(RedisPwd)
	db := rc.SetRedisDb(RedisDb)
	client,err := rc.RedisInit(addr,db,pwd)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
	CacheClient = client
}

func JwtInit() {
	jt := new(jwt.JwtParam)
	ad := jt.SetDefaultAudience(JwtAudience)
	jti := jt.SetDefaultJti(JwtJti)
	iss := jt.SetDefaultIss(JwtIss)
	sk := jt.SetDefaultSecretKey(JwtSecretKey)
	rc := jt.SetRedisCache(CacheClient)
	tl := jt.SetTokenLife(JwtTokenLife)
	_ = jt.JwtInit(ad,jti,iss,sk,rc,tl)
}

func QCaptchaInit() {
	qc := new(qq_captcha.QQCaptcha)
	aid := qc.SetAid(QCaptchaAid)
	sk := qc.SetSecretKey(QCaptchaSecreptKey)
	_ = qc.QQCaptchaInit(aid,sk)
}
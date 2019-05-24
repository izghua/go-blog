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
	AppUrl = "http://localhost:8081"

	AppImgUrl =  AppUrl + "/static/uploads/images/"

	DefaultLimit = "20"
	DefaultIndexLimit = "3"

	DbUser = "root"
	DbPassword = "Passw0rd"
	DbPort = "3306"
	DbDataBase = "go-blog"
	DbHost = "127.0.0.1"

	AlarmType = "mail,wechat"

	MailUser = "test@test.com"
	MailPwd = ""
	MailHost = "smtp.mxhichina.com:25"


	HashIdSalt = "i must add a salt what is only for me"
	HashIdLength = 8


	JwtIss = "go-blog"
	JwtAudience = "blog"
	JwtJti = "go-blog"
	JwtSecretKey = "go-blog"
	JwtTokenLife = time.Hour * time.Duration(3)

	RedisAddr = "localhost:6379"
	RedisPwd = ""
	RedisDb = 0


	QCaptchaAid = ""
	QCaptchaSecreptKey = "**"

	BackUpFilePath = "./backup/"
	BackUpDuration = "* * */1 * *"
	BackUpSentTo = "xzghua@gmail.com"

	DataCacheTimeDuration = 720
	ImgUploadUrl = AppUrl +"/console/post/imgUpload"
	ImgUploadDst = "./static/uploads/images/"
	ImgUploadBoth = true // img will upload to qiniu and your server local

	//qiniu
	QiNiuUploadImg = true // if you do not want to upload img to qiniu ,you can set this with false
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

	UserCnt = 2


	// index
	PostIndexKey = "index:all:post:list"
	TagPostIndexKey = "index:all:tag:post:list"
	CatePostIndexKey = "index:all:cate:post:list"
	LinkIndexKey = "index:all:link:list"
	SystemIndexKey = "index:all:system:list"
	PostDetailIndexKey = "index:post:detail"
	ArchivesKey = "index:archives:list"

	// github gitment
	GithubName = "xzghua"
	GithubRepo = "ttest"
	GithubClientId = "e298594e1aae93dbdaa7"
	GithubClientSecret = "02fd3cccdd51d28ec861aa509fd73bb0a5c15ca1"

	Test = "234"
)

type Conf struct {
	AppUrl string `yaml:"appUrl"`
	AppImgUrl string `yaml:"appImgUrl"`
	DefaultLimit string `yaml:"defaultLimit"`
	DefaultIndexLimit string `yaml:"defaultIndexLimit"`

	DbUser string `yaml:"dbUser"`
	DbPassword string `yaml:"dbPassword"`
	DbPort string `yaml:"dbPort"`
	DbDataBase string `yaml:"dbDataBase"`
	DbHost string `yaml:"dbHost"`

	AlarmType string `yaml:"alarmType"`
	MailUser string `yaml:"mailUser"`
	MailPwd string `yaml:"mailPwd"`
	MailHost string `yaml:"mailHost"`


	HashIdSalt string `yaml:"hashIdSalt"`
	HashIdLength int `yaml:"hashIdLength"`


	JwtIss string `yaml:"jwtIss"`
	JwtAudience string `yaml:"jwtAudience"`
	JwtJti string `yaml:"jwtJti"`
	JwtSecretKey string `yaml:"jwtSecretKey"`
	JwtTokenLife time.Duration `yaml:"jwtTokenLife"`

	RedisAddr string `yaml:"redisAddr"`
	RedisPwd string `yaml:"redisPwd"`
	RedisDb int `yaml:"redisDb"`


	QCaptchaAid string `yaml:"qCaptchaAid"`
	QCaptchaSecretKey string `yaml:"qCaptchaSecretKey"`

	BackUpFilePath string `yaml:"backUpFilePath"`
	BackUpDuration string `yaml:"backUpDuration"`
	BackUpSentTo string `yaml:"backUpSentTo"`

	DataCacheTimeDuration int `yaml:"dataCacheTimeDuration"`
	ImgUploadUrl string `yaml:"imgUploadUrl"`
	ImgUploadDst string `yaml:"imgUploadDst"`
	ImgUploadBoth bool `yaml:"imgUploadBoth"`

	//qiniu
	QiNiuUploadImg bool `yaml:"qiNiuUploadImg"`
	QiNiuHostName string `yaml:"qiNiuHostName"`
	QiNiuAccessKey string `yaml:"qiNiuAccessKey"`
	QiNiuSecretKey string `yaml:"qiNiuSecretKey"`
	QiNiuBucket string `yaml:"qiNiuBucket"`
	QiNiuZone string `yaml:"qiNiuZone"`


	CateListKey string `yaml:"cateListKey"`
	TagListKey string `yaml:"tagListKey"`

	Theme int `yaml:"theme"`
	Title string `yaml:"title"`
	Keywords string `yaml:"keywords"`
	Description string `yaml:"description"`
	RecordNumber string `yaml:"recordNumber"`

	UserCnt int `yaml:"userCnt"`


	// index
	PostIndexKey string `yaml:"postIndexKey"`
	TagPostIndexKey string `yaml:"tagPostIndexKey"`
	CatePostIndexKey string `yaml:"catePostIndexKey"`
	LinkIndexKey string `yaml:"linkIndexKey"`
	SystemIndexKey string `yaml:"systemIndexKey"`
	PostDetailIndexKey string `yaml:"postDetailIndexKey"`
	ArchivesKey string `yaml:"archivesKey"`

	// github gitment
	GithubName string `yaml:"githubName"`
	GithubRepo string `yaml:"githubRepo"`
	GithubClientId string `yaml:"githubClientId"`
	GithubClientSecret string `yaml:"githubClientSecret"`

	Test string `yaml:"test"`
}


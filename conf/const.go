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
	//AppUrl = "http://localhost:8081"
	//
	//AppImgUrl =  AppUrl + "/static/uploads/images/"
	//DefaultLimit = "20"
	//DefaultIndexLimit = "3"
	//
	//DbUser = "root"
	//DbPassword = "Passw0rd"
	//DbPort = "3306"
	//DbDataBase = "go-blog"
	//DbHost = "127.0.0.1"
	//
	//AlarmType = "mail,wechat"
	//
	//MailUser = "test@test.com"
	//MailPwd = ""
	//MailHost = "smtp.mxhichina.com:25"
	//
	//
	//HashIdSalt = "i must add a salt what is only for me"
	//HashIdLength = 8
	//
	//
	//JwtIss = "go-blog"
	//JwtAudience = "blog"
	//JwtJti = "go-blog"
	//JwtSecretKey = "go-blog"
	//JwtTokenLife = time.Hour * time.Duration(3)
	//
	//RedisAddr = "localhost:6379"
	//RedisPwd = ""
	//RedisDb = 0
	//
	//
	//QCaptchaAid = ""
	//QCaptchaSecreptKey = "**"
	//
	//BackUpFilePath = "./backup/"
	//BackUpDuration = "* * */1 * *"
	//BackUpSentTo = "xzghua@gmail.com"
	//
	//DataCacheTimeDuration = 720
	//ImgUploadUrl = AppUrl +"/console/post/imgUpload"
	//ImgUploadDst = "./static/uploads/images/"
	//ImgUploadBoth = true // img will upload to qiniu and your server local
	//
	////qiniu
	//QiNiuUploadImg = true // if you do not want to upload img to qiniu ,you can set this with false
	//QiNiuHostName = "http://pl89sz86l.bkt.clouddn.com/"
	//QiNiuAccessKey = "Mk80G9bd_VcsLvMamVXhqxrWiChc76Vz3UGlbJqA"
	//QiNiuSecretKey = "us0URcelzGY-mcoSY1Lw3mkZrTP1QCsegRxArTvZ"
	//QiNiuBucket = "go-blog"
	//QiNiuZone = "HUABEI" //you can use "HUADONG","HUABEI","BEIMEI","HUANAN","XINJIAPO"
	//
	//
	//CateListKey =  "all:cate:sort"
	//TagListKey =  "all:tag"
	//
	//Theme = 0
	//Title = "默认Title"
	//Keywords = "默认关键词,叶落山城秋"
	//Description = "个人网站,https://github.com/izghua/go-blog"
	//RecordNumber = "000-0000"
	//
	//UserCnt = 2
	//
	//
	//// index
	//PostIndexKey = "index:all:post:list"
	//TagPostIndexKey = "index:all:tag:post:list"
	//CatePostIndexKey = "index:all:cate:post:list"
	//LinkIndexKey = "index:all:link:list"
	//SystemIndexKey = "index:all:system:list"
	//PostDetailIndexKey = "index:post:detail"
	//ArchivesKey = "index:archives:list"
	//
	//// github gitment
	//GithubName = "xzghua"
	//GithubRepo = "ttest"
	//GithubClientId = "e298594e1aae93dbdaa7"
	//GithubClientSecret = "02fd3cccdd51d28ec861aa509fd73bb0a5c15ca1"
	//
	//Test = "234"
)

type Conf struct {
	AppUrl string `yaml:"AppUrl"`
	AppImgUrl string `yaml:"AppImgUrl"`
	DefaultLimit string `yaml:"DefaultLimit"`
	DefaultIndexLimit string `yaml:"DefaultIndexLimit"`

	DbUser string `yaml:"DbUser"`
	DbPassword string `yaml:"DbPassword"`
	DbPort string `yaml:"DbPort"`
	DbDataBase string `yaml:"DbDataBase"`
	DbHost string `yaml:"DbHost"`

	AlarmType string `yaml:"AlarmType"`
	MailUser string `yaml:"MailUser"`
	MailPwd string `yaml:"MailPwd"`
	MailHost string `yaml:"MailHost"`


	HashIdSalt string `yaml:"HashIdSalt"`
	HashIdLength int `yaml:"HashIdLength"`


	JwtIss string `yaml:"JwtIss"`
	JwtAudience string `yaml:"JwtAudience"`
	JwtJti string `yaml:"JwtJti"`
	JwtSecretKey string `yaml:"JwtSecretKey"`
	JwtTokenLife time.Duration `yaml:"JwtTokenLife"`

	RedisAddr string `yaml:"RedisAddr"`
	RedisPwd string `yaml:"RedisPwd"`
	RedisDb int `yaml:"RedisDb"`


	QCaptchaAid string `yaml:"QCaptchaAid"`
	QCaptchaSecretKey string `yaml:"QCaptchaSecretKey"`

	BackUpFilePath string `yaml:"BackUpFilePath"`
	BackUpDuration string `yaml:"BackUpDuration"`
	BackUpSentTo string `yaml:"BackUpSentTo"`

	DataCacheTimeDuration int `yaml:"DataCacheTimeDuration"`
	ImgUploadUrl string `yaml:"ImgUploadUrl"`
	ImgUploadDst string `yaml:"ImgUploadDst"`
	ImgUploadBoth bool `yaml:"ImgUploadBoth"`

	//qiniu
	QiNiuUploadImg bool `yaml:"QiNiuUploadImg"`
	QiNiuHostName string `yaml:"QiNiuHostName"`
	QiNiuAccessKey string `yaml:"QiNiuAccessKey"`
	QiNiuSecretKey string `yaml:"QiNiuSecretKey"`
	QiNiuBucket string `yaml:"QiNiuBucket"`
	QiNiuZone string `yaml:"QiNiuZone"`


	CateListKey string `yaml:"CateListKey"`
	TagListKey string `yaml:"TagListKey"`

	Theme int `yaml:"Theme"`
	Title string `yaml:"Title"`
	Keywords string `yaml:"Keywords"`
	Description string `yaml:"Description"`
	RecordNumber string `yaml:"RecordNumber"`

	UserCnt int `yaml:"UserCnt"`


	// index
	PostIndexKey string `yaml:"PostIndexKey"`
	TagPostIndexKey string `yaml:"TagPostIndexKey"`
	CatePostIndexKey string `yaml:"CatePostIndexKey"`
	LinkIndexKey string `yaml:"LinkIndexKey"`
	SystemIndexKey string `yaml:"SystemIndexKey"`
	PostDetailIndexKey string `yaml:"PostDetailIndexKey"`
	ArchivesKey string `yaml:"ArchivesKey"`

	// github gitment
	GithubName string `yaml:"GithubName"`
	GithubRepo string `yaml:"GithubRepo"`
	GithubClientId string `yaml:"GithubClientId"`
	GithubClientSecret string `yaml:"GithubClientSecret"`

	Test string `yaml:"Test"`
}



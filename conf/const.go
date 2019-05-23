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
	AppUrl string `json:"appUrl"`
	AppImgUrl string `json:"appImgUrl"`
	DefaultLimit string `json:"defaultLimit"`
	DefaultIndexLimit string `json:"defaultIndexLimit"`

	DbUser string `json:"dbUser"`
	DbPassword string `json:"db_password"`
	DbPort string `json:"db_port"`
	DbDataBase string `json:"db_data_base"`
	DbHost string `json:"db_host"`

	AlarmType string `json:"alarm_type"`
	MailUser string `json:"mail_user"`
	MailPwd string `json:"mail_pwd"`
	MailHost string `json:"mail_host"`


	HashIdSalt string `json:"hash_id_salt"`
	HashIdLength int `json:"hash_id_length"`


	JwtIss string `json:"jwt_iss"`
	JwtAudience string `json:"jwt_audience"`
	JwtJti string `json:"jwt_jti"`
	JwtSecretKey string `json:"jwt_secret_key"`
	JwtTokenLife time.Duration `json:"jwt_token_life"`

	RedisAddr string `json:"redis_addr"`
	RedisPwd string `json:"redis_pwd"`
	RedisDb int `json:"redis_db"`


	QCaptchaAid string `json:"q_captcha_aid"`
	QCaptchaSecreptKey string `json:"q_captcha_secrept_key"`

	BackUpFilePath string `json:"back_up_file_path"`
	BackUpDuration string `json:"back_up_duration"`
	BackUpSentTo string `json:"back_up_sent_to"`

	DataCacheTimeDuration int `json:"data_cache_time_duration"`
	ImgUploadUrl string `json:"img_upload_url"`
	ImgUploadDst string `json:"img_upload_dst"`
	ImgUploadBoth bool `json:"img_upload_both"`

	//qiniu
	QiNiuUploadImg bool `json:"qi_niu_upload_img"`
	QiNiuHostName string `json:"qi_niu_host_name"`
	QiNiuAccessKey string `json:"qi_niu_access_key"`
	QiNiuSecretKey string `json:"qi_niu_secret_key"`
	QiNiuBucket string `json:"qi_niu_bucket"`
	QiNiuZone string `json:"qi_niu_zone"`


	CateListKey string `json:"cate_list_key"`
	TagListKey string `json:"tag_list_key"`

	Theme int `json:"theme"`
	Title string `json:"title"`
	Keywords string `json:"keywords"`
	Description string `json:"description"`
	RecordNumber string `json:"record_number"`

	UserCnt int `json:"user_cnt"`


	// index
	PostIndexKey string `json:"post_index_key"`
	TagPostIndexKey string `json:"tag_post_index_key"`
	CatePostIndexKey string `json:"cate_post_index_key"`
	LinkIndexKey string `json:"link_index_key"`
	SystemIndexKey string `json:"system_index_key"`
	PostDetailIndexKey string `json:"post_detail_index_key"`
	ArchivesKey string `json:"archives_key"`

	// github gitment
	GithubName string `json:"github_name"`
	GithubRepo string `json:"github_repo"`
	GithubClientId string `json:"github_client_id"`
	GithubClientSecret string `json:"github_client_secret"`

	Test string `json:"test"`
}


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
	GithubLabels string `yaml:"GithubLabels"`

	OtherScript string `yaml:"OtherScript"`

	ThemeJs string `yaml:"ThemeJs"`
	ThemeCss string `yaml:"ThemeCss"`
	ThemeImg string `yaml:"ThemeImg"`
	ThemeFancyboxCss string `yaml:"ThemeFancyboxCss"`
	ThemeFancyboxJs string `yaml:"ThemeFancyboxJs"`
	ThemeHLightCss string `yaml:"ThemeHLightCss"`
	ThemeHLightJs string `yaml:"ThemeHLightJs"`
	ThemeShareCss string `yaml:"ThemeShareCss"`
	ThemeShareJs string `yaml:"ThemeShareJs"`
	ThemeArchivesJs string `yaml:"ThemeArchivesJs"`
	ThemeArchivesCss string `yaml:"ThemeArchivesCss"`
}



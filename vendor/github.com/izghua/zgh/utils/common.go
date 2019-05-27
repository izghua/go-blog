/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-03
 * Time: 21:34
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	TimeLocation, _ = time.LoadLocation("Asia/Chongqing") //当地时间
)

// 返回当前时间格式
func GetDateTime() string {
	return time.Now().In(TimeLocation).Format("2006-01-02 15:04:05")
}


func Md5(s string) string {
	c := md5.New()
	c.Write([]byte(s))
	cipherStr := c.Sum(nil)
	return hex.EncodeToString(cipherStr)
}


// 各种字符串加星
func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str,"@") {
		res := strings.Split(str,"@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := Substr2(str,0,3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			result =  Substr2(str,0,3) + "****" + Substr2(str,7,11)
		} else {
			nameRune := []rune(str)
			lens  := len(nameRune)
			fmt.Println(lens,"长度",str)
			if  lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result =  string(nameRune[:1]) + "**" + string(nameRune[lens - 1 : lens])
			} else if lens > 4 {
				result =  string(nameRune[:2]) + "***" + string(nameRune[lens - 2 : lens])
			}
		}
		return
	}
}

func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}
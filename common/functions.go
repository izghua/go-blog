/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-11
 * Time: 23:14
 */
package common

import "strconv"

func Offset(page string,limit string) (limitInt int,offset int) {
	pageInt,err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	limitInt,err = strconv.Atoi(limit)
	if err != nil {
		limitInt = 20
	}

	return limitInt,(pageInt - 1) * limitInt
}


func GoMerge(arr1 []interface{},arr2 []interface{}) []interface{} {
	for _,val := range arr2 {
		arr1 = append(arr1,val)
	}
	return arr1
}


func GoRepeat(str string, num int) string {
	var i int
	newStr := ""
	if num != 0 {
		for i = 0; i < num; i++ {
			newStr += str
		}
	}
	return newStr
}


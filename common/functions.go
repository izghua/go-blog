/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-11
 * Time: 23:14
 */
package common

import (
	"strconv"
	"time"
)

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

func MyPaginate(count int64,limit int,page int) Paginate {
	res := round(int(count),limit)
	totalPage := res

	lastPage := 0
	if page - 1 <= 0 {
		lastPage = 1
	} else {
		lastPage = page - 1
	}

	currentPage := 0
	if page >= res {
		currentPage = res
	} else {
		currentPage = page
	}

	nextPage := 0
	if page + 1 >= res {
		nextPage = res
	} else {
		nextPage = page + 1
	}

	return Paginate{
		Limit: limit,
		Count: int(count),
		Total: totalPage,
		Last: lastPage,
		Current: currentPage,
		Next:  nextPage,
	}
}


func round(a int,b int) int {
	rem := a % b
	dis := a / b
	if rem > 0 {
		return dis+1
	} else {
		return dis
	}
}


func Rem(divisor int) bool {
	if (divisor+1) % 4 == 0 {
		return true
	} else {
		return false
	}
}

func MDate(times time.Time) string {
	return times.Format("2006-01-02 15:04:05")
}

func MDate2(times time.Time) string {
	return times.Format("01-02")
}
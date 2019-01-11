/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-11
 * Time: 23:14
 */
package common


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


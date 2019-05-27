package QLog

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

func GetLogTextPrefix(calldepth int, format string) string {
	return GetNowTime(format) + BLANK + GetCaller(calldepth)
}

func GetNowTime(format string) (str string) {
	var (
		now time.Time
	)
	now = GetNowUnixTimeOBJ()
	if format == "" {
		y, m, d := now.Date()
		h, i, s := now.Clock()
		str = Itoa(y, 4) + DIAGONAL + Itoa(int(m), 2) + DIAGONAL + Itoa(d, 2) + BLANK + Itoa(h, 2) + COLON + Itoa(i, 2) + COLON + Itoa(s, 2)
	} else {
		str = now.Format(format)
	}
	return
}

func GetNowUnixTimeOBJ() time.Time {
	return time.Now().In(TimeLocation)
}

func GetCaller(calldepth int) (str string) {
	var (
		file string
		line int
		ok   bool
	)
	_, file, line, ok = runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	str = path.Base(file) + `:` + strconv.Itoa(line) + `:`
	return
}

func Itoa(i int, wid int) string {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	return string(b[bp:])
}

func mkdirlog(dir string) (e error) {
	_, er := os.Stat(dir)
	b := er == nil || os.IsExist(er)
	if !b {
		if err := os.MkdirAll(dir, 0766); err != nil {
			if os.IsPermission(err) {
				e = err
			}
		}
	}
	return
}

func format(v ...interface{}) (s string) {
	for key, val := range v {
		if key%2 == 0 {
			s += fmt.Sprintf("%+v", val) + `=`
		} else {
			s += fmt.Sprintf("%+v", val) + " "
		}
	}
	return
}

func strDefault(str *string, defaultStr string) {
	if *str == "" {
		*str = defaultStr
	}
}

func intDefault(str *int, defaultInt int) {
	if *str == 0 {
		*str = defaultInt
	}
}

func int64Default(str *int64, defaultInt64 int64) {
	if *str == 0 {
		*str = defaultInt64
	}
}

func catchError() {
	if err := recover(); err != nil {
		fmt.Println(err, string(debug.Stack()))
	}
}

//判断所给路径文件/文件夹是否存在
func isExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//获取绝对路径
func absolutePath(dir string) (fp string) {
	fp, _ = filepath.Abs(dir)
	return
}

//判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	return err == nil && s.IsDir()
}

//判断所给路径是否为文件
func isFile(path string) bool {
	return !isDir(path)
}

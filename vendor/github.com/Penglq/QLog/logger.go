package QLog

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	YRDLogger    LoggerInterface
	ServerIp     string
	TimeLocation *time.Location
)

type LoggerInterface interface {
	SetConfig(uint8, string, ...ConfigOption)
	SetTextPrefix(...interface{})
	AddTextPrefix(...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	AlertWithLevel(alertLevel string, v ...interface{})
}

type logger struct {
	logLevel  uint8 // 默认为0
	isConsole bool  // false
	isFile    bool  // false
	//文件相关配置
	filePath     string
	filename     string
	fileSuffix   string
	fileMaxSize  int64
	fileMaxNSize int
	fileCurrSize int64 //文件大小，字节

	//日志前缀信息
	textPrefix string
	//报警配置
	alertAppId       string
	alertURL         string
	alertContentType string
	alertFormat      string
	//日期格式
	dateFormat string

	timeLocation *time.Location

	callDep    int
	nSize      int //超过设定文件大小的重命名文件序号
	mu         *sync.RWMutex
	logfile    *os.File
	outFile    io.Writer
	outConsole io.Writer
	_date      *time.Time
}

func init() {
	var (
		addrs []net.Addr
		err   error
		ips   []string
	)
	addrs, err = net.InterfaceAddrs()
	if err != nil {
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ips = append(ips, ipnet.IP.String())
				}
			}
		}
	}
	ServerIp = strings.Join(ips, "/")
}

func GetLogger() LoggerInterface {
	if YRDLogger == nil {
		YRDLogger = &logger{mu: new(sync.RWMutex), nSize: 1, callDep: callDep}
	}
	return YRDLogger
}

func (l *logger) SetConfig(logLevel uint8, zone string, opts ...ConfigOption) {
	var err error
	TimeLocation, err = time.LoadLocation(zone) // 时区
	if err != nil {
		fmt.Println(err)
		TimeLocation, _ = time.LoadLocation("Local") // 时区
	}

	t, _ := time.Parse(_DATEFORMAT, GetNowUnixTimeOBJ().Format(_DATEFORMAT)) //当日零点
	l._date = &t

	l.logLevel = logLevel
	for _, optFunc := range opts {
		optFunc(l)
	}
	l.setLogger()
}

type ConfigOption func(*logger)

func WithConsoleOPT() ConfigOption {
	return func(l *logger) {
		l.isConsole = true
	}
}

func WithFileOPT(filepath, filename, filesuffix string, fileMaxSize int64, fileMaxNSize int) ConfigOption {
	strDefault(&filename, DEFAULTFILENAME)
	strDefault(&filepath, DEFAULTFILEPATH)
	strDefault(&filesuffix, DEFAULTFILESUFFIX)
	int64Default(&fileMaxSize, DEFAULTFILEMAXSIZE)
	return func(l *logger) {
		l.isFile = true
		l.filePath = absolutePath(filepath)
		l.filename = filename
		l.fileSuffix = DOT + filesuffix
		l.fileMaxSize = fileMaxSize
		l.fileMaxNSize = fileMaxNSize
	}
}

// 报警配置
func WithAlertOPT(appId, URL, ContentType string) ConfigOption {
	return func(l *logger) {
		l.alertAppId = appId
		l.alertURL = URL
		l.alertContentType = ContentType
	}
}
func WithCommonOPT(cDep int, dateFormat string) ConfigOption {
	return func(l *logger) {
		l.callDep = cDep
		l.dateFormat = dateFormat
	}
}
func WithcallDepOPT(cDep int) ConfigOption {
	return func(l *logger) {
		if cDep == 0 {
			l.callDep = callDep
		} else {
			l.callDep = cDep
		}
	}
}

func (l *logger) SetTextPrefix(keyvals ...interface{}) {
	l.textPrefix = format(keyvals...)
}

func (l *logger) AddTextPrefix(keyvals ...interface{}) {
	l.textPrefix += format(keyvals...)
}

func (l *logger) setLogger() {
	if l.isConsole {
		l.outConsole = os.Stdout
	}
	if l.isFile {
		mkdirlog(l.filePath)
		l.openFile()
	}
}

func (l *logger) getFileFullName() string {
	return l.filePath + "/" + l.filename + UNDERSCODE + l._date.Format(_DATEFORMAT) + l.fileSuffix
}

func (l *logger) getSizeFileFullName(num string) string {
	return l.filePath + "/" + l.filename + UNDERSCODE + l._date.Format(_DATEFORMAT) + DOT + num + l.fileSuffix
}

func (l *logger) openFile() {
	defer catchError()
	var err error
	l.logfile, err = os.OpenFile(l.getFileFullName(), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0766)
	if err != nil {
		panic("不能打开/创建文件 " + err.Error())
	}
	l.outFile = l.logfile
	fileInfo, err := l.logfile.Stat()
	if err != nil {
		panic("获取fileinfo出错")
	}
	l.fileCurrSize = fileInfo.Size()
}

func (l *logger) Debug(keyvals ...interface{}) {
	l.log(TYPEDEBUG, DEBUG, "", keyvals...)
}
func (l *logger) Info(keyvals ...interface{}) {
	l.log(TYPEINFO, INFO, "", keyvals...)
}
func (l *logger) Warn(keyvals ...interface{}) {
	l.log(TYPEWARN, WARN, "", keyvals...)
}
func (l *logger) Error(keyvals ...interface{}) {
	l.log(TYPEERROR, ERROR, "", keyvals...)
}
func (l *logger) Fatal(keyvals ...interface{}) {
	l.log(TYPEFATAL, FATAL, "", keyvals...)
	os.Exit(1)
}
func (l *logger) AlertWithLevel(alertLevel string, keyvals ...interface{}) {
	l.log(TYPEALERT, ALERT, alertLevel, keyvals...)
}

func (l *logger) log(level string, _level uint8, alertLevel string, keyvals ...interface{}) {
	defer catchError()
	if l.logLevel <= _level {
		s := GetLogTextPrefix(l.callDep+1, l.dateFormat) +
			strings.TrimRight(fmt.Sprint(level, BLANK, "IP=", ServerIp, BLANK, l.textPrefix, format(keyvals...)), BLANK) +
			NEWLINE
		// 判断是否调用alert
		if _level == ALERT {
			go l.alert(s, alertLevel)
		}
		n := l.write(s)
		if l.isFile {
			l.fileCurrSize += int64(n)
			l.fileCheck()
		}
	}
}

func (l *logger) write(v string) int {
	l.mu.RLock()
	defer func() {
		l.mu.RUnlock()
		catchError()
	}()
	var (
		n   int
		err error
	)
	if l.isConsole {
		l.outConsole.Write([]byte(v))
	}
	if l.isFile {
		n, err = l.outFile.Write([]byte(v))
		if err != nil {
			panic("写文件出错")
		}
	}
	return n
}

func (l *logger) fileCheck() {
	defer catchError()

	if l.isMustRenameDate() {
		l.renameDate()
	}

	if l.isMustRenameSize() {
		l.renameSize()
	}

}

func (l *logger) isMustRenameSize() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.fileCurrSize >= l.fileMaxSize {
		return true
	}
	return false
}

func (l *logger) isMustRenameDate() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	t, _ := time.Parse(_DATEFORMAT, GetNowUnixTimeOBJ().Format(_DATEFORMAT))
	if t.After(*l._date) {
		l._date = &t
		return true
	}
	return false
}

//按日期
func (l *logger) renameDate() {
	defer catchError()
	l.mu.Lock()
	defer l.mu.Unlock()
	l.close()
	l.openFile()
}

//按大小
func (l *logger) renameSize() {
	defer catchError()
	l.mu.Lock()
	defer l.mu.Unlock()
	l.close()

	// 检测此文件是否已经存在
	num := 1
	for isExists(l.getSizeFileFullName(strconv.Itoa(num))) {
		//strconv.Itoa(l.nSize)
		//l.nSize++
		num++
	}

	// 删除大于fileMaxNSize的文件
	for ; num > l.fileMaxNSize; num-- {
		os.Remove(l.getSizeFileFullName(strconv.Itoa(num)))
	}

	// 重命名文件
	for ; num > 1; num-- {
		os.Rename(l.getSizeFileFullName(strconv.Itoa(num-1)), l.getSizeFileFullName(strconv.Itoa(num)))
	}

	os.Rename(l.getFileFullName(), l.getSizeFileFullName("1"))
	l.openFile()
	l.nSize++
	l.flush()
}

func (l *logger) flush() {
	l.fileCurrSize = 0
}

func (l *logger) close() {
	l.logfile.Close()
}

package QLog

//默认值
const (
	_DATEFORMAT        = "2006-01-02" //默认文件日期格式
	DEFAULTFILENAME    = "info"       //默认文件前缀
	DEFAULTFILEPATH    = "./"         //默认文件存放路径
	DEFAULTFILESUFFIX  = "log"        //默认文件后缀
	DEFAULTFILEMAXSIZE = 2 << 29      //默认单个文件最大1G
	callDep            = 3            //默认深度
)

const (
	ALL uint8 = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	ALERT
	OFF
)

// alert level
const (
	ALERTALERT    = "notice"   // notice
	ALERTWARNING  = "warning"  // 严重
	ALERTCRITICAL = "critical" // 紧急
)

const (
	DIAGONAL   = "/"
	COLON      = ":"
	BLANK      = " "
	NEWLINE    = "\n"
	DOT        = "."
	UNDERSCODE = "_"
	DASH       = "-"
)

const (
	TYPEDEBUG = "DEBUG"
	TYPEINFO  = "INFO"
	TYPEWARN  = "WARN"
	TYPEERROR = "ERROR"
	TYPEFATAL = "FATAL"
	TYPEALERT = "ALERT"
)
const COLOR_TPL = "\x1b[%dm%s\x1b[0m"
const (
	COLOR_RED = uint8(iota + 91)
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA  //洋红
	COLOR_INTENSITY
)

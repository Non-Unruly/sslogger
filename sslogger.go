// Author : Non-Unruly@GitHub.com
package sslogger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	DBG = 0
	INF = 1
	WAR = 2
	ERR = 3
	DIS = 4
)

var logfd *os.File = nil
var lv int32

func LogInitialize(logPath string, level int32) int32 {
	var err error
	logfd, err = os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("create log file error: ", err)
		return -1
	}
	log.SetOutput(logfd)
	log.Print("")
	log.Print("====================Start==========================")
	log.Print(time.Now().String())
	log.SetFlags(log.Ldate)
	log.SetFlags(log.Ltime)
	return 0
}

func Release() {
	logfd.Close()
}

func Debug(format string, args ...interface{}) string {
	str := output(DBG, true, format, args...)
	return str
}

func Debugnp(format string, args ...interface{}) string {
	str := output(DBG, false, format, args...)
	return str
}

func Info(format string, args ...interface{}) string {
	str := output(INF, true, format, args...)
	return str
}

func Infonp(format string, args ...interface{}) string {
	str := output(INF, false, format, args...)
	return str
}

func Warn(format string, args ...interface{}) string {
	str := output(WAR, true, format, args...)
	return str
}

func Warnnp(format string, args ...interface{}) string {
	str := output(WAR, false, format, args...)
	return str
}

func Error(format string, args ...interface{}) string {
	str := output(ERR, true, format, args...)
	return str
}
func Errornp(format string, args ...interface{}) string {
	str := output(ERR, false, format, args...)
	return str
}

func Disaster(format string, args ...interface{}) string {
	str := output(DIS, true, format, args...)
	return str
}

func Disasternp(format string, args ...interface{}) string {
	str := output(DIS, false, format, args...)
	return str
}

// output 输出日志字符串
// level-日至登记，print-是否打印到屏幕 ，format-格式字符串，args参数
func output(level int32, print bool, format string, args ...interface{}) string {
	if lv > level || logfd == nil {
		return ""
	}

	var text string
	text = ""
	if args != nil {
		text = fmt.Sprintf(format, args...)
	} else {
		text = format
	}

	var prefix string
	switch level {
	case DBG:
		prefix = "[DBG]"
		break
	case INF:
		prefix = "[INF]"
		break
	case WAR:
		prefix = "[WAR]"
		break
	case ERR:
		prefix = "[ERR]"
		break
	case DIS:
		prefix = "[DIS]"
		break
	}
	_, file, line, _ := runtime.Caller(2)
	lst := strings.Split(file, "server/src/")
	if len(lst) == 2 {
		file = lst[1]
	}
	str := fmt.Sprintf("[%s:%d] %s\n", file, line, text)
	log.SetPrefix(prefix)
	log.Print(str)

	if print {
		fmt.Print(str)
	}
	if level == DIS {
		log.Fatal("EXIT")
	}
	return str
}

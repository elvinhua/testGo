/*
 * @title:mylogger
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-10 10:58:26
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-22 14:52:54
 */

/**
*	1.支持往不同的地方输出日志
*	2.日志级别
*		1).Debug
*		2).Trace
*		3).Info
*		4).Warning
*		5).Error
*		6).Fatal
*	3.日志要支持开关控制（测试 or 生产）
*	4.完整的日志记录包含 时间、行号、日志级别、日志信息（请求、参数、返回）
*	5.大数据文件需要切割
 */
package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

// logType :=[6]string{"Debug","Trace","Info","Waring","Error","Fatal"}
type LogType interface {
	Debug(format string, params ...interface{})
	Trace(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warning(format string, params ...interface{})
	Error(format string, params ...interface{})
	Fatal(format string, params ...interface{})
}

const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARING
	ERROR
	FATAL
)

//字符串转LogLevel
func parseLogType(s string) (l LogLevel, err error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOW, err
	}
}

// LogLevel 转字符串
func logLevelToString(l LogLevel) string {
	switch l {
	case DEBUG:
		return "debug"
	case TRACE:
		return "trace"
	case INFO:
		return "info"
	case WARING:
		return "waring"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	default:
		return "debug"

	}
}

func getFileInfo(skip int) (funcName string, filePath string, lineNum int) {
	// 获取请求信息
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("getFileInfo for failure")
	}
	// 将返回的uintptr值转为所在方法名例：main.main
	funcName = runtime.FuncForPC(pc).Name()
	// 将返回的方法名分割
	funcName = strings.Split(funcName, ".")[1]
	// 返回文件所在路径
	filePath = path.Base(file)
	// 返回调用次方法所在行
	lineNum = line
	return
}

/**
 * @name: 获取当前格式化时间
 * @test: test font
 * @msg:
 * @param {*}
 * @return {time.time}
 */
func nowTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

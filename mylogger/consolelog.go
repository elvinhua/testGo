/*
 * @title:往终端写日志
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-10 10:59:22
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-11 19:02:56
 */
package mylogger

import (
	"fmt"
)

// 日志配置
type ConsoleLogger struct {
	Level LogLevel
}

// 构造日志函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogType(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) log(l LogLevel, format string, params ...interface{}) {
	if c.enable(l) {
		logMsg := fmt.Sprintf(format, params...)
		s := logLevelToString(l)
		funcName, filePath, lineNum := getFileInfo(3)
		// fmt.Printf("[%s][%s][filePath:%v;funcName:%s;line:%d]:%s\n", s, nowTime(), funcName, filePath, lineNum, logMsg)
		fmt.Printf("[%s][%s]\n", s, nowTime())
		fmt.Printf("[filePath:%v;funcName:%s;line:%d]:%s\n", funcName, filePath, lineNum, logMsg)
	}
}
func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c ConsoleLogger) Debug(format string, params ...interface{}) {
	c.log(DEBUG, format, params...)
}
func (c ConsoleLogger) Trace(format string, params ...interface{}) {
	c.log(TRACE, format, params...)
}
func (c ConsoleLogger) Info(format string, params ...interface{}) {
	c.log(INFO, format, params...)
}
func (c ConsoleLogger) Warning(format string, params ...interface{}) {
	// fmt.Printf("[War][%s]:%s\n", nowTime(), logMsg)
	c.log(WARING, format, params...)
}
func (c ConsoleLogger) Error(format string, params ...interface{}) {
	// fmt.Printf("[Err][%s]:%s\n", nowTime(), logMsg)
	c.log(ERROR, format, params...)
}
func (c ConsoleLogger) Fatal(format string, params ...interface{}) {
	// fmt.Printf("[Fat][%s]:%s\n", nowTime(), logMsg)
	c.log(FATAL, format, params...)
}

/*
 * @title:
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-10 17:51:22
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-21 16:37:50
 */
package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 定义一个文件结构体
// 级别、文件路径、文件名、日志文件、错误日志文件、最大文件大小
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFile     *os.File
	maxFileSize int64
}

/**
 * @name:创建结构体构造方法
 * @test: test font
 * @msg:
 * @param {*} levelStr 日志级别
 * @param {*} fp 日志路径
 * @param {string} fn 日志文件名
 * @param {int64} fileSize 日志大小
 * @return {*} 返回定义好的结构体
 */
func NewFileLogger(levelStr, fp, fn string, fileSize int64) *FileLogger {
	level, err := parseLogType(levelStr)
	if err != nil {
		panic(err)
	}
	configFileData := &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: fileSize,
	}
	// 将打开的文件放入补录如结构体并返回
	err = configFileData.initFileConfig()
	if err != nil {
		panic(err)
	}
	return configFileData
}

/**
 * @name:创建打开文件的方法用于配置结构体参数
 * @test: test font
 * @msg:
 * @param {*}
 * @return {*}
 */
func (f *FileLogger) initFileConfig() error {
	fullFilePath := path.Join(f.filePath, f.fileName)
	file, err := os.OpenFile(fullFilePath+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	f.fileObj = file
	if err != nil {
		fmt.Println("open the file is unsuccessful", err)
		return err
	}
	// f.errFile = err
	errfile, err := os.OpenFile(fullFilePath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open the error file is unsuccessful", err)
		return err
	}
	f.errFile = errfile
	return nil
}
func (f *FileLogger) cuttingFile(tfile *os.File) bool {
	fileInfo, err := tfile.Stat()
	if err != nil {
		fmt.Println("Failed to query file information. Procedure:", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

/**
 * @name: 切割文件
 * @test: test font
 * @msg:
 * @param {*os.File} newFile
 * @return {*}
 */
func (f *FileLogger) splicFile(file *os.File) (*os.File, error) {
	// 创建一个时间戳
	timeUnix := time.Now().Format("20060102150405000")
	fileData, err := file.Stat()
	if err != nil {
		fmt.Println("edit filename is eror", err)
		return nil, err
	}
	oldFileName := path.Join(f.filePath, fileData.Name())
	newCutFileName := fmt.Sprintf("%s.bak%s", oldFileName, timeUnix)
	file.Close()
	err = os.Rename(oldFileName, newCutFileName)
	// fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	newCutFile, err := os.OpenFile(oldFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("create cut file is unsuccessful", err)
		return nil, err
	}
	return newCutFile, nil
}

// 最终写入日志的方法
func (f *FileLogger) log(l LogLevel, format string, params ...interface{}) {
	if f.enable(l) {
		logMsg := fmt.Sprintf(format, params...)
		s := logLevelToString(l)
		funcName, filePath, lineNum := getFileInfo(3)
		if f.cuttingFile(f.fileObj) {
			newCutFile, err := f.splicFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newCutFile
		}
		fmt.Fprintf(f.fileObj, "[%s][%s]\n", s, nowTime())
		fmt.Fprintf(f.fileObj, "[filePath:%v;funcName:%s;line:%d]:%s\n", funcName, filePath, lineNum, logMsg)
		if l >= ERROR {
			if f.cuttingFile(f.errFile) {
				newCutFile, err := f.splicFile(f.errFile)
				if err != nil {
					return
				}
				f.errFile = newCutFile
			}
			fmt.Fprintf(f.errFile, "[%s][%s]\n", s, nowTime())
			fmt.Fprintf(f.errFile, "[filePath:%v;funcName:%s;line:%d]:%s\n", funcName, filePath, lineNum, logMsg)
		}
	}
}

// 判断当前级别与输入的级别大小
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) Debug(format string, params ...interface{}) {
	f.log(DEBUG, format, params...)
}
func (f *FileLogger) Trace(format string, params ...interface{}) {
	f.log(TRACE, format, params...)
}
func (f *FileLogger) Info(format string, params ...interface{}) {
	f.log(INFO, format, params...)
}
func (f *FileLogger) Warning(format string, params ...interface{}) {
	// fmt.Printf("[War][%s]:%s\n", nowTime(), logMsg)
	f.log(WARING, format, params...)
}
func (f *FileLogger) Error(format string, params ...interface{}) {
	// fmt.Printf("[Err][%s]:%s\n", nowTime(), logMsg)
	f.log(ERROR, format, params...)
}
func (f *FileLogger) Fatal(format string, params ...interface{}) {
	// fmt.Printf("[Fat][%s]:%s\n", nowTime(), logMsg)
	f.log(FATAL, format, params...)
}
func (f *FileLogger) CloseFile() {
	f.fileObj.Close()
	f.errFile.Close()
}

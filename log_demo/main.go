/*
 * @title:日志
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-10 10:17:21
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-11 19:05:34
 */
package main

import "ibuyd.com/Guoyh/test/mylogger"

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
var log mylogger.LogType

func main() {
	// logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println("写入日志失败，错误值为：", err)
	// }
	// // defer logFile.Close()
	// log.SetOutput(logFile)
	log = mylogger.NewFileLogger("info", "./", "testlog", 10*1024)
	// log = mylogger.NewConsoleLog("info")
	id := 10086
	name := "大漂亮"
	for {
		log.Debug("显示Debug日志")
		log.Trace("显示Trace日志")
		log.Info("显示Info日志")
		log.Warning("显示Warning日志,ID：%d,姓名：%s,请求报错", id, name)
		log.Error("显示Error日志")
		log.Fatal("显示Fatal日志")
	}

}

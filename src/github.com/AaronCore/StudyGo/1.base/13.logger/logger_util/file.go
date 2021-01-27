package logger_util

// 往文件写Log

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// FileLogger 构造函数
func NewFileLogger(levelStr, filePath, fileName string, maxFileSize int64) *FileLogger {
	level, err := parseLoggerLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

// 打开日志文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file failed,error：%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".error", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open error log file failed,error：%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) log(level LogLevel, format string, a ...interface{}) {
	if f.enable(level) {
		msg := fmt.Sprintf(format, a...) // 格式化输出
		levelStr := getLoggerString(level)
		funcName, fileName, lineNo := getRunInfo(3)
		if f.checkFileSize(f.fileObj) {

			//f.fileObj.Close()
			//logName := path.Join(f.filePath, f.fileName)
			//newLogName := fmt.Sprintf("%s.bak%s", logName, time.Now().Format("20060102150405000"))
			//os.Rename(logName, newLogName)
			//fileObj, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			//if err != nil {
			//	fmt.Printf("open new log file failed,error：%v\n", err)
			//	return
			//}
			//f.fileObj = fileObj

			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s \n", time.Now().Format("2006-01-02 15:04:05"), levelStr, fileName, funcName, lineNo, msg)
		if level >= ERROR {
			if f.checkFileSize(f.fileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			// 如果记录日志的等级大于ERROR级别，则要在err目录中记录
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s \n", time.Now().Format("2006-01-02 15:04:05"), levelStr, fileName, funcName, lineNo, msg)
		}
	}
}

// 切割日志文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 切割日志文件
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,error：%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, time.Now().Format("20060102150405000"))

	file.Close() // 关闭当前日志文件

	os.Rename(logName, newLogName)
	fileObj, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,error：%v\n", err)
		return nil, err
	}
	return fileObj, err
}

// 判断是否需要切割日志
func (f *FileLogger) checkFileSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,error：%v", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

// 判断是否需要记录日志
func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Warn(format string, a ...interface{}) {
	f.log(WARN, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

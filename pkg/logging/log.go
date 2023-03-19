package logging

import (
	"fmt"
	"jaingke2023.com/BlogService/pkg/settings"
	"log"
	"os"
	"time"
)

const (
	INFO = iota
	ERROR
	DEBUG
	WARNING
	FATAL
)

var (
	logPath     string
	filename    string
	filePerm    os.FileMode
	logFullPath string
	CallDepth   int
	
	file *os.File
	err  error
)

func init() {
	logSection := settings.Cfg.Section("log")
	logPath = logSection.Key("LOG_FILE_PATH").MustString("runtime")
	filename = logSection.Key("LOG_FILE_NAME").MustString("blogService.log")
	CallDepth = logSection.Key("LOG_CALL_DEPTH").MustInt(2)
	filePerm = os.FileMode(logSection.Key("LOG_FILE_PERM").MustUint(0666))
	
	logFullPath = fmt.Sprintf("%s\\%s%s", logPath, time.Now().Format("2006-01-02"), filename)
	fmt.Println(logFullPath)
	
	err = os.MkdirAll(logPath, filePerm)
	if err != nil {
		panic("日志目录创建失败： " + err.Error())
	}
	file, err = os.OpenFile(logFullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, filePerm)
	if err != nil {
		panic("日志文件创建失败： " + err.Error())
	}
}

var logger *log.Logger

var prefixFlag = []string{"[INFO]", "[ERROR]", "[DEBUG]", "[WARNING]", "[FATAL]"}

func init() {
	logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(a ...any) {
	s := fmt.Sprintln(a)
	
	logger.SetPrefix(prefixFlag[INFO])
	logger.Output(CallDepth, s)
}

func Warning(a ...any) {
	s := fmt.Sprintln(a)
	
	logger.SetPrefix(prefixFlag[WARNING])
	logger.Output(CallDepth, s)
}

func Error(a ...any) {
	s := fmt.Sprintln(a)
	
	logger.SetPrefix(prefixFlag[ERROR])
	logger.Output(CallDepth, s)
}
func Fatal(a ...any) {
	
	s := fmt.Sprintln(a)
	
	logger.SetPrefix(prefixFlag[FATAL])
	logger.Output(CallDepth, s)
}

func Debug(a ...any) {
	//_, file, line, ok := runtime.Caller(2)
	//if ok {
	//	logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	//} else {
	//	logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	//}
	s := fmt.Sprintln(a)
	
	logger.SetPrefix(prefixFlag[DEBUG])
	logger.Output(CallDepth, s)
}

package main

import (
	"admin_backend/global"
	"admin_backend/router"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

type FileHook struct {
	ch chan []byte
}

func NewFileHook() *FileHook {
	w := &FileHook{ch: make(chan []byte, 1000)}
	go w.processLogs()
	return w
}

func (w *FileHook) Write(p []byte) (n int, err error) {
	w.ch <- p
	return len(p), nil
}

func (w *FileHook) processLogs() {
	var currentFile *os.File
	var currentDate string
	defer func() {
		if currentFile != nil {
			_ = currentFile.Close()
		}
	}()
	for p := range w.ch {
		today := time.Now().Format("2006_01_02")
		if currentDate != today {
			if currentFile != nil {
				_ = currentFile.Close()
			}
			filename := "./logs/" + today + ".log"
			f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Println("failed to open log file:", err)
				continue
			}
			currentFile = f
			currentDate = today
		}
		_, _ = currentFile.Write(p)
	}
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if _, osPathErr := os.Stat("./logs"); os.IsNotExist(osPathErr) {
		if osMkdirErr := os.Mkdir("./logs", 0777); osMkdirErr != nil {
			log.Fatalln("os mkdir error")
		}
	}

	multiWriter := io.MultiWriter(os.Stdout, NewFileHook())
	log.SetOutput(multiWriter)

	global.New()
	gin.SetMode(gin.ReleaseMode)
	
	engine := gin.Default()
	router.NewHttpRouter(engine)
	address := "0.0.0.0:34332"
	log.Println("http server start ", address)
	if err := engine.Run(address); err != nil {
		log.Println("http server start failed", address)
	}
}

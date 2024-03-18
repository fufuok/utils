package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/xfile"
)

var (
	recorder *xfile.Roller
	filePath = filepath.Join(utils.ExecutableDir(true), "log")
)

func main() {
	initRecorder()
	testRecorder()
}

func initRecorder() {
	filenameMaker := &xfile.TimeBasedFilename{
		FilePath:    filePath,
		FilenameTpl: "time_%s.log",
		TimeTpl:     "150405",
	}
	_ = os.MkdirAll(filenameMaker.FilePath, 0o755)
	fmt.Printf("%+v\n", filenameMaker)
	filename := filenameMaker.MakeFilename("")
	opt := &xfile.Options{
		FilenameMaker:  filenameMaker,
		FlushSizeLimit: xfile.DefaultFlushSizeLimit,
		FlushInterval:  xfile.DefaultFlushInterval,
	}
	var err error
	recorder, err = xfile.NewRoller(filename, opt)
	if err != nil {
		log.Fatalln(err)
	}
}

func testRecorder() {
	log.Println("start")
	dataChan := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
		recorder.Close()
		log.Println("end")
	}()
	go func() {
		defer close(dataChan)
		for {
			select {
			default:
			case <-ctx.Done():
				return
			}
			n := utils.FastIntn(100)
			dataChan <- fmt.Sprintf("time: %s, data: %s\n", time.Now(), utils.RandString(n))
			time.Sleep(time.Duration(n))
		}
	}()
	for s := range dataChan {
		_, _ = recorder.WriteString(s)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/xfile"
)

var (
	recorder *xfile.Roller
	filePath = filepath.Join(utils.ExecutableDir(true), "log")
)

type nameMaker struct {
	i int
}

func (m *nameMaker) MakeFilename(_ string) string {
	m.i++
	return filepath.Join(filePath, strconv.Itoa(m.i%5)+".log")
}

func main() {
	initRecorder()
	testRecorder()
}

func initRecorder() {
	_ = os.MkdirAll(filePath, 0o755)
	filenameMaker := new(nameMaker)
	filename := filenameMaker.MakeFilename("")
	opt := &xfile.Options{
		FilenameMaker: filenameMaker,
		Rebuild:       true,
		FlushInterval: xfile.MinFlushInterval,
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

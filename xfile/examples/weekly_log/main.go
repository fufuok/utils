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
	weekTag  time.Weekday
)

type weeklyMaker struct{}

func (m *weeklyMaker) MakeFilename(name string) string {
	week := time.Now().Weekday()
	if week == weekTag {
		return name
	}
	weekTag = week
	filename := fmt.Sprintf("weekly_%s.log", week.String())
	return filepath.Join(filePath, filename)
}

func main() {
	initRecorder()
	testRecorder()
}

func initRecorder() {
	_ = os.MkdirAll(filePath, 0755)
	filenameMaker := new(weeklyMaker)
	filename := filenameMaker.MakeFilename("")
	opt := &xfile.Options{
		FilenameMaker: filenameMaker,
		Rebuild:       true,
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

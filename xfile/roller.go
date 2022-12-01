package xfile

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	DefaultFlushSizeLimit = 1 << 20
	DefaultFlushInterval  = 1 * time.Second
	MinFlushSizeLimit     = 4096
	MinFlushInterval      = 100 * time.Millisecond
)

var (
	ErrFilename = errors.New("wrong file name")
)

type FilenameMaker interface {
	MakeFilename(name string) string
}

type stdLogger struct{}

func (s *stdLogger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

type Options struct {
	// 文件名生成器, 每秒检查
	FilenameMaker FilenameMaker

	// 日志处理器
	Logger Logger

	// 每次(实例启动时除外)滚动的文件是否删除后重建
	Rebuild bool

	// 刷新到磁盘的大小和时间间隔, 默认 1MiB, 1秒
	// 注意: 由于写文件有缓冲, 如果按秒切割文件, 数据将有可能写入上一秒的文件名中
	FlushSizeLimit int
	FlushInterval  time.Duration
}

type Logger interface {
	Errorf(format string, v ...interface{})
}

type Roller struct {
	name   string
	maker  FilenameMaker
	logger Logger

	file      *os.File
	writer    *bufio.Writer
	rebuild   bool
	firstOpen bool

	flushSizeLimit int
	flushInterval  time.Duration

	mu   sync.Mutex
	stop chan struct{}
}

func NewRoller(filename string, opt *Options) (*Roller, error) {
	if filename == "" {
		return nil, ErrFilename
	}
	r := &Roller{
		name:      filename,
		firstOpen: true,
	}
	if err := r.openNewFile(); err != nil {
		return nil, err
	}
	r.setup(opt)
	r.stop = make(chan struct{}, 1)

	go r.flushTimer()
	return r, nil
}

func (r *Roller) setup(opt *Options) {
	if opt == nil {
		return
	}
	if opt.FilenameMaker != nil {
		r.maker = opt.FilenameMaker
	} else {
		r.maker = new(DefaultFilename)
	}
	r.rebuild = opt.Rebuild
	if opt.FlushSizeLimit < MinFlushSizeLimit {
		if opt.FlushSizeLimit == 0 {
			r.flushSizeLimit = DefaultFlushSizeLimit
		} else {
			r.flushSizeLimit = MinFlushSizeLimit
		}
	} else {
		r.flushSizeLimit = opt.FlushSizeLimit
	}
	if opt.FlushInterval < MinFlushInterval {
		if opt.FlushInterval == 0 {
			r.flushInterval = DefaultFlushInterval
		} else {
			r.flushInterval = MinFlushInterval
		}
	} else {
		r.flushInterval = opt.FlushInterval
	}
	if opt.Logger != nil {
		r.logger = opt.Logger
	} else {
		r.logger = new(stdLogger)
	}
}

func (r *Roller) flushTimer() {
	ticker := time.NewTicker(r.flushInterval)
	defer ticker.Stop()
	for {
		select {
		case <-r.stop:
			return
		case <-ticker.C:
			r.flush()
		}
	}
}

func (r *Roller) flush() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.writer.Size() == 0 {
		return
	}
	if err := r.writer.Flush(); err != nil {
		r.logger.Errorf("Failed to write file: %v", err)
	}
	name := r.maker.MakeFilename(r.name)
	if r.name != name {
		// 滚动文件
		r.name = name
		if err := r.openNewFile(); err != nil {
			r.logger.Errorf("Unable to create new file: %v", err)
		}
	}
}

func (r *Roller) Write(p []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.writer.Write(p)
}

func (r *Roller) WriteString(s string) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.writer.WriteString(s)
}

func (r *Roller) openNewFile() error {
	if r.rebuild && !r.firstOpen {
		r.firstOpen = false
		if err := os.Remove(r.name); err != nil {
			r.logger.Errorf("rebuild file: %s, err: %v", r.name, err)
		}
	}
	file, err := os.OpenFile(r.name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	_ = r.file.Close()
	r.file = file
	r.writer = bufio.NewWriterSize(r.file, r.flushSizeLimit)
	return nil
}

func (r *Roller) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.writer != nil {
		_ = r.writer.Flush()
	}
	_ = r.file.Close()
	close(r.stop)
}

type DefaultFilename struct{}

func (d *DefaultFilename) MakeFilename(name string) string {
	return name
}

type TimeBasedFilename struct {
	// 文件目录
	FilePath string

	// 文件名模板: run-%s.log
	FilenameTpl string

	// 日期时间模板: 060102
	TimeTpl string

	// 时间标识相同则不生成新文件名: 221130
	TimeTag string
}

func (t *TimeBasedFilename) MakeFilename(name string) string {
	tag := time.Now().Format(t.TimeTpl)
	if tag == t.TimeTag {
		return name
	}
	t.TimeTag = tag
	newName := fmt.Sprintf(t.FilenameTpl, tag)
	if t.FilePath != "" {
		newName = filepath.Join(t.FilePath, newName)
	}
	return newName
}

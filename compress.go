package utils

import (
	"compress/gzip"
	"compress/zlib"
	"io/ioutil"
	"sync"

	"github.com/fufuok/utils/pools/bufferpool"
	"github.com/fufuok/utils/pools/readerpool"
)

var (
	gzipWritePool  = newGzipWriterPool()
	zlibWritePool  = newZlibWriterPool()
	gzipReaderPool = sync.Pool{
		New: func() interface{} {
			return new(gzip.Reader)
		},
	}
)

func Gzip(data []byte) ([]byte, error) {
	return GzipLevel(data, gzip.BestSpeed)
}

func GzipLevel(data []byte, level int) (dst []byte, err error) {
	buf := bufferpool.Get()
	idx := getWriterPoolIndex(level)
	zw := gzipWritePool[idx].Get().(*gzip.Writer)
	zw.Reset(buf)
	defer func() {
		bufferpool.Put(buf)
		gzipWritePool[idx].Put(zw)
	}()

	_, err = zw.Write(data)
	if err != nil {
		return
	}
	err = zw.Flush()
	if err != nil {
		return
	}
	err = zw.Close()
	if err != nil {
		return
	}

	dst = CopyBytes(buf.Bytes())
	return
}

func Ungzip(data []byte) (src []byte, err error) {
	rData := readerpool.New(data)
	zr := gzipReaderPool.Get().(*gzip.Reader)
	defer func() {
		readerpool.Release(rData)
		gzipReaderPool.Put(zr)
	}()

	err = zr.Reset(rData)
	if err != nil {
		return
	}
	defer func() {
		_ = zr.Close()
	}()

	src, err = ioutil.ReadAll(zr)
	if err != nil {
		return
	}
	return
}

func Zip(data []byte) ([]byte, error) {
	return ZipLevel(data, zlib.BestSpeed)
}

func ZipLevel(data []byte, level int) (dst []byte, err error) {
	buf := bufferpool.Get()
	idx := getWriterPoolIndex(level)
	zw := zlibWritePool[idx].Get().(*zlib.Writer)
	zw.Reset(buf)
	defer func() {
		bufferpool.Put(buf)
		zlibWritePool[idx].Put(zw)
	}()

	_, err = zw.Write(data)
	if err != nil {
		return
	}
	err = zw.Flush()
	if err != nil {
		return
	}
	err = zw.Close()
	if err != nil {
		return
	}

	dst = CopyBytes(buf.Bytes())
	return
}

func Unzip(data []byte) (src []byte, err error) {
	rData := readerpool.New(data)
	defer readerpool.Release(rData)
	zr, err := zlib.NewReader(rData)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = zr.Close()
	}()

	src, err = ioutil.ReadAll(zr)
	if err != nil {
		return
	}
	return
}

func newZlibWriterPool() (pools []*sync.Pool) {
	for i := 0; i < 12; i++ {
		level := i - 2
		pools = append(pools, &sync.Pool{
			New: func() interface{} {
				zw, _ := zlib.NewWriterLevel(nil, level)
				return zw
			},
		})
	}
	return
}

func newGzipWriterPool() (pools []*sync.Pool) {
	for i := 0; i < 12; i++ {
		level := i - 2
		pools = append(pools, &sync.Pool{
			New: func() interface{} {
				zw, _ := gzip.NewWriterLevel(nil, level)
				return zw
			},
		})
	}
	return
}

func getWriterPoolIndex(level int) int {
	if level < gzip.HuffmanOnly || level > gzip.BestCompression {
		level = gzip.DefaultCompression
	}
	return level + 2
}

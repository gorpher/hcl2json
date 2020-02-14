package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorpher/hcl2json/log"
)

// 使用log的打印
var printf = log.Printf

func ExistFile(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return nil, false
	}
	return f, true
}

func MkdirAll(f string) {
	printf("[debug] mkdirAll() create dir [%s] \n", f)
	info, ok := ExistFile(f)
	if !ok {
		err := os.MkdirAll(f, os.ModePerm)
		log.Fatalln(err)
		printf("[debug] mkdirAll() dir %s created\n", f)
		return
	}
	if ok && !info.IsDir() {
		log.Fatalln(fmt.Errorf("[error] mkdirAll() [%s] is not a dir", f))
	}
	printf("[debug] mkdirAll() skip dir [%s] , this dir exist\n", f)
}

// 删除文件所有后缀名
func RemoveFileExt(filename string) string {
	if filepath.Ext(filename) == "" {
		return filename
	}
	return RemoveFileExt(strings.ReplaceAll(filepath.Clean(filename), filepath.Ext(filename), ""))
}

// 读取文件
func ReaderFile(f, dest string) ([]byte, *os.File, error) {
	printf("[debug] readerFile( file=[%s] dest=[%s] )\n", f, dest)
	in, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, nil, fmt.Errorf("read file err : %v", err)
	}
	out, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("open file err : %v", err)
	}
	return in, out, nil
}

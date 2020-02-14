package converter

import (
	"fmt"
	"github.com/gorpher/hcl2json/hclutils"
	"github.com/gorpher/hcl2json/log"
	"github.com/gorpher/hcl2json/util"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// 使用log的打印
var printf = log.Printf
var mkdirAll = util.MkdirAll
var removeFileExt = util.RemoveFileExt

// hcl后缀文件
var HCLFileExt = []string{".tf", ".hcl"}

// json后缀文件
var JSONFileExt = []string{".json"}

// singleConvert 单文件转换
// reverse 参数为默认false hcl转json 否则json转hcl
// input 输入字节 , out输出
func SingleConvert(reverse bool, input []byte, out io.Writer) error {
	printf("[debug] singleConvert( reverse=[%v] input=... out=... )\n", reverse)
	if reverse {
		return hclutils.ToHcl(input, out)
	}
	return hclutils.ToJson(input, out)
}

// multiConvertV1 多文件转换, 使用filepath.Walk函数遍历根文件夹和子文件夹进行文件转换
// reverse 参数为默认false hcl转json 否则json转hcl
// src 输入目录 , dest输出目录 且必须是绝对路径,dest必须为输出目录
func MultiConvertV1(reverse bool, src, dest string) error {
	printf("[debug] multiConvertV1( reverse=[%v] src=[%s] dest=[%s] )\n", reverse, src, dest)
	if !filepath.IsAbs(src) || !filepath.IsAbs(dest) {
		return fmt.Errorf("src=%s or dest=%s is not absolute path", src, dest)
	}
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断是否符合条件的文件
		canConvert := IsIdealFile(reverse, info)

		// 处理hcl文件
		if !info.IsDir() && canConvert {
			var destFile = ""
			// 判断文件夹是否存在

			rel, err := filepath.Rel(src, filepath.Dir(path))
			log.Fatalln(err)
			printf("[debug]  multiConvertV1() filepath.Rel(path,src) = %s\n", rel)
			baseDir := filepath.Join(dest, rel)
			mkdirAll(baseDir)
			if !reverse {
				destFile = filepath.Join(baseDir, info.Name()+".json")
			} else {
				destFile = filepath.Join(baseDir, removeFileExt(info.Name())+".tf")
			}
			//读取文件流
			path := filepath.Clean(path)
			in, out, e := util.ReaderFile(path, destFile)
			if e != nil {
				return e
			}
			defer out.Close() //nolint
			return SingleConvert(reverse, in, out)
		}
		return nil
	})
}

// multiConvertV2 多文件转换, 并发方式实现
// reverse 参数为默认false hcl转json 否则json转hcl
// src 输入目录 , dest输出目录 且dest必须为输出目录
func MultiConvertV2(reverse bool, src, dest string) error {
	printf("[debug] multiConvertV2( reverse=[%v] src=[%s] dest=[%s])\n", reverse, src, dest)
	dirs, err := ioutil.ReadDir(src)
	if err != nil {
		return fmt.Errorf("[error] multiConvertV2() read file err : %v", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(len(dirs))
	for i := range dirs {
		go func(i int) {
			defer wg.Done()
			info := dirs[i]
			// 判断是否符合条件的文件
			canConvert := IsIdealFile(reverse, info)
			// 处理hcl文件
			if !info.IsDir() && canConvert {
				//读取文件流
				var destFile = ""
				if !reverse {
					destFile = filepath.Join(dest, filepath.Clean(info.Name())+".json")
				} else {
					destFile = filepath.Join(dest, removeFileExt(info.Name())+".tf")
				}
				path := filepath.Join(src, info.Name())
				in, out, e := util.ReaderFile(path, destFile)
				if e != nil {
					log.Fatalf("[error] multiConvertV2() go routine read file byte err: %v\n", e)
				}
				defer out.Close() //nolint
				e = SingleConvert(reverse, in, out)
				if e != nil {
					log.Fatalf("[error] multiConvertV2() convert file err: %v\n", e)
				}
			}
			if info.IsDir() {
				innerSrcDir := filepath.Join(src, info.Name())
				innerDestDir := filepath.Join(dest, info.Name())
				mkdirAll(innerDestDir)
				err := MultiConvertV2(reverse, innerSrcDir, innerDestDir)
				if err != nil {
					log.Fatalf("[error] multiConvertV2() convert dir %s err: %v\n", innerSrcDir, err)
				}
			}
		}(i)
	}
	wg.Wait()
	return nil
}

// IsIdealFile 是否是理想文件 reverse=true 必须是json文件,否则是hcl文件
func IsIdealFile(reverse bool, info os.FileInfo) bool {
	if !info.IsDir() {
		exts := HCLFileExt
		if reverse {
			exts = JSONFileExt
		}
		for i := range exts {
			if exts[i] == filepath.Ext(info.Name()) {
				return true
			}
		}
	}
	return false
}

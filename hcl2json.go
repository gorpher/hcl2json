package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gorpher/hcl2json/converter"
	"github.com/gorpher/hcl2json/log"
	"github.com/gorpher/hcl2json/util"
)

// VERSION is what is returned by the `-v` flag
var Version = "v1.0.0"

// go build .  -o hclcli.exe
// go run hclcli.go -i fixtures/single_json/infra.tf.json -o infra.tf
// go run hclcli.go -i fixtures/single_hcl/infra.tf -o ddd
// go run hclcli.go -i fixtures/ -o ddd
var version = flag.Bool("version", false, "Prints current app version")
var in = flag.String("i", "", "input file or dir, require")
var out = flag.String("o", "", "output file or dir , default print")
var reverse = flag.Bool("reverse", false, "default false,if true json to hcl")
var d = flag.Bool("debug", false, "default false,if true open debug mode")

var printf = log.Printf
var exist = util.ExistFile

// go build .
// hcl2json -i test-fixtures/hcl_outputs/outputs.tf -o outputs.tf.json
// hcl2json -i test-fixtures/json_outputs/outputs.tf.json -o outputs.tf
// hcl2json -i test-fixtures/hcl_variables/variables.tf
// hcl2json -i test-fixtures/json_terraform/terraform.tf.json
// hcl2json -i test-fixtures/ -o test-fixtures-01
// hcl2json -i test-fixtures-01 -reverse -debug
// hcl2json -i test-fixtures/ -o test-fixtures-02 -reverse
// hcl2json -i test-fixtures-02
func main() {
	flag.Parse()
	var err error
	if *version {
		fmt.Println(Version)
		return
	}
	log.DebugMode = *d
	printf("[debug] params in=[%s] out=[%s]\n", *in, *out)
	if *in == "" {
		err = fmt.Errorf("file not exist: usage : %s -i filename\n", os.Args[0])
	}
	log.Fatalln(err)
	// 判断输入是文件夹还是文件
	srcInfo, srcOk := exist(*in)
	if !srcOk {
		err = fmt.Errorf("%s not exist! \n", *in)
	}
	log.Fatalln(err)
	destInfo, destOk := exist(*out)
	// 输入的是目录
	if srcInfo.IsDir() {
		if !filepath.IsAbs(*in) {
			*in, err = filepath.Abs(*in)
			log.Fatalln(err)
		}
		// 判断输出是否是文件夹,如果没有就在当前文件夹下创建新文件夹
		if !destOk {
			if *out == "" {
				*out = *in
			} else if filepath.IsAbs(*out) {

			} else {
				pwd, err := os.Getwd()
				log.Fatalln(err)
				baseDir := filepath.Join(pwd, filepath.Clean(*out))
				util.MkdirAll(baseDir)
				*out, err = filepath.Abs(baseDir)
				log.Fatalln(err)
			}
		} else if destInfo.IsDir() {
			// 是已存在的文件夹
			*out, err = filepath.Abs(*out)
			log.Fatalln(err)
		} else {
			// 是文件
			log.Fatalln(fmt.Errorf("out param %s  is file, but want a dir", *out))
		}
		//这里必须绝对路径
		err = converter.MultiConvertV2(*reverse, *in, *out)
	}

	// 输入的是文件
	if !srcInfo.IsDir() {
		var output io.Writer
		if *out == "" {
			output = os.Stdout
		} else if destOk && destInfo.IsDir() {
			output = os.Stdout
		} else {
			output, err = os.OpenFile(filepath.Clean(*out), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		}
		log.Fatalln(err)
		input, err := ioutil.ReadFile(*in)
		log.Fatalln(err)
		// 高级比较 如果是不理想的文件,就自动反转
		if converter.IsIdealFile(!*reverse, srcInfo) {
			*reverse = !*reverse
		}
		err = converter.SingleConvert(*reverse, input, output)
	}

	//判断输出是文件夹还是文件
	log.Fatalln(err)
}

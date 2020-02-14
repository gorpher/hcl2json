# hcl2json
转换HCL为JSON,反之亦然. 使用最新版本的hcl源码构建.

[![Build Status](https://travis-ci.org/gorpher/hcl2json.svg?branch=master)](https://travis-ci.org/gorpher/hcl2json)

## 语言
[简体中文](/README_zh.md)|[English](/READMD.md)

## 安装 

### Linux

下载64位linux版本,并安装.
`/usr/local/bin`:

```bash
curl -SsL https://github.com/gorpher/hcl2json/releases/download/v1.0.0/hcl2json_v1.0.0_linux_amd64 \
  | sudo tee /usr/local/bin/hcl2json > /dev/null && sudo chmod 755 /usr/local/bin/hcl2json && hcl2json -version
```


### OSX

下载64位Darwin版本,并安装.

`/usr/local/bin`:

```bash
curl -SsL https://github.com/gorpher/hcl2json/releases/download/v0.0.6/hcl2json_v1.0.1_darwin_amd64 \
  | sudo tee /usr/local/bin/hcl2json > /dev/null && sudo chmod 755 /usr/local/bin/hcl2json && hcl2json -version
```


## 使用

单文件转换

这个是一个例子 [`test-fixtures/hcl_outputs/outputs.tf`](test-fixtures/hcl_outputs/outputs.tf) 转换成JSON的栗子:


```bash
$ hcl2json -i test-fixtures/hcl_outputs/outputs.tf

{
  "output": [
    {
      "public_ip": [
        {
          "value": "${tencentcloud_instance.cvm_test[0].public_ip}"
        }
      ]
    }
  ]
}
```

典型使用

```bash
$ hcl2json -i test-fixtures/hcl_outputs/outputs.tf -o outputs.tf.json
```

文件夹下多文件转换

```bash
$ hcl2json -i test-fixtures/ -o test-fixtures-10/
```

## json2hcl

通过`-reverse`标签,我们可以将json转换成hcl :

```bash
$ hclwjson -i test-fixtures/json_outputs/outputs.tf.json -reverse


output "public_ip" {
  value = "${tencentcloud_instance.cvm_test[0].public_ip}"
}
```

## 开发

```bash
# 克隆代码
git clone git@github.com:gorpher/hcl2json.git
# 进入目录
cd json2hcl
# 下载依赖包
go mod tidy
# 开始coding
```

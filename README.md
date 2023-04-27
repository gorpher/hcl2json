# hcl2json
Convert HCL to JSON , and vice versa. Use the latest version hcl source to build.

[![Build Status](https://travis-ci.org/gorpher/hcl2json.svg?branch=master)](https://travis-ci.org/gorpher/hcl2json)

## Language
[简体中文](/README_zh.md)|[English](/READMD.md)

## Install 

### Linux

Here's how it could look for 64 bits Linux, if you wanted `hcl2json` available globally inside
`/usr/local/bin`:

```bash
curl -SsL https://github.com/gorpher/hcl2json/releases/download/v1.0.0/hcl2json_v1.0.0_linux_amd64 \
  | sudo tee /usr/local/bin/hcl2json > /dev/null && sudo chmod 755 /usr/local/bin/hcl2json && hcl2json -version
```


### OSX

Here's how it could look for 64 bits Darwin, if you wanted `hcl2json` available globally inside
`/usr/local/bin`:

```bash
curl -SsL https://github.com/gorpher/hcl2json/releases/download/v0.0.6/hcl2json_v1.0.1_darwin_amd64 \
  | sudo tee /usr/local/bin/hcl2json > /dev/null && sudo chmod 755 /usr/local/bin/hcl2json && hcl2json -version
```


## Use

Single file convert

Here's an example [`test-fixtures/hcl_outputs/outputs.tf`](test-fixtures/hcl_outputs/outputs.tf) being converted to JSON:


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

Typical use would be

```bash
$ hcl2json -i test-fixtures/hcl_outputs/outputs.tf -o outputs.tf.json
```

Multi files convert

```bash
$ hcl2json -i test-fixtures/ -o test-fixtures-10/
```

## json2hcl

As a bonus, the conversion the other way around is also supported via the `-reverse` flag:

```bash
$ hcl2json -i test-fixtures/json_outputs/outputs.tf.json -reverse


output "public_ip" {
  value = "${tencentcloud_instance.cvm_test[0].public_ip}"
}
```

## Development

```bash
git clone git@github.com:gorpher/hcl2json.git
cd hcl2json
go mod tidy
```

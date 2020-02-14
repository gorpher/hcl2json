package hclutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl"
	hclParser "github.com/hashicorp/hcl/hcl/parser"
	hclprinter "github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"io"
)

// ToJson byte array convert **json** out to io.Writer
func ToJson(input []byte, out io.Writer) error {
	v := new(interface{})

	astFile, err := hclParser.Parse(input)
	if err != nil {
		return fmt.Errorf("unable to parse HCL: %s", err)
	}

	err = hcl.DecodeObject(v, astFile)
	if err != nil {
		return fmt.Errorf("unable to decodeobject ast.node: %s", err)
	}

	data, err := json.MarshalIndent(*v, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal json: %s", err)

	}

	_, err = io.Copy(out, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("unable to write fiel to %s: %s", out, err)
	}
	return nil
}

// ToHcl byte array convert **hcl** out to io.Writer
func ToHcl(input []byte, out io.Writer) error {
	ast, err := jsonParser.Parse(input)
	if err != nil {
		return fmt.Errorf("unable to parse JSON: %s", err)
	}
	var buf bytes.Buffer
	if err := hclprinter.Fprint(&buf, ast); err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}
	_, err = io.Copy(out, &buf)
	if err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}
	return nil
}

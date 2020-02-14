package hclutils

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var tests = []struct {
	hclfile  string
	jsonfile string
	wantErr  bool
}{
	{"/hcl_outputs/outputs.tf", "/json_outputs/outputs.tf.json", false},
	{"/hcl_providers/provider.tf", "/json_providers/provider.tf.json", false},
	{"/hcl_resources/resources.tf", "/json_resources/resources.tf.json", false},
	{"/hcl_terraform/terraform.tf", "/json_terraform/terraform.tf.json", false},
	{"/hcl_variables/variables.tf", "/json_variables/variables.tf.json", false},
}

const fixtureDir = "./../test-fixtures"

func TestToHcl(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.hclfile, func(t *testing.T) {
			d, err := ioutil.ReadFile(filepath.Join(fixtureDir, tt.jsonfile))
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			out := &bytes.Buffer{}
			err = ToHcl(d, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToHcl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut == "" {
				t.Errorf("ToHcl() gotOut = {}")
			}
		})
	}
}

func TestToJson(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.jsonfile, func(t *testing.T) {
			d, err := ioutil.ReadFile(filepath.Join(fixtureDir, tt.hclfile))
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			out := &bytes.Buffer{}
			err = ToJson(d, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			o, err := ioutil.ReadFile(filepath.Join(fixtureDir, tt.jsonfile))
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			if gotOut := out.String(); gotOut != string(o) {
				t.Errorf("ToJson() gotOut = %v, want %v", gotOut, string(o))
			}
		})
	}
}

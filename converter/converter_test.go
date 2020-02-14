package converter

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const fixtureDir = "./../test-fixtures"
const testOutput = "test_outputs"

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

func initEnv(t *testing.T) (string, string) {
	path, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	command := exec.Command("cp", "-r", fixtureDir, path)
	err = command.Start()
	if err != nil {
		t.Errorf("cp file error: %s", err)
	}
	err = command.Wait()
	if err != nil {
		t.Errorf("cp file error: %s", err)
	}
	src := filepath.Join(path, filepath.Base(fixtureDir))
	dest := filepath.Join(path, filepath.Base(testOutput))
	return src, dest
}

func TestMultiConvertV1(t *testing.T) {
	src, dest := initEnv(t)
	defer os.RemoveAll(filepath.Dir(src))

	err := MultiConvertV1(false, src, dest)
	if err != nil {
		t.Errorf("converter error: %s", err)
	}
	err = MultiConvertV1(true, src, dest)
	if err != nil {
		t.Errorf("converter error: %s", err)
	}
}

func TestMultiConvertV2(t *testing.T) {
	src, dest := initEnv(t)
	err := MultiConvertV2(false, src, dest)
	if err != nil {
		t.Errorf("converter error: %s", err)
	}
	err = MultiConvertV2(true, src, dest)
	if err != nil {
		t.Errorf("converter error: %s", err)
	}
}

func TestSingleConvert(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.jsonfile, func(t *testing.T) {
			d, err := ioutil.ReadFile(filepath.Join(fixtureDir, tt.hclfile))
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			out := &bytes.Buffer{}
			err = SingleConvert(false, d, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			o, err := ioutil.ReadFile(filepath.Join(fixtureDir, tt.jsonfile))
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			out2 := &bytes.Buffer{}
			err = SingleConvert(true, o, out2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToHCL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

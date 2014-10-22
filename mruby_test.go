package mruby

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var test_source = "p 'Hello world!'"
var test_source2 = "p 'Foo bar!'"
var wrong_source = "wrong" + test_source
var wrong_syntax = `
class_decation_error T
  def self.a
    puts "aa"
  end
end`

var expected_err = "undefined method 'wrongp' for main"
var expected_syntax_err = "syntax error"

var expected_bin = []byte{0x52, 0x49, 0x54, 0x45, 0x30, 0x30, 0x30, 0x32, 0x4a, 0x9f, 0x0, 0x0, 0x0, 0x88, 0x4d, 0x41, 0x54, 0x5a, 0x30, 0x30, 0x30, 0x30, 0x49, 0x52, 0x45, 0x50, 0x0, 0x0, 0x0, 0x45, 0x30, 0x30, 0x30, 0x30, 0x0, 0x0, 0x0, 0x39, 0x0, 0x1, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x80, 0x0, 0x6, 0x1, 0x0, 0x0, 0x3d, 0x0, 0x80, 0x0, 0xa0, 0x0, 0x0, 0x0, 0x4a, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0xc, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x21, 0x0, 0x0, 0x0, 0x1, 0x0, 0x1, 0x70, 0x0, 0x44, 0x42, 0x47, 0x0, 0x0, 0x0, 0x0, 0x25, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x19, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1, 0x45, 0x4e, 0x44, 0x0, 0x0, 0x0, 0x0, 0x8}

func TestCompile(t *testing.T) {
	bin, err := Compile(test_source)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(expected_bin, bin) {
		t.Errorf("Expected: %#v\nGot: %#v\n", expected_bin, bin)
	}
}

func TestFailCompile(t *testing.T) {
	_, err := Compile(wrong_syntax)
	if err == nil || err.Error() != expected_syntax_err {
		t.Errorf("This code should lead to a `%s` error: %s\nInstead we got `%#v`.\n", expected_syntax_err, wrong_syntax, err.Error())
	}
}

func TestCompileDiff(t *testing.T) {
	bin, _ := Compile(test_source)
	bin2, _ := Compile(test_source2)
	if bytes.Equal(bin, bin2) {
		t.Errorf("Binaries shouldn't be equal!\nbin: %#v\nbin2: %#v\n", bin, bin2)
	}
}

func TestRunSource(t *testing.T) {
	err := RunSource(test_source)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = RunSource(wrong_source)
	if err == nil || err.Error() != expected_err {
		t.Errorf("Running `%s` should produce error `%s`, but produced `%v`.", wrong_source, expected_err, err.Error())
	}
}

func TestRunBytecode(t *testing.T) {
	bin, _ := Compile(test_source)
	err := RunBytecode(bin)
	if err != nil {
		t.Fatal(err.Error())
	}
	bin, _ = Compile(wrong_source)
	err = RunBytecode(bin)
	if err == nil || err.Error() != expected_err {
		t.Errorf("Compiling and running `%s` should produce error `%s`, but produced `%v`.", wrong_source, expected_err, err.Error())
	}
}

func TestMTests(t *testing.T) {
	os.Chdir("mtests")

	visit := func(path string, f os.FileInfo, err error) error {
		if path == "." {
			return nil
		}
		bin, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		defer recover()
		err = RunSource(string(bin))
		if err != nil {
			return fmt.Errorf("Tests in `%s` failed!", path)
		}
		return nil
	}

	err := filepath.Walk(".", visit)
	if err != nil {
		t.Fatal("Ruby tests failed. %s", err.Error())
	}
}

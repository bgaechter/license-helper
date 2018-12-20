package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScanDirectory(t *testing.T) {

	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
	}()

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	expected := "request: Version 2.88.1, License: Apache-2.0\n"

	testdata := filepath.Join(pwd, "testdata")
	filepath.Walk(testdata, scanDirectory)

	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	t.Log(expected)
	t.Log(buf.String())
	if strings.Compare(expected, buf.String()) != 0 {
		t.Error("Test Failed")
	}
}

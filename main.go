package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

func openTemplate(filePath string) (string, error) {
	if _, err := os.Stat(filepath.Base(filePath)); err != nil {
		return "", err
	}

	src, err := ioutil.ReadFile(filePath)
	fmt.Println(string(src))

	return string(src), err
}

func iterate(n int) (map[int]string, error) {
	return nil, nil
}

func times(n int) []int {
	return make([]int, n)
}

func test() string {
	return "test\ntest\ntest"
}

func main() {
	tmplLoc := "/tmp/exmaple.tmpl"
	//tmplSrc, err := openTemplate(tmplLoc)
	tmplSrc, err := ioutil.ReadFile(tmplLoc)

	tmpl := template.New("t").Funcs(template.FuncMap{
		"times": times,
		"test":  test,
	})

	t, err := tmpl.Parse(string(tmplSrc))

	if err != nil {
		panic(err)
	}

	// var blankIface interface{}
	// err = tmpl.ExecuteTemplate(os.Stdout, "t", blankIface)

	err = t.Execute(os.Stdout, map[string]string{
		"word": "cool",
	})

	if err != nil {
		fmt.Errorf("%s", err)
	}

	// spew.Dump(tmpl)
	// fmt.Println(tmpl)

	return
}

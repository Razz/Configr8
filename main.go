package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func times(n int) []int {
	var sliceInt []int
	for i := 0; i < n; i++ {
		sliceInt = append(sliceInt, i)
	}
	return sliceInt
}

func debug() string {
	return "Test"
}

func map_env() map[string]string {
	env_vars := make(map[string]string)
	for _, env := range os.Environ() {
		sep := strings.Index(env, "=")
		env_vars[env[0:sep]] = env[sep+1:]
	}
	return env_vars
}

func main() {
	var (
		tmplLoc string
	)

	flag.String(tmplLoc, "tmpl", "Location of Template to be parsed")
	// tmplLoc := "/tmp/exmaple.tmpl"
	tmplSrc, err := ioutil.ReadFile(tmplLoc)

	tmpl := template.New("t").Funcs(template.FuncMap{
		"times": times,
		"debug": debug,
		"env":   map_env,
	})

	t, err := tmpl.Parse(string(tmplSrc))

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, map_env())

	if err != nil {
		fmt.Errorf("%s", err)
	}

	return
}

package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

//Maps Variables for use in template. IE. Env : {EnvVar: EnvKey}

//creates a Slice of DataMaps that will be consolidated into one big map
//ths will it's self be returned into a datamap that looks like:
//"Data" : { DataMap[GlobalVar] : {ScopedVar: Scoped Value}}
//Ie. "Data": {
//       "env":
//         {"user": "jsmith",
//          "shell": "ZSH",
//         }
//        }
type DataMap map[string]map[string]string
type DataMapSlice []DataMap
type DataMapFull map[string]map[string]string

func (dm DataMapSlice) consolidate() DataMapFull {
	endDM := make(DataMapFull)
	for _, i := range dm {
		for k, v := range i {
			endDM[k] = v
		}
	}
	return endDM
}

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

func MapEnv() map[string]string {
	dm := make(map[string]string)
	for _, env := range os.Environ() {
		sep := strings.Index(env, "=")
		dm[env[0:sep]] = env[sep+1:]
	}
	return dm
}

func add(x int, y int) int {
	return x + y
}

func main() {
	var (
		tmplLoc string
		// dest         string
		dataMapSlice DataMapSlice
	)

	flag.StringVar(&tmplLoc, "tmpl", "tmpl", "Location of Template to be parsed?")
	// flag.StringVar(&dest, "dest", "dest", "Where is the parsed template going? Default: stdout")
	flag.Parse()

	tmplSrc, err := ioutil.ReadFile(tmplLoc)

	dataMapSlice = append(dataMapSlice, DataMap{"env": MapEnv()})

	tmpl := template.New("t").Funcs(template.FuncMap{
		"times": times,
		"debug": debug,
		"add":   add,
	})

	t, err := tmpl.Parse(string(tmplSrc))

	dataMaps := dataMapSlice.consolidate()
	spew.Dump(dataMaps)

	err = t.Execute(os.Stdout, dataMaps)

	if err != nil {
		fmt.Errorf("%s", err)
	}
}

package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/criticalstack/configr8/plugins"
)

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

func checkError(err error, display string) bool {
	if err != nil {
		log.Fatal(display)
		return false
	}
	return true
}

func MapEnv() map[string]string {
	dm := make(map[string]string)
	for _, env := range os.Environ() {
		sep := strings.Index(env, "=")
		dm[env[0:sep]] = env[sep+1:]
	}
	return dm
}

func main() {
	var (
		tmplLoc      string
		jsonLoc      string
		dest         string
		dataMapSlice DataMapSlice
	)

	flag.StringVar(&tmplLoc, "tmpl", "", "Location of Template to be parsed?")
	flag.StringVar(&tmplLoc, "t", "", "Location of Template to be parsed?")
	flag.StringVar(&jsonLoc, "json", "", "JSON with template data?")
	flag.StringVar(&jsonLoc, "j", "", "JSON with template data?")
	flag.StringVar(&dest, "dest", "", "Where is the parsed template going? Default: stdout")
	flag.StringVar(&dest, "d", "", "Where is the parsed template going? Default: stdout")
	flag.Parse()

	if tmplLoc == "" {
		log.Fatal("No template provided")
	}

	tmplSrc, err := ioutil.ReadFile(tmplLoc)
	checkError(err, "Error reading Template file")

	if jsonLoc != "" {
		var jsonMap DataMap
		if jsonData, err := ioutil.ReadFile(jsonLoc); checkError(err, "Error opening JSON file") {
			if err := json.Unmarshal(jsonData, &jsonMap); checkError(err, "Check JSON formatting") {
				dataMapSlice = append(dataMapSlice, jsonMap)
			}
		}
	}
	dataMapSlice = append(dataMapSlice, DataMap{"env": MapEnv()})

	tmpl := template.New("t").Funcs(template.FuncMap{
		"times": plugin.Times,
		"debug": plugin.Debug,
		"add":   plugin.Add,
		"multi": plugin.Multi,
		"list":  plugin.List,
	})

	t, err := tmpl.Parse(string(tmplSrc))
	checkError(err, "Error Parsing Template. Please check syantax and try again")

	dataMaps := dataMapSlice.consolidate()
	if dest == "" {
		err = t.Execute(os.Stdout, dataMaps)
	} else {
		if destPath, err := os.Create(dest); checkError(err, "Error creating config file") {
			err = t.Execute(destPath, dataMaps)
			defer destPath.Close()
		}
	}
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// dataMap as a struct exists to help map context and user supplied data into
// the templating interface. Creating dataMapSlice allows us to extend functions
// to it.

type dataMap map[string]map[string]string
type dataMapSlice []dataMap

// Take a bunch of dataMaps, and turn them into one. This func may be expensive,
// but it makes for a really userfriendly way to call supplied vaiables in the
// templates

func (dm *dataMapSlice) consolidate() *dataMap {
	endDM := make(dataMap)
	for _, i := range *dm {
		for k, v := range i {
			endDM[k] = v
		}
	}
	return &endDM
}

// There are a lot of potential errors, this allows us to not have to right the
// if then else error checking that is so common.

func checkError(err error, display string) bool {
	if err != nil {
		fmt.Printf("\n---------------\n%s::\n\t%s\n---------------\n", err, display)
		os.Exit(1)
		return false
	}
	return true
}

// Maps env variables for use later.

func mapEnv() map[string]string {
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
		dataMapSlice dataMapSlice
	)

	flag.StringVar(&tmplLoc, "tmpl", "", "Location of Template to be parsed?")
	flag.StringVar(&tmplLoc, "t", "", "Location of Template to be parsed?")
	flag.StringVar(&jsonLoc, "json", "", "JSON with template data?")
	flag.StringVar(&jsonLoc, "j", "", "JSON with template data?")
	flag.StringVar(&dest, "dest", "", "Where is the parsed template going? Default: stdout")
	flag.StringVar(&dest, "d", "", "Where is the parsed template going? Default: stdout")
	flag.Parse()

	if tmplLoc == "" {
		// We fatal on this because if there is no teplate, whats the point.
		// Might be worth expanding to take form Stdin

		log.Fatal("No template provided")
	}

	tmplSrc, err := ioutil.ReadFile(tmplLoc)
	checkError(err, "Error reading Template file")

	if jsonLoc != "" {
		var jsonMap dataMap
		if jsonData, err := ioutil.ReadFile(jsonLoc); checkError(err, "Error opening JSON file") {
			if err := json.Unmarshal(jsonData, &jsonMap); checkError(err, "Check JSON formatting") {
				dataMapSlice = append(dataMapSlice, jsonMap)
			}
		}
	}
	// Example of dataMap for env-context
	dataMapSlice = append(dataMapSlice, dataMap{"env": mapEnv()})

	// Plugins are the true strngth of configr8. This will grow as plugins are
	// added

	pluginMap := template.FuncMap{
		"debug": plugin.Debug,
		"times": plugin.Times,
		"add":   plugin.Add,
		"multi": plugin.Multi,
		"list":  plugin.List,
	}

	tmpl := template.New("t").Funcs(pluginMap)

	t, err := tmpl.Parse(string(tmplSrc))
	checkError(err, "Error Parsing Template. Please check syantax and try again")

	dataMaps := dataMapSlice.consolidate()
	if dest == "" {
		err = t.Execute(os.Stdout, dataMaps)
		checkError(err, "Error Parsing Template. Please check syantax and try again")
	} else {
		if destPath, err := os.Create(dest); checkError(err, "Error creating config file") {
			err = t.Execute(destPath, dataMaps)
			checkError(err, "Error Parsing Template. Please check syantax and try again")
			defer destPath.Close()
		}
	}
}

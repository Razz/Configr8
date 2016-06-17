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

// DataMap as a struct exists to help map context and user supplied data into
// the templating interface. Creating DataMapSlice allows us to extend functions
// to it.

type DataMap map[string]map[string]string
type DataMapSlice []DataMap

// Take a bunch of DataMaps, and turn them into one. This func may be expensive,
// but it makes for a really userfriendly way to call supplied vaiables in the
// templates

func (dm DataMapSlice) consolidate() DataMap {
	endDM := make(DataMap)
	for _, i := range dm {
		for k, v := range i {
			endDM[k] = v
		}
	}
	return endDM
}

// There are a lot of potential errors, this allows us to not have to right the
// if then else error checking that is so common.

func checkError(err error, display string) bool {
	if err != nil {
		log.Fatal(display)
		return false
	}
	return true
}

// Maps env variables for use later.

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
		// We fatal on this because if there is no teplate, whats the point.
		// Might be worth expanding to take form Stdin

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
	// Example of DataMap for env-context
	dataMapSlice = append(dataMapSlice, DataMap{"env": MapEnv()})

	// Plugins are the true stregth of configr8. This will grow as plugins are
	// added

	pluginMap := template.FuncMap{
		"times": plugin.Times,
		"debug": plugin.Debug,
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
	} else {
		if destPath, err := os.Create(dest); checkError(err, "Error creating config file") {
			err = t.Execute(destPath, dataMaps)
			defer destPath.Close()
		}
	}
}

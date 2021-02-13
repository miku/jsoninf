package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sort"

	jsoniter "github.com/json-iterator/go"
)

var (
	json      = jsoniter.ConfigCompatibleWithStandardLibrary
	pathTypes = make(map[string]string)

	numIssues int
)

func dumpTypes(m map[string]interface{}, parent string, line int) {
	for k, v := range m {
		path := fmt.Sprintf("%s/%s", parent, k)
		if v == nil {
			// We do not care about null values for now.
			continue
		}
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			w := v.(map[string]interface{})
			dumpTypes(w, path, line)
		default:
			kind := fmt.Sprintf("%v", reflect.TypeOf(v).Kind())
			t, ok := pathTypes[path]
			if ok {
				if kind != t {
					log.Printf("line %d: mixed types detected in: %s [%s, %s]", line, path, t, kind)
					numIssues++
				}
			} else {
				pathTypes[path] = kind
			}
		}
	}
}

func main() {
	br := bufio.NewReader(os.Stdin)
	var line int
	for {
		b, err := br.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		line++
		var stub = make(map[string]interface{})
		if err := json.Unmarshal(b, &stub); err != nil {
			log.Fatal(err)
		}
		dumpTypes(stub, "", line)
	}
	var keys []string
	for k := range pathTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s [%s]\n", k, pathTypes[k])
	}
	log.Printf("found %d issues in %d lines", numIssues, line)
}

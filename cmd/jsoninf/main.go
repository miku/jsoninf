package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
)

var pathTypes = make(map[string]string)

func dumpTypes(m map[string]interface{}, parent string, line int) {
	for k, v := range m {
		path := fmt.Sprintf("%s/%s", parent, k)
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
		line++
		b, err := br.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
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
}

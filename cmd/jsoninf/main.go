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

func dumpTypes(m map[string]interface{}, parent string) {
	for k, v := range m {
		path := fmt.Sprintf("%s/%s", parent, k)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			w := v.(map[string]interface{})
			dumpTypes(w, path)
		default:
			kind := fmt.Sprintf("%v", reflect.TypeOf(v).Kind())
			t, ok := pathTypes[path]
			if ok {
				if kind != t {
					log.Printf("mixed types detected: %s [%s, %s]", path, t, kind)
				}
			} else {
				pathTypes[path] = kind
			}
		}
	}
}

func main() {
	br := bufio.NewReader(os.Stdin)
	for {
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
		dumpTypes(stub, "")
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

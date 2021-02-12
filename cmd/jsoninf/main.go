package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

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
		log.Println(stub)
	}
}

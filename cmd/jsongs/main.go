package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

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
		var data interface{}
		if err := json.Unmarshal(b, &data); err != nil {
			log.Fatal(err)
		}
		tree := Read(data)
		fmt.Println(tree)
	}
}

type Type int

const (
	Any Type = iota
	Null
	Number
	Bool
	String
	Array
	Map
)

type Node struct {
	Name     string
	Type     Type
	Children []*Node
}

func Read(v interface{}) *Node {
	root := &Node{}
	root.read(v)
	return root
}

func (node *Node) read(value interface{}) {
	switch t := value.(type) {
	case nil:
		node.Type = Null
	case bool:
		node.Type = Bool
	case string:
		node.Type = String
	case json.Number:
		node.Type = Number
	case []interface{}:
		node.Type = Array
		for _, v := range t {
			child := &Node{}
			child.read(v)
			node.Children = append(node.Children, child)
		}
	case map[string]interface{}:
		node.Type = Map
		for k, v := range t {
			child := &Node{Name: k}
			child.read(v)
			node.Children = append(node.Children, child)
		}
	}
}

func (node *Node) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s", node.Name)
	for _, c := range node.Children {
		fmt.Fprintf(&buf, "  %s", c.Name)
	}
	return buf.String()
}

package main

import (
	"fmt"
	xj "github.com/basgys/goxml2json"
	"strings"
)

func main() {
	// x, _ := os.ReadFile("a.xml")
	// xmlStr := *(*string)(unsafe.Pointer(&x))
	// xml := strings.NewReader(xmlStr)
	// xml 是一个 io.Reader
	xml := strings.NewReader(`
		<?xml version="1.0" encoding="UTF-8"?>
		<body>
			<person>
				<name>John Doe</name>
				<age>30</age>
				<info>
					<address>123 Elm St</address>
					<city>Somewhere</city>
				</info>
			</person>
		</body>
	`)

	json, err := xj.Convert(xml)
	if err != nil {
		panic("出了点小状况...")
	}
	fmt.Println(json.String())
	// {"hello": "world", "person": {"name": "John Doe", "age": "30", "info": {"address": "123 Elm St", "city": "Somewhere"}}}
}

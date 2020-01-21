package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
type: model
extends:
    - ./mongo_meta

users: # collection/table name
    username:
        - string
        - required
    password: # required field
        - string
        - required
    token: # hidden field
        - string
        - hidden
    email: string # nullable field

export:
    - users
`

func main() {
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
}

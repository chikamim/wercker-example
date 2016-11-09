package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	// log.Printf("Received event: %s\n", string(evt))
	var values map[string]string
	json.Unmarshal(evt, &values)
	log.Printf(">value1 = %s\n", values["key1"])
	log.Printf(">value2 = %s\n", values["key2"])
	log.Printf(">value3 = %s\n", values["key3"])
	return values["key1"], nil // Echo back the first key value
	// return nil, errors.New("Something went wrong")
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}

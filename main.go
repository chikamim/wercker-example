package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	var values map[string]string
	json.Unmarshal(evt, &values)

	products, err := GetSpecialMacProducts()
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		if product.IsLanguageVariant {
			message := fmt.Sprintf("%s\n\n%s\n\n%s\n", product.Name, product.URL, product.Description)
			log.Println("special", message)
			err := SNSPublish(values["arn"], values["region"],"Found the Apple special deal product", message)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}

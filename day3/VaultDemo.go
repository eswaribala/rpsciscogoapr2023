package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	token := "s.7a1lmH5XuRV3LsLbmcVE23fh"
	vaultAddr := "http://localhost:8200"

	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}
	client.SetToken(token)

	//writing the data
	inputData := map[string]interface{}{
		"data": map[string]interface{}{
			"first": "ankit",
		},
	}
	output, err := client.Logical().Write("secret/data/abd", inputData)
	fmt.Println(output)

	if err != nil {
		fmt.Println(err)
	}
	response, err := client.Logical().Read("secret/data/abd")
	fmt.Println(response.Data)

	/*
		//deleting the data
		data, err := client.Logical().Read("secret/data/hello")
		if err != nil {
			panic(err)
		}
		fmt.Println(data)

		//deleting the data
		output, err = client.Logical().Delete("secret/metadata/abd")
		fmt.Println(output)
		if err != nil {
			panic(err)
		}

	*/
}

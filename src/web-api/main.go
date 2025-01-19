package main

import (
	"fmt"
	"simple-go-crud/configuration"
)

func main() {
	config := configuration.GetConfiguration()
	fmt.Println(config)
}

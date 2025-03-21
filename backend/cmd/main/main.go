package main

import (
	"fmt"
	"sensor-simulator/internal/pkg/endpoint/opcua"
)

func main() {
	fmt.Println("hello")
	opcua.Start()
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MyComplex struct {
	Value complex128
}

func main() {
	err := errors.New("err exmaple")
	fmt.Printf("err exmaple err=%s\n", err)

	mc := MyComplex{Value: 1 + 2i}
	_, err = json.Marshal(mc)
	if err != nil {
		fmt.Printf("json.Marshal(mc) failed err=%s\n", err)
		return
	}
}

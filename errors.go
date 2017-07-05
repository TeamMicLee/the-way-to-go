package main

import (
	"errors"
	"fmt"
)

var errNotFount error = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFount)
}

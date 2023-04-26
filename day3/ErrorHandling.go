package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalizeName(name string) (string, int, error) {
	handle := func(err error) (string, int, error) {
		return "", 0, err
	}
	if name == "" {
		return handle(errors.New("no name provided"))
	}
	return strings.ToTitle(name), len(name), nil
}
func main() {
	name, size, err := capitalizeName("")
	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fmt.Printf("Capitalized name: %s, length: %d", name, size)
}

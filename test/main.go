package main

import (
	"fmt"
	"os"
	"strings"

	"autex"
)

func main() {
	args := os.Args[1:]
	data, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	str := strings.Fields(string(data))

	str = autex.WordZ(str)
	str = autex.Low(str)
	str = autex.Up(str)

	str = autex.Vow(str)
	str = autex.HandleLead(str)
	str = autex.HandleTrail(str)
	str = autex.HandleQuote(str)
	str = autex.Hexbin(str)

	datax := strings.Join(str, " ")

	err = os.WriteFile(args[1], []byte(datax), 0o644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful computation; Check:", args[1])
}

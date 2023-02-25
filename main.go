package main

import (
	"log"

	"github.com/dytlzl/indigo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

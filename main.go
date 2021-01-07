package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Create("override.tf")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	content_override_tf := `
terraform {
    backend "local" {
    }
}
`

	f.WriteString(content_override_tf)
}

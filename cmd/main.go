package main

import "github.com/nvonpentz/go-graphql-react/internal"

func main() {
	application, err := internal.NewApplicationWithDefaults()
	if err != nil {
		panic(err)
	}

	application.Start()
}

package main

import (
	"fmt"
)

func main() {
	app := newApp()
	if err := app.Execute(); err != nil {
		fmt.Println(err)
	}

}

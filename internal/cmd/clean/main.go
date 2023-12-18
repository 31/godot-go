package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/godot-go/godot-go/internal/cmdutils"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	for _, g := range cmdutils.FindOrPanic("test/demo/lib/libgodotgo-*") {
		fmt.Printf("Removing %v\n", g)
		if err := os.RemoveAll(g); err != nil {
			return err
		}
	}
	return nil
}

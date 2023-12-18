package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

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
	gi, _ := exec.LookPath("goimports")
	var cf string
	for _, suffix := range []string{"", "-10", "-11", "-12"} {
		cf, _ := exec.LookPath("clang-format" + suffix)
		if cf != "" {
			break
		}
	}

	if cf != "" {
		cmdutils.RunOrPanic(exec.Command(cf, "-i", "pkg/ffi/ffi_wrapper.gen.h"))
		cmdutils.RunOrPanic(exec.Command(cf, "-i", "pkg/ffi/ffi_wrapper.gen.c"))
	}
	gens := cmdutils.FindOrPanic("pkg/*.gen.go")
	for _, m := range gens {
		cmdutils.RunOrPanic(exec.Command("go", "fmt", m))
	}
	if gi != "" {
		for _, m := range gens {
			cmdutils.RunOrPanic(exec.Command(gi, "-w", m))
		}
	}
	return nil
}

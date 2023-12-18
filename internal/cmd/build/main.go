package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

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
	testBinaryPath := filepath.Join("test", "demo", "lib", "libgodotgo-test-")
	switch runtime.GOOS {
	case "windows":
		testBinaryPath += "windows-" + runtime.GOARCH + ".dll"
	case "darwin":
		testBinaryPath += "macos-" + runtime.GOARCH + ".framework"
	case "linux":
		testBinaryPath += "linux-" + runtime.GOARCH + ".so"
	default:
		testBinaryPath += runtime.GOOS + "-" + runtime.GOARCH + ".so"
	}
	cmd := exec.Command(
		"go", "build",
		"-gcflags=all=-v -N -l -L -clobberdead -clobberdeadreg -dwarf -dwarflocationlists=false",
		"-tags", "tools",
		"-buildmode=c-shared",
		// "-v",
		// "-x",
		"-trimpath",
		"-o", testBinaryPath,
		"test/main.go",
	)
	cmd.Env = append(
		os.Environ(),
		"CGO_ENABLED=1",
		"CGO_CFLAGS=-g3 -g -gdwarf -DX86=1 -fPIC -O0",
		"CGO_LDFLAGS=-g3 -g",
	)
	cmdutils.RunOrPanic(cmd)
	return nil
}

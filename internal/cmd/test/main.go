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
	godot, err := cmdutils.FindGodotBinary()
	if err != nil {
		return err
	}

	cmd := exec.Command(
		godot,
		"--headless",
		"--verbose",
		"--path", "test/demo/",
		"--quit",
	)
	cmd.Env = append(
		os.Environ(),
		"CI=1",
		"LOG_LEVEL=info",
		"GOTRACEBACK=1",
		cmdutils.GoDebugForGodot,
	)
	cmdutils.RunOrPanic(cmd)
	return nil
}

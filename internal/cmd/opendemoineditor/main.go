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
		"--verbose",
		"--debug",
		"--path", "test/demo/",
		"--editor",
	)
	cmd.Env = append(
		os.Environ(),
		"DISPLAY=:0",
		"LOG_LEVEL=info",
		"GOTRACEBACK=1",
		cmdutils.GoDebugForGodot,
	)
	cmdutils.RunOrPanic(cmd)
	return nil
}

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
		"--editor",
		"--quit",
	)
	cmd.Env = append(
		os.Environ(),
		"CI=1",
		"LOG_LEVEL=info",
		"GOTRACEBACK=1",
		cmdutils.GoDebugForGodot,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("The command failed, but don't be concerned. (%v)", err)
	}

	// Hack until fix lands: https://github.com/godotengine/godot/issues/84460
	f, err := os.OpenFile(
		"test/demo/.godot/extension_list.cfg",
		os.O_WRONLY|os.O_CREATE|os.O_EXCL,
		0o666,
	)
	if err == nil {
		// The file didn't exist.
		_, err := fmt.Fprintf(f, "res://example.gdextension\n")
		if errClose := f.Close(); err == nil {
			err = errClose
		}
		if err != nil {
			return err
		}
	} else {
		// The file did exist, so we got an error.
		// Or we got an error for a different reason.
	}
	return nil
}

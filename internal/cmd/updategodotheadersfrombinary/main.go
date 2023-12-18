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
	cmd := exec.Command(godot, "--dump-extension-api", "--headless")
	cmd.Env = append(os.Environ(), "DISPLAY=:0")
	cmdutils.RunOrPanic(cmd)
	copyOrPanic("godot_headers/extension_api.json", "extension_api.json")

	cmd = exec.Command(godot, "--dump-gdextension-interface", "--headless")
	cmd.Env = append(os.Environ(), "DISPLAY=:0")
	cmdutils.RunOrPanic(cmd)
	copyOrPanic("godot_headers/godot/gdextension_interface.h", "gdextension_interface.h")
	return nil
}

func copyOrPanic(dst, src string) {
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(dst, data, 0o666); err != nil {
		panic(err)
	}
}

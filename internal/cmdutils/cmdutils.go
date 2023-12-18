package cmdutils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const GoDebugForGodot = "GODEBUG=sbrk=1,gctrace=1,asyncpreemptoff=1,cgocheck=0,invalidptr=1,clobberfree=1,tracebackancestors=5"

func FindOrPanic(glob string) []string {
	matches, err := filepath.Glob("pkg/*.gen.go")
	if err != nil {
		panic(err)
	}
	return matches
}

func RunOrPanic(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func FindGodotBinary() (string, error) {
	godot := os.Getenv("GODOTBINARY")
	if godot == "" {
		godot, _ = exec.LookPath("godot")
	}
	if godot == "" {
		return "", fmt.Errorf("unable to find godot in env or path")
	}
	return godot, nil
}

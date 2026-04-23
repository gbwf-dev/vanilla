//go:build dev

package vite

import (
	"io/fs"
	"os"
	"os/exec"
)

var FS fs.FS = os.DirFS("internal/vite/dev")

func init() {
	cmd := exec.Command("node", "node_modules/vite/bin/vite")
	cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout
	cmd.Start()
}

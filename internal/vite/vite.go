package vite

import (
	"embed"
	"io/fs"
	"log"
	"os"
	"os/exec"

	"gbfw/internal/env"
)

//go:embed build/*
var productionFS embed.FS

//go:embed dev/*
var developmentFS embed.FS

type JSRuntime string

const (
	JSRuntimeKey = "JS_RUNTIME"

	JSRuntimeNode JSRuntime = "node"
	JSRuntimeBun  JSRuntime = "bun"
)

func Load() (fs.FS, error) {
	if env.IsDev() {
		cmd := exec.Command(string(env.Getenv(JSRuntimeKey, JSRuntimeNode)), "node_modules/vite/bin/vite", "--host")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			log.Println(err)
		}
		return fs.Sub(developmentFS, "dev")
	}

	return fs.Sub(productionFS, "build")
}

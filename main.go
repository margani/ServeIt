package main

import (
	"embed"
	"main/backend"
)

//go:embed build/*
//go:embed build/_app/*
var build embed.FS

func main() {
	backend.Server(build)
}

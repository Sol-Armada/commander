package main

import (
	"embed"
)

//go:embed all:dist
var staticFiles embed.FS

package shepherd

import (
	"embed"
	"io/fs"
)

//go:embed web
var WebFS embed.FS

//go:embed public/*
var publicFS embed.FS

var PublicFS, _ = fs.Sub(publicFS, "public")

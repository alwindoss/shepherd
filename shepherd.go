package shepherd

import (
	"embed"
)

//go:embed web/*
var WebFS embed.FS

package storage

import (
	"embed"
)

//go:embed private/*
var Private embed.FS

package handlers

import "time"

const (
	imgIcon             = "📷" // alternatives: ["🎨",  "🖼️"] // rotation based on file type?
	defaultPollInterval = 10 * time.Millisecond
	mediaPollInterval   = 500 * time.Millisecond
)

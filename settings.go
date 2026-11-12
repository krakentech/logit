package logit

import (
	"fmt"
	"io"
	"os"
)

type messageTypeIcon struct {
	debug   string
	info    string
	warning string
	error   string
	data    string
}

var (
	isDebug = false
	icons   = messageTypeIcon{
		debug:   "🐛",
		info:    "🧠",
		warning: "⚠️",
		error:   "❌",
		data:    "🤖",
	}
	messageFormat                     = "{{time}} {{typeIcon}} - {{message}}"
	timeFormat                        = "06.02.01-15:04:05"
	w                       io.Writer = os.Stdout
	printLf                           = fmt.Fprintln
	isTesting                         = false
	knownMessageFormatTypes           = []string{
		"time",
		"typeIcon",
		"typeText",
		"message",
	}
)

// SetDebug sets the debug mode for the logger.
func SetDebug(debug bool) {
	isDebug = debug
}

// SetWriter sets the writer for the logger.
func SetWriter(writer io.Writer) {
	w = writer
}

func SetMessageTemplate(template string) {
	messageFormat = template
}

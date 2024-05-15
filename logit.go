package logit

import (
    "fmt"
    "io"
    "os"
    "strings"
    "time"
)

const Version = "v1.1.0"

var (
    isDebug              = false
    outFormat            = OutFormatDefault
    timeFormat           = TimeFormatDefault
    prefix               = ""
    w          io.Writer = os.Stdout
    printLf              = fmt.Fprintln
    IsTesting            = false
)

// SetWriter will update where the logging is written
func SetWriter(newWriter io.Writer) {
    w = newWriter
}

func SetOutFormat(format string) {
    outFormat = format
}

func SetTimeFormat(format string) {
    timeFormat = format
}

func IsDebug(debugMode bool) {
    isDebug = debugMode
}

func printLine(lType logType, msg string) {
    out := outFormat

    logTime := time.Now().Format(timeFormat)
    if IsTesting {
        logTime = "00.00.00-00:00:00"
    }

    out = strings.Replace(out, "{{time}}", logTime, 1)
    out = strings.Replace(out, "{{prefix}}", prefix, 1)
    out = strings.Replace(out, "{{type}}", string(lType), 1)
    out = strings.Replace(out, "{{message}}", msg, 1)

    _, err := printLf(w, out)
    if err != nil {
        fmt.Printf("failed to print line: %s", err.Error())
    }
}

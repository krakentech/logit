package logit

import (
    "fmt"
    "io"
    "os"
    "strings"
    "time"
)

const Version = "v1.0.1"

var (
    isDebug              = false
    outFormat            = OutFormatDefault
    timeFormat           = TimeFormatDefault
    w          io.Writer = os.Stdout
    printLf              = fmt.Fprintln
    isTesting            = false
)

// SetWriter will update where the logging is written
func SetWriter(newWriter io.Writer) {
    w = newWriter
}

// SetOutFormat will change the format at which things are logged ( please check the readme for better understanding )
func SetOutFormat(format string) {
    outFormat = format
}

// SetTimeFormat changes the time format used when logging a message
func SetTimeFormat(format string) {
    timeFormat = format
}

// SetIsDebug will set log-it to debug mode which is the only way to print debug messages
func SetIsDebug(debugMode bool) {
    isDebug = debugMode
}

// SetTestMode will set log-it's test mode and will affect the output (should only be used when writing unit tests)
func SetTestMode(testMode bool) {
    isTesting = testMode
}

func printLine(lType logType, msg string) {
    out := outFormat

    logTime := time.Now().Format(timeFormat)
    if isTesting {
        logTime = "00.00.00-00:00:00"
    }

    out = strings.Replace(out, "{{time}}", logTime, 1)
    out = strings.Replace(out, "{{type}}", string(lType), 1)
    out = strings.Replace(out, "{{message}}", msg, 1)

    _, err := printLf(w, out)
    if err != nil {
        fmt.Printf("failed to print line: %s", err.Error())
    }
}

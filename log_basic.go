package logit

import (
    "encoding/json"
    "fmt"
    "strings"
)

func Debug(format string, a ...interface{}) {
    if isDebug {
        printLine(logTypeDebug, fmt.Sprintf(format, a...))
    }
}

func Info(format string, a ...interface{}) {
    printLine(logTypeInfo, fmt.Sprintf(format, a...))
}

func Warn(format string, a ...interface{}) {
    printLine(logTypeWarn, fmt.Sprintf(format, a...))
}

func Err(format string, a ...interface{}) {
    printLine(logTypeError, fmt.Sprintf(format, a...))
}

func Error(message string, err error) {
    if err == nil {
        err = fmt.Errorf("no error found")
    }

    printLine(logTypeError, fmt.Sprintf("%s: %s", message, err.Error()))
}

func DebugData(title string, data any) {
    if isDebug {
        b, err := json.MarshalIndent(data, "", "\t")
        if err != nil {
            Err("failed to marshal data object: %s", err.Error())
            return
        }

        padLen := 80 - (len(title) + 3)
        printLine(logTypeData, fmt.Sprintf("%s[%s]=", strings.Repeat("=", padLen), title))
        _, err = printLf(w, string(b))
        if err != nil {
            fmt.Printf("failed to print line: %s", err.Error())
        }
        printLine(logTypeData, strings.Repeat("=", 80))
    }
}

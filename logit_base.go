package logitv2

import (
    "encoding/json"
    "fmt"
    "strings"
)

// Debug works just like fmt.Println but will add some extra context
func Debug(format string, a ...interface{}) {
    fmt.Printf("[DEBUG] %t\n", isDebug)
    if isDebug {
        printLine(logTypeDebug, fmt.Sprintf(format, a...))
    }
}

// Info works just like fmt.Println but will add some extra context
func Info(format string, a ...interface{}) {
    printLine(logTypeInfo, fmt.Sprintf(format, a...))
}

// Warn works just like fmt.Println but will add some extra context
func Warn(format string, a ...interface{}) {
    printLine(logTypeWarn, fmt.Sprintf(format, a...))
}

// Err works just like fmt.Println but will add some extra context
func Err(format string, a ...interface{}) {
    printLine(logTypeError, fmt.Sprintf(format, a...))
}

// Error will append the error message to the end of a regular Err log (example: <error log>: <error message> )
func Error(err error, format string, a ...interface{}) {
    msg := fmt.Sprintf(format, a...)
    if err == nil {
        err = fmt.Errorf("no error found")
    }
    printLine(logTypeError, fmt.Sprintf("%s: %s", msg, err.Error()))
}

func DebugData(title string, data any) {
    debugData(title, data, false)
}

func DebugDataFormated(title string, data any) {
    debugData(title, data, true)
}

// DebugData will print lines above and below a JSON version of the value passed to data
func debugData(title string, data any, formated bool) {
    if isDebug {
        b := make([]byte, 0)
        var err error = nil
        if formated {
            b, err = json.MarshalIndent(data, "", "\t")
        } else {
            b, err = json.Marshal(data)
        }

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

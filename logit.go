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

func DebugData(title string, format bool, data any) {
	if isDebug {
		b := make([]byte, 0)
		var err error = nil
		if format {
			b, err = json.MarshalIndent(data, "", "\t")
		} else {
			b, err = json.Marshal(data)
		}

		if err != nil {
			Error("failed to marshal data object: %s", err.Error())
			return
		}

		padLen := 80 - (len(title) + 3)
		printLine(logTypeData, fmt.Sprintf("%s[%s]=", strings.Repeat("=", padLen), title))
		_, _ = printLf(w, string(b))
	}
}

func Info(format string, a ...interface{}) {
	printLine(logTypeInfo, fmt.Sprintf(format, a...))
}

func Warning(format string, a ...interface{}) {
	printLine(logTypeWarning, fmt.Sprintf(format, a...))
}

func Error(format string, a ...interface{}) {
	printLine(logTypeError, fmt.Sprintf(format, a...))
}

func Err(err error, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if err == nil {
		err = fmt.Errorf("no error found")
	}
	printLine(logTypeError, fmt.Sprintf("%s: %s", msg, err.Error()))
}

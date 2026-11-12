package logit

import (
	"fmt"
	"strings"
	"time"
)

type logType int

const (
	logTypeDebug   logType = 0
	logTypeInfo    logType = 1
	logTypeWarning logType = 2
	logTypeError   logType = 3
	logTypeData    logType = 4
)

func printLine(lType logType, msg string) {
	out := messageFormat
	
	for _, formatType := range knownMessageFormatTypes {
		switch formatType {
		case "time":
			replaceTypeTime(&out)
		case "typeIcon":
			replaceTypeIcon(&out, lType)
		case "typeText":
			replaceTypeText(&out, lType)
		case "message":
			replaceMessage(&out, msg)
		}
	}

	_, err := printLf(w, out)
	if err != nil {
		fmt.Printf("failed to print line: %s", err.Error())
	}
}

func replaceTypeTime(out *string) {
	logTime := time.Now().Format(timeFormat)
	if isTesting {
		logTime = "00.00.00-00:00:00"
	}
	*out = strings.Replace(*out, "{{time}}", logTime, -1)
}
func replaceTypeIcon(out *string, logType logType) {
	icon := ""
	switch logType {
	case logTypeDebug:
		icon = icons.debug
	case logTypeInfo:
		icon = icons.info
	case logTypeWarning:
		icon = icons.warning
	case logTypeError:
		icon = icons.error
	case logTypeData:
		icon = icons.data
	}
	*out = strings.Replace(*out, "{{typeIcon}}", icon, -1)
}
func replaceTypeText(out *string, logType logType) {
	t := ""
	switch logType {
	case logTypeDebug:
		t = "DBG"
	case logTypeInfo:
		t = "NFO"
	case logTypeWarning:
		t = "WRN"
	case logTypeError:
		t = "ERR"
	case logTypeData:
		t = "DAT"
	}
	*out = strings.Replace(*out, "{{typeText}}", t, -1)
}
func replaceMessage(out *string, msg string) {
	*out = strings.Replace(*out, "{{message}}", msg, -1)
}

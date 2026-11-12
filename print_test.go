package logit

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type errorWriter struct{}

func (ew errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("simulated write failure")
}

func Test_printLine(t *testing.T) {
	tests := []struct {
		Name      string
		LogType   logType
		Msg       string
		Want      string
		badBuffer bool
	}{
		{
			Name:      "Debug",
			LogType:   logTypeDebug,
			Msg:       "test message",
			Want:      "00.00.00-00:00:00 DBG 🐛 test message\n",
			badBuffer: false,
		}, // Debug
	}

	// Setup
	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			origFormat := messageFormat
			messageFormat = "{{time}} {{typeText}} {{typeIcon}} {{message}}"
			printLine(tt.LogType, tt.Msg)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
			messageFormat = origFormat
		})
	}

	// Force Failure
	t.Run("Force Failure", func(t *testing.T) {
		SetWriter(errorWriter{})
		printLine(logTypeDebug, "test message")
		assert.Equal(t, "", buf.String())
		buf.Reset()
	})

}

func Test_replaceTypeTime(t *testing.T) {
	tests := []struct {
		Name string
		Want string
	}{
		{
			Name: "time",
			Want: "00.00.00-00:00:00 Msg",
		},
	}

	isTesting = true
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			out := "{{time}} Msg"
			replaceTypeTime(&out)
			assert.Equal(t, tt.Want, out)
		})
	}
}

func Test_replaceTypeIcon(t *testing.T) {
	tests := []struct {
		Name    string
		LogType logType
		Want    string
	}{
		{
			Name:    "Debug",
			LogType: logTypeDebug,
			Want:    fmt.Sprintf("%s Msg", icons.debug),
		}, // Debug
		{
			Name:    "Info",
			LogType: logTypeInfo,
			Want:    fmt.Sprintf("%s Msg", icons.info),
		}, // Info
		{
			Name:    "Warning",
			LogType: logTypeWarning,
			Want:    fmt.Sprintf("%s Msg", icons.warning),
		}, // Warning
		{
			Name:    "Error",
			LogType: logTypeError,
			Want:    fmt.Sprintf("%s Msg", icons.error),
		}, // Error
		{
			Name:    "Data",
			LogType: logTypeData,
			Want:    fmt.Sprintf("%s Msg", icons.data),
		}, // Data
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			out := "{{typeIcon}} Msg"
			replaceTypeIcon(&out, tt.LogType)
			assert.Equal(t, tt.Want, out)
		})
	}
}

func Test_replaceTypeText(t *testing.T) {
	tests := []struct {
		Name    string
		LogType logType
		Want    string
	}{
		{
			Name:    "Debug",
			LogType: logTypeDebug,
			Want:    "DBG Msg",
		}, // Debug
		{
			Name:    "Info",
			LogType: logTypeInfo,
			Want:    "NFO Msg",
		}, // Info
		{
			Name:    "Warning",
			LogType: logTypeWarning,
			Want:    "WRN Msg",
		}, // Warning
		{
			Name:    "Error",
			LogType: logTypeError,
			Want:    "ERR Msg",
		}, // Error
		{
			Name:    "Data",
			LogType: logTypeData,
			Want:    "DAT Msg",
		}, // Data
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			out := "{{typeText}} Msg"
			replaceTypeText(&out, tt.LogType)
			assert.Equal(t, tt.Want, out)
		})
	}
}

func Test_replaceMessage(t *testing.T) {
	tests := []struct {
		name string
		out  string
		msg  string
		want string
	}{
		{
			name: "replace message 1",
			out:  "test: {{message}}",
			msg:  "test message",
			want: "test: test message",
		}, // replace message 1
		{
			name: "replace message 2",
			out:  "test: {{message}} | {{message}}",
			msg:  "test message",
			want: "test: test message | test message",
		}, // replace message 2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceMessage(&tt.out, tt.msg)
			assert.Equal(t, tt.want, tt.out)
		})
	}
}

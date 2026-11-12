package logit

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestDebug(t *testing.T) {
	tests := []struct {
		Name    string
		Debug   bool
		Pattern string
		Values  []any
		Want    string
	}{
		{
			Name:    "Debug not enabled",
			Debug:   false,
			Pattern: "debug message",
			Values:  []any{},
			Want:    "",
		}, // Debug not enabled
		{
			Name:    "Debug message no values",
			Debug:   true,
			Pattern: "debug message",
			Values:  []any{},
			Want:    "00.00.00-00:00:00 🐛 - debug message\n",
		}, // Debug message no values
		{
			Name:    "Debug message with values",
			Debug:   true,
			Pattern: "debug message %s",
			Values:  []any{"text value"},
			Want:    "00.00.00-00:00:00 🐛 - debug message text value\n",
		}, // Debug message with values
	}

	// Setup
	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			SetDebug(tt.Debug)
			Debug(tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}
}

func TestDebugData(t *testing.T) {
	tests := []struct {
		Name   string
		Debug  bool
		Title  string
		Data   any
		Format bool
		Want   string
	}{
		{
			Name:  "Debug Not Enabled",
			Debug: false,
			Title: "Test Data",
			Data: &TestData{
				Name: "John Doe",
				Age:  30,
			},
			Format: false,
			Want:   "",
		}, // Debug Not Enabled
		{
			Name:  "Data Formatted",
			Debug: true,
			Title: "Test Data",
			Data: &TestData{
				Name: "John Doe",
				Age:  30,
			},
			Format: true,
			Want:   "00.00.00-00:00:00 🤖 - ====================================================================[Test Data]=\n{\n\t\"name\": \"John Doe\",\n\t\"age\": 30\n}\n",
		}, // Data Formatted
		{
			Name:  "Data Not Formatted",
			Debug: true,
			Title: "Test Data",
			Data: &TestData{
				Name: "John Doe",
				Age:  30,
			},
			Format: false,
			Want:   "00.00.00-00:00:00 🤖 - ====================================================================[Test Data]=\n{\"name\":\"John Doe\",\"age\":30}\n",
		}, // Data Not Formatted
		{
			Name:   "Null Data",
			Debug:  true,
			Title:  "Test Data",
			Data:   nil,
			Format: false,
			Want:   "00.00.00-00:00:00 🤖 - ====================================================================[Test Data]=\nnull\n",
		}, // Null Data
		{
			Name:  "Data With Error",
			Debug: true,
			Title: "Test Data",
			Data: func() chan int {
				return make(chan int)
			}(),
			Format: false,
			Want:   "00.00.00-00:00:00 ❌ - failed to marshal data object: json: unsupported type: chan int\n",
		},
	}

	// Setup
	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			SetDebug(tt.Debug)
			DebugData(tt.Title, tt.Format, tt.Data)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}
}

func TestInfo(t *testing.T) {
	tests := []struct {
		Name    string
		Pattern string
		Values  []any
		Want    string
	}{
		{
			Name:    "Info message no values",
			Pattern: "info message",
			Values:  nil,
			Want:    "00.00.00-00:00:00 🧠 - info message\n",
		}, // Info message no values
		{
			Name:    "Info message with values",
			Pattern: "info message %s",
			Values:  []any{"text value"},
			Want:    "00.00.00-00:00:00 🧠 - info message text value\n",
		}, // Info message with values
	}

	// Setup
	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			Info(tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}
}

func TestWarning(t *testing.T) {
	tests := []struct {
		Name    string
		Pattern string
		Values  []any
		Want    string
	}{
		{
			Name:    "Warning message no values",
			Pattern: "warning message",
			Values:  nil,
			Want:    "00.00.00-00:00:00 ⚠️ - warning message\n",
		}, // Warning message no values
		{
			Name:    "Warning message with values",
			Pattern: "warning message %s",
			Values:  []any{"text value"},
			Want:    "00.00.00-00:00:00 ⚠️ - warning message text value\n",
		}, // Warning message with values
	}

	// Setup
	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			Warning(tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		Name    string
		Pattern string
		Values  []any
		Want    string
	}{
		{
			Name:    "Error message no values",
			Pattern: "error message",
			Values:  nil,
			Want:    "00.00.00-00:00:00 ❌ - error message\n",
		}, // Error message no values
		{
			Name:    "Error message with values",
			Pattern: "error message %s",
			Values:  []any{"text value"},
			Want:    "00.00.00-00:00:00 ❌ - error message text value\n",
		}, // Error message with values
	}

	// Setup
	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			Error(tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}
}

func TestErr(t *testing.T) {
	tests := []struct {
		Name    string
		Err     error
		Pattern string
		Values  []any
		Want    string
	}{
		{
			Name:    "Err with error and no values",
			Err:     fmt.Errorf("test error"),
			Pattern: "error occurred",
			Values:  nil,
			Want:    "00.00.00-00:00:00 ❌ - error occurred: test error\n",
		}, // Err with error and no values
		{
			Name:    "Err with error and values",
			Err:     fmt.Errorf("test error"),
			Pattern: "error occurred %s",
			Values:  []any{"additional info"},
			Want:    "00.00.00-00:00:00 ❌ - error occurred additional info: test error\n",
		}, // Err with error and values
		{
			Name:    "Err with nil error",
			Err:     nil,
			Pattern: "error occurred",
			Values:  nil,
			Want:    "00.00.00-00:00:00 ❌ - error occurred: no error found\n",
		}, // Err with nil error
	}

	// Setup
	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			Err(tt.Err, tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}

}

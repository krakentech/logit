package logitv2

import (
    "bytes"
    "fmt"
    "github.com/stretchr/testify/assert"
    "math"
    "os"
    "testing"
)

type LogItTest struct {
    Name    string
    Debug   bool
    Pattern string
    Values  []any
    Err     error
    Want    string
}

type debugDataTest struct {
    Name string
}

func TestDebug(t *testing.T) {

    tests := []LogItTest{
        {
            Name:    "debug test debug off",
            Pattern: "test value",
            Values:  []any{},
            Err:     nil,
            Want:    "",
        },
        {
            Name:    "debug test only pattern",
            Pattern: "test value",
            Values:  []any{},
            Err:     nil,
            Want:    "00.00.00-00:00:00 🐛 - test value\n",
            Debug:   true,
        },
        {
            Name:    "debug test  pattern and values",
            Pattern: "the number is %d",
            Values:  []any{123},
            Err:     nil,
            Want:    "00.00.00-00:00:00 🐛 - the number is 123\n",
            Debug:   true},
    }

    var buf bytes.Buffer
    SetWriter(&buf)
    SetTestMode(true)

    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            SetIsDebug(tt.Debug)
            Debug(tt.Pattern, tt.Values...)
            assert.Equal(t, tt.Want, buf.String())
            buf.Reset()
        })
    }

}

func TestInfo(t *testing.T) {
    tests := []LogItTest{
        {Name: "info test debug off", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00 🧠 - test value\n"},
        {Name: "info test only pattern", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00 🧠 - test value\n", Debug: true},
        {Name: "info test  pattern and values", Pattern: "the number is %d", Values: []any{123}, Err: nil, Want: "00.00.00-00:00:00 🧠 - the number is 123\n", Debug: true},
    }

    var buf bytes.Buffer
    SetWriter(&buf)
    SetTestMode(true)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            SetIsDebug(tt.Debug)
            Info(tt.Pattern, tt.Values...)
            assert.Equal(t, tt.Want, buf.String())
            buf.Reset()
        })
    }

}

func TestWarn(t *testing.T) {
    tests := []LogItTest{
        {Name: "warn test debug off", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00 🚧 - test value\n"},
        {Name: "warn test only pattern", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00 🚧 - test value\n", Debug: true},
        {Name: "warn test  pattern and values", Pattern: "the number is %d", Values: []any{123}, Err: nil, Want: "00.00.00-00:00:00 🚧 - the number is 123\n", Debug: true},
    }

    var buf bytes.Buffer
    SetWriter(&buf)
    SetTestMode(true)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            SetIsDebug(tt.Debug)
            Warn(tt.Pattern, tt.Values...)
            assert.Equal(t, tt.Want, buf.String())
            buf.Reset()
        })
    }

}

func TestErr(t *testing.T) {
    tests := []LogItTest{
        {Name: "err test debug off", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00 🛑 - test value\n"},
        {Name: "err test only pattern", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00 🛑 - test value\n", Debug: true},
        {Name: "err test  pattern and values", Pattern: "the number is %d", Values: []any{123}, Err: nil, Want: "00.00.00-00:00:00 🛑 - the number is 123\n", Debug: true},
    }

    var buf bytes.Buffer
    SetWriter(&buf)
    SetTestMode(true)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            SetIsDebug(tt.Debug)
            Err(tt.Pattern, tt.Values...)
            assert.Equal(t, tt.Want, buf.String())
            buf.Reset()
        })
    }
}

func TestError(t *testing.T) {
    type LogItTest struct {
        Name    string
        Debug   bool
        Error   error
        Pattern string
        Values  []any
        Err     error
        Want    string
    }

    tests := []LogItTest{
        {Name: "error test debug off", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00 🛑 - test value: no error found\n"},
        {Name: "err test only pattern", Error: fmt.Errorf("im an error"), Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00 🛑 - test value: im an error\n", Debug: true},
        {Name: "err test  pattern and values", Error: fmt.Errorf("im an error"), Pattern: "the number is %d", Values: []any{123}, Err: nil, Want: "00.00.00-00:00:00 🛑 - the number is 123: im an error\n", Debug: true},
    }

    var buf bytes.Buffer
    SetWriter(&buf)
    SetTestMode(true)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            SetIsDebug(tt.Debug)
            Error(tt.Error, tt.Pattern, tt.Values...)
            assert.Equal(t, tt.Want, buf.String())
            buf.Reset()
        })
    }
}

func TestDebugData(t *testing.T) {

    type LogItTest struct {
        Name         string
        Debug        bool
        Formated     bool
        Values       any
        Err          error
        Want         string
        RemoveBuffer bool
    }

    tests := []LogItTest{
        {
            Name:         "data test debug off",
            Values:       "test value",
            Formated:     true,
            Err:          nil,
            Want:         "",
            RemoveBuffer: false,
        },
        {
            Name:         "formated",
            Values:       debugDataTest{Name: "test value"},
            Formated:     true,
            Err:          nil,
            Debug:        true,
            Want:         "00.00.00-00:00:00 🧶 - =====================================================================[formated]=\n{\n\t\"Name\": \"test value\"\n}\n00.00.00-00:00:00 🧶 - ================================================================================\n",
            RemoveBuffer: false,
        },
        {
            Name:         "un-formated",
            Values:       debugDataTest{Name: "test value"},
            Formated:     false,
            Err:          nil,
            Debug:        true,
            Want:         "00.00.00-00:00:00 🧶 - ==================================================================[un-formated]=\n{\"Name\":\"test value\"}\n00.00.00-00:00:00 🧶 - ================================================================================\n",
            RemoveBuffer: false,
        },
        {
            Name:         "fail to marshal data",
            Values:       math.NaN(),
            Err:          nil,
            Debug:        true,
            Want:         "00.00.00-00:00:00 🛑 - failed to marshal data object: json: unsupported value: NaN\n",
            RemoveBuffer: false,
        },
    }

    var buf bytes.Buffer
    SetTestMode(true)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            if tt.RemoveBuffer {
                SetWriter(nil)
            } else {
                SetWriter(&buf)
            }

            SetIsDebug(tt.Debug)
            debugData(tt.Name, tt.Values, tt.Formated)
            assert.Equal(t, tt.Want, buf.String())
            buf.Reset()
        })
    }
}

func TestDebugDataUnFormatted(t *testing.T) {
    var buf bytes.Buffer
    SetTestMode(true)
    SetIsDebug(true)
    SetWriter(&buf)
    want := "00.00.00-00:00:00 🧶 - =============================================================[test unformatted]=\n{\"Name\":\"test value\"}\n00.00.00-00:00:00 🧶 - ================================================================================\n"

    DebugData("test unformatted", debugDataTest{Name: "test value"})
    assert.Equal(t, want, buf.String())
    SetWriter(os.Stdout)
}

func TestDebugDataFormatted(t *testing.T) {
    var buf bytes.Buffer
    SetTestMode(true)
    SetIsDebug(true)
    SetWriter(&buf)
    want := "00.00.00-00:00:00 🧶 - ===============================================================[test formatted]=\n{\n\t\"Name\": \"test value\"\n}\n00.00.00-00:00:00 🧶 - ================================================================================\n"

    DebugDataFormated("test formatted", debugDataTest{Name: "test value"})
    assert.Equal(t, want, buf.String())
    SetWriter(os.Stdout)
}

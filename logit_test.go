package logitv2

import (
    "bytes"
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
)

func TestSetWriter(t *testing.T) {
    tests := []struct {
        name         string
        setNewWriter bool
    }{
        {
            name:         "override",
            setNewWriter: true,
        },
        {
            name:         "no override",
            setNewWriter: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Reset to default
            w = os.Stdout

            origWriter := os.Stdout
            overWriter := &bytes.Buffer{}

            if tt.setNewWriter {
                SetWriter(overWriter)
            }

            if tt.setNewWriter {
                assert.Equal(t, w, overWriter)
            } else {
                assert.Equal(t, w, origWriter)
            }
        })
    }
}

func TestSetOutFormat(t *testing.T) {
    tests := []struct {
        name       string
        overRide   bool
        wantFormat string
    }{
        {
            name:       "format default",
            overRide:   false,
            wantFormat: OutFormatDefault,
        },
        {
            name:       "format override",
            overRide:   true,
            wantFormat: "this is a test",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            w = os.Stdout
            if tt.overRide {
                SetOutFormat(tt.wantFormat)
            }
            assert.Equal(t, outFormat, tt.wantFormat)
        })
    }
}

func TestSetTimeFormat(t *testing.T) {
    tests := []struct {
        name       string
        overRide   bool
        wantFormat string
    }{
        {
            name:       "time default",
            overRide:   false,
            wantFormat: TimeFormatDefault,
        },
        {
            name:       "time override",
            overRide:   true,
            wantFormat: "this is a test",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.overRide {
                SetTimeFormat(tt.wantFormat)
            }
            assert.Equal(t, timeFormat, tt.wantFormat)
        })
    }
}

func TestSetIsDebug(t *testing.T) {
    isDebug = false
    t.Run("is debug false", func(t *testing.T) {
        assert.Equal(t, isDebug, false)
    })
    SetIsDebug(true)
    t.Run("is debug true", func(t *testing.T) {
        assert.Equal(t, isDebug, true)
    })
    isDebug = false
}

func TestSetTestMode(t *testing.T) {
    isTesting = false
    t.Run("is debug false", func(t *testing.T) {
        assert.Equal(t, isTesting, false)
    })
    SetTestMode(true)
    t.Run("is debug true", func(t *testing.T) {
        assert.Equal(t, isTesting, true)
    })
    isTesting = false
}

func Test_printLine(t *testing.T) {

    tests := []struct {
        name   string
        format string
        want   string
    }{
        {
            name:   "format empty",
            format: "",
            want:   "\n",
        },
        {
            name:   "format time",
            format: "{{time}}",
            want:   "00.00.00-00:00:00\n",
        },
        {
            name:   "format type",
            format: "{{type}}",
            want:   "🐛\n",
        },
        {
            name:   "format message",
            format: "{{message}}",
            want:   "test\n",
        },
        {
            name:   "test default",
            format: OutFormatDefault,
            want:   "00.00.00-00:00:00 🐛 - test\n",
        },
    }

    var buf bytes.Buffer
    SetWriter(&buf)
    SetTestMode(true)

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            SetOutFormat(tt.format)
            printLine(logTypeDebug, "test")
            assert.Equal(t, tt.want, buf.String())
            buf.Reset()
        })
    }
}

package logit

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestSetOutFormat(t *testing.T) {
    tests := []struct {
        name      string
        newFormat string
    }{
        {
            name:      "validate it sets",
            newFormat: "this is just a test",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            outFormat = OutFormatDefault
            SetOutFormat(tt.newFormat)
            assert.Equal(t, tt.newFormat, outFormat)
        })
    }

}

func TestSetTimeFormat(t *testing.T) {
    tests := []struct {
        name      string
        newFormat string
    }{
        {
            name:      "validate it sets",
            newFormat: "this is just a test",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            outFormat = OutFormatDefault
            SetTimeFormat(tt.newFormat)
            assert.Equal(t, tt.newFormat, outFormat)
        })
    }
}

package logit

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestErrDecorator(t *testing.T) {

    tests := []struct {
        name      string
        inError   error
        addedText string
        wantErr   string
    }{
        {
            name:      "error nil",
            inError:   nil,
            addedText: "added stuff",
            wantErr:   "",
        },
        {
            name:      "basic usage",
            inError:   fmt.Errorf("im the error"),
            addedText: "added data",
            wantErr:   "added data: im the error",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ErrDecorator(tt.inError, tt.addedText)
            if err != nil {
                assert.Equal(t, err.Error(), tt.wantErr)
            } else {
                assert.Equal(t, tt.wantErr, "")
            }
        })
    }
}

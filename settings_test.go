package logit

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSetDebug(t *testing.T) {

	tests := []struct {
		name string
		want bool
	}{
		{
			name: "true",
			want: true,
		},
		{
			name: "false",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetDebug(tt.want)
			assert.Equal(t, tt.want, isDebug)
		})
	}
}

func TestSetWriter(t *testing.T) {
	t.Run("Writer Set Null", func(t *testing.T) {
		SetWriter(nil)
		assert.Equal(t, nil, w)
	})
	SetWriter(os.Stdout)
}

func TestSetMessageFormat(t *testing.T) {
	t.Run("Message Format Set", func(t *testing.T) {
		SetMessageTemplate("{{}}")
		assert.Equal(t, "{{}}", messageFormat)
	})
	SetMessageTemplate("{{time}} {{typeIcon}} - {{message}}")
}

package calculation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate(t *testing.T) {
	tests := []struct {
		exp  string
		want int
	}{
		{
			exp:  " 2-1 + 2 ",
			want: 3,
		},
		{
			exp:  "(1+(4+5+2)-3)+(6+8)",
			want: 23,
		},
		{
			exp:  "(1+(4+5+2)-3)*(6+8)",
			want: 126,
		},
		{
			exp:  "1-11",
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s = %d", tt.exp, tt.want), func(t *testing.T) {
			assert.Equal(t, tt.want, calculate(tt.exp))
		})
	}
}

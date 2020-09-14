package discord

import (
	"reflect"
	"testing"
)

func Test_formatCharName(t *testing.T) {
	tests := []struct {
		name     string
		wantName []string
	}{
		{
			name:     "bdezuiad a dbzaoid inazi dniza   daz dzadaz   dazd az   dzad  ",
			wantName: []string{"bdezuiad", "a", "dbzaoid", "inazi", "dniza", "daz", "dzadaz", "dazd", "az", "dzad"},
		},
		{
			name:     "   the big    TbBoos   ! ",
			wantName: []string{"the", "big", "TbBoos", "!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotName := formatCharName(tt.name); !reflect.DeepEqual(gotName, tt.wantName) {
				t.Errorf("formatCharName() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}

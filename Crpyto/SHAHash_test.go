package Crypto

import (
	"fmt"
	"testing"
)

func TestNewHash(t *testing.T) {
	type args struct {
		length int32
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			"test1",
			args{length: 1},
			"1",
		},
		{
			"test2",
			args{length: 2},
			"2",
		},
		{
			"test7",
			args{length: 7},
			"7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := fmt.Sprintf("%d", len(NewHash(tt.args.length)))
			if gotResult != tt.wantResult {
				t.Errorf("NewHash() length = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

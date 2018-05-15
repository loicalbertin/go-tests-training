package basic

import "testing"

func TestSeeCoverage(t *testing.T) {
	type args struct {
		e SomeEnum
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestE1", args{E1}},
		{"TestE3", args{E3}},
		{"TestE4", args{E4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SeeCoverage(tt.args.e)
		})
	}
}

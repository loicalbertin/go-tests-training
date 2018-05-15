package basic

import (
	"testing"
)

func TestEnumAsString(t *testing.T) {
	t.Parallel()
	type args struct {
		e SomeEnum
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestE1", args{E1}, "E1", false},
		{"TestE3", args{E3}, "E3", false},
		{"TestE4", args{E4}, "E4", false},
		{"TestUnexpected", args{5}, "", true},
	}
	for _, tt := range tests {
		// Here is a subtest

		// This is tricky due to the t.Parallel below see
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		// for a better understanding
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// (sub)tests could be run in parallel
			t.Parallel()
			got, err := EnumAsString(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("EnumAsString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EnumAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

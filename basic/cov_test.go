package basic

import (
	"testing"
)

func TestEnumAsString(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

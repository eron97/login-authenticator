package crypto

import (
	"reflect"
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Valid Password",
			args:    args{"mysecretpassword"},
			want:    `#string`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == "#string" && reflect.TypeOf(got).Kind() != reflect.String {
				t.Errorf("HashPassword() = %v, want a string", got)
			}
		})
	}
}

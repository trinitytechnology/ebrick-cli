package utils

import "testing"

func TestIsValidVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid version",
			args: args{
				version: "v1.0.0",
			},
			want: true,
		},
		{
			name: "valid version with single digit",
			args: args{
				version: "v1.2.3",
			},
			want: true,
		},
		{
			name: "valid version with multiple digits",
			args: args{
				version: "v12.34.56",
			},
			want: true,
		},
		{
			name: "invalid version missing 'v'",
			args: args{
				version: "1.0.0",
			},
			want: false,
		},
		{
			name: "invalid version missing patch number",
			args: args{
				version: "v1.0",
			},
			want: false,
		},
		{
			name: "invalid version with extra characters",
			args: args{
				version: "v1.0.0-beta",
			},
			want: true,
		},
		{
			name: "invalid version with letters",
			args: args{
				version: "v1.a.0",
			},
			want: false,
		},
		{
			name: "invalid version with special characters",
			args: args{
				version: "v1.0.0!",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidVersion(tt.args.version); got != tt.want {
				t.Errorf("IsValidVersion(%s) = %v, want %v", tt.args.version, got, tt.want)
			}
		})
	}
}

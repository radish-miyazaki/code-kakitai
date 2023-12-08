package ulid

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid",
			args: args{
				s: NewULID(),
			},
			want: true,
		},
		{
			name: "Invalid: string is empty",
			args: args{
				s: "",
			},
			want: false,
		},
		{
			name: "Invalid: string length is 32",
			args: args{
				s: "01234567890123456789012345678901",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

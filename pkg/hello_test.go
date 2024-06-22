package pkg

import "testing"

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "simple",
			want: "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelloWorld(); got != tt.want {
				t.Errorf("HelloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}

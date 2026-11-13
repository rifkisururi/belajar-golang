package greeting

import "testing"

func TestMessage(t *testing.T) {
	tests := []struct{ name, input, want string }{
		{"named learner", "Rifki", "Hello, Rifki!"},
		{"default learner", "", "Hello, Gopher!"},
		{"trim whitespace", "  Rifki  ", "Hello, Rifki!"},
		{"blank learner", " \t\n", "Hello, Gopher!"},
		{"unicode learner", "世界", "Hello, 世界!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Message(tt.input); got != tt.want {
				t.Errorf("Message(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

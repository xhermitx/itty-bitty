package url

import (
	"errors"
	"github.com/xhermitx/itty-bitty/internal/db"
	"github.com/xhermitx/itty-bitty/internal/utils"
	"testing"
)

func TestService_ValidateURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		want        string
		expectedErr error
	}{
		// TODO: Add test cases.
		{
			name:        "url without http(s)",
			url:         "example.com",
			expectedErr: utils.ErrInvalidURL,
		},
		{
			name:        "url without .com, .edu, etc",
			url:         "https://example",
			want:        "",
			expectedErr: utils.ErrInvalidURL,
		},
		{
			name: "valid URL",
			url:  "https://example.com",
			want: "https://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Service{
				db: db.New(),
			}
			got, err := u.ValidateURL(tt.url)
			if !errors.Is(err, tt.expectedErr) || got != tt.want {
				t.Errorf("ValidateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

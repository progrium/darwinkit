package declparse

import (
	"regexp"
	"testing"
)

func TestAST_Strings(t *testing.T) {
	for _, tt := range tests {
		if tt.ParseOnly {
			continue
		}
		t.Run(tt.s, func(t *testing.T) {
			got := normalizeWhitespace(tt.n.String())
			want := normalizeWhitespace(tt.s)
			if got != want {
				t.Errorf("String()\n  got: %s\n want: %s", got, want)
			}
		})
	}
}

func normalizeWhitespace(s string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(s, " ")
}

package version

import "testing"

func TestShouldNotifyUpdate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		current string
		latest  string
		want    bool
	}{
		{
			name:    "newer patch version",
			current: "0.11.0",
			latest:  "v0.11.1",
			want:    true,
		},
		{
			name:    "same version",
			current: "0.11.0",
			latest:  "v0.11.0",
			want:    false,
		},
		{
			name:    "older latest",
			current: "0.11.0",
			latest:  "v0.10.9",
			want:    false,
		},
		{
			name:    "prerelease does not notify for stable current",
			current: "1.0.0",
			latest:  "v1.0.0-rc.1",
			want:    false,
		},
		{
			name:    "prerelease numeric identifiers are compared numerically",
			current: "1.0.0-rc.2",
			latest:  "v1.0.0-rc.10",
			want:    true,
		},
		{
			name:    "invalid latest tag",
			current: "0.11.0",
			latest:  "latest",
			want:    false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ShouldNotifyUpdate(tt.current, tt.latest); got != tt.want {
				t.Fatalf("ShouldNotifyUpdate(%q, %q) = %v, want %v", tt.current, tt.latest, got, tt.want)
			}
		})
	}
}

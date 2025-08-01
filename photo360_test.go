package photo360

import (
	"testing"
)

func TestNewPhoto360(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "valid ephoto360 URL",
			url:     "https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html",
			wantErr: false,
		},
		{
			name:    "invalid URL - not photo360",
			url:     "https://example.com/test",
			wantErr: true,
		},
		{
			name:    "empty URL uses default",
			url:     "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPhoto360(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPhoto360() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetName(t *testing.T) {
	p360, _ := NewPhoto360("")
	p360.SetName("Test Name")

	if len(p360.InputText) != 1 || p360.InputText[0] != "Test Name" {
		t.Errorf("SetName() failed, got %v, want [Test Name]", p360.InputText)
	}
}

func TestSetNames(t *testing.T) {
	p360, _ := NewPhoto360("")
	names := []string{"First", "Second", "Third"}
	p360.SetNames(names)

	if len(p360.InputText) != 3 {
		t.Errorf("SetNames() failed, got length %d, want 3", len(p360.InputText))
	}

	for i, name := range names {
		if p360.InputText[i] != name {
			t.Errorf("SetNames() failed at index %d, got %s, want %s", i, p360.InputText[i], name)
		}
	}
}

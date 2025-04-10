package traffic

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestLightController_switchLight(t *testing.T) {
	tests := map[string]struct {
		// Named input parameters for target function.
		currLight    Light
		wantLight    Light
		wantDuration time.Duration
	}{
		"switch to red": {
			currLight:    LightGreen,
			wantLight:    LightRed,
			wantDuration: 2 * time.Second,
		},
		"switch to yellow": {
			currLight:    LightRed,
			wantLight:    LightYellow,
			wantDuration: time.Second,
		},
		"switch to green": {
			currLight:    LightYellow,
			wantLight:    LightGreen,
			wantDuration: 3 * time.Second,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewLightController(2*time.Second, time.Second, 3*time.Second)
			gotLight, gotDuration := c.switchLight(tt.currLight)
			if diff := cmp.Diff(gotLight, tt.wantLight); diff != "" {
				t.Errorf("switchLight() = %v, want %v", gotLight, tt.wantLight)
			}
			if diff := cmp.Diff(gotDuration, tt.wantDuration); diff != "" {
				t.Errorf("switchLight() = %v, want %v", gotDuration, tt.wantDuration)
			}
		})
	}
}

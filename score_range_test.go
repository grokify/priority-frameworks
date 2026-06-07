package priorityframeworks

import (
	"testing"
)

func TestCVSSScoreRange(t *testing.T) {
	sr := CVSSScoreRange()

	tests := []struct {
		score   float64
		wantID  string
		wantErr bool
	}{
		{10.0, "critical", false},
		{9.5, "critical", false},
		{9.0, "critical", false},
		{8.9, "high", false},
		{7.5, "high", false},
		{7.0, "high", false},
		{6.9, "medium", false},
		{5.0, "medium", false},
		{4.0, "medium", false},
		{3.9, "low", false},
		{2.0, "low", false},
		{0.1, "low", false},
		{0.0, "informational", false},
		{-0.1, "", true}, // below min
		{10.1, "", true}, // above max
	}

	for _, tt := range tests {
		level, err := sr.LevelFromScore(tt.score)
		if tt.wantErr {
			if err == nil {
				t.Errorf("LevelFromScore(%v) expected error, got nil", tt.score)
			}
			continue
		}
		if err != nil {
			t.Errorf("LevelFromScore(%v) unexpected error: %v", tt.score, err)
			continue
		}
		if level == nil {
			t.Errorf("LevelFromScore(%v) returned nil level", tt.score)
			continue
		}
		if level.ID != tt.wantID {
			t.Errorf("LevelFromScore(%v) = %s, want %s", tt.score, level.ID, tt.wantID)
		}
	}
}

func TestPercentageScoreRange(t *testing.T) {
	sr := PercentageScoreRange(Severity())
	if sr == nil {
		t.Fatal("PercentageScoreRange returned nil")
	}

	tests := []struct {
		score  float64
		wantID string
	}{
		{100.0, "critical"},
		{75.0, "critical"},
		{74.9, "high"},
		{50.0, "high"},
		{49.9, "medium"},
		{25.0, "medium"},
		{24.9, "low"},
		{0.0, "low"},
	}

	for _, tt := range tests {
		level, err := sr.LevelFromScore(tt.score)
		if err != nil {
			t.Errorf("LevelFromScore(%v) unexpected error: %v", tt.score, err)
			continue
		}
		if level == nil {
			t.Errorf("LevelFromScore(%v) returned nil level", tt.score)
			continue
		}
		if level.ID != tt.wantID {
			t.Errorf("LevelFromScore(%v) = %s, want %s", tt.score, level.ID, tt.wantID)
		}
	}
}

func TestPercentageScoreRangeNilFramework(t *testing.T) {
	sr := PercentageScoreRange(nil)
	if sr != nil {
		t.Error("PercentageScoreRange(nil) should return nil")
	}
}

func TestMustLevelFromScorePanic(t *testing.T) {
	sr := CVSSScoreRange()

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustLevelFromScore should panic on out-of-range score")
		}
	}()

	sr.MustLevelFromScore(15.0) // should panic
}

package priorityframeworks

import (
	"errors"
	"fmt"
)

// ErrScoreOutOfRange is returned when a score is outside the defined ranges.
var ErrScoreOutOfRange = errors.New("score out of range")

// ErrNilFramework is returned when a ScoreRange has no associated Framework.
var ErrNilFramework = errors.New("score range has no framework")

// ErrUnknownLevel is returned when a matched RangeEntry references a level ID
// that does not exist in the associated Framework.
var ErrUnknownLevel = errors.New("range references unknown level")

// ScoreRange maps numeric score ranges to framework levels.
// This enables conversion from scoring systems (like CVSS) to priority levels.
type ScoreRange struct {
	// Framework is the target framework for level lookups.
	Framework *Framework

	// Ranges defines the score thresholds for each level.
	// Each entry maps a level ID to its minimum score (inclusive).
	// Levels are evaluated from highest to lowest score.
	Ranges []RangeEntry

	// Min is the minimum valid score (inclusive).
	Min float64

	// Max is the maximum valid score (inclusive).
	Max float64
}

// RangeEntry defines a minimum score threshold for a level.
type RangeEntry struct {
	// LevelID is the framework level this range maps to.
	LevelID string

	// MinScore is the minimum score (inclusive) for this level.
	MinScore float64
}

// LevelFromScore returns the level for the given score.
// Returns nil and ErrNilFramework if Framework is not set.
// Returns nil and ErrScoreOutOfRange if score is outside Min/Max bounds.
// Returns nil and ErrUnknownLevel if the matched range references a level
// that is not in the Framework.
// Returns nil, nil if no matching range is found.
func (sr *ScoreRange) LevelFromScore(score float64) (*Level, error) {
	if sr.Framework == nil {
		return nil, ErrNilFramework
	}
	if score < sr.Min || score > sr.Max {
		return nil, ErrScoreOutOfRange
	}

	// Find the first range where score >= MinScore
	// Ranges should be ordered from highest to lowest MinScore
	for _, r := range sr.Ranges {
		if score >= r.MinScore {
			if level := sr.Framework.Parse(r.LevelID); level != nil {
				return level, nil
			}
			return nil, fmt.Errorf("%w: %q", ErrUnknownLevel, r.LevelID)
		}
	}

	return nil, nil
}

// MustLevelFromScore returns the level for the given score.
// Panics if score is out of range. Returns nil if no match found.
func (sr *ScoreRange) MustLevelFromScore(score float64) *Level {
	level, err := sr.LevelFromScore(score)
	if err != nil {
		panic(err)
	}
	return level
}

// CVSSScoreRange returns a ScoreRange for CVSS v3 scores mapped to Severity levels.
// CVSS ranges: None (0), Low (0.1-3.9), Medium (4.0-6.9), High (7.0-8.9), Critical (9.0-10.0)
func CVSSScoreRange() *ScoreRange {
	return &ScoreRange{
		Framework: Severity(),
		Min:       0.0,
		Max:       10.0,
		Ranges: []RangeEntry{
			{LevelID: "critical", MinScore: 9.0},
			{LevelID: "high", MinScore: 7.0},
			{LevelID: "medium", MinScore: 4.0},
			{LevelID: "low", MinScore: 0.1},
			{LevelID: "informational", MinScore: 0.0},
		},
	}
}

// PercentageScoreRange returns a ScoreRange for percentage scores (0-100)
// mapped to a 4-level framework using quartiles.
func PercentageScoreRange(f *Framework) *ScoreRange {
	if f == nil || len(f.Levels) < 4 {
		return nil
	}
	return &ScoreRange{
		Framework: f,
		Min:       0.0,
		Max:       100.0,
		Ranges: []RangeEntry{
			{LevelID: f.Levels[0].ID, MinScore: 75.0},
			{LevelID: f.Levels[1].ID, MinScore: 50.0},
			{LevelID: f.Levels[2].ID, MinScore: 25.0},
			{LevelID: f.Levels[3].ID, MinScore: 0.0},
		},
	}
}

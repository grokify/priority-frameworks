// Package priorityframeworks provides pluggable prioritization systems.
//
// Built-in frameworks include Severity, Priority (P#), IETF RFC 2119, MoSCoW,
// and a General requirement framework. Users can choose which framework fits
// their organization's practices.
//
// The position in the Levels slice implies priority order (index 0 = highest).
package priorityframeworks

// Framework represents a prioritization system with ordered levels.
type Framework struct {
	// ID is the unique identifier (e.g., "severity", "moscow").
	ID string `json:"id" yaml:"id"`

	// Name is the display name (e.g., "Severity", "MoSCoW").
	Name string `json:"name" yaml:"name"`

	// Description explains when to use this framework.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Levels are ordered from highest to lowest priority.
	// Index 0 is the highest priority level.
	Levels []Level `json:"levels" yaml:"levels"`
}

// Level represents a single priority level within a framework.
type Level struct {
	// ID is the canonical identifier (e.g., "critical", "must", "p0").
	ID string `json:"id" yaml:"id"`

	// Name is the display name (e.g., "Critical", "MUST", "P0").
	Name string `json:"name" yaml:"name"`

	// Aliases are alternative names that parse to this level.
	Aliases []string `json:"aliases,omitempty" yaml:"aliases,omitempty"`

	// Actionable indicates whether items at this level require action.
	// For example, "Critical" and "Must have" are actionable; "Informational" is not.
	Actionable bool `json:"actionable" yaml:"actionable"`

	// Color is the suggested display color (hex code).
	Color string `json:"color,omitempty" yaml:"color,omitempty"`
}

// IndexOf returns the index of the level with the given ID or name.
// Returns -1 if not found.
func (f *Framework) IndexOf(idOrName string) int {
	for i, l := range f.Levels {
		if l.ID == idOrName || l.Name == idOrName {
			return i
		}
		for _, alias := range l.Aliases {
			if alias == idOrName {
				return i
			}
		}
	}
	return -1
}

// Parse returns the Level matching the given string (ID, Name, or Alias).
// Returns nil if not found.
func (f *Framework) Parse(s string) *Level {
	idx := f.IndexOf(s)
	if idx < 0 {
		return nil
	}
	return &f.Levels[idx]
}

// Default returns the default level (middle of the range).
func (f *Framework) Default() *Level {
	if len(f.Levels) == 0 {
		return nil
	}
	mid := len(f.Levels) / 2
	return &f.Levels[mid]
}

// Highest returns the highest priority level (index 0).
func (f *Framework) Highest() *Level {
	if len(f.Levels) == 0 {
		return nil
	}
	return &f.Levels[0]
}

// Lowest returns the lowest priority level (last index).
func (f *Framework) Lowest() *Level {
	if len(f.Levels) == 0 {
		return nil
	}
	return &f.Levels[len(f.Levels)-1]
}

// ActionableLevels returns only levels where Actionable is true.
func (f *Framework) ActionableLevels() []Level {
	var result []Level
	for _, l := range f.Levels {
		if l.Actionable {
			result = append(result, l)
		}
	}
	return result
}

// Compare compares two level identifiers within this framework.
// Returns: 1 if a > b (higher priority), -1 if a < b, 0 if equal.
// Returns 0 if either level is not found.
func (f *Framework) Compare(a, b string) int {
	idxA := f.IndexOf(a)
	idxB := f.IndexOf(b)
	if idxA < 0 || idxB < 0 {
		return 0
	}
	// Lower index = higher priority
	switch {
	case idxA < idxB:
		return 1
	case idxA > idxB:
		return -1
	default:
		return 0
	}
}

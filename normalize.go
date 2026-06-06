package priorityframeworks

// NormalizedPriority represents a normalized priority bucket (1-4).
// This allows comparison across different frameworks.
type NormalizedPriority int

const (
	NormalizedCritical NormalizedPriority = 4 // Highest
	NormalizedHigh     NormalizedPriority = 3
	NormalizedMedium   NormalizedPriority = 2
	NormalizedLow      NormalizedPriority = 1 // Lowest
)

// String returns the display name for the normalized priority.
func (p NormalizedPriority) String() string {
	switch p {
	case NormalizedCritical:
		return "Critical"
	case NormalizedHigh:
		return "High"
	case NormalizedMedium:
		return "Medium"
	case NormalizedLow:
		return "Low"
	default:
		return "Unknown"
	}
}

// Normalize converts a level index to a normalized priority (1-4).
// This enables sorting and comparison across different frameworks.
//
// The normalization maps the level's position in the framework to one of
// four buckets: Critical (4), High (3), Medium (2), Low (1).
func Normalize(f *Framework, levelID string) NormalizedPriority {
	idx := f.IndexOf(levelID)
	if idx < 0 {
		return NormalizedMedium // Default
	}
	return NormalizeIndex(idx, len(f.Levels))
}

// NormalizeIndex converts a level index and total count to normalized priority.
func NormalizeIndex(index, total int) NormalizedPriority {
	if total <= 0 {
		return NormalizedMedium
	}
	// Map to 4 buckets
	ratio := float64(index) / float64(total)
	switch {
	case ratio < 0.25:
		return NormalizedCritical
	case ratio < 0.5:
		return NormalizedHigh
	case ratio < 0.75:
		return NormalizedMedium
	default:
		return NormalizedLow
	}
}

// CompareAcross compares levels from potentially different frameworks.
// Returns: 1 if a > b (higher priority), -1 if a < b, 0 if equal.
func CompareAcross(fA *Framework, levelA string, fB *Framework, levelB string) int {
	normA := Normalize(fA, levelA)
	normB := Normalize(fB, levelB)
	switch {
	case normA > normB:
		return 1
	case normA < normB:
		return -1
	default:
		return 0
	}
}

// MapTo maps a level from one framework to the closest equivalent in another.
// Returns nil if the source level is not found.
func MapTo(src *Framework, srcLevel string, dst *Framework) *Level {
	norm := Normalize(src, srcLevel)
	// Find the level in dst that maps to the same normalized priority
	for i, l := range dst.Levels {
		if NormalizeIndex(i, len(dst.Levels)) == norm {
			return &l
		}
	}
	// Fallback to default
	return dst.Default()
}

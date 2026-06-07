package priorityframeworks

// LevelCounts tracks counts of items at each priority level.
// This is useful for dashboards, reports, and aggregations.
type LevelCounts struct {
	// Framework is the associated framework (optional, for validation).
	Framework *Framework

	// Counts maps level IDs to their count.
	Counts map[string]int
}

// NewLevelCounts creates a new LevelCounts for the given framework.
// If framework is provided, initializes all levels to zero.
func NewLevelCounts(f *Framework) *LevelCounts {
	lc := &LevelCounts{
		Framework: f,
		Counts:    make(map[string]int),
	}
	if f != nil {
		for _, level := range f.Levels {
			lc.Counts[level.ID] = 0
		}
	}
	return lc
}

// Add increments the count for the given level ID.
// If the level doesn't exist in the map, it's created with count 1.
func (lc *LevelCounts) Add(levelID string, count int) {
	if lc.Counts == nil {
		lc.Counts = make(map[string]int)
	}
	lc.Counts[levelID] += count
}

// Increment adds 1 to the count for the given level ID.
func (lc *LevelCounts) Increment(levelID string) {
	lc.Add(levelID, 1)
}

// Get returns the count for the given level ID.
// Returns 0 if the level is not found.
func (lc *LevelCounts) Get(levelID string) int {
	if lc.Counts == nil {
		return 0
	}
	return lc.Counts[levelID]
}

// Total returns the sum of all counts.
func (lc *LevelCounts) Total() int {
	total := 0
	for _, count := range lc.Counts {
		total += count
	}
	return total
}

// ActionableTotal returns the sum of counts for actionable levels only.
// Requires Framework to be set; returns 0 if Framework is nil.
func (lc *LevelCounts) ActionableTotal() int {
	if lc.Framework == nil {
		return 0
	}
	total := 0
	for _, level := range lc.Framework.Levels {
		if level.Actionable {
			total += lc.Counts[level.ID]
		}
	}
	return total
}

// HigherThan returns the sum of counts for levels higher than the given level.
// Requires Framework to be set; returns 0 if Framework is nil.
func (lc *LevelCounts) HigherThan(levelID string) int {
	if lc.Framework == nil {
		return 0
	}
	idx := lc.Framework.IndexOf(levelID)
	if idx < 0 {
		return 0
	}
	total := 0
	for i, level := range lc.Framework.Levels {
		if i < idx {
			total += lc.Counts[level.ID]
		}
	}
	return total
}

// HigherThanOrEqual returns the sum of counts for levels >= the given level.
// Requires Framework to be set; returns 0 if Framework is nil.
func (lc *LevelCounts) HigherThanOrEqual(levelID string) int {
	if lc.Framework == nil {
		return 0
	}
	idx := lc.Framework.IndexOf(levelID)
	if idx < 0 {
		return 0
	}
	total := 0
	for i, level := range lc.Framework.Levels {
		if i <= idx {
			total += lc.Counts[level.ID]
		}
	}
	return total
}

// LowerThan returns the sum of counts for levels lower than the given level.
// Requires Framework to be set; returns 0 if Framework is nil.
func (lc *LevelCounts) LowerThan(levelID string) int {
	if lc.Framework == nil {
		return 0
	}
	idx := lc.Framework.IndexOf(levelID)
	if idx < 0 {
		return 0
	}
	total := 0
	for i, level := range lc.Framework.Levels {
		if i > idx {
			total += lc.Counts[level.ID]
		}
	}
	return total
}

// IsEmpty returns true if all counts are zero.
func (lc *LevelCounts) IsEmpty() bool {
	return lc.Total() == 0
}

// Reset sets all counts to zero.
func (lc *LevelCounts) Reset() {
	for k := range lc.Counts {
		lc.Counts[k] = 0
	}
}

// Merge adds counts from another LevelCounts.
func (lc *LevelCounts) Merge(other *LevelCounts) {
	if other == nil || other.Counts == nil {
		return
	}
	if lc.Counts == nil {
		lc.Counts = make(map[string]int)
	}
	for levelID, count := range other.Counts {
		lc.Counts[levelID] += count
	}
}

// Clone returns a deep copy of the LevelCounts.
func (lc *LevelCounts) Clone() *LevelCounts {
	clone := &LevelCounts{
		Framework: lc.Framework,
		Counts:    make(map[string]int),
	}
	for k, v := range lc.Counts {
		clone.Counts[k] = v
	}
	return clone
}

// Slice returns counts in framework level order.
// Requires Framework to be set; returns nil if Framework is nil.
func (lc *LevelCounts) Slice() []int {
	if lc.Framework == nil {
		return nil
	}
	counts := make([]int, len(lc.Framework.Levels))
	for i, level := range lc.Framework.Levels {
		counts[i] = lc.Counts[level.ID]
	}
	return counts
}

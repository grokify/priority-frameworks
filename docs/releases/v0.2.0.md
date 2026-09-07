# Release Notes v0.2.0

This release adds score-to-level mapping, level counting, and separates IETF requirements from prohibitions for cleaner prioritization.

## Highlights

- Score-to-level mapping with built-in CVSS support
- Level counting for dashboards and reports
- IETF RFC 2119 split into requirements and prohibitions frameworks

## Breaking Changes

### IETF Framework Split

`IETF()` now returns only requirement levels (MUST, SHOULD, MAY) for prioritization. Prohibition levels have moved to `IETFProhibitions()`.

```go
// Before v0.2.0
ietf := pf.IETF() // 5 levels: MUST, SHOULD, MAY, SHOULD NOT, MUST NOT

// v0.2.0+
ietf := pf.IETF()              // 3 levels: MUST, SHOULD, MAY
prohib := pf.IETFProhibitions() // 2 levels: MUST NOT, SHOULD NOT
```

**Rationale:** Requirements and prohibitions serve different purposes:
- **Requirements** (MUST/SHOULD/MAY) are backlog items to prioritize and implement
- **Prohibitions** (MUST NOT/SHOULD NOT) are constraints to validate during compliance checks

## New Features

### ScoreRange

Map numeric scores to priority levels:

```go
// CVSS scores to Severity
sr := pf.CVSSScoreRange()
level, _ := sr.LevelFromScore(7.5) // "High"

// CVSS thresholds:
// Critical: 9.0+
// High: 7.0+
// Medium: 4.0+
// Low: 0.1+
// Informational: 0

// Percentage-based
sr := pf.PercentageScoreRange(pf.Severity())
level, _ := sr.LevelFromScore(85.0) // "Critical" (75%+)
```

### LevelCounts

Track counts per priority level for dashboards:

```go
counts := pf.NewLevelCounts(pf.Severity())
counts.Add("critical", 2)
counts.Add("high", 5)
counts.Add("medium", 10)

counts.Total()              // 17
counts.ActionableTotal()    // 17 (excludes Informational)
counts.HigherThan("medium") // 7 (Critical + High)
counts.Slice()              // [2, 5, 10, 0, 0]
```

### IETFProhibitions Framework

New framework for compliance validation:

```go
prohib := pf.IETFProhibitions()
// Levels: MUST NOT, SHOULD NOT

// Use for constraint checking, not backlog prioritization
for _, constraint := range constraints {
    level := prohib.Parse(constraint.Level)
    if level.ID == "must-not" {
        // Critical compliance violation
    }
}
```

## Migration Guide

If you were using IETF's MUST NOT or SHOULD NOT levels:

```go
// Before
ietf := pf.IETF()
level := ietf.Parse("MUST NOT") // worked

// After
prohib := pf.IETFProhibitions()
level := prohib.Parse("MUST NOT") // use prohibitions framework
```

## Links

- [Documentation](https://github.com/grokify/priority-frameworks#readme)
- [Changelog](https://github.com/grokify/priority-frameworks/blob/main/CHANGELOG.md)

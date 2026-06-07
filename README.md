# Priority Frameworks

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/priority-frameworks/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/priority-frameworks/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/priority-frameworks/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/priority-frameworks/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/priority-frameworks/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/priority-frameworks/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/priority-frameworks
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/priority-frameworks
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/priority-frameworks
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/priority-frameworks
 [viz-svg]: https://img.shields.io/badge/visualization-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fpriority-frameworks
 [loc-svg]: https://tokei.rs/b1/github/grokify/priority-frameworks
 [repo-url]: https://github.com/grokify/priority-frameworks
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/priority-frameworks/blob/main/LICENSE

Pluggable prioritization systems for Go applications.

## Installation

```bash
go get github.com/grokify/priority-frameworks
```

## Overview

This package provides a unified interface for working with different prioritization frameworks. Instead of hardcoding a single priority system, applications can allow users to choose the framework that fits their organization's practices.

## Built-in Frameworks

| Framework | Levels | Use Case |
|-----------|--------|----------|
| **Severity** | Critical, High, Medium, Low, Informational | Security, incidents, bugs |
| **Priority (P#)** | P0, P1, P2, P3, P4 | Engineering work prioritization |
| **IETF RFC 2119** | MUST, MUST NOT, SHOULD, SHOULD NOT, MAY | Technical specifications |
| **MoSCoW** | Must have, Should have, Could have, Won't have | Agile, product management |
| **General** | Required, Recommended, Optional, Avoid | General-purpose requirement levels |

## Usage

```go
import pf "github.com/grokify/priority-frameworks"

// Get a built-in framework (returns nil if not found)
framework := pf.Get("severity")
if framework == nil {
    log.Fatal("unknown framework")
}

// Parse a level (returns nil if not found)
level := framework.Parse("Critical")
if level == nil {
    log.Fatal("unknown level")
}
fmt.Println(level.Name)       // "Critical"
fmt.Println(level.Actionable) // true
fmt.Println(level.Color)      // "#7f1d1d"

// Compare levels within a framework
result := framework.Compare("critical", "low") // 1 (critical > low)

// Get highest/lowest priority levels
highest := framework.Highest() // Critical
lowest := framework.Lowest()   // Informational

// Get only actionable levels
actionable := framework.ActionableLevels()

// List all built-in frameworks
ids := pf.AllBuiltinIDs()    // ["severity", "priority", "ietf", "moscow", "general"]
frameworks := pf.All()       // []*Framework
```

## Cross-Framework Normalization

Compare and map levels across different frameworks:

```go
severity := pf.Severity()
moscow := pf.MoSCoW()

// Normalize to a common scale (1-4)
norm := pf.Normalize(severity, "critical") // NormalizedCritical (4)

// Compare across frameworks
result := pf.CompareAcross(severity, "critical", moscow, "must") // 0 (equal)

// Map a level to another framework
mapped := pf.MapTo(severity, "critical", moscow)
fmt.Println(mapped.Name) // "Must have"
```

## Score Ranges

Map numeric scores to priority levels (e.g., CVSS scores to severity):

```go
// Built-in CVSS score range
sr := pf.CVSSScoreRange()
level, err := sr.LevelFromScore(7.5)
if err != nil {
    log.Fatal(err)
}
fmt.Println(level.Name) // "High"

// CVSS thresholds: Critical (9.0+), High (7.0+), Medium (4.0+), Low (0.1+), Info (0)

// Custom percentage-based range
sr := pf.PercentageScoreRange(pf.Severity())
level, _ := sr.LevelFromScore(85.0)  // "Critical" (75%+)
```

## Level Counts

Track counts of items at each priority level:

```go
counts := pf.NewLevelCounts(pf.Severity())

// Add counts
counts.Increment("critical")
counts.Add("high", 5)
counts.Add("medium", 10)

// Query counts
fmt.Println(counts.Total())           // 16
fmt.Println(counts.ActionableTotal()) // 16 (excludes Informational)
fmt.Println(counts.HigherThan("medium"))      // 6 (Critical + High)
fmt.Println(counts.HigherThanOrEqual("high")) // 6 (Critical + High)

// Merge from another source
otherCounts := pf.NewLevelCounts(pf.Severity())
otherCounts.Add("critical", 3)
counts.Merge(otherCounts)

// Get counts in framework order
slice := counts.Slice() // [4, 5, 10, 0, 0] (Critical to Informational)
```

## Custom Frameworks

Create your own framework:

```go
custom := &pf.Framework{
    ID:          "custom",
    Name:        "Custom Framework",
    Description: "My organization's priority system",
    Levels: []pf.Level{
        {ID: "urgent", Name: "Urgent", Actionable: true, Color: "#dc2626"},
        {ID: "normal", Name: "Normal", Actionable: true, Color: "#ca8a04"},
        {ID: "backlog", Name: "Backlog", Actionable: false, Color: "#6b7280"},
    },
}
```

## Design

- **Position = Priority**: Level position in the `Levels` slice determines priority (index 0 = highest)
- **Simple data types**: Frameworks are plain structs, easily serializable to JSON/YAML
- **Actionable flag**: Distinguishes between levels that require action vs. informational
- **Aliases**: Multiple names can map to the same level (e.g., "CRITICAL", "Crit", "S1")

## Contributing

Contributions are welcome. Please see the [CHANGELOG](CHANGELOG.md) for recent changes.

## License

MIT - see [LICENSE](LICENSE) for details.

# priority-frameworks

Pluggable prioritization systems for Go applications.

## Overview

This package provides a unified interface for working with different prioritization frameworks. Instead of hardcoding a single priority system, applications can allow users to choose the framework that fits their organization's practices.

## Built-in Frameworks

| Framework | Levels | Use Case |
|-----------|--------|----------|
| **Severity** | Critical, High, Medium, Low, Informational | Security, incidents, bugs |
| **Priority (P#)** | P0, P1, P2, P3, P4 | Engineering work prioritization |
| **IETF RFC 2119** | MUST, MUST NOT, SHOULD, SHOULD NOT, MAY | Technical specifications |
| **MoSCoW** | Must have, Should have, Could have, Won't have | Agile, product management |
| **General** | Required, Recommended, Optional, Avoid | General-purpose requirements |

## Usage

```go
import pf "github.com/grokify/priority-frameworks"

// Get a built-in framework
framework := pf.Get("severity")

// Parse a level
level := framework.Parse("Critical")
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

## License

MIT

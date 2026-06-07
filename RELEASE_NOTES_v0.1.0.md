# Release Notes v0.1.0

Initial release of priority-frameworks, a Go library for standardized priority level definitions with cross-framework normalization.

## Highlights

- Go library for standardized priority level definitions with 5 built-in frameworks and cross-framework normalization

## Features

### Core Types

- `Framework` and `Level` types for defining custom priority systems
- `GetLevelByID()` and `GetLevelByNameOrAlias()` lookup functions
- Support for level aliases (e.g., "Sev1" for "Critical")

### Built-in Frameworks

| Framework | Levels | Use Case |
|-----------|--------|----------|
| `Severity` | Critical, High, Medium, Low, Info | Security vulnerabilities, incident response |
| `Priority` | P0, P1, P2, P3, P4 | Engineering task prioritization |
| `IETF` | MUST, SHOULD, MAY | RFC 2119 requirement levels |
| `MoSCoW` | Must Have, Should Have, Could Have, Won't Have | Product prioritization |
| `General` | High, Medium, Low | Simple 3-level prioritization |

### Cross-Framework Normalization

- `Normalize()` - Convert any level to a 0-100 scale for comparison
- `CompareAcross()` - Compare priorities from different frameworks
- `MapTo()` - Find equivalent levels between frameworks

## Installation

```bash
go get github.com/grokify/priority-frameworks
```

## Quick Start

```go
import pf "github.com/grokify/priority-frameworks"

// Use built-in frameworks
sev := pf.Severity()
level, _ := sev.GetLevelByID("critical")

// Compare across frameworks
result := pf.CompareAcross(pf.Severity(), "high", pf.Priority(), "P1")
// result == 0 (equivalent priority)

// Map between frameworks
equiv := pf.MapTo(pf.Severity(), "critical", pf.Priority())
// equiv.ID == "P0"
```

## What's Next

- Additional frameworks (CVSS, DREAD, STRIDE)
- JSON/YAML serialization
- Framework composition and extension

## Links

- [Documentation](https://github.com/grokify/priority-frameworks#readme)
- [Changelog](https://github.com/grokify/priority-frameworks/blob/main/CHANGELOG.md)

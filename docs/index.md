# priority-frameworks

Pluggable prioritization systems for Go applications. Instead of every
consumer hardcoding its own Critical/High/Medium/Low (or P0-P4, or
MoSCoW) definition, applications share one canonical framework — and
avoid the drift that comes from N near-duplicate copies of the same
level order and accepted spellings.

## Built-in frameworks

Levels are ordered by implementation priority (index 0 = highest) — this
ordering drives code actions: items to implement come before items to
avoid.

| Framework | Levels (highest → lowest) | Use case |
|-----------|---------------------------|----------|
| Severity | Critical, High, Medium, Low, Informational | Security, incidents, bugs |
| Priority (P#) | P0, P1, P2, P3, P4 | Engineering work prioritization |
| IETF RFC 2119 | MUST, SHOULD, MAY | Requirements — what to implement |
| IETF Prohibitions | MUST NOT, SHOULD NOT | Compliance — constraints to validate |
| MoSCoW | Must have, Should have, Could have, Won't have | Agile, product management |
| General | Required, Recommended, Optional, Avoid | General-purpose requirement levels |

## Library

- **Parse** — `Framework.Parse(s)`/`IndexOf(s)` accept a level's ID, Name,
  or any of its `Aliases` (e.g. `"CRITICAL"`, `"Crit"`, `"S1"` all resolve
  to Severity's `critical`).
- **Emit full or abbreviated** — `Level.Name` is the canonical display form
  (e.g. `"Critical"`); `Level.Abbreviation` is a distinct, opt-in short form
  (e.g. `"CRIT"`). `Framework.AbbreviationFor(idOrName)` returns the
  abbreviation with a fallback to `Name` when a level has none (e.g.
  Priority's `"P0"` is already minimal). A caller picks whichever fits a
  given call site — including both, in different places, within the same
  application.
- **Compare and rank** — `Framework.Compare(a, b)`, `Highest()`, `Lowest()`,
  `Default()`, `ActionableLevels()`.
- **Cross-framework normalization** — `Normalize`, `CompareAcross`, `MapTo`
  translate a level's position onto a common 0-1 scale so two different
  frameworks (e.g. Severity and MoSCoW) can be compared or converted.
- **Score ranges** — `ScoreRange.LevelFromScore` maps a numeric score (e.g.
  a CVSS score) to a level; `CVSSScoreRange()` and `PercentageScoreRange()`
  are built in.
- **Level counts** — `LevelCounts` tracks and aggregates counts of items at
  each level, with `Total()`, `ActionableTotal()`, `HigherThan(level)`, and
  `Merge()` for combining counts from multiple sources.
- **Custom frameworks** — `Framework`/`Level` are plain structs; define
  your own alongside or instead of the built-ins.

## Installation

```bash
go get github.com/grokify/priority-frameworks
```

See the [README](https://github.com/grokify/priority-frameworks#readme) for
full usage examples, and [Releases](releases/v0.4.0.md) for version history.

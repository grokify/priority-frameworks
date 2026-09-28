# AGENTS.md

Project-specific instructions for coding agents. See
`~/go/src/github.com/grokify/.github/CLAUDE.md` for org-wide conventions.

## What this is

A small, dependency-free library of pluggable prioritization frameworks
(Severity, Priority P#, IETF RFC 2119, MoSCoW, General). It exists so
consumers (e.g. `prism-capability`, PubGuard) share one canonical
definition of level order, accepted spellings, and abbreviations instead of
each maintaining a near-duplicate copy that drifts over time.

## Development

```bash
go build ./...
go test ./...
golangci-lint run
```

## Hard invariants

- **Zero dependencies.** `go.mod` has no `require` block; keep it that way.
  This is a leaf library other repos depend on — it must never depend on a
  consumer or pull in anything beyond the standard library.
- **`Name` is the canonical display form; `Abbreviation` is a distinct,
  opt-in short form — not a replacement.** Both are always populated
  intentionally (or `Abbreviation` left empty when `Name` is already
  minimal, e.g. `"P0"`, `"MUST"`). A caller picks whichever fits a given
  call site; `Framework.AbbreviationFor` falls back to `Name` when a level
  has no distinct abbreviation. Never remove `Name` in favor of always
  emitting `Abbreviation`, or vice versa.
- **`Abbreviation` is stored upper-case, as the one canonical value to
  build from.** A caller wanting different casing (title case, lower case)
  applies that transform itself. Do not add per-case fields.
- **`Aliases` only affects parsing (`IndexOf`/`Parse`); it never affects
  emission.** Adding a new accepted input spelling to `Aliases` must not
  change what `Name`/`Abbreviation` return for that level.
- **This is where severity/priority logic gets fixed, not in a consumer.**
  If a consumer needs a new accepted spelling or abbreviation for a
  built-in framework (e.g. PubGuard needing `"CRIT"` for Severity), add it
  here and let consumers pick up the new version — never special-case
  parsing/emission logic in a downstream repo. That defeats the reason this
  library exists.
- **Level order (`Framework.Levels` index) is part of the public API.**
  Reordering, inserting, or removing a level for a built-in framework is a
  breaking change (it shifts `Rank`/`IndexOf`/`Normalize` results for every
  consumer) and requires a major version bump. Adding a new alias or
  `Abbreviation` to an existing level, or adding an entirely new framework,
  is additive.

## Release workflow

Use `schangelog` (TOON output, more token-efficient than raw `git log`)
rather than raw git commands for commit analysis:

1. Parse commits: `schangelog parse-commits --since=<tag>`
2. Add a release entry to `CHANGELOG.json` (see existing entries for the
   category shape: `highlights`, `added`, `fixed`, `changed`, `breaking`,
   `tests`, `documentation`)
3. Validate: `schangelog validate CHANGELOG.json`
4. Regenerate: `schangelog generate CHANGELOG.json -o CHANGELOG.md`
5. Write `docs/releases/vX.Y.Z.md` (follow the format of prior release notes)
6. Add the new release page to `mkdocs.yml` nav and update `docs/index.md`
7. Update `README.md` for any new user-facing feature

Do not tag until commits are pushed and CI passes (per the org pre-push /
release-tagging checklist).

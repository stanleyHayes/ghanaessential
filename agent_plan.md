# GhanaEssential execution ledger

Status: Public beta — verified 2026-09-01
Canonical hostname: `essential.digitalghana.dev`

## Product definition gate

- [x] Problem, users and non-goals approved in the portfolio brief and product boundary ADR.
- [x] Four official institutional sources reviewed; factual contact metadata only.
- [x] Five-record versioned fixture with deterministic freshness and source invariants approved.
- [x] Safety/privacy review complete: no dispatch claims, incident intake, location, medical or personal data.
- [x] Read-only web/API and identical offline export boundary approved.

## Live task board

| ID | Task | Status | Owner | Dependency | Evidence |
|---|---|---|---|---|---|
| P-0.1 | Product definition and source review | Done | Codex | — | NADMO, NAS, GNFS and Ghana Police official sources reviewed 2026-09-01 |
| P-0.2 | Domain contracts and fixtures | Done | Codex | P-0.1 | Five source-complete contacts; staleness, national-112 and no-personal-data invariants |
| P-1.1 | Implementation | Done | Codex | P-0.2 | Read-only Go API, identical offline export and Next.js contact UI pass the complete local quality gate |
| P-2.1 | Production release | Done | Codex | P-1.1 | CI, canonical TLS, API/CORS/offline export, browser safety/SEO and provider rollback/restore evidence recorded |

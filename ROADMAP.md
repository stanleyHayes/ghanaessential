# GhanaEssential roadmap

This roadmap is **directional, not a commitment**. It contains no delivery dates and no promises. It is a readable summary of the execution ledger; the machine-readable state of record is [`agent_plan.md`](agent_plan.md), and that file wins wherever this page disagrees with it.

Lifecycle words are used strictly: `proposed`, `building`, `beta`, `stable`, `externally blocked`, `retired`, `deferred`. Nothing is marked done here without linked evidence in the ledger. The product boundary that constrains everything below is [ADR-0001](docs/adr/0001-product-boundary.md).

Current state: **public beta**, verified 2026-09-01. Web and API are live. Nothing here has reached `stable`.

---

## Now — shipped in the current beta

Product definition gate, all approved before implementation:

- [x] Problem, users and non-goals approved in the portfolio brief and the product boundary ADR.
- [x] Four official institutional sources reviewed — NADMO, the National Ambulance Service, the Ghana National Fire Service and the Ghana Police Service — recorded in [`docs/governance/source-register.json`](docs/governance/source-register.json) as factual contact metadata only.
- [x] Five-record versioned fixture with deterministic freshness and source invariants approved.
- [x] Safety and privacy review complete: no dispatch claims, no incident intake, no location, medical or personal data.
- [x] Read-only web and API boundary approved, with an identical offline export.

Delivery, from the live task board in [`agent_plan.md`](agent_plan.md):

- [x] **P-0.1 Product definition and source review.** Four official sources reviewed 2026-09-01.
- [x] **P-0.2 Domain contracts and fixtures.** Five source-complete contacts with staleness, national-112 and no-personal-data invariants, in [`data/contacts.json`](data/contacts.json) and [`tests/contact.test.mjs`](tests/contact.test.mjs).
- [x] **P-1.1 Implementation.** Read-only Go API ([`cmd/api/main.go`](cmd/api/main.go)), identical offline export, and the Next.js contact UI ([`app/`](app)), all passing the complete local quality gate.
- [x] **P-2.1 Production release.** CI, canonical TLS, API and CORS behaviour, offline export, browser safety and SEO checks, and provider rollback/restore evidence recorded in [`docs/runbooks/release-evidence.md`](docs/runbooks/release-evidence.md).

What that means in practice:

- [x] `GET /health`, `GET /v1/contacts` and `GET /v1/offline.json` live at `api-essential.digitalghana.dev`.
- [x] Contact site live at `essential.digitalghana.dev` with a visible disclaimer, `tel:` links, per-record official source links and a visible checked date.
- [x] A fail-closed freshness gate: [`scripts/validate.rb`](scripts/validate.rb) aborts if the fixture is more than 30 days past review.
- [x] A proven rollback path on both providers, recorded with deployment identifiers.

---

## Next — the open gates

These are the real blockers. Several are decisions or recurring operational work rather than code.

| Gate | What it unblocks | Blocking dependency |
|---|---|---|
| Fixture re-review within the 30-day window | Continued publication of all five records | Every source URL in the source register must be re-opened and `checkedAt` updated; otherwise `pnpm check` aborts with `contact fixture is stale` and CI goes red. Currently a manual release step in [`docs/runbooks/operations.md`](docs/runbooks/operations.md). |
| Automated source-liveness check | Freshness that does not depend on someone remembering | No script in [`scripts/`](scripts) re-requests the registered source URLs; drift is detected only by hand at release. |
| Negative tests for the validator | Trusting the safety gate itself | [`scripts/validate.rb`](scripts/validate.rb) enforces nine abort conditions and none of them is proven to fail on bad input. |
| Broader contact coverage | More than five national-level records | The beta is pinned at exactly five `VERIFIED` records by the validator, the Go fixture test and the Node test. Any increase needs a recorded source and licence decision per record, plus the invariants updated in the same change. |
| Human review and publish workflow | `stable` | The pre-release source re-check, downgrade rule and rollback are specified in the operations runbook but are not an implemented reviewer workflow. |
| Local development script | Contributor onboarding | [`package.json`](package.json) exposes `typecheck`, `test`, `build` and `check` only; there is no `dev` script. |

---

## Later — deferred by design

- **A generated SDK or published client package.** [`contracts/README.md`](contracts/README.md) is explicit: the fixture is the contract until demonstrated consumers justify a generated SDK. Until then, consumers pin the dataset version and read the JSON.
- **Shared portfolio packages.** [ADR-0001](docs/adr/0001-product-boundary.md) accepts that some configuration and small primitives are repeated across products until two proven consumers justify a versioned shared package. Independent failure is preferred to premature sharing.
- **Cross-product integration.** Any consumption of another Digital Ghana product happens only through a versioned public contract or a pinned dataset artifact — never a shared database.

---

## Explicitly out of scope

These are stated non-goals, not a backlog. They come from [ADR-0001](docs/adr/0001-product-boundary.md), [`contracts/README.md`](contracts/README.md), [`docs/runbooks/operations.md`](docs/runbooks/operations.md) and the invariants in [`tests/contact.test.mjs`](tests/contact.test.mjs).

- **Emergency dispatch.** This product never summons help. In an emergency, call 112.
- **Incident reporting or intake of any kind.** No form, no endpoint, no writable store.
- **User location, coordinates, medical details or personal data.** Asserted by test, not by policy alone.
- **Any availability, pickup or response-time guarantee.** Official publication is not proof that a call will connect. "Source checked" must never be described as anything more than that.
- **Medical or safety advice**, or positioning as a replacement for calling 112.
- **Any claim of government operation, endorsement, affiliation or official status.**
- **A comprehensive national directory.** The dataset stays a deliberately small, source-linked verified subset that declares its own boundaries.
- **Mutation, GraphQL or location endpoints.** The contract is three read-only GET routes.
- **Silently rewriting a record whose source has disappeared or changed.** Such a record is downgraded before release and stays reviewable.
- **Publishing data with an unknown or blocked licence decision.** The validator rejects it.

---

## How this roadmap changes

A roadmap item moves only when the evidence moves with it: the task board in [`agent_plan.md`](agent_plan.md), this page, and the dated evidence under [`docs/runbooks/`](docs/runbooks) are updated in the same reviewed change. A passing build is not evidence of deployment. See [`CONTRIBUTING.md`](CONTRIBUTING.md).

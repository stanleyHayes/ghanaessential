# GhanaEssential

`GhanaEssential` is an independent, source-linked directory of Ghana's emergency and essential-service contact numbers, with a visible verification date on every record and a versioned offline export.

![Licence Apache-2.0](https://img.shields.io/badge/licence-Apache--2.0-blue.svg)
![Quality](https://img.shields.io/github/actions/workflow/status/stanleyHayes/ghanaessential/quality.yml?branch=main&label=quality)
![Live at essential.digitalghana.dev](https://img.shields.io/badge/live-essential.digitalghana.dev-brightgreen)
![Runtime Go 1.23 and Next.js 16.3.3](https://img.shields.io/badge/runtime-Go%201.23%20%7C%20Next.js%2016.3.3-black)

> ### ⚠ Safety notice — read this first
>
> **In an emergency, call 112. Do not use this project to get help.**
>
> GhanaEssential is a **directory of contact numbers that official institutions have already published**. It does not dispatch help, does not contact any agency on your behalf, does not accept incident reports, and does not collect your location or any medical detail. It is not an emergency service and must never be presented as one. The dataset carries that warning in the payload itself: *"If life or property is in immediate danger, call 112. Network and agency availability can change; this directory does not dispatch help."*
>
> "Verified" here means exactly one thing: **the named official source published this number on the checked date.** It is not proof that a call will connect, that a line is staffed, or that anyone will respond.

---

## Live now

| Surface | URL | State |
|---|---|---|
| Web | [essential.digitalghana.dev](https://essential.digitalghana.dev) | live (Vercel) |
| API | [api-essential.digitalghana.dev](https://api-essential.digitalghana.dev/health) | live (Render) |

The API runs on Render's free tier and sleeps when idle. **The first request after a quiet period can take 30–60 seconds.** Re-run it and the second request is fast.

```sh
curl -s https://api-essential.digitalghana.dev/health
```

```json
{"contacts":5,"reviewedAt":"2026-09-01","status":"ok","version":"2026.09.01-beta.1"}
```

Download the same versioned fixture the website renders from — the filename carries the dataset version:

```sh
curl -sOJ https://api-essential.digitalghana.dev/v1/offline.json
# → ghanaessential-2026.09.01-beta.1.json
```

Five records. That number is the point: it is small, and it says so.

---

## The problem this solves

### For developers

If you have ever needed Ghanaian emergency contact data in software, you have paid at least one of these costs:

- **You hardcoded `112` and hoped.** Ghana's emergency numbers are spread across four separate institutional websites — NADMO, the National Ambulance Service, the Ghana National Fire Service and the Ghana Police Service — each publishing its own set (`112`, `191`, `192`, `193`, `18555`, plus direct agency lines). There is no single machine-readable endpoint, so the numbers get pasted into a constants file once and never looked at again.
- **You could not tell how old your copy was.** A scraped or copied list has no retrieval date attached, so nobody can answer "was this true when we shipped, and is it true now?" For safety-adjacent data that is not a style problem; it is the whole problem.
- **You could not cite anything.** When a number is queried, the useful answer is a URL and a date, not "we got it from somewhere". Without per-record provenance, a correction is an argument rather than a diff.
- **You had no offline story.** The moment emergency contacts matter most is the moment connectivity is least reliable. Rebuilding a bundled, versioned, verifiable snapshot is work every project redoes.
- **You had no safe default for stale data.** Most hand-maintained lists fail *open*: they keep serving whatever they last had, indefinitely.

What this gives you instead: a read-only HTTP API and an identical downloadable JSON fixture where **every record carries a stable `id`, a `sourceId`, a `sourceUrl`, a `status` and a `checkedAt` date**, the dataset carries a single `version` string you can pin, and the build **fails closed** — `scripts/validate.rb` aborts with `contact fixture is stale` if the fixture is more than 30 days past its review date, and both test suites reject any record that is not `VERIFIED` and source-complete. You can pin a version, cite where a number came from, and reproduce the answer later.

### For the community

Emergency contact information is the least glamorous and most consequential kind of public reference data. When the only convenient copies live in undated blog posts, screenshots and private spreadsheets, wrong numbers propagate quietly and nobody can correct them at the source.

This exists as open, source-linked, independent infrastructure so that:

- **Provenance is visible.** Every record names the authority, the exact page, and the date it was checked — on the website, in the API, and in the offline file. See [`docs/governance/source-register.json`](docs/governance/source-register.json).
- **Freshness is enforced, not promised.** Staleness is a build failure, and a record whose source disappears or changes must be downgraded before release. That rule is in [ADR-0001](docs/adr/0001-product-boundary.md) and in the [operations runbook](docs/runbooks/operations.md).
- **There is no lock-in.** Apache-2.0 code, a published contract, an exportable dataset, and a hostname that hides the provider rather than the other way round. If this project stops, the data and the code remain forkable.
- **Corrections are a first-class path.** Anyone can propose a correction with evidence; a human reviewer approves publication. See [`CONTRIBUTING.md`](CONTRIBUTING.md).

### What this is not

Drawn from the stated non-goals in [ADR-0001](docs/adr/0001-product-boundary.md), [`agent_plan.md`](agent_plan.md) and [`contracts/README.md`](contracts/README.md):

- **Not an emergency dispatch system.** Nothing here summons help. Call 112.
- **Not an incident-reporting channel.** There is no intake endpoint, no form, and no writable store of any kind.
- **Not an availability or response-time guarantee.** Official publication is not proof of network availability, pickup or response.
- **Not a collector of personal, location or medical data.** [`tests/contact.test.mjs`](tests/contact.test.mjs) asserts that no record carries a `person`, `patient`, `incident`, `location`, `coordinates` or `email` field.
- **Not a government service.** Not operated by, endorsed by or affiliated with the Government of Ghana or any of its agencies.
- **Not a comprehensive national directory.** It is a deliberately small, source-linked verified subset of five national-level contacts.
- **Not a medical or safety adviser, and not a replacement for calling 112.**
- **Not a mutable API.** No GraphQL, no mutations, no location endpoint. The fixture is the contract.

---

## Quickstart

Prerequisites: **Node.js 22** and **pnpm 10.17.1** (pinned by `packageManager` in [`package.json`](package.json)), **Go 1.23** ([`go.mod`](go.mod)), and **Ruby 3.4** for the validator — the versions [`.github/workflows/quality.yml`](.github/workflows/quality.yml) pins.

```sh
git clone https://github.com/stanleyHayes/ghanaessential.git
cd ghanaessential
pnpm install
pnpm check          # full local quality gate: validator, Go, TypeScript, tests, build
```

Run the read-only API locally using the same commands the deployment uses ([`render.yaml`](render.yaml)):

```sh
go build -trimpath -o bin/ghanaessential-api ./cmd/api
./bin/ghanaessential-api                        # PORT defaults to 8080
curl -s http://localhost:8080/health
```

The binary reads [`data/contacts.json`](data/contacts.json) relative to the working directory, so start it from the repository root. `ALLOWED_ORIGIN` defaults to `https://essential.digitalghana.dev`. The web app has no `dev` script yet — `package.json` exposes `typecheck`, `test`, `build` and `check` only, and adding one is a [good first contribution](CONTRIBUTING.md#good-first-contributions).

---

## API

Three read-only endpoints, defined in [`contracts/README.md`](contracts/README.md) and implemented in [`cmd/api/main.go`](cmd/api/main.go).

| Method | Path | Returns |
|---|---|---|
| `GET` | `/health` | Liveness, dataset `version`, record count and `reviewedAt` |
| `GET` | `/v1/contacts` | The complete dataset |
| `GET` | `/v1/offline.json` | The identical payload, served as a download named `ghanaessential-<version>.json` |

```sh
curl -s https://api-essential.digitalghana.dev/v1/contacts
```

The first of the five records, formatted for readability:

```json
{
  "version": "2026.09.01-beta.1",
  "reviewedAt": "2026-09-01",
  "warning": "If life or property is in immediate danger, call 112. Network and agency availability can change; this directory does not dispatch help.",
  "contacts": [
    {
      "id": "national-emergency-112",
      "service": "National Emergency",
      "category": "General emergency",
      "numbers": ["112"],
      "availability": "24/7",
      "status": "VERIFIED",
      "sourceId": "nadmo-contact",
      "sourceUrl": "https://nadmo.gov.gh/contact",
      "checkedAt": "2026-09-01"
    }
  ]
}
```

| Field | Type | Meaning |
|---|---|---|
| `version` | string | Immutable dataset version. Pin this. |
| `reviewedAt` | date | When the dataset as a whole was reviewed. |
| `warning` | string | Safety text that consumers are expected to display. |
| `contacts[].id` | string | Stable opaque identifier. |
| `contacts[].service`, `category` | string | Institution name and service class. |
| `contacts[].numbers` | string[] | Published numbers, primary first. |
| `contacts[].availability` | string | As stated by the source, not as measured. |
| `contacts[].status` | string | `VERIFIED` only; anything else fails the gate. |
| `contacts[].sourceId`, `sourceUrl` | string | Entry in the source register, and the exact page. |
| `contacts[].checkedAt` | date | When that page was last re-opened. |

**CORS.** The API echoes `Access-Control-Allow-Origin` only for an exact match on `ALLOWED_ORIGIN`; other browser origins are denied. Server-side clients and `curl` are unaffected. Responses carry `X-Content-Type-Options: nosniff`.

---

## Data and provenance

Four official institutional sources, recorded in [`docs/governance/source-register.json`](docs/governance/source-register.json). Its expected shape is documented in [`source-register.schema.json`](docs/governance/source-register.schema.json); schema validation is not yet automated — [`scripts/validate.rb`](scripts/validate.rb) only checks that the register is non-empty and that no entry carries a `blocked` or `unknown` publication decision:

| Source ID | Authority | Page | Reviewed |
|---|---|---|---|
| `nadmo-contact` | National Disaster Management Organisation | [nadmo.gov.gh/contact](https://nadmo.gov.gh/contact) | 2026-09-01 |
| `nas-home` | National Ambulance Service | [nas.gov.gh](https://www.nas.gov.gh/) | 2026-09-01 |
| `gnfs-home` | Ghana National Fire Service | [gnfs.gov.gh](https://gnfs.gov.gh/) | 2026-09-01 |
| `police-official` | Ghana Police Service | [police.gov.gh](https://police.gov.gh/en/index.php/publication-officials/) | 2026-09-01 |

- **Licence boundary.** Every source is registered as *"official factual contact information; source rights retained"* with a `metadata-link-only` publication decision. Apache-2.0 covers this repository's own code and configuration; it does **not** relicense the referenced source pages. See [`NOTICE`](NOTICE).
- **Versioning.** The dataset version (`2026.09.01-beta.1`) moves independently of the API path version (`/v1`), and published versions are immutable.
- **Verification.** Each record carries `checkedAt`. [`scripts/validate.rb`](scripts/validate.rb) aborts if `reviewedAt` is more than 30 days old, if the dataset does not contain exactly five contacts, if any record is not `VERIFIED`, or if a source entry is `blocked` or `unknown`.
- **Corrections.** Open an issue or pull request with the affected `id`, the current value, the proposed value, the authoritative source URL and its publication date. Automation may draft; a human reviewer approves publication. See [`CONTRIBUTING.md`](CONTRIBUTING.md).

---

## Project layout

```text
ghanaessential/
├── app/                  # Next.js 16 contact UI, metadata, manifest, robots, sitemap, OG image
├── cmd/api/              # read-only Go API (main.go) and its fixture safety test
├── contracts/            # the published read-only contract surface
├── data/contacts.json    # the canonical versioned fixture — web, API and offline export all read this
├── docs/
│   ├── adr/              # 0001 product boundary and safety non-goals
│   ├── governance/       # source register + JSON Schema
│   └── runbooks/         # operations baseline and dated release/rollback evidence
├── infra/                # Vercel project configuration
├── scripts/validate.rb   # dependency-free source, freshness and safety gate
├── tests/                # Node test runner checks over the fixture
└── .github/workflows/    # quality.yml — runs `pnpm check` on PRs and pushes to main
```

---

## Verification

`pnpm check` is the whole gate, in order:

```sh
ruby scripts/validate.rb   # required files, source register, five VERIFIED records, 30-day freshness, secret scan
go test ./...              # fixture safety invariants
go vet ./...
pnpm typecheck             # tsc --noEmit
pnpm test                  # node --test tests/*.test.mjs
pnpm build                 # next build
```

CI runs exactly this on every pull request and every push to `main` ([`.github/workflows/quality.yml`](.github/workflows/quality.yml)). Dated production and rollback evidence for the current release is in [`docs/runbooks/release-evidence.md`](docs/runbooks/release-evidence.md).

---

## Status and roadmap

**Public beta**, verified 2026-09-01. Web and API are live. The product has not reached `stable`.

[`ROADMAP.md`](ROADMAP.md) sets out what has shipped, the open gates and their blocking dependencies, and what is explicitly out of scope. It is directional, not a commitment; the machine-readable state of record is [`agent_plan.md`](agent_plan.md).

---

## Contributing and policy

- [`CONTRIBUTING.md`](CONTRIBUTING.md) — setup, verification, commit and PR conventions, data corrections, first contributions.
- [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md) — Contributor Covenant v2.1; reports go to the private channel documented there.
- [`SECURITY.md`](SECURITY.md) — vulnerability reporting. Do not open a public issue containing exploit details.
- [`AGENTS.md`](AGENTS.md) — coordination rules for automated and human contributors.
- [`LICENSE`](LICENSE) and [`NOTICE`](NOTICE) — Apache License 2.0, inbound = outbound.

Part of the [Digital Ghana](https://digitalghana.dev) portfolio of source-linked public-interest infrastructure.

---

## Independence

GhanaEssential is an independent open-source project. It is **not operated by, endorsed by, or affiliated with the Government of Ghana**, any ministry, agency or public institution. Records that reference official sources do so by citation only, and nothing here should be treated as an official record, as evidence of official status, or as a channel to any emergency service.

**If life or property is in immediate danger, call 112.**

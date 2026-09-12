# Contributing to GhanaEssential

GhanaEssential welcomes contributions to the code, the documentation, the data provenance and the safety boundaries of this product.

This repository owns one product: a small, source-linked directory of Ghana's published emergency and essential-service contacts, its read-only API, its offline export and its deployment. Portfolio-wide standards, the product registry and cross-product architecture decisions live in [stanleyHayes/digitalghana](https://github.com/stanleyHayes/digitalghana).

> **Before anything else, read the safety boundary.** This product is a directory of already-published contact numbers. It does not dispatch help, does not accept incident reports, and does not collect location, medical or personal data. A change that weakens that boundary — or that could lead a reader to believe the site can summon help — will be rejected regardless of how well it is written. The boundary is stated in [ADR-0001](docs/adr/0001-product-boundary.md) and enforced by [`scripts/validate.rb`](scripts/validate.rb) and [`tests/contact.test.mjs`](tests/contact.test.mjs).

## Your first contribution

1. **Read the ledger.** [`agent_plan.md`](agent_plan.md) is the execution ledger and state of record; [`ROADMAP.md`](ROADMAP.md) is the readable summary. Check whether the work is already claimed or blocked on an external gate.
2. **Read the product boundary.** [ADR-0001](docs/adr/0001-product-boundary.md) and [`contracts/README.md`](contracts/README.md) define what this product will and will not do. Most rejected ideas are rejected there, not in review.
3. **Set up locally** with the sequence below and confirm a clean checkout passes `pnpm check` *before* you change anything. If it does not pass, that is itself a bug worth reporting.
4. **Open an issue first for anything non-trivial** — a data correction, a new field, a contract change, or anything touching the safety wording. Evidence is easier to agree on before code exists.
5. **Keep the change small and evidenced.** One concern per pull request.

### Prerequisites

| Tool | Version | Where it is pinned |
|---|---|---|
| Node.js | 22 | `actions/setup-node` in [`.github/workflows/quality.yml`](.github/workflows/quality.yml) |
| pnpm | 10.17.1 | `packageManager` field in [`package.json`](package.json) |
| Go | 1.23 | [`go.mod`](go.mod) and `actions/setup-go` in CI |
| Ruby | 3.4 | `ruby/setup-ruby` in CI |
| Git | any recent version | — |

### Local setup

```sh
git clone https://github.com/stanleyHayes/ghanaessential.git
cd ghanaessential
pnpm install
pnpm check
```

To run the read-only API locally, use the same commands the deployment uses ([`render.yaml`](render.yaml)):

```sh
go build -trimpath -o bin/ghanaessential-api ./cmd/api
./bin/ghanaessential-api
curl -s http://localhost:8080/health
```

The binary reads [`data/contacts.json`](data/contacts.json) relative to the working directory, so start it from the repository root. `PORT` defaults to `8080`; `ALLOWED_ORIGIN` defaults to `https://essential.digitalghana.dev`.

### Verification

Run this before opening a pull request:

```sh
pnpm check
```

It is the whole gate, in order — the same sequence CI runs:

```sh
ruby scripts/validate.rb   # required files, source register, five VERIFIED records, 30-day freshness, secret scan
go test ./...              # fixture safety invariants
go vet ./...
pnpm typecheck             # tsc --noEmit
pnpm test                  # node --test tests/*.test.mjs
pnpm build                 # next build
```

[`.github/workflows/quality.yml`](.github/workflows/quality.yml) runs `pnpm check` on every pull request and every push to `main`.

`scripts/validate.rb` is dependency-free. It checks that every required file exists, that the source register is non-empty and contains no `blocked` or `unknown` publication decision, that the fixture contains exactly five records, that every record is `VERIFIED`, that `reviewedAt` is no more than 30 days old, and that no unresolved template token or private key has been committed.

> The fixture gate **fails closed**. If `data/contacts.json` has not been re-reviewed within 30 days, `pnpm check` aborts with `contact fixture is stale` and CI goes red. That is intended behaviour, not a flake: it is how this product avoids serving safety-adjacent data it can no longer vouch for.

> Note: the validator reads project files with the default external encoding. On a machine whose default locale is not UTF-8 (an older system Ruby on macOS, for example) it can abort with `invalid byte sequence in US-ASCII`. Run it under Ruby 3.4, or prefix the command with `RUBYOPT="-E UTF-8"`. CI is unaffected.

## Good first contributions

Each of these is a real, currently open gap in this repository or its ledger:

1. **Add a local development script.** [`package.json`](package.json) exposes `typecheck`, `test`, `build` and `check` only. There is no way to run the Next.js app in development without knowing the framework command. Add a `dev` script and document it in the README quickstart.
2. **Add a source-liveness checker.** Every `sourceUrl` in [`data/contacts.json`](data/contacts.json) is re-opened by hand before release, as described in [`docs/runbooks/operations.md`](docs/runbooks/operations.md). A dependency-free script under `scripts/` that re-requests each registered source URL and reports status drift would turn a manual release step into a repeatable check.
3. **Add negative tests for the validator.** [`scripts/validate.rb`](scripts/validate.rb) enforces nine separate abort conditions, and none of them has a test proving it actually fails. Add a companion script that feeds it a stale `reviewedAt`, a non-`VERIFIED` record, a sixth contact and a `blocked` source decision, and asserts each one aborts.
4. **Assert the safety copy in a test.** [`tests/contact.test.mjs`](tests/contact.test.mjs) proves the data is safe but nothing proves the *presentation* is. Add a case asserting the dataset `warning` is non-empty, mentions `112`, and is rendered by [`app/page.tsx`](app/page.tsx).
5. **Document a `checkedAt`-aware freshness display.** The website shows the review date, but there is no published guidance for downstream consumers of the offline export on how old is too old. Propose wording for `contracts/README.md` that states the 30-day rule as a consumer-facing expectation.
6. **Propose a coverage increase with sources.** The beta is pinned at exactly five national-level contacts. A well-evidenced proposal for the next records — with the official source URL, the publication date and a recorded licence decision for each — is more valuable than code. See "Data corrections" below; the same evidence bar applies.

## Change expectations

- **Never weaken the safety boundary.** No dispatch claim, no incident intake, no location, medical or personal data, no implication that a published number guarantees a connection or a response time.
- **Never describe "source checked" as proof of availability.** The only claim this product makes is that a named official source published a number on a stated date.
- Keep lifecycle state honest. The permitted words are `proposed`, `building`, `beta`, `stable`, `externally blocked`, `retired` and `deferred`. A passing build is not evidence of deployment.
- A status transition updates [`agent_plan.md`](agent_plan.md), [`ROADMAP.md`](ROADMAP.md) and the evidence under [`docs/runbooks/`](docs/runbooks/) in the same reviewed change.
- Preserve stable identifiers. A contact `id` is a public contract; changing one is a breaking change and needs a migration note.
- The dataset version is immutable once published. A data change means a new version, not an edit in place.
- Include tests proportional to risk, and update public documentation in the same pull request.
- Never commit credentials, private exports, personal data or provider environment files.
- Do not imply government endorsement or official status anywhere — in code, data, copy or metadata.

## Commit and pull-request conventions

Commits follow the existing history in this repository: a lowercase `type: imperative summary` subject, no scope, no trailing full stop, short enough to read in `git log --oneline`.

Types in use here: `feat`, `docs`, `chore`.

```text
chore: establish GhanaEssential foundation
feat: build safety-gated GhanaEssential beta
docs: record GhanaEssential public beta evidence
```

A pull request should state:

- the scope of the change and which surface it affects (web, API, fixture, docs);
- the source, ADR or evidence file it relies on;
- the verification actually performed — paste the commands you ran;
- migration and rollback impact, including whether any contact `id` or the dataset version changes;
- whether the change touches the safety boundary in any way, and if so, why it is still safe.

## Data corrections

Corrections need evidence, not confidence. A published emergency number being wrong is the highest-severity defect this project can have, so the bar is deliberately high. Provide:

- the affected contact `id`;
- the current published value;
- the proposed value;
- the authoritative source, with a stable `https://` URL — an official institutional page, not a news article or a screenshot;
- the source publication or effective date;
- whether the correction changes the dataset version, and whether any historical record is affected.

Rules that apply to every correction:

- Automation may draft a correction; a **human reviewer approves publication**.
- A record whose source has disappeared or changed must be **downgraded before release**, never silently rewritten to whatever seems right.
- Unknown licence status is recorded as `unknown` and blocks publication; it is never assumed open.
- New sources are added to [`docs/governance/source-register.json`](docs/governance/source-register.json) with an authority, title, stable URL, licence note, publication decision and review date, and must match the shape documented in [`source-register.schema.json`](docs/governance/source-register.schema.json) — this is reviewed by hand, not by CI.

The full release procedure, including the pre-release source re-check and the rollback path, is in [`docs/runbooks/operations.md`](docs/runbooks/operations.md).

## Review expectations

- A maintainer reviews every pull request; expect questions about provenance, safety framing and evidence before questions about style.
- Changes that alter published state (a number, a dataset version, a lifecycle word, a hostname) need the evidence in the diff, not in the conversation.
- CI must be green. A red validator is treated as a blocking defect, not a flake.
- Discussion stays on evidence and public benefit. Conduct expectations are in [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md).
- Maintainers may edit, reject or remove contributions that violate the policies above.
- Security vulnerabilities do not go in pull requests or public issues — follow [`SECURITY.md`](SECURITY.md).

## Licensing

Unless explicitly stated otherwise, contributions intentionally submitted for inclusion are provided under the Apache License 2.0 under that licence's inbound = outbound terms (section 5). See [`LICENSE`](LICENSE) and [`NOTICE`](NOTICE).

Contributions must not include third-party data or documents without recorded permission. The Apache-2.0 repository licence does not relicense the referenced official source pages or any imported dataset; those retain their own rights.

## Independence

GhanaEssential is an independent open-source project. It is not operated by, endorsed by, or affiliated with the Government of Ghana or any of its agencies. Do not submit changes that imply otherwise.

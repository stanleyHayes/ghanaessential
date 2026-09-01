# Release evidence — public beta — 2026-09-01

- Source commit: `5e4e68c4b1b652f961ed04069557ef023ef533d4`; fixture `2026.09.01-beta.1`.
- GitHub Quality run `33526951922`: success.
- Vercel project `ghanaessential`; restored deployment `dpl_EabK9Fwxrqzzx9n18ciu6xfkxEzh`; canonical `https://essential.digitalghana.dev`.
- Render service `srv-dabf3dfqj5pc739q2d0g`; restored deploy `dep-dabf4r610ojc73a4pcc0`; custom domain `cdm-dabf3fvavr4c73fvqda0`; DNS `rec_3bff9c0939a46d97aea095d7`; canonical `https://api-essential.digitalghana.dev`.

## Verification

- Source/freshness gate, Go tests/vet, fixture safety tests, TypeScript and production build passed.
- Five records are source-complete, checked 2026-09-01 and contain no personal, location, medical or incident data.
- Canonical web/API TLS, health, complete contact response, identical offline download, allowed-origin and denied-origin CORS passed.
- Browser proved visible disclaimer, ten `tel:` links, official source links, zero prohibited native controls, zero horizontal overflow, Outfit/Geist Mono/Newsreader, canonical/JSON-LD/OG/Twitter metadata.
- Favicon/manifest/robots/sitemap returned 200; OG PNG is exactly 1200x630.

## Rollback proof

- Vercel restored `dpl_EgmppPsaZuoqYZFz67ZtRcy4FExj`, canonical smoke passed, then restored intended `dpl_EabK9Fwxrqzzx9n18ciu6xfkxEzh`, canonical smoke passed.
- Render restored `dep-dabf3efqj5pc739q2g50`, producing `dep-dabf4jvavr4c73fvt6ig`; health passed after propagation. It restored intended `dep-dabf4au7bikc73an76kg`, producing `dep-dabf4r610ojc73a4pcc0`; health passed.

Known limit: official publication does not guarantee network availability, pickup or response time. Records fail closed after 30 days without re-review.

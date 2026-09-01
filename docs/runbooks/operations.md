# Operations baseline

- Web: `https://essential.digitalghana.dev` on Vercel project `ghanaessential`.
- API: `https://api-essential.digitalghana.dev` on Render service `srv-dabf3dfqj5pc739q2d0g`.
- Health: `GET /health`; offline export: `GET /v1/offline.json`.
- State is an immutable fixture; there is no writable store or user data to back up.
- Owner: Digital Ghana maintainers through the public repository.

Before release, run `pnpm check`, require green CI, re-open every official source, and reject any fixture over 30 days old. A missing/changed source must downgrade its record before publication. Smoke canonical web, API, CORS, tel links, offline export, metadata and OG image. Use Vercel rollback or Render `POST /v1/services/{serviceId}/rollback`, then restore and recheck the intended release.

Never accept incident reports, user location or medical data. Never describe “source checked” as proof of telephone availability or response time.

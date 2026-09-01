# ADR-0001: Independent product boundary

Status: Accepted

## Decision

GhanaEssential owns its repository, release lifecycle, data, credentials, contracts and operational evidence. It may consume another Digital Ghana product only through a versioned public contract or pinned dataset artifact.

The canonical web hostname is `essential.digitalghana.dev` and the read-only API is `api-essential.digitalghana.dev` as reserved in the portfolio registry.

The product is a small directory of emergency and essential contacts sourced from official institutional pages. It is not an emergency dispatch system, availability guarantee, government service, medical adviser or replacement for calling 112. Every record carries its source, checked date and verification state. A record older than 30 days or with a failed source check must be downgraded before release.

The offline export is the same versioned fixture. No user location, medical details or incident reports are collected.

## Consequences

Failures remain isolated, histories remain understandable, and a portfolio-wide platform outage is not created by convenience. Some configuration and small primitives may be repeated until two proven consumers justify a versioned shared package.

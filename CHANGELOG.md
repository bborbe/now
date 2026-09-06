# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

- fix: `make build` refuses to stamp a version onto a tree that is not that version's tag (`check-version-tag`, escape hatch `ALLOW_UNTAGGED_BUILD=1`). `VERSION` defaults to the newest tag repo-wide, so an operator-run build from an untagged or older tree silently republishes under the newest tag. The guard compares `git describe --exact-match HEAD` against `$(VERSION)` and exits non-zero on mismatch.

## 1.3.0

- Use timezones information from alpine
- Refactoring

## 1.2.1

- Remove duplicate

## 1.2.0

- Use global list of locations for now

## 1.1.1

- Fix Content-Type

## 1.1.0

- Use fixed timezones

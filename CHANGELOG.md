# Changelog

## [1.0.2](https://github.com/liblaf/ddns/compare/v1.0.1...v1.0.2) (2025-09-14)


### 🐛 Bug Fixes

* **deps:** update module github.com/cloudflare/cloudflare-go/v6 to v6.0.1 ([#30](https://github.com/liblaf/ddns/issues/30)) ([7776a2e](https://github.com/liblaf/ddns/commit/7776a2ed78cb7022f65d0aa7ea0903705f54a4ae))
* **deps:** update module github.com/spf13/viper to v1.21.0 ([#28](https://github.com/liblaf/ddns/issues/28)) ([0761565](https://github.com/liblaf/ddns/commit/076156586c322037cfdd37bf9a5bad93bcedcc35))

## [1.0.1](https://github.com/liblaf/ddns/compare/v1.0.0..v1.0.1) - 2025-09-04

### ⬆️ Dependencies

- **deps:** update module github.com/spf13/cobra to v1.10.1 (#26) - ([93fb5f0](https://github.com/liblaf/ddns/commit/93fb5f0fbac67f449acd6c9413c7a8cb999b1e3f))
- **deps:** update module github.com/cloudflare/cloudflare-go/v5 to v6 (#23) - ([6f8ff17](https://github.com/liblaf/ddns/commit/6f8ff176c5cf1cf2a657fe9fe018af1d728a8189))

### ❤️ New Contributors

- [@renovate[bot]](https://github.com/apps/renovate) made their first contribution in [#27](https://github.com/liblaf/ddns/pull/27)
- [@liblaf-bot[bot]](https://github.com/apps/liblaf-bot) made their first contribution

## [1.0.0](https://github.com/liblaf/ddns/compare/v0.0.3..v1.0.0) - 2025-08-23

### 💥 BREAKING CHANGES

- upgrade cloudflare-go from v4 to v5 - ([3d6127c](https://github.com/liblaf/ddns/commit/3d6127c7af8792e07f0ff6e05083f42cfae0bc1c))

### ⬆️ Dependencies

- **deps:** update module github.com/cloudflare/cloudflare-go/v4 to v5 (#16) - ([4c6b544](https://github.com/liblaf/ddns/commit/4c6b5449a475c3cd37788ed004d5b1bdda8d8913))

### 👷 Build System

- Set up build system and update dependencies - ([98645c5](https://github.com/liblaf/ddns/commit/98645c541014e854cae25eecae09d695a36dbf97))

### ❤️ New Contributors

- [@liblaf-bot[bot]](https://github.com/apps/liblaf-bot) made their first contribution in [#20](https://github.com/liblaf/ddns/pull/20)

## [0.0.3](https://github.com/liblaf/ddns/compare/v0.0.2..v0.0.3) - 2025-08-10

### ⬆️ Dependencies

- **deps:** update module github.com/paulsonoflars/gotgbot/v2 to v2.0.0-rc.33 (#11) - ([4213b92](https://github.com/liblaf/ddns/commit/4213b928a1a02bb1b0c6968eaf0094c96775f723))
- **deps:** update module github.com/samber/oops to v1.19.0 (#13) - ([6efe96c](https://github.com/liblaf/ddns/commit/6efe96c58330507b7ba4290b9c5215e4cdf33e83))
- **deps:** update module github.com/spf13/viper to v1.20.1 (#10) - ([2304f84](https://github.com/liblaf/ddns/commit/2304f84aa83728419092ab4c9ed9c209a4e89833))

### ❤️ New Contributors

- [@mergery[bot]](https://github.com/apps/mergery) made their first contribution

## [0.0.2](https://github.com/liblaf/ddns/compare/v0.0.1..v0.0.2) - 2025-03-30

### ⬆️ Dependencies

- **deps:** update module github.com/rs/zerolog to v1.34.0 (#8) - ([3bb0d23](https://github.com/liblaf/ddns/commit/3bb0d235e358895696628b3b018153659e734b70))

## [0.0.1](https://github.com/liblaf/ddns/compare/v0.0.0..v0.0.1) - 2025-03-23

### ⬆️ Dependencies

- **deps:** update module github.com/cloudflare/cloudflare-go/v4 to v4.2.0 (#7) - ([2225362](https://github.com/liblaf/ddns/commit/2225362097497cc846b360e2a70f51a71e8df48e))
- **deps:** update module github.com/spf13/viper to v1.20.0 (#4) - ([df244b8](https://github.com/liblaf/ddns/commit/df244b8cecdebe664c5164e0718caf023a936268))

### ❤️ New Contributors

- [@renovate[bot]](https://github.com/apps/renovate) made their first contribution in [#7](https://github.com/liblaf/ddns/pull/7)

## [0.0.0] - 2025-03-09

### ✨ Features

- **cmd:** enhance configuration and environment variable handling - ([bd6e558](https://github.com/liblaf/ddns/commit/bd6e55881b5bcaa46d1bac87261a78deb031d74a))
- enhance logging and notification handling - ([d02bc3b](https://github.com/liblaf/ddns/commit/d02bc3be8a23c5f4b38ac528d4781fd398cffadb))
- add Cloudflare DNS management and Telegram notifications - ([4497175](https://github.com/liblaf/ddns/commit/44971751253a12bdbd0e751038afa195b84c9384))

### 🐛 Bug Fixes

- **notify:** trim trailing whitespace in notification messages - ([83daebe](https://github.com/liblaf/ddns/commit/83daebee02d010a0c7ec0fef1b68f2b3d7d1ee40))

### ♻ Code Refactoring

- **notify:** move GetLabel function to cloudflare package - ([ff335c4](https://github.com/liblaf/ddns/commit/ff335c4f962a5e8bb85f5b091b99a371bf1c9404))
- **notify:** centralize record formatting logic - ([f0cf245](https://github.com/liblaf/ddns/commit/f0cf24598e872896783ec842d2e5726f901f16bf))

### 🔧 Continuous Integration

- update artifact paths from ./dist/ to ./bin/ - ([3f15ab6](https://github.com/liblaf/ddns/commit/3f15ab6e0dc62a4b81ad0ab97e566f0fbdcde1ba))
- rename workflows and update PAT secret reference - ([9b99236](https://github.com/liblaf/ddns/commit/9b99236042656dfdefbb679880518a45f46aeceb))
- add release workflow and update Justfile for cross-platform builds - ([e59793c](https://github.com/liblaf/ddns/commit/e59793c61356620057d4ff5ba666753642fcf855))

### ❤️ New Contributors

- [@github-actions[bot]](https://github.com/apps/github-actions) made their first contribution in [#2](https://github.com/liblaf/ddns/pull/2)
- [@liblaf](https://github.com/liblaf) made their first contribution

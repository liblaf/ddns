# Changelog

## [0.0.2](https://github.com/liblaf/ddns/compare/v0.0.1...v0.0.2) (2025-03-23)


### ⬆️ Dependencies

* **deps:** update module github.com/rs/zerolog to v1.34.0 ([#8](https://github.com/liblaf/ddns/issues/8)) ([3bb0d23](https://github.com/liblaf/ddns/commit/3bb0d235e358895696628b3b018153659e734b70))

## [0.0.1](https://github.com/liblaf/ddns/compare/v0.0.0..v0.0.1) - 2025-03-20

### ⬆️ Dependencies

- **deps:** update module github.com/cloudflare/cloudflare-go/v4 to v4.2.0 (#7) - ([2225362](https://github.com/liblaf/ddns/commit/2225362097497cc846b360e2a70f51a71e8df48e))
- **deps:** update module github.com/spf13/viper to v1.20.0 (#4) - ([df244b8](https://github.com/liblaf/ddns/commit/df244b8cecdebe664c5164e0718caf023a936268))

### ❤️ New Contributors

- @renovate[bot] made their first contribution in [#7](https://github.com/liblaf/ddns/pull/7)

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

- @github-actions[bot] made their first contribution in [#2](https://github.com/liblaf/ddns/pull/2)
- @liblaf made their first contribution

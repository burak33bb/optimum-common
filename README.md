# optimum-common
This library will serve as the **shared SDK** for Optimum projects, consolidating configuration, logging, claims parsing, rate-limiting, and other cross-cutting utilities into one place.


## Optimum Common — TODO

### **Initial Scope**

- **Auth**
    - [x] Centralize JWT claims model (`Claims`)
    - [x] Parsing helpers (`ParseUnverified`, JWKS verifier)
    - [ ] Remove duplicate parsing logic from CLI and Proxy
    - [ ] Unit tests for claim parsing & coercion
 

- **Rate Limiting**
    - [ ] Pure functions for per-sec, per-hour, daily, and message-size checks
    - [ ] `UsageData` type with pluggable storage (in-memory, file-backed)
    - [ ] Shared test suite to ensure consistent enforcement

- **API Models**
    - [ ] Define shared JSON request/response structs (e.g., `PublishRequest`, `SubscribeRequest`)
    - [ ] Replace per-repo definitions in CLI and Proxy

- **Versioning**
    - [ ] Header constant for `X-CLI-Version`
    - [ ] Helper to add version header from ldflags (`Version`, `CommitHash`)
    - [ ] Middleware-agnostic validator for version checks

- **HTTP Client Helpers**
    - [ ] JSON POST/GET helpers with auth header + version injection
    - [ ] Configurable timeouts and retry logic

- **Config Loader**
    - [ ] Support for flags, environment variables, and YAML config files
    - [ ] Override priority: flags > env > YAML
    - [ ] Example usage in CLI and Proxy

- **Logging**
    - [ ] Standard `AppLogger` interface with structured fields
    - [ ] Pluggable backends (Zap, Zerolog, etc.)
    - [ ] Consistent log format across services

- **Utilities**
    - [ ] IP detection (`publicIP`, `outboundIP`, private range detection)
    - [ ] Hashing helpers (XXHash, SHA-256, message ID generation)
    - [ ] TTL map, RW map/slice, broadcaster patterns
    - [ ] File helpers (atomic write, safe temp files)

- **P2P**
    - [ ] Identity key persistence/load for libp2p components
    - [ ] Shared multiaddr builders

- **Telemetry**
    - [ ] Prometheus registry creation helper
    - [ ] HTTP `/metrics` mux setup

---

## **Integration Plan**

1. **Build & publish** `optimum-common` as a standalone Go module.
2. **Replace** duplicated code in:
    - CLI (`mump2p-cli/internal/auth`, rate-limit logic, utils)
    - Proxy (`pkg/proxy/middleware/auth`, rate-limit checks, utils)
    - Gateway (`internal/utils`, IP helpers, broadcaster)
3. **Enforce imports** in all repos so shared logic lives only in `optimum-common`.
4. **Add CI tests** in `optimum-common` to cover all packages — run in downstream repos via `go test ./...` after vendoring.
5. **Document** example usage and integration steps in `/docs`.



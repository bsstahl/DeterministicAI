# Pete's Pottery Paradise

This demo finds the revenue-maximizing production plan for Pete's Pottery
Paradise. Implementations are grouped by language.

## Problem

- Small vase: 1 oz. of clay, 1 oz. of glaze, $3 revenue.
- Large vase: 4 oz. of clay, 2 oz. of glaze, $9 revenue.
- Inventory: 24 oz. of clay and 16 oz. of glaze.
- Production quantities must be non-negative integers.

The model maximizes `3 * small + 9 * large` subject to:

```text
small + 4 * large <= 24   (clay)
small + 2 * large <= 16   (glaze)
```

The optimum is 8 small vases and 4 large vases, for $60 revenue.

## Go implementation

The Go module and its Docker-based development environment are in [`go/`](go/).
Open that folder in VS Code and choose **Reopen in Container** to use the
provided development container. From its terminal:

```sh
go run ./cmd/petes-pottery-paradise
go test ./...
```

The container checks out the matching OR-Tools Go sources, generates its
protobuf bindings, and installs the C++ runtime for the official Go CP-SAT
binding. A container-only Go workspace uses that local checkout because the
upstream OR-Tools source module is too large to download through the Go module
proxy. No Go or OR-Tools installation is required on the host.

## JavaScript implementation

The [`js/`](js/) version uses the community-maintained `or-tools-wasm` package
to run the CP-SAT model in WebAssembly. Google does not provide or support this
JavaScript binding; see the [JavaScript demo README](js/README.md) for the
dependency caveat and setup instructions. Node.js, npm, and packages are
contained in its Dev Container.

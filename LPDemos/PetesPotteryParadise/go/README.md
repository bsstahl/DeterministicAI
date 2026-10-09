# Pete's Pottery Paradise (Go)

This Go implementation uses Google OR-Tools CP-SAT to choose an integer-valued production plan that maximizes revenue. Go and OR-Tools do not need to be installed on the host; use the provided Docker development environment or open the project in its VS Code Dev Container.

## Set up the development environment

From the repository root in PowerShell (the directory containing
`LPDemos`), build the development image once:

```powershell
docker build -t petes-pottery-paradise-go-dev -f LPDemos/PetesPotteryParadise/go/.devcontainer/Dockerfile .
```

Alternatively, open `LPDemos/PetesPotteryParadise/go` as the VS Code workspace
and choose **Reopen in Container**. VS Code builds the same development image.

The Dev Container explicitly mounts this Go project at `/workspace`, which
matches the Go workspace configured in the image.

If reopening fails with an error mentioning `distro-services/ubuntu.sock`
while mounting `wayland-0`, disable **Dev Containers: Mount Wayland Socket**
in VS Code's **User** settings (`"dev.containers.mountWaylandSocket": false`)
and retry. This command-line application does not need WSLg GUI forwarding.
The setting applies to all Dev Containers; projects that need GUI forwarding
require working WSLg and Docker Desktop WSL integration before re-enabling it.

## Quality checks and build

The project has a single quality gate. It checks `gofmt` formatting, runs
`go vet`, Staticcheck, and the race-enabled test suite. The build fails if any
check fails:

```powershell
docker run --rm -v "${PWD}\LPDemos\PetesPotteryParadise\go:/workspace" petes-pottery-paradise-go-dev sh -c "make build"
```

`make check` runs the checks without compiling a build. `make build` is the
default VS Code build task and runs the checks before building. `make run`
also runs the checks before starting the optimizer. In VS Code, use **Run
Build Task** or **Run Task** to invoke these targets. The Go extension runs
Staticcheck and `go vet` on save to surface diagnostics while editing.

In the Dev Container terminal, the same commands are simply:

```sh
make check
make build
make run
```

Staticcheck 2025.1.1 is installed at a pinned version in the Dev Container
image. Run the quality gate there for consistent analyzer results.

For debugging, you can run a focused test directly, for example:

```sh
go test -v -count=1 ./internal/pottery
```

Direct `go test`, `go build`, and `go run` commands do not run the full quality gate. Use `make check`, `make build`, or `make run` for the normal workflow.

## Run the optimizer

From the repository root in PowerShell, run the optimizer through the quality gate:

```powershell
docker run --rm -v "${PWD}\LPDemos\PetesPotteryParadise\go:/workspace" petes-pottery-paradise-go-dev sh -c "make run"
```

With the default inputs, the program prints a plan of 8 small vases and
4 large vases, using 24 oz. of clay and 16 oz. of glaze for $60 revenue.

## Try changing the inputs

Clay and glaze inventories are passed to `pottery.Solve` in ounces. The CLI
currently calls it with 24 oz. of clay and 16 oz. of glaze. Vase resource use
and revenue values remain constants at the top of
`internal/pottery/production.go`. For example, change the small vase revenue
from `$3` to `$2`:

```go
smallVaseRevenue = int64(2)
```

Then run the optimizer again. It should choose 6 large vases and no small
vases, using 24 oz. of clay and 12 oz. of glaze for $54 revenue. The changed
revenue constant is used both in the objective and the reported revenue, so
the displayed result reflects the new input.

Run the tests after changing parameters as well. The existing regression test
expects the original `$3` value and will fail until you restore it or update
the expected plan to match the modified problem.

# Pete's Pottery Paradise (JavaScript)

This implementation models the integer production problem with Google OR-Tools
CP-SAT through the community-maintained `or-tools-wasm` package. It uses the
same vase values, inventory constraints, and default inputs as the other language versions.

## Important dependency caveat

Google does not publish or support a JavaScript OR-Tools binding. This demo
uses [`or-tools-wasm`](https://github.com/Axelwickm/or-tools-wasm), an
independent WebAssembly port exposing CP-SAT to JavaScript. Its API and solver
package are not covered by Google's OR-Tools support. The dependency is pinned
to an exact version in `package.json` and `package-lock.json`; review updates
and test them before upgrading.

The demo runs the solver locally through WebAssembly. It disables the
package's informational cloud notice and does not configure cloud execution.

## Run

Open this `js/` folder in VS Code and choose **Reopen in Container**. Node.js,
npm, and npm dependencies are installed and used only inside the container;
the `node_modules` directory is stored in a Docker volume, not the host folder.
From the container terminal:

```sh
npm start
npm test
```

The default run prints a plan of 8 small vases and 4 large vases, using 24 oz.
of clay and 16 oz. of glaze for $60 revenue.

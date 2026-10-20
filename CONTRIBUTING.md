# Contributing to openscalar

Thanks for your interest. openscalar is small and early, so the process is
deliberately light.

## How to contribute

- **Bugs and small fixes**: open a pull request directly.
- **New features or design changes**: open an issue first so we can agree on
  the approach before you spend time on code. The [design notes](#design-notes)
  below explain the constraints you are working within.
- **Questions**: open a discussion or an issue.
- **Security issues**: see [SECURITY.md](SECURITY.md). Do not open a public
  issue.

AI-assisted contributions are welcome. You are responsible for understanding
what the code does and for the validation you report in the PR.

## Development setup

Requires [mise](https://mise.jdx.dev/).

```sh
mise install        # installs the pinned Go toolchain and dev tools
lefthook install    # installs git hooks (gofmt/vet on commit, tests on push)
go build ./...
go test -race ./...
```

CI runs the same checks (`go vet`, `gofmt`, `staticcheck`, `go test -race`)
on Ubuntu and macOS. Windows is not yet verified; reports and fixes are
welcome.

## Commit messages

We use [Conventional Commits](https://www.conventionalcommits.org/). Release
notes are generated from commit messages, so the format matters.

```
<type>(<scope>): <summary>

<optional body: why, not what>
```

- `type` is one of `feat`, `fix`, `refactor`, `test`, `docs`, `ci`, `chore`.
- `scope` is the package or area: `codex`, `tui`, `goal`, `deps`, and so on.
  Omit it when the change is repository-wide.
- `summary` is imperative, lower-case, no trailing period, 50 characters or
  fewer.
- Use the body to explain the reason for the change when it is not obvious
  from the diff. Wrap at 72 characters.
- Add `!` after the type or scope for a breaking change, and describe it in
  the body.

Examples:

```
feat(codex): parse token_count events into Session
fix(tui): keep selection when rows are refreshed
refactor(codex): split Parse into per-event handlers
test(codex): add testdata for compacted sessions
docs: explain context ratio in README
ci: run staticcheck in lint job
chore(deps): update bubbletea to v1.4.0
```

`feat` and `fix` commits appear in release notes; `chore(deps)`, `docs` and
`test` do not. If a PR is squash-merged, make sure the resulting commit
message follows this format.

## Pull requests

- Keep each PR to one change. Unrelated fixes belong in separate PRs.
- Lead the description with the problem being solved and the user-visible
  effect, then how you validated it (test output, a screenshot of the TUI, a
  sample session file).
- Add or update tests for behaviour you change. Table-driven tests with
  fixtures under `testdata/` are the convention here.
- Make sure CI is green before requesting review.

## Design notes

These are the rules the codebase follows. If a change needs to break one,
say so in the issue or PR and explain why.

- `cmd/openscalar` only handles arguments and wiring. All logic lives under
  `internal/`.
- `internal/agent/<provider>` (currently `codex`) parses that agent's session
  logs. It never imports `internal/tui`.
- `internal/tui` renders data; it does not read files or compute metrics.
- No `common` or `util` packages. Name packages after what they contain.
- Interfaces are introduced when the second implementation arrives, not
  before.
- Parsers take an `io.Reader` so they work on files and on test strings
  alike.
- Errors are wrapped with context using `%w`. Do not `panic`.
- Exported identifiers have doc comments that start with the identifier's
  name.
- `gofmt`, `go vet` and `staticcheck` must be clean.

## Code of conduct

Be kind and assume good intent. Maintainers may close contributions that are
disrespectful or off-topic.

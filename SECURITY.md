# Security Policy

## Reporting a Vulnerability

Please do not report security issues in public GitHub issues or pull requests.

Use GitHub's private vulnerability reporting instead:
https://github.com/OkkyYukiya/openscalar/security/advisories/new

Include, where possible:

- A description of the issue and its impact
- Steps to reproduce, or a minimal proof of concept
- The affected version or commit

We will acknowledge your report within 7 days and keep you informed as we
work on a fix. Once a fix is released, we will credit you in the advisory
unless you prefer to remain anonymous.

openscalar is a volunteer-maintained open source project. There is no bug
bounty program at this time.

## Supported Versions

Only the latest release receives security fixes.

## Scope and Trust Model

openscalar reads session logs written by local AI coding agents
(currently Codex; Claude Code is planned) in order to measure context
usage and detect degradation.

What it does:

- Reads files under the agent's session directory (for example
  `~/.codex/sessions/`) with read-only access
- Renders the results in a terminal UI, or prints them as text or JSON

What it does not do:

- Write to, modify, or delete session files
- Send any data over the network
- Execute agent commands or act on an agent's behalf

Session logs contain the full content of your conversations with the
agent, which may include secrets or proprietary code. openscalar treats
these files as sensitive local data and never transmits them. If a future
feature changes any of the guarantees above (for example, launching agents
or cleaning up old sessions), it will be opt-in and documented here.

## Out of Scope

The following are generally not considered vulnerabilities in openscalar:

- Issues in the agents themselves (Codex, Claude Code) or in their log
  formats. Please report those upstream.
- Reports consisting only of automated scanner output without a
  demonstrated impact
- Attacks that require an already-compromised local user account, since
  openscalar runs with the permissions of the user who invokes it

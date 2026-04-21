# gh-triage

**This is a fork of [k1LoW/gh-triage](https://github.com/k1LoW/gh-triage).**

### Changes from upstream

- **Always fetch all notifications (`all=true`)**: The upstream version only fetches unread notifications by default, which caused `gh triage` to silently process zero notifications when the GitHub Web UI Inbox showed threads that were already read. This fork sets `All: true` in `ListNotifications` so all inbox notifications (read + unread) are fetched — matching the behaviour of the GitHub Web UI Inbox(All).
- **`unread` field reflects actual state**: The upstream hardcodes `unread = true` for every notification. This fork replaces that with `n.GetUnread()` so the `unread` field accurately reflects whether a notification is unread. This makes conditions like `done: "!unread"` work as expected.

`gh-triage` is a tool that helps you manage and triage GitHub issues, pull requests, and discussions through notifications. It fetches all notifications from your Inbox (including already-read ones, equivalent to `all=true` in the GitHub API), so conditions like `done: "!unread"` can match read notifications as well.


## Install this fork (m-mahiro/gh-triage)

```bash
# Clone the repository and install from source
git clone https://github.com/m-mahiro/gh-triage.git
cd gh-triage
gh extension install .
```

If you already have `k1LoW/gh-triage` (or another version) installed, remove it first:

```bash
gh extension remove triage
git clone https://github.com/m-mahiro/gh-triage.git
cd gh-triage
gh extension install .
```

## Build locally with Go

You can build the extension entrypoint yourself. Build the binary in the repository root with the name `gh-triage` (or `gh-triage.exe` on Windows).

### macOS / Linux

```bash
git clone https://github.com/m-mahiro/gh-triage.git
cd gh-triage
go build -o gh-triage .
gh extension remove triage 2>/dev/null || true
gh extension install .
gh triage --help
```

### Windows (PowerShell)

```powershell
git clone https://github.com/m-mahiro/gh-triage.git
cd gh-triage
go build -o gh-triage.exe .
gh extension remove triage
gh extension install .
gh triage --help
```

> [!NOTE]
> On Windows, the executable must be `gh-triage.exe`.

## FAQ: `...gh-triage: Is a directory`

If `gh triage` fails with an error like:

```text
--: line 1: .../gh-triage/gh-triage: Is a directory
```

it usually means a directory named `gh-triage` exists where the executable should be.  
The extension entrypoint must be a file named `gh-triage` (or `gh-triage.exe` on Windows), not a directory.

- Remove or rename the conflicting `gh-triage` directory.
- Rebuild the binary in the repository root.
- Reinstall with `gh extension install .`.

## If you do not have Go installed

Prebuilt binaries for Windows/macOS/Linux are attached to the [Releases](https://github.com/m-mahiro/gh-triage/releases) page as Assets.  
Download the binary for your platform and use it instead of building with Go locally.

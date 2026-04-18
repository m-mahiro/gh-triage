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

## Version update notification

When you run `gh triage`, it checks the latest release in `m-mahiro/gh-triage`.
If a newer version is available, it prints an upgrade notice:

```bash
gh extension upgrade triage
```

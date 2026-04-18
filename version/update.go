package version

import (
	"strings"
	"golang.org/x/mod/semver"
)

const (
	ReleaseOwner = "m-mahiro"
	ReleaseRepo  = "gh-triage"
)

func ShouldNotifyUpdate(current, latest string) bool {
	cv, ok := normalizeSemver(current)
	if !ok {
		return false
	}
	lv, ok := normalizeSemver(latest)
	if !ok {
		return false
	}
	return semver.Compare(lv, cv) > 0
}

func normalizeSemver(v string) (string, bool) {
	v = strings.TrimSpace(v)
	if v == "" {
		return "", false
	}
	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}
	if !semver.IsValid(v) {
		return "", false
	}
	return v, true
}

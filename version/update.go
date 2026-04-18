package version

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	ReleaseOwner = "m-mahiro"
	ReleaseRepo  = "gh-triage"
)

var versionPattern = regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)(?:-([0-9A-Za-z.-]+))?(?:\+[0-9A-Za-z.-]+)?$`)

type parsedVersion struct {
	major int
	minor int
	patch int
	pre   string
}

func ShouldNotifyUpdate(current, latest string) bool {
	cv, ok := parseVersion(current)
	if !ok {
		return false
	}
	lv, ok := parseVersion(latest)
	if !ok {
		return false
	}
	return compareVersion(lv, cv) > 0
}

func parseVersion(v string) (parsedVersion, bool) {
	matches := versionPattern.FindStringSubmatch(strings.TrimSpace(v))
	if matches == nil {
		return parsedVersion{}, false
	}
	major, err := strconv.Atoi(matches[1])
	if err != nil {
		return parsedVersion{}, false
	}
	minor, err := strconv.Atoi(matches[2])
	if err != nil {
		return parsedVersion{}, false
	}
	patch, err := strconv.Atoi(matches[3])
	if err != nil {
		return parsedVersion{}, false
	}
	return parsedVersion{
		major: major,
		minor: minor,
		patch: patch,
		pre:   matches[4],
	}, true
}

func compareVersion(a, b parsedVersion) int {
	if a.major != b.major {
		if a.major > b.major {
			return 1
		}
		return -1
	}
	if a.minor != b.minor {
		if a.minor > b.minor {
			return 1
		}
		return -1
	}
	if a.patch != b.patch {
		if a.patch > b.patch {
			return 1
		}
		return -1
	}
	if a.pre == b.pre {
		return 0
	}
	if a.pre == "" {
		return 1
	}
	if b.pre == "" {
		return -1
	}
	if a.pre > b.pre {
		return 1
	}
	return -1
}

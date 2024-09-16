package utils

import "regexp"

func IsValidVersion(version string) bool {
	// Regular expression for version validation: "v*.*.*"
	regex := `^v\d+\.\d+\.\d+(-[a-zA-Z0-9._]+)?$`
	matched, _ := regexp.MatchString(regex, version)
	return matched
}

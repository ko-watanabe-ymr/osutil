//go:build openbsd && dragonfly && netbsd
// +build openbsd,dragonfly,netbsd

package osutil

func GetVersion() string {
	return ""
}

func GetDisplay() string {
	return Name
}

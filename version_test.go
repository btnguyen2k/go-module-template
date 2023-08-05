package mymodule

import (
	"bufio"
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/btnguyen2k/consu/semver"
)

func extractGoVersion(input string) *semver.Semver {
	re := regexp.MustCompile(`^go\s*(\d+(\.\d+)+)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) == 0 {
		return nil
	}
	goVersionStr := matches[1]
	if verTokens := strings.Split(goVersionStr, "."); len(verTokens) < 3 {
		goVersionStr += ".0"
	}
	goVersion := semver.ParseSemver(goVersionStr)
	return &goVersion
}

func extractReleaseVersion(input string) *semver.Semver {
	reSemver := regexp.MustCompile(`^#+.*?[\s:-]v?((0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?)`)
	matches := reSemver.FindStringSubmatch(input)
	if len(matches) == 0 {
		return nil
	}
	releaseVersion := semver.ParseSemver(matches[1])
	return &releaseVersion
}

func openReleaseNotesFile() (*os.File, error) {
	releaseNoteFilenames := []string{
		"RELEASE-NOTES.md", "RELEASE_NOTES.MD", "RELEASE-NOTES",
		"RELEASE_NOTES.md", "RELEASE_NOTES.MD", "RELEASE_NOTES",
		"release-notes.md", "release-notes",
		"release_notes.md", "release_notes",
	}
	for _, filename := range releaseNoteFilenames {
		file, err := os.Open(filename)
		if err == nil || !os.IsNotExist(err) {
			return file, err
		}
	}
	return nil, os.ErrNotExist
}

func TestGoVersion(t *testing.T) {
	testName := "TestGoVersion"
	file, err := os.Open("./go.mod")
	if err != nil {
		t.Fatalf("%s failed: cannot open <go.mod> file, error %s", testName, err)
	}
	defer file.Close()

	var goVersionMod *semver.Semver
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		goVersionMod = extractGoVersion(line)
		if goVersionMod != nil {
			break
		}
	}
	if goVersionMod == nil {
		t.Fatalf("%s failed: cannot find go version from <go.mod> file", testName)
	}

	if goVersionMod.Error() != nil {
		t.Fatalf("%s failed: cannot parse go version from <go.mod> file, error %s", testName, goVersionMod.Error())
	}

	goVersionRuntime := extractGoVersion(runtime.Version())
	if goVersionRuntime == nil || goVersionRuntime.Error() != nil {
		t.Fatalf("%s failed: cannot parge go version from runtime <%s>", testName, runtime.Version())
	}

	if goVersionMod.Compare(*goVersionRuntime) > 0 {
		t.Fatalf("%s failed: go.mod is expecting Go version <%s> but Go runtime is <%s>", testName, goVersionMod.String(), goVersionRuntime.String())
	}
}

func TestReleaseVersion(t *testing.T) {
	testName := "TestReleaseVersion"
	if os.Getenv("CHECK_RELEASE_VERSION") == "" {
		t.Skipf("%s skipped: environment variable CHECK_RELEASE_VERSION is not set", testName)
	}

	file, err := openReleaseNotesFile()
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("%s failed: no release notes file found", testName)
		}
		t.Fatalf("%s failed: cannot open release notes file, error %s", testName, err)
	}
	defer file.Close()

	var releaseVersion *semver.Semver
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		releaseVersion = extractReleaseVersion(line)
		if releaseVersion != nil {
			break
		}
	}
	if releaseVersion == nil {
		t.Fatalf("%s failed: cannot locate release version in release notes file", testName)
	}

	if releaseVersion.Error() != nil {
		t.Fatalf("%s failed: cannot parse release version from release notes file, error %s", testName, releaseVersion.Error())
	}

	moduleVersion := semver.ParseSemver(Version)
	if releaseVersion.Error() != nil {
		t.Fatalf("%s failed: cannot parse module version <%s>, error %s", testName, Version, releaseVersion.Error())
	}

	if releaseVersion.Compare(moduleVersion) != 0 {
		t.Fatalf("%s failed: release version in release notes <%s> is not in-sync with module version <%s>", testName, releaseVersion.String(), moduleVersion.String())
	}
}

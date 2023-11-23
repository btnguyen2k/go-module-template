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

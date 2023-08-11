package manifest_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"code.cloudfoundry.org/bosh-compile-action/pkg/manifest"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestCanCreateManifest(t *testing.T) {
	m := manifest.Manifest{
		Packages: []manifest.Package{
			{Name: "go-application", Dependencies: []string{"golang"}},
			{Name: "java-application", Dependencies: []string{"maven"}},
			{Name: "maven", Dependencies: []string{"java"}},
			{Name: "java", Dependencies: []string{}},
			{Name: "golang", Dependencies: []string{}},
		},
	}

	type test struct {
		node string
		want []string
	}

	tests := []test{
		{node: "go-application", want: []string{"golang", "go-application"}},
		{node: "java-application", want: []string{"java", "maven", "java-application"}},
		{node: "maven", want: []string{"java", "maven"}},
		{node: "java", want: []string{"java"}},
		{node: "golang", want: []string{"golang"}},
	}

	for _, tc := range tests {
		got, err := m.Dependencies(tc.node)
		assert.NoError(t, err)
		assert.Equal(t, tc.want, got)
	}
}

func TestDetermineTopLevelPackages(t *testing.T) {
	m := manifest.Manifest{
		Packages: []manifest.Package{
			{Name: "go-application", Dependencies: []string{"golang"}},
			{Name: "java-application", Dependencies: []string{"maven"}},
			{Name: "maven", Dependencies: []string{"java"}},
			{Name: "java", Dependencies: []string{}},
			{Name: "golang", Dependencies: []string{}},
		},
	}

	got, err := m.TopLevelPackages()
	assert.NoError(t, err)
	assert.Equal(t, len(got), 2)
	expected := []string{"go-application", "java-application"}
	assert.Equal(t, expected, got)
}

func TestCanParseRealManifest(t *testing.T) {
	bytes, err := os.ReadFile(filepath.Join("test_data", "release.MF"))
	assert.NoError(t, err)

	m := manifest.Manifest{}

	err = yaml.Unmarshal(bytes, &m)
	assert.NoError(t, err)

	type test struct {
		node string
		want []string
	}

	fmt.Printf("found %+v\n", m)
	t.Logf("found %+v", m)

	tests := []test{
		{node: "golangapiserver", want: []string{"golang", "autoscaler-src", "golangapiserver"}},
		{node: "db", want: []string{"java", "maven", "db"}},
		{node: "maven", want: []string{"java", "maven"}},
		{node: "java", want: []string{"java"}},
		{node: "golang", want: []string{"golang"}},
	}

	for _, tc := range tests {
		got, err := m.Dependencies(tc.node)
		assert.NoError(t, err)
		assert.Equal(t, tc.want, got)
	}
}

func TestCanDetermineBestBuildListFromRealManifest(t *testing.T) {
	bytes, err := os.ReadFile(filepath.Join("test_data", "release.MF"))
	assert.NoError(t, err)

	m := manifest.Manifest{}

	err = yaml.Unmarshal(bytes, &m)
	assert.NoError(t, err)

	got, err := m.TopLevelPackages()
	assert.NoError(t, err)

	expected := []string{
		"changeloglockcleaner",
		"common",
		"db",
		"eventgenerator",
		"golangapiserver",
		"metricsforwarder",
		"metricsgateway",
		"metricsserver",
		"operator",
		"scalingengine",
		"scheduler",
	}
	assert.Equal(t, expected, got)
}

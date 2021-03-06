package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/jenkins-x/jx/pkg/tests"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestReplacePlaceholders(t *testing.T) {
	f, err := ioutil.TempDir("", "test-extract-domain")
	assert.NoError(t, err)

	testData := path.Join("test_data", "replace_placeholders")
	_, err = os.Stat(testData)
	assert.NoError(t, err)

	util.CopyDir(testData, f, true)

	assert.NoError(t, err)
	o := ImportOptions{}
	o.Out = tests.Output()
	o.Dir = f
	o.AppName = "bar"

	o.replacePlaceholders("github.com", "foo")

	// root file
	testFile, err := util.LoadBytes(f, "file.txt")
	assert.Equal(t, "/home/jenkins/go/src/github.com/foo/bar", string(testFile), "replaced placeholder")

	// dir1
	testDir1 := path.Join(f, "dir1")
	testFile, err = util.LoadBytes(testDir1, "file.txt")
	assert.Equal(t, "/home/jenkins/go/src/github.com/foo/bar", string(testFile), "replaced placeholder")

	// dir2
	testDir2 := path.Join(f, "dir2")
	testFile, err = util.LoadBytes(testDir2, "file.txt")
	assert.Equal(t, "/home/jenkins/go/src/github.com/foo/bar", string(testFile), "replaced placeholder")

}

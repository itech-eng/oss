package filesystem

import (
	"testing"

	"github.com/itech-eng/oss/tests"
)

func TestAll(t *testing.T) {
	fileSystem := New("/tmp")
	tests.TestAll(fileSystem, t)
}

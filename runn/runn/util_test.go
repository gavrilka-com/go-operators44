package runn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_runCommand(t *testing.T) {
	//test for RunCommand
	out, err := RunCommand("uname -s")
	if err != nil {
		t.Error("Error in RunCommand")
	}
	assert.Equal(t, "Linux\n", out)
}

func Test_runCommandVerbose(t *testing.T) {
	//test for RunCommand
	var out, err = RunCommandVerbose("echo \"test_echo1\" && sleep 1 && echo \"test_echo2\"")
	if err != nil {
		t.Error("Error in RunCommand")
	}
	assert.Equal(t, "test_echo1\ntest_echo2\n", out)
}

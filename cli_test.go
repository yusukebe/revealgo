package revealgo

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunFlagExitCode(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := CLI{OutStream: outStream, ErrStream: errStream}

	flags := []struct {
		flag     string
		expected int
	}{
		{flag: "version", expected: ExitCodeOK},
		{flag: "help", expected: ExitCodeOK},
		{flag: "blahblah", expected: ExitCodeError},
	}

	for _, f := range flags {
		command := fmt.Sprintf("revealgo --%s", f.flag)
		args := strings.Split(command, " ")
		status := cli.Run(args)
		if status != f.expected {
			t.Errorf("got %v\n want %v", status, f.expected)
		}
	}
}

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := CLI{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("revealgo --version", " ")

	cli.Run(args)

	expected := fmt.Sprintf("revealgo version %s\n", Version)
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("Output=%q, want %q", outStream.String(), expected)
	}
}

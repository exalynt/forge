package playwright

import (
	"errors"
	"os"
	"os/exec"
)

// execNode runs Node with this process's input and output, and exits as it
// does, since Windows can't replace a running process.
func execNode(node string, args []string) error {
	c := exec.Command(node, args...)
	c.Stdin, c.Stdout, c.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := c.Run()
	var exit *exec.ExitError
	if errors.As(err, &exit) {
		os.Exit(exit.ExitCode())
	}
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

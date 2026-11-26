//go:build !windows

package playwright

import (
	"os"
	"syscall"
)

// execNode replaces this process with Node, so the agent talks to the server
// directly and stopping it stops the server.
func execNode(node string, args []string) error {
	return syscall.Exec(node, append([]string{node}, args...), os.Environ())
}

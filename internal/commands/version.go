package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v3"
	"golang.org/x/mod/semver"
)

// VersionCommand prints this binary's version, or compares it with the
// latest forge.
func VersionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Print this binary's version",
		Description: `--check asks Go for the latest forge, and exits 1 when this binary is older,
or 2 when it can't tell, such as without Go or a network connection.`,
		Flags:  []cli.Flag{jsonFlag, &cli.BoolFlag{Name: "check", Usage: "compare with the latest forge"}},
		Action: runVersion,
	}
}

func runVersion(ctx context.Context, cmd *cli.Command) error {
	out := cmd.Root().Writer
	v := version()
	if !cmd.Bool("check") {
		if cmd.Bool("json") {
			writeJSON(out, map[string]string{"version": v})
		} else {
			_, _ = fmt.Fprintln(out, v)
		}
		return nil
	}
	latest, err := latestVersion(ctx)
	if err != nil {
		return err
	}
	older := false
	switch {
	case v == "dev" || !semver.IsValid(v):
		return fmt.Errorf("this binary is %s, so it can't be compared with the latest, %s: install it with go install %s/cmd/forge@latest", v, latest, module)
	case semver.Compare(v, latest) < 0:
		older = true
	}
	if cmd.Bool("json") {
		writeJSON(out, map[string]any{"version": v, "latest": latest, "upToDate": !older})
	} else if older {
		_, _ = fmt.Fprintf(out, "forge %s is older than the latest, %s. Update it with:\n\n  go install %s/cmd/forge@latest\n\nthen run forge update in each repository.\n", v, latest, module)
	} else {
		_, _ = fmt.Fprintf(out, "forge %s is the latest.\n", v)
	}
	if older {
		return errStale
	}
	return nil
}

// latestVersion asks Go for forge's latest version.
func latestVersion(ctx context.Context) (string, error) {
	goBin, err := exec.LookPath("go")
	if err != nil {
		return "", errors.New("can't check: Go isn't installed, or isn't on PATH")
	}
	c := exec.CommandContext(ctx, goBin, "list", "-m", "-json", module+"@latest")
	// Outside any module, so the current folder's go.mod can't get in the way.
	c.Dir = os.TempDir()
	var stderr strings.Builder
	c.Stderr = &stderr
	data, err := c.Output()
	if err != nil {
		return "", fmt.Errorf("can't check: go list couldn't reach %s (%s)", module, strings.TrimSpace(stderr.String()))
	}
	var m struct{ Version string }
	if err := json.Unmarshal(data, &m); err != nil || m.Version == "" {
		return "", errors.New("can't check: go list didn't say which version is the latest")
	}
	return m.Version, nil
}

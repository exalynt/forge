// Package playwright sets up Playwright MCP once per machine, so the qa, ux and
// use skills can drive a browser, and runs it for an agent.
//
// Setup installs @playwright/mcp and Chrome, and records in setup.json where
// Node and the package are and the flags the server gets. Agents register
// forge browser server, which reads setup.json and runs the package with
// Node. Re-running setup after a Node switch, or with other flags, so needs
// no new registration: the servers use it the next time they start.
package playwright

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//go:embed chrome.sh
var chromeScript []byte

//go:embed keep-focus.ps1
var keepFocusScript []byte

// Servers are the names to register, each with the same command: playwright,
// the one the skills use by default, and four more for when a skill needs
// several browsers at once.
var Servers = []string{"playwright", "playwright-1", "playwright-2", "playwright-3", "playwright-4"}

// Setup is what setup.json records.
type Setup struct {
	// Forge is the binary setup ran as, which the servers are registered with.
	Forge string `json:"forge"`
	// Node and Package, its cli.js, are absolute, so the server starts even
	// where the agent's environment doesn't have Node on PATH, as with nvm
	// outside an interactive shell.
	Node    string   `json:"node"`
	Package string   `json:"package"`
	Flags   []string `json:"flags"`
	Browser string   `json:"browser"`
	// Launcher is chrome.sh, which Playwright runs in place of Chrome, or empty
	// where it isn't used: a headless browser, another browser, or Windows.
	Launcher string `json:"launcher,omitempty"`
}

// Args are the arguments the server runs Node with. It always gets
// --isolated, so each session starts with a fresh, in-memory profile.
func (s Setup) Args() []string {
	args := []string{s.Package, "--isolated"}
	if s.Launcher != "" {
		args = append(args, "--executable-path", s.Launcher)
	}
	return append(args, s.Flags...)
}

// Dir is where setup writes setup.json and the launcher.
func Dir() (string, error) {
	if d := os.Getenv("XDG_DATA_HOME"); filepath.IsAbs(d) {
		return filepath.Join(d, "forge", "playwright"), nil
	}
	if runtime.GOOS == "windows" {
		d, err := os.UserCacheDir() // %LocalAppData%
		if err != nil {
			return "", err
		}
		return filepath.Join(d, "forge", "playwright"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", "forge", "playwright"), nil
}

// ErrNotSetUp is returned when setup.json doesn't exist.
var ErrNotSetUp = errors.New("browsers aren't set up: run forge browser setup")

// Load reads setup.json.
func Load() (Setup, string, error) {
	dir, err := Dir()
	if err != nil {
		return Setup{}, "", err
	}
	path := filepath.Join(dir, "setup.json")
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return Setup{}, path, ErrNotSetUp
	}
	if err != nil {
		return Setup{}, path, err
	}
	var s Setup
	if err := json.Unmarshal(data, &s); err != nil {
		return Setup{}, path, fmt.Errorf("can't read %s: %w: run forge browser setup again", path, err)
	}
	return s, path, nil
}

// Install installs @playwright/mcp globally with npm, and Chrome if it's
// the browser and missing, then writes the launcher and setup.json. flags are
// passed on to the server. Commands' output goes to stdout and stderr.
func Install(ctx context.Context, stdout, stderr io.Writer, flags []string) (Setup, string, error) {
	npm, err := exec.LookPath("npm")
	if err != nil {
		return Setup{}, "", errors.New("npm not found: install Node.js first")
	}
	node, err := exec.LookPath("node")
	if err != nil {
		return Setup{}, "", errors.New("node not found: install Node.js first")
	}
	if node, err = filepath.Abs(node); err != nil {
		return Setup{}, "", err
	}
	forge, err := os.Executable()
	if err != nil {
		return Setup{}, "", err
	}
	run := func(name string, args ...string) error {
		c := exec.CommandContext(ctx, name, args...)
		c.Stdout, c.Stderr = stdout, stderr
		if err := c.Run(); err != nil {
			return fmt.Errorf("%s %s: %w", filepath.Base(name), strings.Join(args, " "), err)
		}
		return nil
	}

	_, _ = fmt.Fprintln(stdout, "Installing @playwright/mcp globally")
	if err := run(npm, "install", "-g", "@playwright/mcp@latest"); err != nil {
		return Setup{}, "", err
	}
	browser, headless := readFlags(flags)
	if browser == "chrome" {
		if path, ok := Chrome(); ok {
			_, _ = fmt.Fprintf(stdout, "Google Chrome is already installed: %s\n", path)
		} else {
			npx, err := exec.LookPath("npx")
			if err != nil {
				return Setup{}, "", errors.New("npx not found, to install Chrome: install Node.js with npm")
			}
			_, _ = fmt.Fprintln(stdout, "Installing Google Chrome (may ask for sudo)")
			if err := run(npx, "-y", "playwright", "install", "--with-deps", "chrome"); err != nil {
				return Setup{}, "", err
			}
		}
	}

	var root bytes.Buffer
	c := exec.CommandContext(ctx, npm, "root", "-g")
	c.Stdout, c.Stderr = &root, stderr
	if err := c.Run(); err != nil {
		return Setup{}, "", fmt.Errorf("npm root -g: %w", err)
	}
	pkg := filepath.Join(strings.TrimSpace(root.String()), "@playwright", "mcp", "cli.js")
	if _, err := os.Stat(pkg); err != nil {
		return Setup{}, "", fmt.Errorf("not found after install: %s", pkg)
	}

	dir, err := Dir()
	if err != nil {
		return Setup{}, "", err
	}
	s := Setup{Forge: forge, Node: node, Package: pkg, Flags: flags, Browser: browser}
	if s.Flags == nil {
		s.Flags = []string{}
	}
	// The launcher arranges Chrome's windows. A headless browser has none, and
	// another browser opens as the server would open it anyway.
	if browser == "chrome" && !headless && runtime.GOOS != "windows" {
		if s.Launcher, err = WriteLauncher(dir); err != nil {
			return Setup{}, "", err
		}
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return Setup{}, "", err
	}
	path := filepath.Join(dir, "setup.json")
	if err := writeFile(path, append(data, '\n'), 0o644); err != nil {
		return Setup{}, "", err
	}
	return s, path, nil
}

// readFlags returns the browser the server's flags choose, chrome when they
// don't, and whether they make it headless.
func readFlags(flags []string) (browser string, headless bool) {
	browser = "chrome"
	for i, f := range flags {
		switch {
		case f == "--headless":
			headless = true
		case f == "--browser" && i+1 < len(flags):
			browser = flags[i+1]
		case strings.HasPrefix(f, "--browser="):
			browser = strings.TrimPrefix(f, "--browser=")
		}
	}
	return browser, headless
}

// Chrome returns where Google Chrome is, and whether it was found.
func Chrome() (string, bool) {
	if p, err := exec.LookPath("google-chrome"); err == nil {
		return p, true
	}
	paths := []string{"/Applications/Google Chrome.app"}
	if runtime.GOOS == "windows" {
		for _, env := range []string{"ProgramFiles", "ProgramFiles(x86)", "LocalAppData"} {
			if d := os.Getenv(env); d != "" {
				paths = append(paths, filepath.Join(d, "Google", "Chrome", "Application", "chrome.exe"))
			}
		}
	}
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p, true
		}
	}
	return "", false
}

// WriteLauncher writes chrome.sh, and keep-focus.ps1 beside it, into dir where
// they differ from this binary's, and returns chrome.sh's path.
func WriteLauncher(dir string) (string, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	if err := writeFile(filepath.Join(dir, "keep-focus.ps1"), keepFocusScript, 0o644); err != nil {
		return "", err
	}
	path := filepath.Join(dir, "chrome.sh")
	return path, writeFile(path, chromeScript, 0o755)
}

// writeFile writes data to path unless it's there already, by renaming a
// temporary file over it: a browser opening runs chrome.sh while it's being
// replaced, and bash reads a script as it goes.
func writeFile(path string, data []byte, perm fs.FileMode) error {
	if old, err := os.ReadFile(path); err == nil && bytes.Equal(old, data) {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+"-*")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(f.Name()) }()
	if _, err := f.Write(data); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Chmod(perm); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}

// Serve runs @playwright/mcp as setup.json records, in place of this process
// where the platform allows it. Its stdout is the agent's connection to the
// server, so nothing else may write there. It rewrites the launcher first, so
// a forge update brings a new one without running setup again.
func Serve() error {
	s, _, err := Load()
	if err != nil {
		return err
	}
	if s.Launcher != "" {
		if _, err := WriteLauncher(filepath.Dir(s.Launcher)); err != nil {
			return err
		}
	}
	return execNode(s.Node, s.Args())
}

// Statuses a check reports.
const (
	OK      = "ok"
	Missing = "missing"
)

// Item is one thing Check looked at.
type Item struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Status string `json:"status"`
	Fix    string `json:"fix,omitempty"`
}

// Check reports whether what setup recorded is still there. The launcher
// isn't checked, since each server rewrites it as it starts.
func Check() ([]Item, error) {
	s, path, err := Load()
	if errors.Is(err, ErrNotSetUp) {
		return []Item{{Name: "setup", Path: path, Status: Missing, Fix: "run forge browser setup"}}, nil
	}
	if err != nil {
		return nil, err
	}
	items := []Item{{Name: "setup", Path: path, Status: OK}}
	exists := func(name, p, fix string) {
		it := Item{Name: name, Path: p, Status: OK}
		if _, err := os.Stat(p); err != nil {
			it.Status, it.Fix = Missing, fix
		}
		items = append(items, it)
	}
	exists("forge", s.Forge, "run forge browser setup, then register the command it prints in place of the old one")
	exists("node", s.Node, "run forge browser setup, to use the Node on PATH now")
	exists("package", s.Package, "run forge browser setup, to install it for the Node on PATH now")
	if s.Browser == "chrome" {
		p, ok := Chrome()
		it := Item{Name: "chrome", Path: p, Status: OK}
		if !ok {
			it.Status, it.Fix = Missing, "run forge browser setup, to install it"
		}
		items = append(items, it)
	}
	return items, nil
}

// +build windows

package open

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	cmd      = "url.dll,FileProtocolHandler"
	runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)

func cleaninput(input string) string {
	r1 := strings.NewReplacer("&", "^&")
	r1.Replace(input)
	r2 := strings.NewReplacer("/", "\\")
	return r2.Replace(input)
}

func open(input string) *exec.Cmd {
	return exec.Command(runDll32, cmd, cleaninput(input))
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command("cmd", "/C", "start", "", appName, cleaninput(input))
}

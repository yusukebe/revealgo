package revealgo

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func openBrowser(port int) error {
	<-time.After(100 * time.Millisecond)
	url := fmt.Sprintf("http://localhost:%d/", port)
	var args []string
	var cmd string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func addExtention(path string, ext string) string {
	if strings.HasSuffix(path, fmt.Sprintf(".%s", ext)) {
		return path
	}
	path = fmt.Sprintf("%s.%s", path, ext)
	return path
}

func detectContentType(path string) string {
	switch {
	case strings.HasSuffix(path, ".css"):
		return "text/css"
	case strings.HasSuffix(path, ".js"):
		return "application/javascript"
	case strings.HasSuffix(path, ".svg"):
		return "image/svg+xml"
	}
	return ""
}

package pkg

import (
	"fmt"
	"os/exec"
	"runtime"
)

func ShutdownSystemWithDelay(seconds int) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("shutdown", "/s", "/t", fmt.Sprintf("%d", seconds))

	case "linux", "darwin":
		cmd = exec.Command("shutdown", "-h", fmt.Sprintf("+%d", seconds/60))

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Run()
}

package pkg

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
)

type AfterFinishFunc func(int) error

func AskAfterFinishBehavior() (AfterFinishFunc, error) {
	options := []string{
		"keep the pc running",
		"shut down",
	}

	prompt := &survey.Select{
		Message: "What should the program do after completion?",
		Options: options,
	}

	var selected string
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	fmt.Println("\nYou selected:")
	fmt.Printf("✓ %s\n", selected)

	switch selected {
	case "keep the pc running":
		return nil, nil
	case "shut down":
		return shutDown, nil
	default:
		return nil, nil
	}
}

func shutDown(seconds int) error {
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

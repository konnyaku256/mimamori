package exec_command

import (
	"fmt"
	"os/exec"
)

func ExecCommand(command Command) error {
	out, err := exec.Command(command.Name, command.Arguments...).CombinedOutput()
	fmt.Printf("Command result:\n%s :Error:\n%v\n", out, err)

	return err
}
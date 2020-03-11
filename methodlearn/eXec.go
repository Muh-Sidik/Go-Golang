package methodlearn

import (
	"fmt"
	"os/exec"
)

func testCommand()  {
	var cmd1, _ = exec.Command("ls").Output()
	fmt.Printf("-> ls\n%s\n", string(cmd1))

	var cmd2, _ = exec.Command("whoami").Output()
	fmt.Printf("-> whoami\n%s\n", string(cmd2))

	var cmd3, _ = exec.Command("git", "config", "user.email").Output()
	fmt.Printf("-> git config user.email\n%s\n", cmd3)


}
package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func GetOperationDelay() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the delay in seconds:")
	scanner.Scan()
	return scanner.Text()
}

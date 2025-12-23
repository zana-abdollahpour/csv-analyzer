package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func GetSearchTerm() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the term you wish to search for:")
	scanner.Scan()
	return scanner.Text()
}

package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Int int input test
func Int() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Println("Type a number.")
	stdin.Scan()
	text := stdin.Text()
	i, _ := strconv.Atoi(text)
	fmt.Println("Next is " + strconv.Itoa(i+1) + ".")
}

// SplitInt splitted int input test
func SplitInt() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Println("Type numbers separated with space.")
	stdin.Scan()
	text := stdin.Text()
	strs := strings.Fields(text)
	// strs := strings.Split(text, " ")
	ints := []int{}
	for _, c := range strs {
		i, err := strconv.Atoi(c)
		if err == nil {
			ints = append(ints, i)
		}
	}
	fmt.Println(ints)
}

package printversion

import "fmt"

const (
	version = "1.1.1"
)

//PrintVersion Print version
func PrintVersion() {
	fmt.Println(version)
}

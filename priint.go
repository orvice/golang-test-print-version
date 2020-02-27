package printversion

import "fmt"

const (
	version = "1.1.0"
)

//PrintVersion Print version
func PrintVersion() {
	fmt.Println(version)
}

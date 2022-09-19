package Utils

import (
	"fmt"
	"strings"
)

func CreatePath(path string) string {

	splitPath := strings.Split(path, "/")

	fmt.Println(splitPath)

	return path
}

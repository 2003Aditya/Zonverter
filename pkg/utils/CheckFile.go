// CheckFile.go is responsible to check if the given input file is correct or not if not the program will break
package utils

import (
	"fmt"
	"os"
)

func CheckFile(filepath string) bool {
    _, err := os.Stat(filepath)
    if err == nil {
        return true
    }

    if os.IsNotExist(err) {
        return false
    }

    fmt.Println("error with file:", err)
    return false

}

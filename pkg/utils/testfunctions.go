package utils

import "fmt"

func PrintSomething(someMessage string) error {
	fmt.Printf("YOUR MESSAGE: %s", someMessage)
	return nil
}

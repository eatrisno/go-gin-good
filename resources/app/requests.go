package app

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MarkErrors(errors []validator.ValidationErrors) {

	for _, e := range errors {
		fmt.Printf("Error: %s\n", e.Error())
	}
}

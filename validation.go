package main

import "fmt"

// declare custom variable type
type inputValidation string
type inputValidationResult bool

// receiver function name validate for inputValidation
// with return pointer of inputValidationResult
func (inp inputValidation) validate() *inputValidationResult {
	// init custom variable inputValidationResult
	inpResFalse := inputValidationResult(false)
	inpResTrue := inputValidationResult(true)

	if inp == "[]" || inp == "{}" {
		// return pointer address of inpResTrue
		return &inpResTrue
	}
	if inp == "[" || inp == "{" || inp == "]" || inp == "}" || inp == "[}" || inp == "{]" {
		// return pointer address of inpResFalse
		return &inpResFalse
	}
	// return empty pointer
	return nil
}

// receiver function name print for inputValidation
// with return inputValidation
func (inp inputValidation) print() inputValidation {
	fmt.Print("Input: '", inp, "' ")
	return inp
}

// receiver function name print for inputValidationResult pointer
func (rslt *inputValidationResult) print() {
	if rslt == nil {
		// pointer == nil, print the pointer address
		fmt.Println("Return: ", rslt)
	} else {
		// pointer != nil, print the pointer value
		fmt.Println("Return: ", *rslt)
	}
}

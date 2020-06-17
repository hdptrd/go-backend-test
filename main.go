package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// test 01
	input1 := inputString("19374048")
	input2 := inputString("aiueobcd")
	idx := input1.findIndexOfChar("7")
	str := input2.getStringStartFromIndex(idx)
	fmt.Println("=============================")
	fmt.Println("========== TEST 01 ==========")
	fmt.Println("=============================")
	fmt.Print("Input 1: ")
	input1.print()
	fmt.Println()
	fmt.Print("Input 2: ")
	input2.print()
	fmt.Println()
	fmt.Println("Result: ", str)
	fmt.Println()

	// test 02
	fmt.Println("=============================")
	fmt.Println("========== TEST 02 ==========")
	fmt.Println("=============================")
	inputValidation("[]").print().validate().print()
	inputValidation("{}").print().validate().print()
	inputValidation("{").print().validate().print()
	inputValidation("[").print().validate().print()
	inputValidation("}").print().validate().print()
	inputValidation("]").print().validate().print()
	inputValidation("[}").print().validate().print()
	inputValidation("{]").print().validate().print()
	inputValidation("haloo").print().validate().print()
	fmt.Println()

	// test 03
	key := "1234privy5678"
	message := "Selamat Datang"
	fmt.Println("=============================")
	fmt.Println("========== TEST 03 ==========")
	fmt.Println("=============================")
	fmt.Println("key: ", key)
	fmt.Println("Original message: ", message)
	encryptedString := encryptString(key, message)
	fmt.Println("Encrypted message: ", encryptedString)
	decryptedString := decryptString(key, encryptedString)
	fmt.Println("Decrypted message: ", decryptedString)
	fmt.Println()

	// test 04
	slices := []string{"one", "two", "three", "one", "four", "five", "two"}
	result := fungsiA(slices)
	fmt.Println("=============================")
	fmt.Println("========== TEST 04 ==========")
	fmt.Println("=============================")
	fmt.Println(slices)
	fmt.Println(result)
	fmt.Println()

	// test 05
	fmt.Println("=============================")
	fmt.Println("========== TEST 05 ==========")
	fmt.Println("=============================")
	os.Setenv("ENVIRONMENT", "DEVELOPMENT")
	os.Setenv("DEVELOPMENT_PORT", "8081")
	os.Setenv("DEVELOPMENT_APP_NAME", "app name dev")
	os.Setenv("DEVELOPMENT_DEBUG_MODE", "true")
	os.Setenv("PRODUCTION_PORT", "8081")
	os.Setenv("PRODUCTION_APP_NAME", "app name prod")
	os.Setenv("PRODUCTION_DEBUG_MODE", "true")

	envEnvironment := os.Getenv("ENVIRONMENT")
	envDebugMode, _ := strconv.ParseBool(os.Getenv(envEnvironment + "_DEBUG_MODE"))

	cfg := config{
		envEnvironment: os.Getenv("ENVIRONMENT"),
		envAppName:     os.Getenv(envEnvironment + "_APP_NAME"),
		envPort:        os.Getenv(envEnvironment + "_PORT"),
		envDebugMode:   envDebugMode,
	}

	http.HandleFunc("/ping", loggingHandler(pingHandler, cfg))
	http.HandleFunc("/time", loggingHandler(timeHandler, cfg))
	http.HandleFunc("/", loggingHandler(rootHandler, cfg))
	log.Println("listening on port", cfg.envPort, "...")
	log.Fatal(http.ListenAndServe(":"+cfg.envPort, nil))
}

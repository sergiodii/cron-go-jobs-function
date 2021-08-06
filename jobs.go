package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func test(envFileName string) {
	err := godotenv.Load(envFileName)
	if err != nil {
		fmt.Println("\033[31mError: Shared Package dont founded " + envFileName + " file\033[0m")
	}
	fmt.Println("TESTE JOBS: " + os.Getenv("DB_HOST"))
}

func main() {
	var execute string = ""
	fmt.Println(execute)
}

func Job() map[string]interface{} {
	jobs := make(map[string]interface{})
	jobs["Test"] = test
	return jobs
}

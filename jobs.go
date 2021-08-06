package main

import (
	"fmt"
	"os"
)

func test() {
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

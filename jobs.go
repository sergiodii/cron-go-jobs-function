package main

import "fmt"

func test() {
	fmt.Println("Ola Mundo, from jobs_functions")
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

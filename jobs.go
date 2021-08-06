package jobs_functions

import "fmt"

func test() {
	fmt.Println("Ola Mundo, from jobs_functions")
}

func Job() map[string]interface{} {
	jobs := make(map[string]interface{})
	jobs["Test"] = test
	return jobs
}

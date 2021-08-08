package jobs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Test(envFileName string) {
	err := godotenv.Load(envFileName)
	if err != nil {
		fmt.Println("\033[31mError: Shared Package dont founded " + envFileName + " file\033[0m")
	}
	fmt.Println("TESTE JOBS: " + os.Getenv("DB_HOST"))
}

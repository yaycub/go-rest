package connect

import "fmt"

func Connect() {
	dbUrl := GoDotEnvVariable("DB_URL")

	fmt.Printf(dbUrl)
}

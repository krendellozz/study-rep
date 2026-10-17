package main

import (
	"fmt"
	"study/feature1"
	"study/feature2"
	simpleconnection "study/feature_postres/simple_connection"
)

// "postgres//YourUserName:YourPassword@YourHostName:5432/YorDatabaseName"
func main() {
	fmt.Println("Hello, ")
	feature1.Feature1()
	feature2.Feature2()
	simpleconnection.CheckConnection()
}

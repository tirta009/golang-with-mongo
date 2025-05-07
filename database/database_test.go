package database

import (
	"fmt"
	"testing"
)

func TestInitClient(t *testing.T) {

	InitClient()
	fmt.Println("Connected to MongoDB")

}

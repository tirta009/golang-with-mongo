package config

import (
	"fmt"
	"testing"
)

func TestInitClient(t *testing.T) {

	InitDatabase()
	fmt.Println("Connected to MongoDB")

}

package main

import (
	"fmt"

	"github.com/aman1117/backend/utils"
)

func main() {

	db := utils.GetDB()
	fmt.Println(db)
}

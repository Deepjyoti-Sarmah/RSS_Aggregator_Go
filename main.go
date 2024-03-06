package main

import (
	"fmt"
	"os"
  "log"

  "github.com/joho/godotenv"
)

func main() {
  fmt.Println("RSS Aggrigator in Golang")
  
  godotenv.Load(".env")

  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT is not found in the environment")
  }

  fmt.Println("PORT", portString)
}

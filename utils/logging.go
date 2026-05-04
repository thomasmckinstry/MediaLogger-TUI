package utils

import (
	"log"
	"os"
)

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func DebugLog(msg string, output any) {
	if len(os.Getenv("DEBUG")) > 0 {
		log.Println(msg, output)
	}
}

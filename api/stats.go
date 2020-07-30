package api

import (
	"log"
)

func InitStats(mKey string) {
	var err error
	err = sumPixelStats()
	if err != nil {
		log.Fatal("cannot sum Pixel Stats:", err)
	}
	log.Println("sum Pixel Stats OK")
}

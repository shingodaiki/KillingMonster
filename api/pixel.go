package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mssola/user_agent"
)

// 1 pixel transparent GIF in 43 bytes
var gif = []byte{
	0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x01, 0x00, 0x01, 0x00, 0x80, 0x00, 0x00, 0xFF, 0xFF, 0xFF,
	0x00, 0x00, 0x00, 0x21, 0xF9, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00, 0x00,
	0x01, 0x00, 0x01, 0x00, 0x00, 0x02, 0x02, 0x44, 0x01, 0x00, 0x3B,
}

func getTimeUTC() string {
	now := time.Now().UTC()
	unix := now.Unix()
	return now.Format("2 Jan 2006, 3:04 PM") + " / " + strconv.FormatInt(unix, 10)
}

func getAddressFromReq(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		return "xff-" + xff
	}

	return r.RemoteAddr
}

func getBrowserFromReq(r *http.Request) string {
	ua := user_agent.New(r.Header.Get("User-Agent"))
	os := ua.OSInfo().Name
	bot := ua.Bot()
	mobile := ua.Mobile()
	browser, _ := ua.Browser()
	browser = browser + " / " + os
	if bot {
		browser = "[BOT] " + browser
	}
	if mobile {
		browser = browser + " (mobile)"
	}
	return browser
}

func pixel(w http.ResponseWriter, r *http.Request) {
	log.Println("timeutc:", getTimeUTC())
	log.Println("address:", getAddressFromReq(r))
	log.Println("browser:", getBrowserFromReq(r))

	w.Header().Set("Cache-Control", "max-age=0, no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	w.Header().Set("Content-Type", "image/gif")
	w.Write(gif)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", r.RemoteAddr)
}

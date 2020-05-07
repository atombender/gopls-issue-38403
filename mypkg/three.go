package mypkg

import (
	"log"
	"net/url"
)

func Three(u *url.URL) {
	log.Printf("url: %s", u.String())
}

package mypkg

import (
	net_url "net/url"
)

func Two() *net_url.URL {
	var url net_url.URL
	url.Scheme = "https"
	url.Host = "example.com"
	return &url
}

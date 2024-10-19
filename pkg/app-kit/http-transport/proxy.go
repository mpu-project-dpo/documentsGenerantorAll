package http_transport

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (c *ProxyConfig) NewReverseProxy(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(c.Host)
	if err != nil {
		http.Error(w, "Error parsing URL", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/rpraveenkumar1203/Golang/LoadBalancer/handler"
)

type Server interface {
	Address() string
	ISalivve() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type Load_Balancer struct {
	Port    string
	RRC     int
	Servers []Server
}

type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

func NewSimpleServer(addr string) *SimpleServer {

	serverUrl, err := url.Parse(addr)
	handler.Error(err)

	return &SimpleServer{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func LoadBalancer(port string, Servers []Server) *Load_Balancer {

	return &Load_Balancer{
		Port:    port,
		RRC:     0,
		Servers: Servers,
	}

}

package main

import (
	"fmt"
	"net/url"
	"os"

	"api_gateway/server"
)

const (
	defaultPort              = "8080"
	
	defaultAuthHost			 = "auth_subsystem"
	defaultAuthPort			 = "8080"

	defaultStreamHost		 = "stream_subsystem"
	defaultStreamPort		 = "8080"
)

func main() {
	var (
		addr   = envString("PORT", defaultPort)

		authHost = envString("AUTH_HOST", defaultAuthHost)
		authPort = envString("AUTH_PORT", defaultAuthPort)

		streamHost = envString("STREAM_HOST", defaultStreamHost)
		streamPort = envString("STREAM_PORT", defaultStreamPort)
	)
	
	authURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", authHost, authPort),
	}

	streamURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", streamHost, streamPort),
	}
	
	srv := server.New(authURL, streamURL)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"tailscale.com/tsnet"
)

var (
	hostname   = flag.String("hostname", "", "tailscale host name")
	targetAddr = flag.String("target", "", "target address")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if *hostname == "" {
		flag.PrintDefaults()
		log.Fatal("No hostname given")
	}

	if *targetAddr == "" {
		flag.PrintDefaults()
		log.Fatal("No target address given")
	}

	s := new(tsnet.Server)
	defer s.Close()

	s.Hostname = *hostname

	confDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(confDir, "tsproxy-"+*hostname)
	if err := os.MkdirAll(dir, 0700); err != nil {
		log.Fatal(err)
	}
	s.Dir = dir

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

	_, port, err := net.SplitHostPort(*targetAddr)
	if err != nil {
		log.Fatal(err)
	}
	ln, err := s.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	log.Printf("Accepted connection from %v", conn.RemoteAddr())

	remote, err := net.Dial("tcp", *targetAddr)
	if err != nil {
		log.Printf("tsproxy: dial error: %v, closing %v",
			err, conn.RemoteAddr())
		return
	}
	defer remote.Close()

	go func() {
		defer remote.Close()
		io.Copy(remote, conn)
	}()

	io.Copy(conn, remote)
}

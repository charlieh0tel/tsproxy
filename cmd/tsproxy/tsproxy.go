package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/netip"
	"os"
	"path/filepath"
	"strings"

	"tailscale.com/tsnet"
)

var (
	hostname   = flag.String("hostname", "", "tailscale host name (default: first label of the target host)")
	targetAddr = flag.String("target", "", "target address, host:port")
)

// defaultHostname returns the tailnet host name to use for target when none
// is given: the first DNS label of the target's host.  An IP address has no
// such label, so it is an error.
func defaultHostname(targetHost string) (string, error) {
	if _, err := netip.ParseAddr(targetHost); err == nil {
		return "", fmt.Errorf("target %q is an IP address; give -hostname", targetHost)
	}
	label, _, _ := strings.Cut(targetHost, ".")
	if label == "" {
		return "", fmt.Errorf("target host %q has no name; give -hostname", targetHost)
	}
	return label, nil
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if *targetAddr == "" {
		flag.PrintDefaults()
		log.Fatal("No target address given")
	}

	targetHost, port, err := net.SplitHostPort(*targetAddr)
	if err != nil {
		log.Fatal(err)
	}

	if *hostname == "" {
		*hostname, err = defaultHostname(targetHost)
		if err != nil {
			log.Fatal(err)
		}
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

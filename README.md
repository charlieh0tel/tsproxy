# tsproxy

tsproxy is a proxy that can expose TCP services accessible from the
proxy machine to your tailnet.

## Usage

```
go build ./cmd/tsproxy
./tsproxy -hostname name-on-tailnet -target host:port

```

You will neeed to authorize the new host using the URL.

## Acknowledgements

Inspired in large measure by James Tucker's
https://github.com/raggi/teltailnet/ .



# tsproxy

tsproxy is a proxy that can expose TCP services accessible from the
proxy machine to your tailnet.

## Install

A Debian package is published in the
[charlieh0tel apt repository](https://github.com/charlieh0tel/apt-repo#adding-the-repository-to-a-host):

```
sudo apt-get install tsproxy
```

Or build from source:

```
go build ./cmd/tsproxy
```

## Usage

```
tsproxy -hostname name-on-tailnet -target host:port
```

You will need to authorize the new host using the URL.

tsproxy listens on the tailnet at the target's port number.  Its
Tailscale state is kept in `~/.config/tsproxy-NAME`.

## Running under systemd

The package installs a systemd user unit, `tsproxy@.service`, that runs
one proxy per tailnet host name.  The instance name is the host name;
the target comes from `~/.config/tsproxy/NAME.env`:

```
mkdir -p ~/.config/tsproxy
echo 'TARGET=host:port' > ~/.config/tsproxy/name-on-tailnet.env
systemctl --user enable --now tsproxy@name-on-tailnet.service
```

The first start logs the authorization URL; read it with
`journalctl --user -u tsproxy@name-on-tailnet.service`.

## Acknowledgements

Inspired in large measure by James Tucker's
https://github.com/raggi/teltailnet/ .

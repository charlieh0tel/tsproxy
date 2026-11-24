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
tsproxy -target host:port [-hostname name-on-tailnet]
```

tsproxy joins the tailnet as `name-on-tailnet` and listens there on the
target's port number.  Without `-hostname` the name is the first label
of the target host: `-target printer.example.com:9100` joins as
`printer`.  A target given as an IP address needs `-hostname`.

You will need to authorize the new host using the URL.

Tailscale state is kept in `~/.config/tsproxy-NAME`.

## Running under systemd

The package installs a systemd user unit, `tsproxy@.service`, that runs
one proxy per target.  The instance name is the target:

```
systemctl --user enable --now tsproxy@printer.example.com:9100.service
```

The first start logs the authorization URL; read it with
`journalctl --user -u tsproxy@printer.example.com:9100.service`.

User units start at boot only if lingering is enabled for the account:
`loginctl enable-linger`.

An optional `~/.config/tsproxy/INSTANCE.env`, where `INSTANCE` is the
instance name, sets environment variables for that proxy, such as
`TS_AUTHKEY` to join without the URL.  For a different tailnet name or
an IP target, override the command in a drop-in
(`systemctl --user edit tsproxy@INSTANCE.service`):

```
[Service]
ExecStart=
ExecStart=/usr/bin/tsproxy --hostname name-on-tailnet --target %i
```

## Acknowledgements

Inspired in large measure by James Tucker's
https://github.com/raggi/teltailnet/ .

# conntrack-parse
stupid tool that I use to make /proc/net/nf_conntrack more readable.

e.g.

```bash
podman run --rm -i localhost/conntrack < /proc/net/nf_conntrack
# or
go run ./cmd/cli < /proc/net/nf_conntrack
```

# NAT46 Multitool

Linux daemon to perform different types of NAT46 (IPv4 to IPv6) translation. Inspired by snid (https://github.com/AGWA/snid/).

Provides:

- HTTP backend (determine which host to forward to by DNS request based on 'Host' header in request)
- SNI backend (determine which host to forward to by DNS request based on TLS SNI)
- Full IP backend (forwards all ports on a certain IPv4 address to a mapped IPv6 address in the config)
# iplisttodns

Small Go utility to extract a DNS filter list from IP/domain JSON files. Get the list file from https://russia.iplist.opencck.org/

## What it does

- Always includes these matching TLDs:
  - `ru`
  - `рф`
  - `moscow`
  - `москва`
  - `рус`
  - `by`
  - `su`
- Reads all domain names from the `domains` arrays in the input JSON.
- For domains with a non-matching TLD, adds unique TLD+1 values (for example `vk.com`, `yandex.net`).
- Prints one value per line.

## Requirements

- Go 1.20+ (or any recent Go version)

## Usage

Run with flag:

```bash
go run extract_dns.go -input ip-list.json
```

Run with positional filename:

```bash
go run extract_dns.go ip-list.json
```

## Example output

```text
ru
рф
moscow
москва
рус
by
su
mycdn.me
vk.com
yandex.com
yandex.net
vkuser.net
```

## Makefile commands

Use:

```bash
make help
```

Common targets:

- `make run` - run using `INPUT` (default: `ip-list.json`)
- `make test` - run Go tests
- `make run-main` - run with `ip-list.json`
- `make run-main -B` - refresh `ip-list.json` and run with it
- `make build` - build binary to `bin/extract_dns`
- `make fmt` - format Go source files
- `make clean` - remove `bin/`

Override input file:

```bash
make run INPUT=ip-list.json
```

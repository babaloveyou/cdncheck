## cf-check
Check an IP is Owned by CDN .

## Install
```
▶ go get -u github.com/babaloveyou/cdncheck
```

## Usage
```
▶ echo "uber.com" | cf-check
```

The goal is that you don't need to do a port scan if it's proven that the IP is owned by Cloudflare.

```
▶ subfinder -silent -d uber.com | filter-resolved | cf-check | sort -u | naabu -silent -verify | httprobe
```

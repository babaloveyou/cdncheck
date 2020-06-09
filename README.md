## cdncheck
Check an IP is Owned by Cloudflare or Incapsula  or Sucurl or Akamai
				

## Install
```
▶ go get -u github.com/babaloveyou/cdncheck
```

## Usage
```
▶ echo "uber.com" | cdncheck
```

The goal is that you don't need to do a port scan if it's proven that the IP is owned by  Cloudflare or Incapsula  or Sucurl or Akamai.

```
▶ subfinder -silent -d uber.com | filter-resolved | cdncheck | sort -u | naabu -silent -verify | httprobe
```

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	concurrency := 20
	flag.IntVar(&concurrency, "c", 20, "Set the concurrency level")
	flag.Parse()
	jobs := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			for host := range jobs {
				addr, err := net.LookupIP(host)
				if err != nil {
					continue
				}
				if isCloudflare(addr[0]) {
					fmt.Println(addr[0], ":inCloudflare")
				} else if isIncapsula(addr[0]) {
					fmt.Println(addr[0], ":inIncapsula")
				} else if isSucuri(addr[0]) {
					fmt.Println(addr[0], ":inSucurl")
				} else if isAkamai(addr[0]) {
					fmt.Println(addr[0], ":inAkamai")
				} else {
					fmt.Println(addr[0], ":inNormal")
				}
			}
			wg.Done()
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		jobs <- sc.Text()
	}

	close(jobs)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	wg.Wait()
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips, nil

	default:
		return ips[1 : len(ips)-1], nil
	}
}

func isCloudflare(ip net.IP) bool {
	cidrs := []string{"173.245.48.0/20", "103.21.244.0/22", "103.22.200.0/22", "103.31.4.0/22", "141.101.64.0/18", "108.162.192.0/18", "190.93.240.0/20", "188.114.96.0/20", "197.234.240.0/22", "198.41.128.0/17", "162.158.0.0/15", "104.16.0.0/12", "172.64.0.0/13", "131.0.72.0/22"}
	for i := range cidrs {
		hosts, err := hosts(cidrs[i])
		if err != nil {
			continue
		}

		for _, host := range hosts {
			if host == ip.String() {
				return true
			}
		}
	}
	return false
}
func isIncapsula(ip net.IP) bool {
	cidrs := []string{"199.83.128.0/21", "198.143.32.0/19", "149.126.72.0/21", "103.28.248.0/22", "45.64.64.0/22", "185.11.124.0/22", "192.230.64.0/18", "107.154.0.0/16", "45.60.0.0/16", "45.223.0.0/16"}
	for i := range cidrs {
		hosts, err := hosts(cidrs[i])
		if err != nil {
			continue
		}

		for _, host := range hosts {
			if host == ip.String() {
				return true
			}
		}
	}
	return false
}

func isSucuri(ip net.IP) bool {
	cidrs := []string{"185.93.228.0/24", "185.93.229.0/24", "185.93.230.0/24", "185.93.231.0/24", "192.124.249.0/24", "192.161.0.0/24", "192.88.134.0/24", "192.88.135.0/24", "193.19.224.0/24", "193.19.225.0/24", "66.248.200.0/24", "66.248.201.0/24", "66.248.202.0/24", "66.248.203.0/24"}
	for i := range cidrs {
		hosts, err := hosts(cidrs[i])
		if err != nil {
			continue
		}

		for _, host := range hosts {
			if host == ip.String() {
				return true
			}
		}
	}
	return false
}

func isAkamai(ip net.IP) bool {
	cidrs := []string{"104.101.221.0/24", " 184.51.125.0/24", " 184.51.154.0/24", " 184.51.157.0/24", " 184.51.33.0/24", " 2.16.36.0/24", " 2.16.37.0/24", " 2.22.226.0/24", " 2.22.227.0/24", " 2.22.60.0/24", " 23.15.12.0/24", " 23.15.13.0/24", " 23.209.105.0/24", " 23.62.225.0/24", " 23.74.29.0/24", " 23.79.224.0/24", " 23.79.225.0/24", " 23.79.226.0/24", " 23.79.227.0/24", " 23.79.229.0/24", " 23.79.230.0/24", " 23.79.231.0/24", " 23.79.232.0/24", " 23.79.233.0/24", " 23.79.235.0/24", " 23.79.237.0/24", " 23.79.238.0/24", " 23.79.239.0/24", " 63.208.195.0/24", " 72.246.0.0/24", " 72.246.1.0/24", " 72.246.116.0/24", " 72.246.199.0/24", " 72.246.2.0/24", " 72.247.150.0/24", " 72.247.151.0/24", " 72.247.216.0/24", " 72.247.44.0/24", " 72.247.45.0/24", " 80.67.64.0/24", " 80.67.65.0/24", " 80.67.70.0/24", " 80.67.73.0/24", " 88.221.208.0/24", " 88.221.209.0/24", " 96.6.114.0/24"}
	for i := range cidrs {
		hosts, err := hosts(cidrs[i])
		if err != nil {
			continue
		}

		for _, host := range hosts {
			if host == ip.String() {
				return true
			}
		}
	}
	return false
}

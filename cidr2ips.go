// Converts a provided CIDR to a range of IPs printed to stdout for use by other tools
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"regexp"
	"time"
)

func main() {
	// define and set default command parameter flags
	var cidrFlag = flag.String("cidr", "", "\nRequired: CIDR block (ex: 192.168.0.0/16) to expand to a list of IP addresses\n")
	var randFlag = flag.Bool("randomize", false, "\nOptional: randomize the order of the IP addresses provided as output\n")
	var fullFlag = flag.Bool("full", false, "\nOptional: provide the network and broadcast addresses in the output; by default only usable addresses are included\n")
	var hFlag = flag.Bool("help", false, "\nPrint usage information\n")

	// usage function that's executed if a required flag is missing or user asks for help (-h)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %s [--randomize] --cidr <cidr-block>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nExample: %s --cidr 172.17.24.0/24 --randomize\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println()
	}
	flag.Parse()

	// provide help (-h)
	if *hFlag == true {
		flag.Usage()
		os.Exit(0)
	}

	// the --cidr flag is required
	if *cidrFlag == "" {
		fmt.Fprintf(os.Stderr, "\nMissing required --cidr argument\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// regex to match valid CIDR patterns
	r, _ := regexp.Compile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.)" +
		"{3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\\/" +
		"([0-9]|[1-2][0-9]|3[0-2]))$")

	// check to see if the provider CIDR is valid and if not provide error and exit
	match := r.MatchString(*cidrFlag)
	if !match {
		fmt.Fprintf(os.Stderr, "\ninvalid network CIDR: %s\n\n", *cidrFlag)
		flag.Usage()
		os.Exit(1)
	}

	//
	ips, err := hosts(*cidrFlag, *fullFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n\n")
	}

	// shuffle the order of IP addresses if requested (--random)
	if *randFlag == true {
		shuffleIPs(ips)
	}

	// print the IP addresses in the CIDR range to stdout
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// increment the IP address by 1
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// randomly shuffles slice of strings
func shuffleIPs(slice []string) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// hosts takes cidr string provided by --cidr flag and returns a slice of all IP addressees in the range
// minus the network number and broadcast address of the provided cidr
func hosts(cidr string, fullFlag bool) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	if fullFlag == true {
		// include network and broadcast addresses
		return ips, nil
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

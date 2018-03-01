# cidr2ips - CIDR to IP list 

This utility takes a network CIDR range (192.168.0.0/16) and expands it to a list of IP addresses that are printed to standard out, one IP per line. By default the IP addresses are printed sequentially, but if you would like to randomize the output you can use the --randomize flag.

## Usage

```
Usage: ./cidr2ips [--randomize] --cidr <cidr-block>

Example: ./cidr2ips --cidr 172.17.24.0/24 --randomize

  -cidr string
    
        Required: CIDR block (ex: 192.168.0.0/16) to expand to a list of IP addresses
    
  -help
    
        Print usage information
    
  -randomize
    
        Optional: randomize the order of the IP addresses provided as output
```


## Example

```
./cidr2ips --cidr 192.168.43.0/28 
192.168.43.1
192.168.43.2
192.168.43.3
192.168.43.4
192.168.43.5
192.168.43.6
192.168.43.7
192.168.43.8
192.168.43.9
192.168.43.10
192.168.43.11
192.168.43.12
192.168.43.13
192.168.43.14
```

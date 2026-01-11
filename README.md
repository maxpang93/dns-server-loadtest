# dns-server-loadtest
Exercise to optimize DNS server implementation

## Load Testing
1. git clone `dnsblast`: https://github.com/jedisct1/dnsblast
2. run `./dnsblast 127.0.0.1 <total queries> <queries per second> <dns port>`


## Test Results
```bash
./dnsblast 127.0.0.1 50000 1000 10080
Sent: [50000] - Received: [48152] - Reply rate: [979 pps] - Ratio: [96.30%]
```

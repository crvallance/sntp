sntp
====

### A implementation of NTP server with Golang
##### What this?
- Based on [RFC 2030](http://tools.ietf.org/html/rfc2030)
- Multiple equipments sync time from local
- Designed for multiple equipments which can't connect to internet and need synchronization time
- Compatible with [ntpdate](http://www.eecis.udel.edu/~mills/ntp/html/ntpdate.html) service on the linux
- NTP client is fork from [beevik](https://github.com/beevik/ntp/)，a better client
- Originally from http://www.btfak.com

#### Usage manual
##### 1. install Golang

Please reference  [Go install](https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/01.1.md) chapter of open-source Golang book "build-web-application-with-golang".

#### Building
* `go get github.com/crvallance/sntp`
* Export your Go path
* `go build -o sntpbin`
* Rename sntpbin as needed e.g. sntpbin.exe

#### Usage
* Server side:  
* `sntpbin.exe 8888` where 8888 is a port of your choosing
* Client side:  
* `NTP-time-set.sh 8888` with matching port

##### License
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).

//package main
//about: btfak.com
//create: 2013-9-25
//update: 2016-08-22

package main

import (
	"fmt"
	"github.com/btfak/sntp/netapp"
	"github.com/btfak/sntp/netevent"
	"os"
	"strconv"
)

func main() {
	var handler = netapp.GetHandler()

	if len(os.Args) == 2 {
		portarg := os.Args[1]
		port, err := strconv.Atoi(portarg)
		if err != nil {
			fmt.Printf("Couldn't parse `%s` as a port number\n", portarg)
			os.Exit(1)
		}
		netevent.Reactor.ListenUdp(port, handler)
		netevent.Reactor.Run()
	} else {
		fmt.Printf("Must specify port number: `%s` <portnumber>\n", os.Args[0])
	}
}

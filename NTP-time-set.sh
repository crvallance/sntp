#!/usr/bin/env bash
# This is exceptionally brute force and ugly.  YMMV
if [ -z "$1" ]
  then
    echo "No port supplied"
    exit 1
fi
HOST_IP=`sudo grep lease /var/lib/dhcp/dhcpd.leases | tail -1 | cut -d ' ' -f 2`
# echo $HOST_IP
sudo iptables -t nat -A OUTPUT -pudp --dport 123 -j DNAT --to-destination $HOST_IP:$1
sudo sntp -S $HOST_IP
sleep 5
sudo iptables -t nat -D OUTPUT -pudp --dport 123 -j DNAT --to-destination $HOST_IP:$1
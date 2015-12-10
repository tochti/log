+++
title = "Config Network"
date = "2015-11-25T09:07:27+01:00"
draft = false
+++

Allgemein
---------

~~~bash
# Verschieden Beispiele zur manuellen konfirugartion von Netzwerk-Interfaces
$ ip link show dev eth0
$ ip link set eth0 up
$ ip link set eth0 down
$ ip addr add 192.168.1.2/24 broadcast 192.168.1.255 dev eth0
$ ip route add default via 192.168.1.1

$ resolvconf -a eth0 < tmp_resolv.conf_file
~~~

 ~~~bash
# Scanne nach Wlan-Access-Poits
$ iw dev wlan0 scan
$ iwlist wlan0 scanning
~~~


Chromebook
----------

Deaktivier die wlan Verbindung über Chrome Einstellungen WLAN Deaktivieren

~~~bash
sudo wpa_supplicant -Wext -iwlan0 -c/home/chronos/user/$profil
sudo ifconfig wlan0 192.168.1.11/24 broadcast 192.168.1.255 netmask 255.255.255.0 up
sudo route add default gw 192.168.1.1
sudo vim /etc/resolv.conf
  |nameserver 192.168.1.1
  |nameserver 8.8.8.8

~~~

??? Aktivier die wlan Verbindung wieder über die Chrome Einsetllungen ansonsten lässt einen Chrome nicht ins Internet da er von irgendwo her die Information bekommt das keine Verbindung zu irgendeinem Netzwerk besteht.

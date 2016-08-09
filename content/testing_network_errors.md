* Verwerfe Zufällig IP Packages - iptables -A INPUT -s <SOURCE_IP_ADDRESS> -m statistic --mode random --probability 0.1 -j DROP
* https://wiki.linuxfoundation.org/networking/netem

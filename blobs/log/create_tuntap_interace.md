ip tuntap add name tap0 mode tap
ip link set up dev tap0

Beachte, das Tun/Tap Device wird erst aktiviert (UP) wenn ein Programm sich mit dem Device verbunden hat. Damit ist nicht gemeint das ein Programm aut der IP Adresse oder ähnliche lauscht.

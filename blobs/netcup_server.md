# netcup server #

Login als root und erzeuge neuen Benutzer

```
useradd -m tochti
passwd tochti
usermod -a -G wheel
```

Schalte Gruppe wheel in sudoes frei.

Übertrage public key zu server ssh-copy-id -i .ssh/$key tochti@$host

Bearbeite sshd config so das root sich nicht mehr anmelden. Weiter ist die Anmeldung nur noch über Zertifikat möglich.

```
AllowUsers tochti
PermitRootLogin no
PasswordAuthentication no
ChallengeResponseAuthentication no
```

```
systemctl restart sshd
```

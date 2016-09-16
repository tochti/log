
# OpenSSL #

## Zertifikat mit SNI Feldern ##

SNI: Server Name Indication ist eine Erweiterung des x.509 Standards. Wann benötigt man diese Felder? Benutzt man zum Beispiel den golang http-client und baut eine Verbindung zu einem https Server auf überprüft ob der im Feld subjectAltName angegebene Wert mit dem host in der URL übereinstimmt.

Bekommt man von einem Programm denn Error "Failed to tls handshake with x.x.x.x x509: cannot validate certificate for x.x.x.x because it doesn't contain any IP SANs" gezeigt heißt in diesem Fall das das Zertifikat keine subjectAltName hinterlegt hat. 

Es gibt zwei Felder in dennen der Host hinterlegt sein sollte CommanName und in subjectAltName.

Gibt eine IP Adresse im subjectAltName an hat diese das folgende Format IP:XXX.XXX.XXX.XXX möchte man eine Domain angeben dann benutzt man DNS:domain.com es gibt noch weiter möglichkeiten z. Bsp. um mehrer IP-Adressen oder Domains anzugeben.


Beim erzugen der CSR und des Zertifikats muss nun die Konfigurations Datei angegeben werden ebenso muss openssl mitgeteilt werden das eine extension verwendet werden soll

Wir sind nun wieder bei einem Beispiel angekommen das heißt es ist wieder Elliot time.

Elliot gönnt seinem Blog ein Zertifikat mit SNI Feldern damit auch der Panonidste http client das Zertifikat akzeptiert.

Dazu benötigen wir folgende Konfigurations Datei (openssl.cnf)

```
[ req ]
default_bits       = 4096
default_md         = sha512
prompt             = no

distinguished_name = req_distinguished_name

req_extensions     = v3_req

[ req_distinguished_name ]
countryName            = "USA"
stateOrProvinceName    = "New York"
localityName           = "Brooklyn"
postalCode             = ""
streetAddress          = "Straße3027 West 12th Street "
organizationName       = "fsociety"
organizationalUnitName = "Dark Army"
commonName             = "r1ng0.311iot.com"
emailAddress           = "e@corp.com"  

[ v3_req ]
subjectAltName  = DNS:r1ng0.3lliot.com
```

Erstelle neuen Certificate-Sign-Request, blog-server.key.pem muss Private- und Public-Key enthalten.
```bash
$ openssl req -new -x509 -extensions v3_req -key blog-server.key.pem -out blog-server.csr -config openssl.cnf
```

Möchte man nochmals sichergehen das auch alle Daten im CSR enthalten sind kann der folgende Befehl abhilfe bringen
```bash
$ openssl req -text -in blog-server.key.pem
[...]
        Attributes:
        Requested Extensions:
            X509v3 Subject Alternative Name: 
                DNS:r1ng0.3lliot.com
[...]
```
Findet sich in der Ausgabe X509v3 Subject Alternative Name passt alles.

Die Mr. Robot CA kann den CSR nun signieren, natrülich nach dem Sie überprüft hat ob Elliot auch der Rechtmäßige Besitzer des Servers und der Domain ist.
```bash
$ openssl x509 -CAkey mrrobot-ca-private.key.pem -days 365 -in blog-server.csr -out blog-server.cert.pem
```

Auch hier kann man überprüfen ob die benötigen Felder vorhanden sind.
```bash
$ openssl x509 -text -in blog-server.cert.pem
[...]
        X509v3 extensions:
            X509v3 Subject Alternative Name: 
                DNS:r1ng0.3lliot.com
[...]
```
Done!

### Quellen ###
* https://www.prshanmu.com/2009/03/generating-ssl-certificates-with-x509v3-extensions.html
* https://en.wikipedia.org/wiki/Server_Name_Indication
* http://wiki.cacert.org/FAQ/subjectAltName
* https://www.openssl.org/docs/manmaster/apps/x509v3_config.html




## Prüfe ob eine CA der Herausgeber eines Zertifikats ist ##

```bash
$ openssl verify -verbose -CAfile cacert.pem  server.crt
```

## asn.1 ##
Es kommt machmal vor das in RFC der Begriff asn.1 auftaucht. Diese ist eine Beschreibungsprache um Daten Strukturen zu beschreiben. Steht aber sonst in keinem Zusammenhang mit Verschlüsslung/Signieren/TLS

https://de.wikipedia.org/wiki/Abstract_Syntax_Notation_One





# OpenSSL #

## Hinweise ##

Erzeugt ein Private-Key und ein Public-Key beide gespeichert in mykey.pem
```bash
openssl genrsa -out mykey.pem 1024
```

Liste den Public-Key aus mykey.pub dieser wird nicht nachträglich erzeugt das ist natürlich gegen die Idee des Verfahren. Es sind beide Schlüssel in der Datei mykey.pem
```bash
openssl rsa -in mykey.pem -pubout > mykey.pub
```

## Übersichten zu OpenSSL Befehlen ##
* https://www.digitalocean.com/community/tutorials/openssl-essentials-working-with-ssl-certificates-private-keys-and-csrs
* http://ngs.ac.uk/ukca/certificates/advanced


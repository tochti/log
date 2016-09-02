# Public-Key-Kryptoverfahren #

Datum: 30-08-20161

## Einleitung ##

### Warum schreibe ich diesen Text ###
Es ging mir nicht darum einzelnen Verfahren Technisch zu Beschreiben oder gar erschöfpend zu Beschreiben dazu fehlt mir bei weitem das Wissen. Es ging darum eine Übersicht über die Zusammenhänge zwischen Verfahren und Funktionen zu geben ebenso aufzuzeigen warum diese unterschiedliche Verfahren und Funktionen existieren. Weiter möchte ich anderen die Denkfehler ersparen wie ich sie gemacht. Die ganze Reise began damit das ich nicht verstanden habe warum man zum Erzeugen eines CSR (Certificate Signing Request) ein Private Key benötigt und woher der Public Key darin plötzlich herkamm obwohl ich diesen doch nie explizit erzeugt habe.

Der Text richtet sich an Anfänger die einen Übersicht über die Zusammenhänge der Themen Verschlüsslung/Signierung und TLS bekommen möchten. Da die Kapitel aufeinander aufbauen ist es zu empfehlen das Dokument von oben nach unten zu lesen.

## Einführung ##

Wie so oft in dieser Welt gab es eine Aufgabe welche gelöst werden mochte. Also lasst uns beginnen mit dieser Aufgabe. 
```
Wie können Informationen übertragen werden so das diese nur von Berechtigten Personen gelesen werden können.
```

Wie immer im Leben geht nichts über ein gutes Beispiel an dem man sich abarbeiten kann. Wählen wir zwei Protagonisten. Nennen wir Sie ganz zufällig 

Elliot und Darlene

Es könnte nun sein das diese zwei Protagonisten unteranderem von einer sehr mächtigen Chinesischen Hacker Gruppe, nennen wir Sie mal Dark Army, verfolgt werden. Es könnte nun äußerste Überlebenswichtig sein das die Nachrichten die zwischen Eliot und Darlenne ausgetauscht werden geheim bleiben um dies zu erreichen gibt es nun unterschiedliche Verfahren.

Unsere neuen Freunde werden uns, in den nächste Kapitel, eine kurze Einsicht gewähren warum das Public-Key-Kryptoverfahren entwickelt wurde. Es wird nicht weiter darauf eingegangen wie wichtig dies Erfindung für unsere moderne Welt und die Demokratie ist.

### Kapitel 1 - Symmetrische Verschlüsslung ###

Sagen wir Elliot und Darlenne wollen sich treffen um gemütlich bei Kaffee und Kuchen über die Weltherrschaft sprechen es ist nur etwas störend wenn man dabei von einer Dark Army beschossen wird. Die Lösung die Information zu dem Treffpunkt muss gut verschlüsselt werden. Zum Glück haben Elliot und Darlene bei ihrem letzen Treffen eine ziemlich ziemlich große Zahl ausgetauscht dies können beide nun verwenden um ihr Treffen zu organisieren.

Dazu ein Beispiel aus der Praxis
```bash
$ echo "23:00Uhr 3027 W 12th St, Brooklyn, NY 11224, USA" > meetingpoint.txt
$ openssl aes-256-cbc -e -a -in meetingpoint.txt -out meetingpoint.txt.aes // -a erzeugt eine base64 Ausgabe
enter aes-256-cbc encryption password:
Verifying - enter aes-256-cbc encryption password:
```
Alternative könnte man Programme wie gpg, pgp und aescrypt verwendet werden.

Die Datei meetingpoint.txt.aes kann nun an Darlene übermittelt werden. Diese kann die Nachricht mit dem bekannten Passwort wieder entschlüsseln.

```bash
$ openssl aes-256-cbd -d -a -in meetingpoint.txt.aes 
enter aes-256-cbc decryption password:
```

Nachdem das Treffen vom FBI unfreundlich unterbrochen wurde und Elliot seinen Laptop liegen lassen musst. Gehen beide davon aus das der Geheimeschlüssel nicht mehr sicher ist. Auf der Flucht konnten die beiden jedoch keinen neuen Schlüssel vereinbaren. Wie kann nun die wichtige Informationen zwischen den beiden ausgetauscht werden? Dazu mehr im Kapitel 2 zuvor noch ein paar Erleuterungen und weiter Informationen.

Was können wir aus dem Beispiel lernen.

* Geheimerschlüssel muss beiden Parteien bekannt sein

#### Funktionen ####

E(sk, m) = c
D(sk, c) = m

```golang
funny
```

#### Vor-/Nachteile ####

* Schnellere Ver-/Entschlüsslung als bei Public-Key
* Höhere Sicherheit bei gleicher Schlüsselänge als bei Public-Key
* Geheimschlüssel kann nicht über unsichere Kommunikationsweg vereinbart werden.
* Bekommt der Angreifer den Geheimschlüssel kann die gesamte Kommunikation gelesen werden.

Das Thema Verschlüsslung ist sehr Komplex und enthält extrem viele Fallen.

#### Quellen ####
* Übersicht was die einzelnen Verfahren bei AES bedeuten - http://stackoverflow.com/questions/1220751/how-to-choose-an-aes-encryption-mode-cbc-ecb-ctr-ocb-cfb
* https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Electronic_codebook_.28ECB.29



### Kapitel 2 - Asymmetrische Verschlüsslung ###

Die Panik ist groß es gibt wichtige Informationen Elliot betreffend die Dark Army ist ihm auf der Spur Darlene muss diese Elliot unbedingt mitteilen ohne das die Dark Army etwas davon erfährt. Jedoch gibt es keinen Geheimenschlüsseln mehr um Nachrichten zu verschlüsseln. Auch kann kein neuer Schlüssel über einen sicheren Kommunikationskannal ausgetauscht werden. Da davon auszugehen ist das die Dark Army sämtlichen Netzwerkverkehr von Elliot mitliest.

Auftritt Asymmetrische Verschlüsslung. Hier gibt es einen Bekannten Schlüssel (Public-Key) und einen Geheimenschlüssel (Private-Key). Es gilt das der Private-Key der einzigste ist der Nachrichten entschlüsseln kann die vom Public-Key erzeugt wurden das gleiche gilt auch für den vertauschten Fall bei dem der Public-Key der einzigste ist der Nachrichten entschlüsseln kann die vom Private-Key verschlüsselt wurden. 

Darlene erzeugt einen Schlüsselpaar (Privat-Key/Public-Key) dazu verwendet Sie RSA (später dazu mehr) die beiden Schlüssel werden im PEM Format abgelegt und BASE64 Encoded
```bash
$ openssl genrsa -out darlene.pem 4096
```

Das selbe macht Elliot.
```bash
$ openssl genrsa -out elliot.pem 4096
```

Elliot ebenso wie Darlene stellen ihren Public-Key auf einer Webseite/Forum/etc. bereit.
```bash
$ openssl rsa -in elliot.pem -pubout -out elliot-public-key.pem
$ openssl rsa -in darlene.pem -pubout -out darlene-public-key.pem
```

Darlene verschlüsselt nun die Nachricht an Elliot mit dessen Public-Key.
```bash
$ echo "Dark Army is watching you. Send me your Adresse that we can meet us" > msg.txt
$ openssl rsautl -inkey elliot-public-key.pem -pubin -encrypt -in msg.txt -out msg.txt.encrypt
```
Denn Inhalt der msg.txt.encrypt stellt Darlene der Weltöffentlichkeit bereit da nur Elliot diese Nachricht wieder entschlüsseln kann stellt das kein Problem dar.

Um die neusten Ereignise zu Erfahren kann Elliot die Nachricht mit seinem Privaten Key wieder entschlüsseln.
```bash
$ openssl rsautl -inkey elliot.pem -decrypt -in msg.txt.encrypt
```

Elliot kann nun auf die selbe Art Darlene Antworten
```bash
$ echo "New York City USA" > address.txt
$ openssl rsautl -inkey darlene-public-key.pem -pubin -encrypt -in address.txt -out address.txt.encrypt
```

Das Problem ist hier nur. Das Elliot nicht Elliot ist stattdessen steckt hinter dem Public-Key die Dark Army. Zack, das wars, das Böse hat gewonnen und das wollen wir doch alle nicht! Da muss es doch eine alternative geben? Richtig Zertifikate. Weiter im Kapitel 3.

#### Funktionen ####
E(public-key, msg) = c
D(private-key, c) = msg

#### Quellen ####
* Verschlüsseln von Daten - https://www.devco.net/archives/2006/02/13/public_-_private_key_encryption_using_openssl.php



### Kapitel 3 - Zertifikate ###

Jetzt da der Sieg des Bösen kurz bevor steht, hoffen wir auf die Rettung durch Zertifikate.

#### Signatur / Signieren ####

Bevor wir über Zertifikate sprechen wollen müssen wir einen kurzen Exkurse über das Signieren eines Dokuments abhalten.

Setze ich im analogen Leben eine Unterschrift unter ein Dokument zeige ich so das ich damit einverstanden bin was in dem Dokument geschrieben ist und was die Vorrausetzung für den Teil 1 diese Satz ist, ich es auch gelesen habe. Wenn das Dokument von größere Bedeutung ist kann es auch vorkommen das ein Notar ebenfalls seine Unterschrift darunter setzen muss. Wenn ein Vertrag zwischen zwei Parteien geschlossen wird unterzeichnen beide Parteien den Vertrag und geben so Ihr Einverständnis zu dem Geschrieben. Für die Unterschrift im analogen Leben gibt es eine Gegenpart in der digitalen Welt.

Die Idee Unterscheidet sich nicht sehr von der des Verschlüsselns einer Datei nur das hier der Weg andersherum ist.

Aber lass uns das ganze einmal wieder abarbeiten unter Zuhilfenahme eines Beispiels. Kommen wir somit zurück zu unseren 2 Protagonisten Elliot und Darlene.

Die Beiden haben Vereinbart die Welt zu Teilen genau 50 - 50 sollte der Weltherrschaftsplan erfolgreich verlaufen. Da Elliot aber eine eher gespalten Persönlichkeit hat und daher sehr Sprunghaft ist muss dies Vertraglich geregelt werden.

```bash
$ cat vertrag.txt
Die Vertragsparteien Elliot und Darlene stimmen hiermit zu das die Welt zu zwei gleichen Teilen unter Ihnen aufgeteilt werden soll wenn der Plan zur Übernahme der Weltherrschaft erfolgreich verläuft.
```

Um die Unterschrift zu leisten wird wie zuvor ein Private-Key und ein Public-Key beider Personen benötigt

```bash
$ openssl genrsa -out elliot.pem 4096
$ openssl genrsa -out darlene.pem 4096
```

Nun Unterzeichnet Elliot ebenso wie Darlene den Vertrag mit ihrem Private-Key
```bash
$ openssl dgst -sign elliot.pem -out vertrag.elliot.sign vertrag.txt
$ openssl dgst -sign darlene.pem -out vertrag.darlene.sign vertrag.txt
```

Damit Darlene, nach getaner Arbeit, anspruch auf ihre hälfte hat und zeigen kann das Elliot tatsächlich die zweite Person war welche den Vertrag unterschrieben hat muss sie Kopien des vertrag.txt, vertrag.elliot.sign und den Public-Key von Elliot an sich nehmen.

Die Welt ist unterworfen, der Tag nach der Siegesfeier bricht an. Darlene ist Geschäftstüchtig und erhebt sogleich, in Katrigerstimmung, anspruch auf ihren Anteil und zeigt sogleich das sie Unterzeichnerin des Vertrags ist.

Dazu verwendet sie ihren Priavte-Key und ihre Signatur. Sie führt die Signatur über vertrag.txt nochmals durch und vergleicht die Ausgabe mit der damals erstellten Signatur da vertrag.txt noch immer der selbe ist sollte die Ausgabe ebenfalls die selbe sein.
```bash
$ openssl dgst -preverify darlene.pem -signature vertrag.darlene.sign vertrag.txt
Verified OK
```

Elliot, noch ordentlich Verkatert, ist fest davon überzeugt das er noch nie etwas von diesem Vertrag gehört hätte. Abwarten denkt sich Darlene und legt folgende Befehle vor

Unter verwendung des Public-Key von elliot und dessen Signatur kann gezeigt werden das auch er an dem Vertrag beteiltigt war. Die Signatur in vertag.elliot.sign wird entschlüsselt und die Ausgabe muss mit dem Inahlt der Datei vertrag.txt übereinstimmen.
```bash
$ openssl dgst -verify elliot-public-key.pem -signature vertrag.elliot.sign vertrag.txt
Verified OK
```

Aus mathematischen Gründen wird die Signatur mindesten Doppelt so groß wie die Daten die Verschlüsselt wurden das kann ziemilich ausufern bei größeren Dateien. Daher wird meistens ein Hase der Datei erstellt und dieser wird dan signiert. Das spart nicht nur Platz sonder hat auch noch den Vorteil das mit großer Wahrscheinlichkeit die Daten welche Signiert wurde nicht verändert wurden.

Am Ende des Tages ist das Signieren eines Dokumentes nichts anders als das erzeugen einer Kopie eben diese Dokuments und die verschslüsslunge des selbigen durch ein Private-Key. In der Praxis wird das Dokument nicht kopiert stattdessen wird ein Hash des Dokuments erzeugt und dieser wird durch den Private-Key Verschlüsselt. Die Signaturen und das Dokument können auch in einer Datei liegen welche auch nochmals Verschlüsselt ist oder in base64 Format vorliegt. Es gibt wie so oft in der Kryptowelt unterschiedlichste Kombinationen und Standards

Mit dem folgendem Befehl erstellen wir eine Signatur über den SHA-512 Hash der Datei vertrag.txt. Die Signatur wird in der Datei vertrag.darlene.sha512.sign gespeichert.
```bash
$ openssl dgst -sign darlene.pem -sha512 -out vertrag.darlene.sha512.sign vertrag.txt
Verified OK
```

Möchte man zeigen das Darlene die den Vertrag unterzeichnet hat benötigt man hierfür den Public-Key von Darlene. Es muss mitangegeben werden welches hashing Verfahren zuvor benutzt wurde
```bash
$ openssl dgst -verify darlene-public-key.pem -sha512 -signature vertrag.darlene.sha512.sign vertrag.txt
Verified OK
```

Am Ende steht nun noch die Frage offen woher weiß man das der Public-Key der zum Verifizieren benutzt wurde überhaupt darlene gehört. Es wäre möglich das jemand den Vertrag mit einem anderen Schlüssel-Paar unterschrieben hat und daraufhin behauptet Darelene hätte den Vertrag unterschrieben. Das gleiche Problem hatten wir auch schon zuvor mit Elliot dieser hätte ebenfalls behaupten können das der Key welcher zum überprüfen verwendet wurde überhaupt nicht sein Public-Key gewesen sei.

Das führt uns zu zwei neuen Begriff dem Zertifikat und der Certificate Authority (abkz. CA)

TODO(tim):Hier noch ein Ablaufdiagram wie das Signieren funktioniert

##### Quellen #####
* http://security.stackexchange.com/questions/57336/certificate-request-why-does-the-requester-have-to-create-a-private-key
* Signieren von Dokumenten - https://raymii.org/s/tutorials/Sign_and_verify_text_files_to_public_keys_via_the_OpenSSL_Command_Line.html


#### Zertifikate ####

Bevor wir uns um die Begriffe Zertifikat und Certificate Authority widmen gibt es noch immer das Problem das das Böse kurz davor steht zu gewinnen. Also lass es uns endlich bekämpfen.

Wie hätte Darlene zum einen Sicherstellen können ob der Public-Key Elliot gehört und wie hätte sie dafür sorgen können das nur der richtige Elliot die verschlüsselten Daten entschlüsseln kann und so die Standort Daten nicht in die falschen Hände gelangen.

Elliot schreibt gerne Tagebuch (er muss wie man hört irgendwelche Probleme mit seinem Vater verarbeiten, wie man hört). Elliot ist ein verdamter Cyberpunk daher macht er das natürlich online und betreibt ein Blog dazu betreibt er einen HTTPS Server. Das heißt die Kommunikation findet Verschlüsselt statt und ebenso liegt auf dem Server ein Zertifikat (was das nun ist wird sich noch erschließen keine Angst). Da Darlene ebenfalls ziemlich Cyber ist weiß sie das zu nutzen und besorgt sich das Zertifikat des Servers.

Dazu verbinden wir uns mit dem Server von Elliot und Fragen nach dem Zertifikat
```bash
openssl s_client -showcerts -connect https://r1ng0.3ll1ot.com:443 </dev/null
```

Beachte das durch eine Man-In-The-Middle-Attack ein falsche Certificate eingeschmuckelt werden kann.
Es gibt immer noch ein Certificate in der Kette das man überprüfen kann doch irgendwann ist man oben angekommen diese Zertifkat muss irgend wie zu einem kommen ohne das es Manipuliert wurde im Fall von Browser fest ein gebaut.

Public-Key aus dem Zertifkat des HTTPS Servers von Elliot und verschlüsseln die Daten darauß Prüfen das es von einer CA unterzeichet wurde.

Definition:
```
Data = [Subject-Informations, Public-Key]
cert = Sign(Private-Key, Data)
```

Was will diese Definition einem Mitteilen?
Zu erste einmal kann darauß ablesen das es zwei Parteien gibt ein Unterzeichner, zwei Schlüssel den Private-Key des Unterzeichners und den Public-Key für das Etwas dem die Informationen zugeordnet sind (diese kann z.Bsp. sein eine Person oder ein HTTPS Server) und Informationen (z. Bsp. der Name, die Anschrift oder eine IP Adresse oder eine Domain).

Ein Zertifikat ist ein eine Kombination von Informationen zu einem zu Idendifizierendem Etwas und desen Public-Key. Diese Kombination wird von einer zweiten Partie als korrekt Bestätigt. Das zeigt Sie der Welt in dem Sie eine Digital Unterschrift unter diese Kombination setzt. Da durch kann man sicherstellen wenn etwas mit diesem Public-Key verschlüsselt kann diese nur wieder von dem Etwas entschlüsselt werden. Ob das nun eine Person ist oder ein HTTPS Server spielt dabei keine Rolle.

Es beweist das ein Öffentlicher Schlüssel zu einem Benutzer gehört.

##### Certificate Sign Request (CSR) #####

Mögliche Formate
* PKCS#10 https://tools.ietf.org/html/rfc2986
* PEM https://tools.ietf.org/html/rfc1424

Unterschiede:
Note 4 - This document is not compatible with the certification
   request syntax for Privacy-Enhanced Mail, as described in RFC 1424
   [5].  The syntax here differs in three respects: It allows a set of
   attributes; it does not include issuer name, serial number, or
   validity period; and it does not require an "innocuous" message to be
   signed.  This document is designed to minimize request size, an
   important feature for certification authorities accepting requests on
   paper.


###### Quelle ######

* https://en.wikipedia.org/wiki/Certificate_signing_request#Structure
* http://stackoverflow.com/questions/21297139/how-do-you-sign-certificate-signing-request-with-your-certification-authority
* https://tools.ietf.org/html/rfc2986



##### Quellen #####

* http://security.stackexchange.com/questions/48802/how-to-validate-a-client-certificate
* Was ist ein Zertifikat - ftp://ftp.pgpi.org/pub/pgp/6.5/docs/german/IntroToCrypto.pdf



### Kapitel 4 - Diffe-Hellmann-Schlüsselaustausch ###

Einweg Public-Keys 

```Latex
b^x mod p
```

#### Quellen ####
https://de.wikipedia.org/wiki/Diffie-Hellman-Schl%C3%BCsselaustausch

### Kapitel 5 - Das Ultimum Komination a'la Certificate ###




## RSA ##
* https://www.emc.com/collateral/white-papers/h11300-pkcs-1v2-2-rsa-cryptography-standard-wp.pdf
* https://tools.ietf.org/html/rfc3447






# Sonstiges #

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





# TLS #

RFC2246: https://tools.ietf.org/html/rfc2246

## RSA ##

Ansynchrones Verfahren für Schlüsselautausch

## Verbindungs aufbau ##

Client sendet eine Liste von chiper suits die er beherrscht
Server sendet das vom ihm ausgewählte Verfahren an den Client + x.509 Zertifikat. Das erste Zertifikat das übermittelt wird sollte das Server Zertifikat sein zu diesem sollte ein Privat Key auf dem Server existieren

Public Key der CA entpacke Server Zertifikat prüfe Commen Name ob passt.

## CCA (Client Certificate Authentication) ##

chiper suite TLS

TLS_RSA_WITH_3DES_EDE_CBC_SHA

Schlüsselaustausch, bspw.: RSA, DH (auch ADH, ECDH), PSK, SRP
Authentifizierung, bspw.: RSA, DSA (auch ECDSA), PSK
Hashfunktion (MD5, SHA)
Verschlüsselung (keine, RC4, DES, 3DES, IDEA, AES)


es wird findet zu erste ein schlüsselautausch statt um darüber können dinge verschlüsselt werden was diese dient dazu einen schlüssel austausch für ein symetisches verschlüsslung. Hashfunktionen werden verwednet um sicherzustellen das nicht verändert wurde.

Der Hash algorithmus wird verwendet um eine hash über alle gesendetn Init Message zu erzeugen 

Wir CCA verwendet wird der pre-master-key (momentan gehe ich davon aus das das das errechnete Geheimnis ist welches im Diffe-Hellman-Schlüsselaustausch erzeugt wird) mit dem auf dem server hinterlegten public-key unterschrieben und an den client gesendet. 

Momentan gehe ich davon aus das nur Client Anfragen erlaubt sind welche auf dem Server hinterlegt sind das heißt es gibt ein abgleich der subject informationen des certificates.

## Protokoll Aufbau ##

Zwei Schichten

Record Protocol

## Quellen ##
[1] https://eprint.iacr.org/2013/538.pdf





# Stichwörter #
PKCS Public Key Cryptography Standards - https://en.wikipedia.org/wiki/PKCS





# Quellen #
[1] https://eprint.iacr.org/2013/538.pdf

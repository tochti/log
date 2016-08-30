# Public-Key-Kryptoverfahren #

Datum: 30-08-20161

## Einleitung ##

### Warum schreibe ich diesen Text ###
Es ging mir nicht darum einzelnen Verfahren Technisch zu Beschreiben oder gar erschöfpend zu Beschreiben dazu fehlt mir bei weitem das Wissen. Es ging darum eine Übersicht über die Zusammenhänge zwischen Verfahren und Funktionen zu geben ebenso aufzuzeigen warum diese unterschiedliche Verfahren und Funktionen existieren. Weiter möchte ich anderen die Denkfehler ersparen wie ich sie gemacht. Die ganze Reise began damit das ich nicht verstanden habe warum man zum Erzeugen eines CSR (Certificate Signing Request) ein Private Key benötigt und woher der Public Key darin plötzlich herkamm obwohl ich diesen doch nie explizit erzeugt habe.

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

Die Welt ist unterworfen, der Tag nach Siegesfeier bricht an. Darlene ist Geschäftstüchtig und erhebt sogleich in Katrigerstimmung anspruch auf ihren Teil und zeigt sogleich das Unterzeichnerin des Vertrags ist.

```bash
$ openssl dgst -preverify darlene.pem -signature vertrag.darlene.sign vertrag.txt
Verified OK
```

Elliot, noch ordentlich Verkatert, ist fest davon überzeugt das er noch nie etwas von diesem Vertrag gehört hätte. Abwarten denkt sich Darlene und legt folgende Befehle vor

```bash
$ openssl dgst -verify elliot-public-key.pem -signature vertrag.elliot.sign vertrag.txt
Verified OK
```

Hier haben wir ein Verbindung

##### CA #####

Aus mathematischen Gründen wird die Signatur mindesten Doppelt so groß wie die Daten die Verschlüsselt wurden da diese zu bei großen Daten zu noch größeren Daten führt erstellt man die Signatur meist über den Hash der Daten. 

Am Ende des Tages ist das Signieren eines Dokumentes nichts anders als das man eine kopie eines Dokuments anlegt und dieses Verschlüsselt. In der Praxis wird das Dokument nicht kopiert stattdessen wird ein Hash des Dokuments erzeugt und dieser wird durch den Private-Key Verschlüsselt. Die Signaturen und das Dokument können auch in einer Datei liegen welche auch nochmals Verschlüsselt ist oder in base64 Format vorliegt. Es gibt wie so oft in der Kryptowelt unterschiedlichste Kombinationen und Standards


* Hier ein Beispiel Ablauf *

* Hier noch ein Ablaufdiagram wie das Signieren funktioniert *

##### Quellen #####
* http://security.stackexchange.com/questions/57336/certificate-request-why-does-the-requester-have-to-create-a-private-key
* Signieren von Dokumenten - https://raymii.org/s/tutorials/Sign_and_verify_text_files_to_public_keys_via_the_OpenSSL_Command_Line.html

#### Zertifikate ####

Noch immer ist das Böse kurz davor zu gewinnen. Also lasse es uns endlich bekämpfen.

Gehen wir hier zu erste eine Alternative zu der Verwendeten Methode aus dem Kapitel zuvor durch.

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

## Zertifikate ##


## RSA ##
* https://www.emc.com/collateral/white-papers/h11300-pkcs-1v2-2-rsa-cryptography-standard-wp.pdf
* https://tools.ietf.org/html/rfc3447

## Stichwörter ##
PKCS Public Key Cryptography Standards - https://en.wikipedia.org/wiki/PKCS

## Beachte ##

Erzeugt ein Private-Key und ein Public-Key beide gespeichert in mykey.pem
```bash
openssl genrsa -out mykey.pem 1024
```

Liste den Public-Key aus mykey.pub dieser wird nicht nachträglich erzeugt das ist natürlich gegen die Idee des Verfahren. Es sind beide Schlüssel in der Datei mykey.pem
```bash
openssl rsa -in mykey.pem -pubout > mykey.pub
```

## Sonstiges ##

### asn.1 ###
Es kommt machmal vor das in RFC der Begriff asn.1 auftaucht. Diese ist eine Beschreibungsprache um Daten Strukturen zu beschreiben.

https://de.wikipedia.org/wiki/Abstract_Syntax_Notation_One


### OpenSSL ###

## Häufig verwendet Befehle ##

* https://www.digitalocean.com/community/tutorials/openssl-essentials-working-with-ssl-certificates-private-keys-and-csrs

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


## Protokoll Aufbau ##

Zwei Schichten

Record Protocol


## Quellen #
[1] https://eprint.iacr.org/2013/538.pdf

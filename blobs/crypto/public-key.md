+++
title = "Crypto"
date = 2016-09-08
type = "post"
categories = ["crypto"]
+++
# Public-Key-Kryptoverfahren #

Datum: 30-08-20161

## Einleitung ##

### Warum schreibe ich diesen Text ###
Es ging mir nicht darum einzelnen Verfahren Technisch zu Beschreiben oder gar erschöfpend zu Beschreiben dazu fehlt mir bei weitem das Wissen. Es ging darum eine Übersicht über die Zusammenhänge zwischen Verfahren und Funktionen zu geben ebenso aufzuzeigen warum diese unterschiedliche Verfahren und Funktionen existieren. Weiter möchte ich anderen die Denkfehler ersparen wie ich sie gemacht. Die ganze Reise began damit das ich nicht verstanden habe warum man zum Erzeugen eines CSR (Certificate Signing Request) ein Private Key benötigt und woher der Public Key darin plötzlich herkamm obwohl ich diesen doch nie explizit erzeugt habe.

Der Text richtet sich an Anfänger die einen Übersicht über die Zusammenhänge der Themen Verschlüsslung/Signierung und TLS bekommen möchten. Da die Kapitel aufeinander aufbauen ist es zu empfehlen das Dokument von oben nach unten zu lesen. Das Ziel ist es eine Basis-Verständnis zu erzeugen damit weiter Vorgänge daraus abgeleitet werden können.

Damit die Beispiele nachvollziehen kann benötigt man openssl und muss wissen wie man eine Konsole öffnet und darin Befehle ausführt.

## Einführung ##

Wie so oft ist der Anfang eine "Herausforderung" welche gelöst werden möchte. Die Herausforderung war in diesem Fall
 
```
Wie können Informationen übertragen werden so das diese nur von einer berechtigten Person gelesen werden kann. Die Informationen werden dabei über einen öffentlichen Weg übermittelt. D.h. die Nachricht kann jeder einsehen die Informationen dürfen aber nicht drauß ersichtlich sein.
```

Wie man diese "Herausforderung" lösen kann. Darauf gehen wir in dem restlichen Text ein.

Wie immer im Leben geht nichts über ein gutes Beispiel an dem man sich abarbeiten kann. Wählen wir zwei Protagonisten. Nennen wir Sie ganz zufällig 

Elliot und Darlene

Es könnte nun sein das diese zwei Protagonisten, unteranderem von einer sehr mächtigen Chinesischen Hacker Gruppe, nennen wir Sie mal Dark Army, verfolgt werden. Es könnte nun äußerste Überlebenswichtig sein das die Nachrichten die zwischen Elliot und Darlenne ausgetauscht werden geheim bleiben um dies zu erreichen gibt es nun unterschiedliche Verfahren.

Unsere neuen Freunde werden uns, in den nächste Kapitel, eine kurze Einsicht gewähren warum das Public-Key-Kryptoverfahren entwickelt wurde. Es wird nicht weiter darauf eingegangen wie wichtig dies Erfindung für unsere moderne Welt und die Demokratie ist.

### Kapitel 1 - Symmetrische Verschlüsslung ###

Sagen wir Elliot und Darlenne wollen sich treffen um gemütlich bei Kaffee und Kuchen über die Weltherrschaft zu sprechen es ist nur etwas störend wenn man dabei von einer Dark Army beschossen wird. Die Lösung die Information zu dem Treffpunkt muss gut verschlüsselt werden. Zum Glück haben Elliot und Darlene bei ihrem letzen Treffen eine ziemlich ziemlich große Zahl ausgetauscht dies können beide nun verwenden um ihr Treffen zu organisieren.

Dazu ein Beispiel aus der Praxis
```bash
$ echo "23:00Uhr 3027 W 12th St, Brooklyn, NY 11224, USA" > meetingpoint.txt
$ openssl aes-256-cbc -e -a -in meetingpoint.txt -out meetingpoint.txt.aes // -a erzeugt eine base64 Ausgabe
enter aes-256-cbc encryption password:
Verifying - enter aes-256-cbc encryption password:
```
Alternative könnte man Programme wie gpg, pgp und aescrypt verwenden.

Die Datei meetingpoint.txt.aes kann nun an Darlene übermittelt werden. Diese kann die Nachricht mit dem bekannten Passwort wieder entschlüsseln.

```bash
$ openssl aes-256-cbd -d -a -in meetingpoint.txt.aes 
enter aes-256-cbc decryption password:
```

Nachdem das Treffen vom FBI unfreundlich unterbrochen wurde und Elliot seinen Laptop liegen lassen musst. Gehen beide davon aus das der Geheimeschlüssel nicht mehr sicher ist. Auf der Flucht konnten die beiden jedoch keinen neuen Schlüssel vereinbaren. Wie kann nun die wichtige Informationen zwischen den beiden ausgetauscht werden? Dazu mehr im Kapitel 2 zuvor noch ein paar Erleuterungen und weiter Informationen.

Was können wir aus dem Beispiel lernen.

* Geheimerschlüssel muss beiden Parteien bekannt sein

#### Funktionen ####

```
E(sk, m) = c
D(sk, c) = m
```

#### Vor-/Nachteile ####

* Schnellere Ver-/Entschlüsslung als bei Public-Key
* Höhere Sicherheit bei gleicher Schlüsselänge als bei Public-Key
* Geheimschlüssel kann nicht über unsichere Kommunikationsweg vereinbart werden.
* Bekommt der Angreifer den Geheimschlüssel kann die gesamte Kommunikation gelesen werden.

#### Quellen ####
* Übersicht was die einzelnen Verfahren bei AES bedeuten - http://stackoverflow.com/questions/1220751/how-to-choose-an-aes-encryption-mode-cbc-ecb-ctr-ocb-cfb
* https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Electronic_codebook_.28ECB.29


### Kapitel 2 - Asymmetrische Verschlüsslung ###

Die Panik ist groß, es gibt wichtige Informationen Elliot betreffend, die Dark Army ist ihm auf der Spur. Darlene muss dies Elliot unbedingt mitteilen ohne das die Dark Army etwas davon erfährt. Jedoch gibt es keinen Geheimenschlüsseln mehr um Nachrichten zu verschlüsseln. Auch kann kein neuer Schlüssel über einen sicheren Kommunikationskannal ausgetauscht werden. Da davon auszugehen ist das die Dark Army sämtlichen Netzwerkverkehr von Elliot mitliest.

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
$ echo "Dark Army is watching you. Send me your adresse that we can meet us" > msg.txt
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

Setze ich im analogen Leben eine Unterschrift unter ein Dokument zeige ich so das ich damit einverstanden bin was in dem Dokument geschrieben ist und was die Vorrausetzung für den Teil 1 diese Satz ist, ich es auch gelesen habe. Wenn das Dokument von größere Bedeutung ist kann es auch vorkommen das ein Notar ebenfalls seine Unterschrift darunter setzen muss. Wenn ein Vertrag zwischen zwei Parteien geschlossen wird unterzeichnen beide Parteien den Vertrag und geben so ihr Einverständnis zu dem Geschrieben. Für die Unterschrift im analogen Leben gibt es eine Gegenpart in der digitalen Welt.

Die Idee Unterscheidet sich nicht sehr von der des Verschlüsselns einer Datei nur das hier der Weg andersherum ist.

Aber lass uns das ganze einmal wieder abarbeiten unter zuhilfenahme eines Beispiels. Kommen wir somit zurück zu unseren 2 Protagonisten Elliot und Darlene.

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

Damit Darlene, nach getaner Arbeit, anspruch auf ihre hälfte hat und zeigen kann das Elliot tatsächlich die zweite Person war welche den Vertrag unterschrieben hat muss sie eine Kopien des vertrag.txt, vertrag.elliot.sign und den Public-Key von Elliot an sich nehmen.

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

Aus mathematischen Gründen wird die Signatur mindesten Doppelt so groß wie die Daten die Verschlüsselt wurden das kann ziemilich ausufern bei größeren Dateien. Daher wird meistens ein Hash der Datei erstellt und dieser wird dan signiert. Das spart nicht nur Platz sonder hat auch noch den Vorteil das mit großer Wahrscheinlichkeit die Daten welche Signiert wurden nicht verändert wurden.

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

Am Ende steht nun noch die Frage offen woher weiß man das der Public-Key der zum Verifizieren benutzt wurde überhaupt Darlene gehört. Es wäre möglich das jemand den Vertrag mit einem anderen Schlüssel-Paar unterschrieben hat und daraufhin behauptet Darelene hätte den Vertrag unterschrieben. Das gleiche Problem hatten wir auch schon zuvor mit Elliot dieser hätte ebenfalls behaupten können das der Key welcher zum überprüfen verwendet wurde überhaupt nicht sein Public-Key gewesen sei.

Das führt uns zu zwei neuen Begriff dem Zertifikat und der Certificate Authority (abkz. CA)

TODO(tim):Hier noch ein Ablaufdiagram wie das Signieren funktioniert

##### Quellen #####
* http://security.stackexchange.com/questions/57336/certificate-request-why-does-the-requester-have-to-create-a-private-key
* Signieren von Dokumenten - https://raymii.org/s/tutorials/Sign_and_verify_text_files_to_public_keys_via_the_OpenSSL_Command_Line.html


#### Zertifikate ####

Bevor wir uns um die Begriffe Zertifikat und Certificate Authority widmen gibt es noch immer das Problem das das Böse kurz davor steht zu gewinnen. Also lass es uns endlich bekämpfen.

Wie hätte Darlene zum einen Sicherstellen können ob der Public-Key Elliot gehört und wie hätte sie dafür sorgen können das nur der richtige Elliot die verschlüsselten Daten entschlüsseln kann und so die Standort Daten nicht in die falschen Hände gelangen.

Elliot schreibt gerne Tagebuch (er muss wie man hört irgendwelche Probleme mit seinem Vater verarbeiten, wie man hört). Elliot ist ein verdamter Cyberpunk daher macht er das natürlich online und betreibt ein Blog dazu betreibt er einen HTTPS Server. Das heißt die Kommunikation findet Verschlüsselt statt und ebenso liegt auf dem Server ein Zertifikat (was das nun ist wird sich noch erschließen keine Angst). Da Darlene ebenfalls ziemlich Cyber ist weiß sie das zu nutzen und besorgt sich das Zertifikat des Servers.

Dazu verbinden wir uns mit dem Server von Elliot und Fragen nach dem Zertifikat. Das erste Zertifiakt welches wir zurück bekommen sollte das Zertifikat des Servers sein.
```bash
$ openssl s_client -showcerts -connect https://r1ng0.3ll1ot.com:443 </dev/null
---
CONNECTED(00000003)
---
Certificate chain
 0 s:/C=UA/ST=New York/L=New York City/O=E-Corp/OU=AI/CN=r1ng0.3ll1ot.com/emailAddress=ai@e-corp.com
   i:/C=UA/ST=New York/L=New York/O=Mr. Robot INC/OU=AI/CN=Mr. Robot INC/emailAddress=ai@mrrobot.com
-----BEGIN CERTIFICATE-----
MIIFnDCCA4QCCQCsulmnsw57WDANBgkqhkiG9w0BAQsFADCBjzELMAkGA1UEBhMC
VUExETAPBgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEWMBQGA1UE
CgwNTXIuIFJvYm90IElOQzELMAkGA1UECwwCQUkxFjAUBgNVBAMMDU1yLiBSb2Jv
dCBJTkMxHTAbBgkqhkiG9w0BCQEWDmFpQG1ycm9ib3QuY29tMB4XDTE2MDkwODEz
MjgxOVoXDTE2MTAwODEzMjgxOVowgY8xCzAJBgNVBAYTAlVBMREwDwYDVQQIDAhO
ZXcgWW9yazEWMBQGA1UEBwwNTmV3IFlvcmsgQ2l0eTEPMA0GA1UECgwGRS1Db3Jw
MQswCQYDVQQLDAJBSTEZMBcGA1UEAwwQcjFuZzAuM2xsMW90LmNvbTEcMBoGCSqG
SIb3DQEJARYNYWlAZS1jb3JwLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCC
AgoCggIBAJsLnia0gVyt/kUtUEchtfnlRBOTb8jKh+q3Iu8x1TV2UMk19A1ciCIE
fvC5cSyEeNbRYlWZSHK4KLuANdgCZ7UbEOoC6JMNSpzuaWwqqrCT8k0v1r+UYy2o
KV2pcQkjHaWxH++RUueouPVPTkhGlXV+6vyovIhoLQ1W569VrCxYJqbIGd5gHESU
1ZHMetRbx/pFxfI4WPFTfUAkV/s+pCkbW7FH3qygG2qSKeUNPZcvMRYNmt3lXHaD
cBadG0/IKnC7qBVyyK8z9dgha8o4S+M8xuuRvyKLHhVM/ts7va4HMGPkpC6+z0i5
VCnjFDzNaVH80GRvWR1x6b1LZ4WPNY3JzqXUET4Jg17T5kpaFu9InsUOpLfXmTKT
3gIViGeZa/guvfTgfwnAtAcnXUzJ2xcZ9SHyVoryM6NN8lr9Q+GYH0UN/3x3MfXg
La15s4vz/mUcIpMXe3PLQVF5dK8y53X/o0w5ltL3gbWdzlRUpWZP9qrrMM/F3WN4
O4m3oqK54o9BW79LPvVJgJx05tQWWGRRCMC5jtAZukFmp8x0rrgGACP4X6siOx7H
A5jkNR6BFdtUIMiatgs/xkUSkuNnIX15FPXGiPebELNgwU5pOLDDRxkLjkKYI/2a
vknKFy2ItJmiHlYSOoL7LdYoUTUYT3G4p1rWKwjOdySeYTeTIkYxAgMBAAEwDQYJ
KoZIhvcNAQELBQADggIBAN3wJExT5vAGNbdaGy6t9Z7Y7i2lrXLzSt+IaNSGxNTR
bRS2+AnV+klolq7HH+1wT59DlV14GnyZmxRwm8PV+5OntmnpOsyVa0Ma0eaCfHf2
GDNh6mJnd1W/DKiwgPprai2ymHpWaDn3J4jK0stNP3tlk9TmD2PR8wnR7mP5nyj9
gVKygcy77iaCH3OEWX3fwOE3BMSnLyL+efPZg6SmlWNzXs2LKNKMYHmXOtz8rhWA
fn1SflKYOqFjV9tfcse72eyPLALwh/tIuELLpSxmRwOepgmCiQRS1UegRRTjCEB/
WXnETbzpQED5K9krmkdWpK8+kduyWSP2F9wY5Fo/FaOyU2hkmNLdA8Cfo797D4SW
7hZw0YIinw+zXp4AsHrSKukhtkoKe0sMTm/543AOccm5gAMytCrzqqGD+c6w7+O9
N2XF98IMsRSI8YdNrq3zY6dU4lgrQGhhXMJFh+0Z3TubeoVD9hV+80qfkCgUbUAK
ryMTJiCN0E8LctiL90er5EChza/mq82n9ISZUU1xmKe+lOOdguasjNFYlN6zhxjx
UcArwXsDQn1H+IvrLssVr59jkPnmJynuelBXoS5jQsUJ3SH8nsOGhdNs80l56q7/
88n0SnV2R9gNS4TQpdogv+8tzoiw5ZZzAdBE+ZgP9h1n/SHzirCtQuxzuXznkNp/
-----END CERTIFICATE-----
---
Server certificate
subject=/C=UA/ST=New York/L=New York City/O=E-Corp/OU=AI/CN=r1ng0.3ll1ot.com/emailAddress=ai@e-corp.com
issuer=/C=UA/ST=New York/L=New York/O=Mr. Robot INC/OU=AI/CN=Mr. Robot INC/emailAddress=ai@mrrobot.com
---
No client certificate CA names sent
Peer signing digest: SHA512
Server Temp Key: ECDH, P-256, 256 bits
---
SSL handshake has read 2355 bytes and written 433 bytes
---
New, TLSv1/SSLv3, Cipher is ECDHE-RSA-AES256-GCM-SHA384
Server public key is 4096 bit
Secure Renegotiation IS supported
Compression: NONE
Expansion: NONE
No ALPN negotiated
SSL-Session:
    Protocol  : TLSv1.2
    Cipher    : ECDHE-RSA-AES256-GCM-SHA384
    Session-ID: 7E289CD32F1B023752AE7FCF5BDCD4C9CA79E0253F0D82EE3EDF96F41A80415A
    Session-ID-ctx: 
    Master-Key: 12AE6D70555A3DF53380B9BCAC6CC5DB1F92BC7B17D13B22E0EA56828E4A90981A3E135ABECF0FDE9E25C4E9E57AA5EB
    Key-Arg   : None
    PSK identity: None
    PSK identity hint: None
    SRP username: None
    TLS session ticket lifetime hint: 300 (seconds)
    TLS session ticket:
    0000 - 00 c3 18 b8 6c 70 df d5-12 f6 cc 1b 71 f1 64 5f   ....lp......q.d_
    0010 - 66 61 eb de 52 b6 4d b4-73 f6 03 76 bf 4b 77 48   fa..R.M.s..v.KwH
    0020 - 3f 3b 5b d9 dd 37 f6 8d-23 15 a8 3a ca 89 e4 50   ?;[..7..#..:...P
    0030 - 08 b6 71 e8 e7 20 cc d1-51 29 03 46 2d a4 09 2c   ..q.. ..Q).F-..,
    0040 - ab 9e 37 c7 63 a7 2b 5f-0c 08 99 dc 98 3c 69 8d   ..7.c.+_.....<i.
    0050 - e5 a2 b6 f9 66 9b a1 05-d4 1e 2d 8a 45 cd 00 10   ....f.....-.E...
    0060 - ce 3d 69 8c f3 b8 ce 5c-09 a4 dc a1 74 37 2d 97   .=i....\....t7-.
    0070 - 7f 0d c6 dd 00 11 67 38-d4 08 3c 84 e8 a9 c7 55   ......g8..<....U
    0080 - 67 ff 39 43 08 59 14 6f-fd f4 a8 ae 81 c8 cf c3   g.9C.Y.o........
    0090 - 23 5a 1f 47 44 2a 5d 3d-3f 9f 94 66 4a 94 05 3e   #Z.GD*]=?..fJ..>

    Start Time: 1473341388
    Timeout   : 300 (sec)
    Verify return code: 0 (ok)
-client -CAfile mr-robot-cert.pem -showcerts -connect https://r1ng0.3ll1ot.com:443
```
Die Informationen deuten schon einmal darauf hin das diese Zertifikat zu Elliot gehören könnte.

Darlene muss nun herausfinden ob das zurückbekommen Zertifikat tatsächlich Elliot gehört. Hier kommt nun eine weiter Person/Unternehmen ins Spiel. Eine Dritte unabhänig Istanz die vertrauenswürdig ist. Diese muss bestätigen das Elliot Elliot ist. Nennen wir diesen Dritten in unserem Beispiel Mr. Robot INC.


Diese Dritte Person auch CA genannt bestätigt das eine Person die ist für die sich ausgibt und hält das in einem Dokument, dem Zertifikat fest. Was ist nun diese Zertifikat.

```Anmerkung
Ein paar Beispiel für CAs aus dem wirklichen Leben
* Let's Encrypt https://letsencrypt.org/
* VeriSign https://www.verisign.com/
* https://www.startssl.com/
```

Schauen wir uns dazu einen Funktion an die eine Zertifikat erstellt.


Nochmals kurz wie wude eine Signatur definiert
```
signatur = sign(Private-Key, Data)
```

Wie man gleich sieht ein Zertifikat ist nichts anders. Es kommt auf die Daten an die übergeben werden
```
data = [Subject-Informationen, Public-Key]
cert = sign(private-key, data)
```

D.h. es gibt eine Funktion sign die aus private-key, Subject-Informationen (z. Bsp. Server IP, Domain, Personen Daten) und public-key ein Zertifikat erzeugt.

Was macht die Funktion?
```
cert(private-key, subject-information, public-key) = 
	1) Erstelle einen Hash aus subject-information und public-key
	2) Signiere Hash mit private-key
	3) Verarbeite Hash, public-key, subject-information und Signatur in das x509 Format
```

Ist Private-Key und Public-Key von der selben Person ist das Zertifikat ein Self-Signed-Certificate. Erstellt eine CA ein Zertifikat stammt der Private-Key von der CA und der Public-Key kommt von Antragsteller.

Damit die Idee da hinter deutlicher wird hier noch eine Funktion die ein Zertifiakt überprüft.
```
verify(public-key, certificate) =
	// Anmerkung der public-key ist der gegenschlüssel zu dem in der cert Funktion verwendet private-key.
	1) Erstelle Hash aus der Subject-Informationen und Public-Key welche beide im certificate hinterlegt sind.
	2) Benutze public-key um den Hash aus der Signatur auszulesen.
	3) Vergleiche Hash1 mit Hash2
		3.1) Gleich, dann ist Zertifikat in Ordnung
		3.2) Ungleich, irgendetwas stimmt mit dem Zertifikat nicht
```

Ein Zertifikat ist ein eine Kombination von Informationen zu einem zu Idendifizierendem Etwas und desen Public-Key. Diese Kombination wird von einer zweiten Partie als korrekt Bestätigt. Das zeigt Sie der Welt in dem Sie eine Digital Unterschrift unter diese Kombination setzt. Da durch kann man sicherstellen wenn etwas mit diesem Public-Key verschlüsselt kann diese nur wieder von dem Etwas entschlüsselt werden. Ob das nun eine Person ist oder ein HTTPS Server spielt dabei keine Rolle. Es beweist das ein Öffentlicher Schlüssel zu einem Benutzer gehört.

```Anmerkung
Der public-key der Hier verwendet wird ist im Fall eines Self-Signed-Certificate der gleiche Public-Key wie er verwendet wurde zum erstellen des Certificates in der cert Funktion. Ist es jedoch nicht der Fall ist der public-key der Öffentliche Schlüssel zum private-key welcher verwendet wurde in der cert Funktion um ein Zertifikat zu erstellen. Dieser Public-Key muss von einer CA zurverfügung gestellt werden. Meistens kann man diesen auf der Webseite der CA herunterladen. Im Fall des Browser werden alle Möglichen Public-Keys von unterschiedlichen CA mitgeliefert. Man könnte an die verify Funktion anstatt des public-key auch ein Zertifkate übergeben dann würde aus diesem Zertifikat der Public-Key verwendet werden. Um den Hash aus der Signatur zu lesen.

Ein Public-Key kann wie folgt aus einem eim Zertifikat ausgelesen werden.
$ openssl x509 -pubkey -in mr-robot-cert.pem
```
Verwenden wir einmal wieder ein Beispiel um das ganze zu konkretisenren. Wir werden zu einer CA und erstellen ein Zertifkat mit dem wir andere Zertifikate unterschreiben können.

##### CA #####

Eine CA benötigt ebenso wie alle anden auch ein Public-/Private-Key den die CA macht auch nicht anderes als das in den Kapiteln zuvor beschrieben wurde.

Nehmen wir also an wir wollten eine CA Gründen wir sind schließlich äußerst Vertraunenswürdig. Wir haben ein Unternehmen mit dem Namen Mr. Robot Inc. Wir erzeugen uns also als erstes ein Schlüssel-Paar.

```bash
genrsa -out mr-robot-keys.pem -aes256 4096
```

Nun benötigen wir ein Zertifikat das bezeugt wer wir sind. Mit diese Zertifikat können wir dann andere Zertifikat unterzeichnen.

Da wir die oberste Instanz sind heißt das wir bestätigen einfach selbst das wir sind wer wir sind. (Self-Signed-Certificate). Wie oben gezeigt gehören zu jedem Zertifiakt ein public-key und Informationen in denen stehen wem das Zertifkat gehört. Damit jeder die Infomationen auslesen kann gibt es einen Standard der definiert wo welche Daten stehen und die Informationen bezeichnet werden den sogenanten x509 Standard. 

Mit dem folgenden Befehl erzeuge wir nun ein Zertifikat das mit unserem Private-Key Signiert ist.
```bash
openssl req -x509 -new -days 365 -in mr-robot-keys.pem -out mr-robot-cert.pem 
```

Die 2 Befehle oben können auch in einem zusammengefasst werden aber damit man sieht das hier kein Vodoo stattfindet wurden die Schritte einzeln ausgeführt.

Das ganze in einem Befehl
```bash
openssl req -newkey rsa:4096 -nodes -keyout mr-robot-keys.pem -x509 -days 365 -out mr-robot-cert.pem
```

```Anmerkung
Was ist jetzt der Unterschied zwischen unseren CA und all den anderen da draußen. Das wird schnell klar wenn jemand mit Chrome auf einen Server surft der ein von uns ausgestellten Zertifikat verwendet. Der Browser zeigt an das man dem Zertifiakt nicht trauen kann. Warum das. Wird auf eine Webseite gesurft die eine Zertifiakt verwendet um zu zeigen das sie auch die Seite ist die man angesurft hat überprüft der Browser das angeboten Zertifiakt in dem er aus einer Fest in den Browser integriert List von CA-Public-Keys das für das Server-Zertifikat passende auswählt und damit verifiziert. Da haben wir den unterschied, wir sind nich so vertrauenswürdig das wir in diese Browser CA-List aufgenommen werden. Und ja google entscheidet für wem wir Vertrauen und wem nicht. Die Liste von Public-Keys kann man unter Einstellungen -> Zertifikate angeschaut werden.
```

Zurück zur Geschichte. Mr. Robot INC hat vor vielen Jahren, als die Welt von Elliot und Darlene noch in Ordnung war, bestätigt (mit dem oben erstellten Zertifikat) das die Angaben und der Public-Key welcher im Zertifiakt von Elliots Server hinterlegt sind korrekt sind und zu Elliot gehören. Diese hilft Darlene sicherzustellen das die Nachricht (Errinern Sie sich noch an die Nachricht die die unglücklicherweise nicht von Elliot kamm) nur mit dem Private-Key von Elliot wieder entschlüsslt werden kann. Andersherum kann man Verifizieren das eine Nachricht die mit dem Zertifikat entschlüsslet werden kann nur von dem Besitzer des Private-Keys kommen kann.

Darlene hatte zuvor schon das Zertifiakt von Elliots Server geholt. Nun haben wir gelernt das man zum Verifzieren des Zertifikats das Public-Root-Zertifikat der Mr. Robot Inc benötigt. Diese besorgt sich Darlene und speichert es in mr-robot-cert.pem ab.

```Anmerkung
Wenn man ein Public-Root-Zertifikat (Public-Root-Key) einer CA benötigt kann man dazu im Browser seiner wahl nachschauen. Dort gibt es eine Liste von mehr oder weniger vertrauenswürdigen CA Zertifikaten.
```

Darlene kopieren sich nun das das Zertifikat das sie zuvor vom Server geholt hat in eine Datei mit dem Namen elliot-cert.pem. Das ganze sieht dann wie folgt aus (Vergleiche mit Ausgabe oben).

```bash
$ cat elliot-cert.pem
-----BEGIN CERTIFICATE-----
MIIFnDCCA4QCCQCsulmnsw57WDANBgkqhkiG9w0BAQsFADCBjzELMAkGA1UEBhMC
VUExETAPBgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEWMBQGA1UE
CgwNTXIuIFJvYm90IElOQzELMAkGA1UECwwCQUkxFjAUBgNVBAMMDU1yLiBSb2Jv
dCBJTkMxHTAbBgkqhkiG9w0BCQEWDmFpQG1ycm9ib3QuY29tMB4XDTE2MDkwODEz
MjgxOVoXDTE2MTAwODEzMjgxOVowgY8xCzAJBgNVBAYTAlVBMREwDwYDVQQIDAhO
ZXcgWW9yazEWMBQGA1UEBwwNTmV3IFlvcmsgQ2l0eTEPMA0GA1UECgwGRS1Db3Jw
MQswCQYDVQQLDAJBSTEZMBcGA1UEAwwQcjFuZzAuM2xsMW90LmNvbTEcMBoGCSqG
SIb3DQEJARYNYWlAZS1jb3JwLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCC
AgoCggIBAJsLnia0gVyt/kUtUEchtfnlRBOTb8jKh+q3Iu8x1TV2UMk19A1ciCIE
fvC5cSyEeNbRYlWZSHK4KLuANdgCZ7UbEOoC6JMNSpzuaWwqqrCT8k0v1r+UYy2o
KV2pcQkjHaWxH++RUueouPVPTkhGlXV+6vyovIhoLQ1W569VrCxYJqbIGd5gHESU
1ZHMetRbx/pFxfI4WPFTfUAkV/s+pCkbW7FH3qygG2qSKeUNPZcvMRYNmt3lXHaD
cBadG0/IKnC7qBVyyK8z9dgha8o4S+M8xuuRvyKLHhVM/ts7va4HMGPkpC6+z0i5
VCnjFDzNaVH80GRvWR1x6b1LZ4WPNY3JzqXUET4Jg17T5kpaFu9InsUOpLfXmTKT
3gIViGeZa/guvfTgfwnAtAcnXUzJ2xcZ9SHyVoryM6NN8lr9Q+GYH0UN/3x3MfXg
La15s4vz/mUcIpMXe3PLQVF5dK8y53X/o0w5ltL3gbWdzlRUpWZP9qrrMM/F3WN4
O4m3oqK54o9BW79LPvVJgJx05tQWWGRRCMC5jtAZukFmp8x0rrgGACP4X6siOx7H
A5jkNR6BFdtUIMiatgs/xkUSkuNnIX15FPXGiPebELNgwU5pOLDDRxkLjkKYI/2a
vknKFy2ItJmiHlYSOoL7LdYoUTUYT3G4p1rWKwjOdySeYTeTIkYxAgMBAAEwDQYJ
KoZIhvcNAQELBQADggIBAN3wJExT5vAGNbdaGy6t9Z7Y7i2lrXLzSt+IaNSGxNTR
bRS2+AnV+klolq7HH+1wT59DlV14GnyZmxRwm8PV+5OntmnpOsyVa0Ma0eaCfHf2
GDNh6mJnd1W/DKiwgPprai2ymHpWaDn3J4jK0stNP3tlk9TmD2PR8wnR7mP5nyj9
gVKygcy77iaCH3OEWX3fwOE3BMSnLyL+efPZg6SmlWNzXs2LKNKMYHmXOtz8rhWA
fn1SflKYOqFjV9tfcse72eyPLALwh/tIuELLpSxmRwOepgmCiQRS1UegRRTjCEB/
WXnETbzpQED5K9krmkdWpK8+kduyWSP2F9wY5Fo/FaOyU2hkmNLdA8Cfo797D4SW
7hZw0YIinw+zXp4AsHrSKukhtkoKe0sMTm/543AOccm5gAMytCrzqqGD+c6w7+O9
N2XF98IMsRSI8YdNrq3zY6dU4lgrQGhhXMJFh+0Z3TubeoVD9hV+80qfkCgUbUAK
ryMTJiCN0E8LctiL90er5EChza/mq82n9ISZUU1xmKe+lOOdguasjNFYlN6zhxjx
UcArwXsDQn1H+IvrLssVr59jkPnmJynuelBXoS5jQsUJ3SH8nsOGhdNs80l56q7/
88n0SnV2R9gNS4TQpdogv+8tzoiw5ZZzAdBE+ZgP9h1n/SHzirCtQuxzuXznkNp/
-----END CERTIFICATE-----
```

Nun können wir das Zertifikat mit dem Root-Zertifikat von Mr. Robot Inc verifzieren.

```bash
$ openssl verify -CAfile mr-robot-cert.pem elliot-cert.pem 
elliot-cert.pem: OK
```

Ein fröhliches OK bestätigt Darlene das das Zertfikat zu Elliot gehört. Erleichterung wir haben einen Weg gefunden um mit Elliot zu kommunizieren.

Wir können nun die Frage nach einem neuen Treffpunkt an Elliot übermitteln und wir können sichergehen das ihn nur Elliot lesen kann und somit auch nur dieser auf diese Fragen antworten kann und nicht die Dark Army. Elliot könnte natürlich den selben Weg gehen und nach einem Zertifikat suche das zu Darlene gehört und diesen Public-Key verwenden um Geheime Nachrichten an Darlene zu senden.

```bash
$ openssl x509 -pubkey -in elliot-cert.pem
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAmwueJrSBXK3+RS1QRyG1
+eVEE5NvyMqH6rci7zHVNXZQyTX0DVyIIgR+8LlxLIR41tFiVZlIcrgou4A12AJn
tRsQ6gLokw1KnO5pbCqqsJPyTS/Wv5RjLagpXalxCSMdpbEf75FS56i49U9OSEaV
dX7q/Ki8iGgtDVbnr1WsLFgmpsgZ3mAcRJTVkcx61FvH+kXF8jhY8VN9QCRX+z6k
KRtbsUferKAbapIp5Q09ly8xFg2a3eVcdoNwFp0bT8gqcLuoFXLIrzP12CFryjhL
4zzG65G/IoseFUz+2zu9rgcwY+SkLr7PSLlUKeMUPM1pUfzQZG9ZHXHpvUtnhY81
jcnOpdQRPgmDXtPmSloW70iexQ6kt9eZMpPeAhWIZ5lr+C699OB/CcC0ByddTMnb
Fxn1IfJWivIzo03yWv1D4ZgfRQ3/fHcx9eAtrXmzi/P+ZRwikxd7c8tBUXl0rzLn
df+jTDmW0veBtZ3OVFSlZk/2quswz8XdY3g7ibeiornij0Fbv0s+9UmAnHTm1BZY
ZFEIwLmO0Bm6QWanzHSuuAYAI/hfqyI7HscDmOQ1HoEV21QgyJq2Cz/GRRKS42ch
fXkU9caI95sQs2DBTmk4sMNHGQuOQpgj/Zq+ScoXLYi0maIeVhI6gvst1ihRNRhP
cbinWtYrCM53JJ5hN5MiRjECAwEAAQ==
-----END PUBLIC KEY-----
-----BEGIN CERTIFICATE-----
MIIFkjCCA3oCCQCsulmnsw57VzANBgkqhkiG9w0BAQsFADCBjzELMAkGA1UEBhMC
VUExETAPBgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEWMBQGA1UE
CgwNTXIuIFJvYm90IElOQzELMAkGA1UECwwCQUkxFjAUBgNVBAMMDU1yLiBSb2Jv
dCBJTkMxHTAbBgkqhkiG9w0BCQEWDmFpQG1ycm9ib3QuY29tMB4XDTE2MDkwODEy
NTgxOVoXDTE2MTAwODEyNTgxOVowgYUxCzAJBgNVBAYTAkFVMREwDwYDVQQIDAhO
ZXcgWW9yazEWMBQGA1UEBwwNTmV3IFlvcmsgQ2l0eTEPMA0GA1UECgwGRS1Db3Jw
MQswCQYDVQQLDAJBSTEPMA0GA1UEAwwGRWxsaW90MRwwGgYJKoZIhvcNAQkBFg1h
aUBlLWNvcnAuY29tMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAmwue
JrSBXK3+RS1QRyG1+eVEE5NvyMqH6rci7zHVNXZQyTX0DVyIIgR+8LlxLIR41tFi
VZlIcrgou4A12AJntRsQ6gLokw1KnO5pbCqqsJPyTS/Wv5RjLagpXalxCSMdpbEf
75FS56i49U9OSEaVdX7q/Ki8iGgtDVbnr1WsLFgmpsgZ3mAcRJTVkcx61FvH+kXF
8jhY8VN9QCRX+z6kKRtbsUferKAbapIp5Q09ly8xFg2a3eVcdoNwFp0bT8gqcLuo
FXLIrzP12CFryjhL4zzG65G/IoseFUz+2zu9rgcwY+SkLr7PSLlUKeMUPM1pUfzQ
ZG9ZHXHpvUtnhY81jcnOpdQRPgmDXtPmSloW70iexQ6kt9eZMpPeAhWIZ5lr+C69
9OB/CcC0ByddTMnbFxn1IfJWivIzo03yWv1D4ZgfRQ3/fHcx9eAtrXmzi/P+ZRwi
kxd7c8tBUXl0rzLndf+jTDmW0veBtZ3OVFSlZk/2quswz8XdY3g7ibeiornij0Fb
v0s+9UmAnHTm1BZYZFEIwLmO0Bm6QWanzHSuuAYAI/hfqyI7HscDmOQ1HoEV21Qg
yJq2Cz/GRRKS42chfXkU9caI95sQs2DBTmk4sMNHGQuOQpgj/Zq+ScoXLYi0maIe
VhI6gvst1ihRNRhPcbinWtYrCM53JJ5hN5MiRjECAwEAATANBgkqhkiG9w0BAQsF
AAOCAgEAYk+tjzIKriTFbE6GQzQocBFHYBkEj0zOY8k+mAPsznTsXBR2BYLGI3Kh
szeOgqVxuqfAgLOmKrSQeWxms3VGldaip+zOs2bEWmaL29xGFzsjFeopoK3pKVoA
81TWyV28Mp+CO+2IL36s7J6Mo9Ks8Ww0NeW3sqIvcHGYY1qfZEejPyyqMOt/y16t
KfLX+Wwl86TV875ShZlqyKlEGg3WgrTuLzas4m3agQq98ELKBSGTwZ7bQBR5XDom
43lSl1EifNuPYvxVYJ8DUjTTt/66dHxgdUfhGqA1sQh0prtbbWhfuLFu+ePCUg/8
4vHdA+pYqxe1hixX/fWyzqk2GuPyxv7OOWpdjSkBLaTkE8ac632nieJw+iJyyRHo
gxeHDGwiup7THV971PpfB7j2yEVjGe6O0FmvDDhgZpcSXnelZdp7XLSoRXPiECk8
ueZkptYslsIEIzScO1h58Wl20PmfLaZGHmh3Ap9sKanBLj0xcqne77/mkWRtj6Po
6d+1CDYGGxW8K3zzcTfQVk5ULUR9hD/U3MA2kC/ItTMoXqYvBSkH17dede/qvdzP
h/IbsdChkA9WjfLTXsREFWFvGToAuLWO/6ly/ZJN7h+s8fMdwoZuXVDLSt4y8ElE
LZi5X4FIVIpyNgNQ0gNsyh1mW6ifSE1YBmnM/AnzlGtlRkXDskY=
-----END CERTIFICATE-----
```

Darlene kopiert sich nun den Public-Key in die Datei elliot-public-key.pem und benutzt diesen Key um die Nachricht zu verschlüsseln.
```bash
$ echo "Dark Army is watching you. Send me your Adresse that we can meet us" > msg.txt
$ openssl rsautl -inkey elliot-public-key.pem -pubin -encrypt -in msg.txt -out msg.txt.encrypt
```

Alternative könnte Darlene es sich sparen den public-key zu extrahieren und das Zertifikat direkt zum Verschlüsseln verwenden dabei wird nicht anders gemacht als der darin enthalten Public-Key verwedet wird
```
$ echo "Dark Army is watching you. Send me your Adresse that we can meet us" > msg.txt
openssl rsautl -inkey elliot-cert.pem -certin -encrypt -in msg.txt -out msg.txt.encrypt 
```

##### Certificate Sign Request (CSR) #####

Gehen wir einige Jahre zurück und schauen wie Elliot zu seinem Zertifikat gekommen ist. Das, man kann es so sagen, die Welt gerettet hat. Alles begann mit einem CSR.

Wie kommen die Daten und der Private Schlüssel zu einer CA das diese ein Zertifiakt darauß erzeugen kann. 

Ist ein CSR nicht schon ein Zertifikat?
Da sind wir wieder. Ein CSR ist nichts anders als ein Zertifikat wir hatten ein Zertifikat wie folgt definiert

```
cert = sign(private-key; infos, public-key)
```

Nun signieren wir unsere Information mit unserem Private-Key
```bash
openssl x509 -signkey mr-robot-keys.pem -in mr-robot.csr -req -days 365 -out mr-robot-cert.pem
```

Anfrage nicht im x509 Format sonder in pkcs#7
```bash
openssl req -new -in mr-robot-keys.pem -out mr-robot-cert.pem 
```

Zertifikat wird von Mr. Robot Inc Signiert
```bash
openssl x509 -req -CAkey mr-robot-keys.pem -CA mr-robot-cert.pem -CAcreateserial -in elliot.csr -out elliot-cert.pem
```

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


```bash
openssl req -new -in mr-robot-keys.pem -out mr-robot.csr 
```

Wie wird überprüft ob die Person die den CSR and die CA sendet tatsächlich die Person ist für die sie sich ausgibt?Nehmen wir hier das Beispiel das eine Person für einen Server ein Zertifikat bekommen möchte. Dan wird in dem CSR die Server Adresse hinterlegt in Form von IP oder der Domain. Nun kann die CA verlangen das man auf dem Server ein Programm installiert das auf einen Request mit Token bestätigt das der Server zur Anfrage der Person X gehört. Somit kann die CA die Anfrage bestätigen. Das ist so ungefähr wie Let's Encrypt das macht. Es gibt viele weiter Möglichkeiten z. Bsp. könnte eine CA verlangen das die Person einen Ausweis vorlegt mit diesem Ausweis wird überprüft ob der Server auf diese Person registiert ist und so bestätigen das Server von der Person betrieben wird.

Es gibt die unterschiedlichsten Verfahren von ziemlich einfach bist Paranoid daher gibt es auch unterschieldiche Klassen welches ein Zertifikat erhalten kann. Eine Klasse zeigt wie sehr man einem Zertifiakt vertrauen kann d.h. wie sicher man sich sein kann das die Person mit der man Kommuniziert die richtige Person ist. Davon abgesehen das der Zertifikats Besitzer sein Private-Key verloren hat oder das er gestohlen wurde.

Beispiel für unterschiedliche Zertifikats Klassen bei VeriSign werden folgende Klassen verwendet diese können aber von CA zu CA unterschiedlich sein.

VeriSign uses the concept of classes for different types of digital certificates [3]:

* Class 1 for individuals, intended for email.
* Class 2 for organizations, for which proof of identity is required.
* Class 3 for servers and software signing, for which independent verification and checking of identity and authority is done by the issuing certificate authority.
* Class 4 for online business transactions between companies.
* Class 5 for private organizations or governmental security.

Beachte das durch eine Man-In-The-Middle-Attack ein falsche Certificate eingeschmuckelt werden kann.
Es gibt immer noch ein Certificate in der Kette das man überprüfen kann doch irgendwann ist man oben angekommen diese Zertifkat muss irgend wie zu einem kommen ohne das es Manipuliert wurde im Fall von Browser fest ein gebaut.



###### Quelle ######

* https://en.wikipedia.org/wiki/Certificate_signing_request#Structure
* http://stackoverflow.com/questions/21297139/how-do-you-sign-certificate-signing-request-with-your-certification-authority
* https://tools.ietf.org/html/rfc2986



##### Quellen #####

* http://security.stackexchange.com/questions/48802/how-to-validate-a-client-certificate
* Was ist ein Zertifikat - ftp://ftp.pgpi.org/pub/pgp/6.5/docs/german/IntroToCrypto.pdf




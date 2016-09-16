### Kapitel 4 - Diffe-Hellmann-Schlüsselaustausch ###

Einweg Public-Keys 

```Latex
b^x mod p
```

#### Quellen ####
https://de.wikipedia.org/wiki/Diffie-Hellman-Schl%C3%BCsselaustausch



## RSA ##
* https://www.emc.com/collateral/white-papers/h11300-pkcs-1v2-2-rsa-cryptography-standard-wp.pdf
* https://tools.ietf.org/html/rfc3447


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

Server prüft nur ob Zertifikat von CA signiert wurde.

## Protokoll Aufbau ##

Zwei Schichten

Record Protocol

## Quellen ##
[1] https://eprint.iacr.org/2013/538.pdf

# Stichwörter #
PKCS Public Key Cryptography Standards - https://en.wikipedia.org/wiki/PKCS





# Quellen #
[1] https://eprint.iacr.org/2013/538.pdf
--

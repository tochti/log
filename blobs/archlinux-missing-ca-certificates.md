# Fehlende CA Zertifikate #

Damit git, curl, etc. wissen welchen Zertifikaten Sie vertrauen können müssen benötigen Sie eine Liste der zu Vertrauenden CA Zertifikaten. Eine solche Liste kann aus den in Archlinux mitgelieferneden CA Zertifikaten.

```bash
$ trust extract --foramt=pem-bundle ca-certificates.crt
```

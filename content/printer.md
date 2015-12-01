Drucker
------

CUPS
....

backends /usr/lib/cups/backend

Zeigt all backens an
# lpinfo -v

Canon
.....

cnijfilter-common-i386 ?
cnijfilter-mg5200

Suche Drucker
# cnijnetprn --search auto

Drucker anlegen
# sudo /usr/sbin/lpadmin -p canonmg5200 -m canonmg5200.ppd -v cnijnet:/00-1E-8F-67-4C-D8 -E
# sudo rc.d restart cupsd

Gehe auf http://localhost:631/ Bearbeite den Drucker ohne etwas zu ändern.

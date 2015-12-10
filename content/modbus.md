+++
date = "2015-11-25T09:07:27+01:00"
draft = true
title = "Modbus"
+++

connect_id = 1

Beim setzten des Modbus Registers muss darauf geachtet werden das die Datenbank in der sich das Array befindet in dem Daten abgelegt werden oder aus dem Daten gelesen werden nicht optimized ist. Um diese zu überprüfen kann man in den Eigenschaften der Datenbank nachschauen. Dazu geht man mit der Rechtenmaustaste auf dei Gewünschte Datenbank und müsste dann auf der ersten Seite den richtigen Wert sehen.

Ebenfalls ist zu beachten das mein beim auswählen des Registers nicht eine Element des Arrays auswählt sonder nur den Pointer das heißt man muss auf das Array klicken und dann "keinen Eintrag" auswählen.

$DB
  $Array 0 to X type word
    0: 
    1:
    ...

pymodbus
=======

wenn man mittels der pymodbus Funktion read_holding_register auf eine Array zugreifen möchte muss man zuerst die Start Adresse angeben diese ist relative zu sehen nicht absolut das heißt einfach ausgedrückt man muss den Element-Index angeben diese Funktioniert allerdings nur wenn man bei Array den type word angegeben hat ansonsten muss man sich ausrechnen welcher Index benötigt wird.

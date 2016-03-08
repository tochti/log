+++
date = "2016-01-05T01:51:37+01:00"
draft = false
title = "Enable hdmi video + sound"

+++

Das Ziel ist es Filme mittels Notebook mit dem TV als Ausgabemedium anzuschauen.
Zu beantworten gilt es also folgendes, "wie verdammt nochmal bekomme ich Bild und Ton über das HDMI Kabel".

Bild
----

Starten wir mit dem Bild mittels xrandr kann man zwischen Verschieden Videoeingängen hin und her schalten. Mann kann einiges damit anstellen. Für unseres Problem reichen folgende Befehle

Find die Bezeichnung für den HDMI Ausgang

~~~bash
xrandr

eDP1 connected (normal left inverted right x axis y axis)
   1366x768      60.00 +
   1280x720      60.00  
   1024x768      60.00  
   1024x576      60.00  
   960x540       60.00  
   800x600       60.32    56.25  
   864x486       60.00  
   640x480       59.94  
   720x405       60.00  
   680x384       60.00  
   640x360       60.00  
HDMI1 connected 1920x1080+0+0 (normal left inverted right x axis y axis) 160mm x 90mm
   1920x1080     60.00*+  50.00    59.94    30.00    25.00    24.00    29.97    23.98  
   1920x1080i    60.00    50.00    59.94  
   1680x1050     59.88  
   1400x1050     59.95  
   1600x900      59.98  
   1280x1024     60.02  
   1440x900      59.90  
   1360x768      60.02  
   1280x800      59.91  
   1152x864      59.97  
   1280x720      59.81    60.00    50.00    59.94  
   1440x576i     50.00  
   1024x768      60.00  
   800x600       60.32  
   720x576       50.00  
   720x480       60.00    59.94  
   640x480       60.00    59.94  
   720x400       70.08  
VIRTUAL1 disconnected (normal left inverted right x axis y axis)
~~~

Mit dem folgenden Befehl schalten wir den Laptop Display aus und den HDMI1 Ausgang ein. Die Auflösung für den HDMI1 Ausgang wird dabei automatisch bestimmt. Wenn ich bei mir das Display eDP1 nicht auschalte bekomme ich keine richtige Auflösung auf dem TV Bildschirm.

~~~bash
xrandr --output eDP1 --off --output HDMI1 --auto
~~~

Damit der Bildschirm nicht ständig in Standbymodus fährt und der Bildschirmschoner dazwischen funkt schalten wir mit folgendem Befehl beides aus.

~~~bash
xset s off -dpms
~~~

Kurze Erklärung des Befehls
~~~bash
// turn off screensaver
xset s off
// turn off standby 
xset -dpms
~~~

Aktivieren kann man beides wieder wie folgt

~~~bash
xset s on +dpms
~~~

https://wiki.archlinux.de/title/DPMS

Ton
----

Ton über das HDMI Kabel zu bekommen hat bei mir jedoch ein wenig länger gedauert da folgende Fehlermeldung nicht verschwinden wollte.

~~~
aplay: set_params:1239: Kanalanzahl nicht unterstützt
~~~

Aber immer langsam mit den jungen Pferden. Prinzipel kann man mittels den unten stehenden Commands Informationen über die Soundkarten einholen.

~~~bash
aplay -l
aplay -L
cat /proc/asound/cards
~~~

Nun kann man in .asoundrc die standard Soundcard ändern. Da bei mir wie oben beschrieben eine Fehlermeldung auftauchte wurde meine konfiguration ein wenig aufwändiger.

https://wiki.gentoo.org/wiki/ALSA#APlay_SPDIF.2FHDMI_output_has_incorrect_speaker_channels

Ich bin nun wiefolgt vorgegangen zum einen musste ich bei meinem HDMI Sounddeveice nicht nur die "Hardware-Adresse" hw:1,3 (Karte 1, Gerät 3) angeben sonder auch die Kannalanzahl bei mir 2 (Stero) das Ergebnis für meine Hardware sieht man unter pcm.myhdmi. Für Leute die keine Probleme mit dem obenen beschrieben Fehler haben kann es sein das die konfiguration anderst aussieht. Unter https://wiki.ubuntuusers.de/.asoundrc kann man eine kurze Beschreibung finden.

Der zweite Schritt den ich gemacht habe befindet sich unter pcm.!default dieser Abschnitt sorgt dafür das ich mittels der Umgebungsvariable ALSAPCM das standard Sounddeveice wechseln kann.

Final .asoundrc 
~~~
pcm.!default {
    type plug
    slave.pcm {
        @func getenv
        vars [ ALSAPCM ]
        default "hw"
    }
}

pcm.myhdmi {
	type plug slave {
    		pcm "hw:1,3"
    		channels 2 
	}

}
~~~

Nun ist es möglich mittles der Umgebungsvariable ALSAPCM=myhdmi die standard Soundkarte auf HDMI1 umzustellen. So kann man zum Beispiel firefox über die Konsole starten welcher dann für die Soundausgabe den TV verwendet.

~~~bash
$(ALSAPCM=myhdmi firefox)
~~~

Noch ein kurzer Hinweis. Zu beginn habe ich es nicht hinbekommen überhaupt einen Ton mit aplay auf den TV-Lautsprächern auszugeben. Daher dachte ich es wäre irgendwas Prinzipeles an der alsa konfiguration falsch. Jedoch habe ich dann mittels speaker-test es schlußendlich hinbekommen einen Ton auszugeben da hier die Kannal Angaben (-c 2) richtig verarbeitet wurde. Dies hat aus irgendwelchen Gründen bei aplay -c 2 nicht richtig funktioniert. Also am besten beides testen.

~~~bash
speaker-test -D hw:1,3 -c 2
~~~


ALSA
====

*Benutzer muss in Gruppe audio sein.*

Aktiviere eine Soundcard

~~~bash
amixer set Master on
~~~

Informationen über die Soundcards

~~~bash
cat /proc/asound/cards
aplay -l
aplay -L
~~~

Lautsprecher Testen

~~~bash
speaker-test -D hw:1,3 -c 2
~~~

Dabei besteht die einfache Möglichkeit die Kanalanzahl zu übergeben (-c 2). Interessant bei folgender Fehler 

~~~
aplay: set_params:1239: Kanalanzahl nicht unterstützt
~~~

so kann getestet werden welche Kanalzahl die Richtig ist

HP ProBook 640 G1
-----------------

Eine Möglichkeit ist Soundcarde zu konfigurieren ist unter https://wiki.ubuntuusers.de/.asoundrc/ die .asoundrc  Konfiguration unter dem Punkt Stereo in .asoundrc zu kopieren und die Card/Device Parameter anzupassen siehe dazu Informationen Soundcards

Quellen:
*) https://wiki.archlinux.de/title/Alsa
*) https://wiki.ubuntuusers.de/.asoundrc/
*) http://www.alsa-project.org/main/index.php/Asoundrc
*) http://www.alsa-project.org/alsa-doc/alsa-lib/pcm_plugins.html
*) https://wiki.archlinux.org/index.php/Advanced_Linux_Sound_Architecture

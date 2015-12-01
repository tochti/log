+++
title = "Fehler beim laden der vbox* Module"
date = "2015-11-25T09:07:27+01:00"
draft = true
+++

Es konnten die vbox module nicht geladen werden.

{% highlight text linenos %}
Jul 23 18:50:29 euler systemd-modules-load[175]: Inserted module 'kvm'
Jul 23 18:50:29 euler systemd-modules-load[175]: Failed to find module 'vboxdrv'
Jul 23 18:50:29 euler systemd-modules-load[175]: Failed to find module 'vboxvideo'
Jul 23 18:50:29 euler systemd-modules-load[175]: Failed to find module 'vboxnetadp'
Jul 23 18:50:29 euler systemd-modules-load[175]: Failed to find module 'vboxnetflt'
Jul 23 18:50:29 euler systemd[1]: systemd-modules-load.service: main process exited, code=exited, status=1/FAILURE
Jul 23 18:50:29 euler systemd[1]: Failed to start Load Kernel Modules.
{% endhighlight %}

Wie man unten sieht befinden diese sich in dem extramodules Verzeichnis und sind abhängig von bestimmten anden module.

{% highlight bash %}
[wirrwarr@euler modules]$ ls -l 3.15.5-2-ARCH/
insgesamt 3908
drwxr-xr-x  3 root root   4096 23. Jul 18:27 build
lrwxrwxrwx  1 root root     25 11. Jul 07:59 extramodules -> ../extramodules-3.15-ARCH
{% endhighlight %}

{% highlight bash %}
[wirrwarr@euler modules]$ sudo ls -l /lib/modules/extramodules-3.15-ARCH/
insgesamt 316
-rw-r--r-- 1 root root 143082 16. Jul 22:48 vboxdrv.ko.gz
-rw-r--r-- 1 root root 106971 16. Jul 22:48 vboxguest.ko.gz
-rw-r--r-- 1 root root   4303 16. Jul 22:48 vboxnetadp.ko.gz
-rw-r--r-- 1 root root  12587 16. Jul 22:48 vboxnetflt.ko.gz
-rw-r--r-- 1 root root  10479 16. Jul 22:48 vboxpci.ko.gz
-rw-r--r-- 1 root root  21436 16. Jul 22:48 vboxsf.ko.gz
-rw-r--r-- 1 root root   1536 16. Jul 22:48 vboxvideo.ko.gz
-rw-r--r-- 1 root root     14 11. Jul 07:59 version
{% endhighlight %}

Um alle Module korrekt laden zu können muss folgender Befehl ausgeführt werden.
{% highlight bash %}
[wirrwarr@euler modules]$ depmod -a 
{% endhighlight %}

Erzeugt alle Abhänigkeiten und trägt diese in Datei /lib/modules/<Kernelversion>/modules.dep ein.
Da per default der Ordner /lib/modules/<Kernelversion/kernel von modprobe nach Modulen durchsucht wird fällt das Verzeichnis /lib/modules/<Kernelversion>/extramodules nicht unter die durchsuchten Verzeichnisse. Allerdings verwendet modprobe die modules.dep Datei und da dort die vbox* Module drin stehen nachdem das depmod -a command ausgeführt wurde können diese gefunden werden.

Depmode benutze alle Verzeichnis unter /lib/modules/<Kernelversion> um die modules.dep Datei zu erstellen.

This file is used by modprobe to know the order in which to load modules (they are loaded right to left, and removed left to right).

Auszug aus Datei:
{% highlight text %}
extramodules/vboxsf.ko.gz: extramodules/vboxguest.ko.gz
extramodules/vboxnetflt.ko.gz: extramodules/vboxdrv.ko.gz
extramodules/vboxvideo.ko.gz: kernel/drivers/gpu/drm/drm.ko.gz kernel/drivers/i2c/i2c-core.ko.gz
extramodules/vboxnetadp.ko.gz: extramodules/vboxdrv.ko.gz
extramodules/vboxguest.ko.gz:
extramodules/vboxpci.ko.gz: extramodules/vboxdrv.ko.gz
extramodules/vboxdrv.ko.gz
{% endhighlight %}

Ein weiter Programm um Module zu verwlaten ist kmod (Archlinux erstatz für module-init-tools)

Quelle
http://wiki.ubuntuusers.de/Kernelmodule

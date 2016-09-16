# haproxy #

## Benutzerverwaltung ##

Möchte man das nur bestimmte Benutzer Zugang auf ein vom HAProxy verwaltetes Backend haben. Kann man dies unterandem über die HTTP Basic Authentication handhaben. Die Standard Lösung wäre userlist zu verwenden das funktioniert prima bis zu einer besimmten Anzahl von Benutzern. Da man für jeden neu hinzugefügten/entferneten Benutzer haproxy neustarten muss und das ziemlich lang dauert zu Problemen führen. Daher kann man alternative mittels Lua Script das ganze ncoh aufbohren. 

Beispiel userlist
Beispiel luascript
Beispiel luascript ssl


## Lua ##

## links ##
* https://serversforhackers.com/using-ssl-certificates-with-haproxy
* https://www.haproxy.com/doc/aloha/7.0/haproxy/tls.html
* https://raymii.org/s/snippets/haproxy_ssl_backends.html
* https://www.digitalocean.com/community/tutorials/an-introduction-to-haproxy-and-load-balancing-concepts
* https://serversforhackers.com/load-balancing-with-haproxy
* http://notes.iopush.net/configure-haproxy-in-reverse-proxy-with-http-authentification/
* http://blog.haproxy.com/2013/09/16/howto-transparent-proxying-and-binding-with-haproxy-and-aloha-load-balancer/

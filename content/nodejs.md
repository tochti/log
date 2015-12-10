+++
date = "2015-12-03T17:04:15+01:00"
draft = true
title = "node.js"
+++


Basics
------

Beschäftigt man sich zum ersten mal mit Node.js fallen schnell Begriffe wie non-blocking, event-driven, async und EventLoop. Dann liest man auf der anderen Seite aber wieder das JavaScript single thread ist also nur ein Befehl nachdem andern ausführen kann. Wie passt das alles zusammen? Zunächst einmal ist es so das wir uns hier im Land der Bits und Bytes befinden das heißt es gibt immer einen weg etwas zum machen wie es nicht vorgesehen ist eine weg um die Idee herum. Es ist fast alles möglich. Der Versuch ist hier die Grundlegenden Konzepte zu begutachten.

Als ich mich zum erstenmal mit Node.js auseiander gesetzt hatte wollte ich ein simple Programm schreiben nichts non-blocking, nichts async oder sonstiges. Da stellte sich mir dir Frage wie mach ich das ist in Node.js ist nicht alles non-blocking, event-driven, etc.? Wie sich heruasstellt ist es gar nicht so schwer sync Code zu schreiben, man muss dazu nämlich garnichts beachten solange man nicht die Node.js Bibliotheken verwendet wie zum Beispiel fs oder http. Erst wenn man die von Node.js mitgelieferten Bibliotheken, welche meistens in C++ geschrieben sind, verwenden wird die Sache "non-blocking" und eine "EventLoop". 

Wollen wir nichts weitermachen als ein paar Paar Roboter Tanzen zu lassen passiert noch nicht viel anders als JavaScript

```JavaScript
var Robo = function(n) {
  return {
    name: n,
    dance: function() {
      console.log(n + ' is dancing the happy dance!');
    }
  }
};

['Klaus', 'Peter', 'Hans'].forEach(function(n) {
  r = new Robo(n);
  r.dance();
});
```

Aber in der Werbung sprechen sie doch alle von async. Wie läuft das nun alles ab?

Hier ein paar Beispiele und meine Gedanken dazu.

Wenn man nun schon einmal etwas mit anderen Programmiersprachen zutun hatte wie z. Bsp. [Go](golang.org) ist man vielleicht ein wenig vertraut mit dem Konzept [Cuncurrency](https://www.youtube.com/watch?v=cN_DpYBzKso) (Nebenläufigkeit). In Go gibt es den Befehl go um eine sogenannte goroutine zu starten eine Funktionen die mit diesem Befehl gestart wird läuft dann gleichzeitig es wird nicht gewartet bis die Funktionen einen Reutrnwert zurück gibt es wird direkt mit dem nächsten Befehl weitergemacht. Also fragte ich mich ist das bei Node.js ebenso? Ich habe keine von keinem Befehel ähnlich dem go Befehl gehört. Aber es könnte ja sein das einfach alles so funktioniert, wer weiss.

```JavaScript
for (var x = 0; x < 100; x++) {
  console.log('1 - ' + x);
}

for (var x = 0; x < 100; x++) {
  console.log('2 - ' + x);
} 
```

Dann könnte es ja passieren das die console.log(..) ausgaben vermischt werden heißt einmal 1 - 1, 1 - 2, 1 - 3, 2 - 1, 1 - 4. Das passiert jedoch nicht. Auch dann nicht wenn man 1000000 Zeilen auf der Konsole ausgeben würde, der ablauf bleibt synchron. Vielleicht ist es so das nur Funktionen asynchron ausgeführt werden.

```JavaScript
function one() {
  for (var x = 0; x < 100; x++) {
    console.log('1 - ' + x);
  }
}

function two() {
  for (var x = 0; x < 100; x++) {
    console.log('2 - ' + x);
  }
}

one();
two();
```

es beleibt dabei alles schön in einer Reihenfolge.

Zurück blieben einige ?? also erstmal was essen.

```JavaScript
setTimeout(function() {
  console.log('funky bacon!');
}, 2);

for (var i = 100000000; i < 1000000000; i++) {
  if ((i % 5000000) === 0) {
    console.log(Math.log(i));
  }
}
```

setTimeout sollte nach 2 millisekunden funky bacon! auf der Konsole ausgeben macht es aber nicht. Obwohl setTimeout eine von Node.js mitgelieferte non-blocking Funktion ist. Wird funky bacon! immer als letztes auf der Konsole ausgegeben. Aber wir kommen der Sache schon näher. Es wird immerhin schon einmal was ausgegeben auch wenn es viel zu spät ausgegeben wird. Es sieht so aus alle erst der JavaScript Code ausgeführt wird und danach die beantragten Events eingelöst werden.

Noch ein interessantes Beispiel.

```JavaScript
setTimeout(function() {
  console.log('funky bacon!');
}, 2);

for (var i = 100000000; i < 1000000000; i++) {
  if ((i % 5000000) === 0) {
    console.log(Math.log(i));
  }
}

return;

console.log('after return');

```

hier wird funky bacon! als letzte ausgegeben und nicht after return. Es ist sogar so das after return überhaupt nicht ausgegeben wird. Es scheint so als würde EventLoop erst ganz zum Schluß angestart werden. EventLoop jetzt also doch? Schaut man in node.cc sieht man das mittels uv_run_loop eine EventLoop gestart wird. Verwendet man EventLoop sieht Code normalerweis irgendwie so aus

```

// new loop object
loop = new EventLoop()

// config loop
loop.add_event('x', callback)

// Start loop 
loop.run

```

so oder so ähnlich kann man sich das bei Node.js auch vorstellen von Node.js wird eine neues EventLoop Objekt erzeugt. Dann kommt der Programmiere mit JavaScript und konfiguriert die Loop und fügt Events hinzu. Diese events werden aber mittels libuv Bibliothek hinzugefügt auch haben diese events nicht wirklich was mit der events JavaScript Bibliothek zu tun die mit Node.js ausgeliefert wird. Hierzu zwei Beispiele

```JavaScript
setTimeout(function() {
  console.log('Kung Fu!');
}, 2000);
```

```JavaScript
var EventEmitter = require('events');
var events = new EventEmitter();

events.on('Kung', function() {
  console.log('FU!');
});
```

Auch konnte ich Aussagen finden wie "Node.js wird dann beenden wenn keine Events mehr vorhanden sind". Was uns direkt wieder zu unseren zwei Beispielen und unserer oberen Aussage führt. Zu erst warum Beendet sich das erste Programm erst nach 2 Sekunden und das zweite Beende sich sofort, müsste es nicht für immer weiterlaufen da niemand den Event Kung auslöst? Genau aus dem oben genannten Grund events hat nicht wirklich etwas zutun mit der EventLoop die in Node.js gestart wrid. Für die EventLoop, die mittels uv_run_loop gestart wird, können nur Events registriert werden die in [libuv definiert](https://nikhilm.github.io/uvbook/basics.html) sind.Diese Events haben meistens etwas mit Events zu tun die man beim Betriebsystemkernel hinterlegen kann. Das ist nun der Punkt an dem es wirklich non-blocking und asynchron wird.

Hier noch ein paar Wort zu js events von Node.js diese sind nicht asynchron. Vereinfacht gesagt ist es ein Hash in dem Namen mit Arrays von Funktionen verbunden werden. Wird ein Events mittel emit ausgelöst werden alles hinterlegten Funktionen ausgeführt.

```JavaScript
var EventEmitter = require('events');
var emitter = new EventEmitter();

function run(tmp) {
  return function() {
    var no = tmp;
    [1, 2, 3, 4].forEach(function(x) {
      console.log(no + ': ' + x);
    });
  };
}

emitter.on('go', run(1));
emitter.on('go', run(2));
emitter.on('go', run(3));

emitter.emit('go');
```

```JavaScript
var EventEmitter = require('events');
var emitter = new EventEmitter();

function run1() {
  console.log('Enter #1');
  emitter.on('go1', function() {
    console.log('it\'s me #1');
  });
  console.log('Leave #1');
}

function run2() {
  console.log('Enter #2');
  emitter.emit('go1');
  console.log('Leave #2');
}

run1();
run2();
```

Alles ziemlich synchron hier.

Das interessante ist das JavaScript verwendet wurde was wie oben schon geschrieben per Default nicht die beste Unterstüzung an multiproccessing oder threads hat. Das heißt es kann immer nur ein Befehl nacheinadner ausgeführt werden das führt dazu das auch immer nur ein Befehl in einem bestimmten Moment auf den Speicher zugreifen kann. Das heißt man muss sich nicht wie so oft bei Nebeläufigkeit und Pallalism den Kopfzerbrechen über [race conditions](https://de.wikipedia.org/wiki/Race_Condition).


[Node.js Tutorial with Ryan Dahl, creator of Node.js](https://www.youtube.com/watch?v=eqlZD21DME)

Loop
----

http://blog.carbonfive.com/2013/10/27/the-javascript-event-loop-explained/

EventEmitter
https://strongloop.com/strongblog/node-js-event-loop/


Error handling
--------------

https://www.joyent.com/developers/node/design/errors

Pattern
--------
http://book.mixu.net/node/ch7.html

Style Guide
-----------
https://github.com/felixge/node-style-guide#2-spaces-for-indentation

Garbage Collector
-----------------
https://strongloop.com/strongblog/node-js-performance-garbage-collection/

Projekte
---------
Loopback http://loopback.io


Weiter Dokumente
-----------------
https://strongloop.com/strongblog/
http://techblog.netflix.com/
https://medium.com/@tjholowaychuk/callbacks-vs-coroutines-174f1fe66127#.j2a91c7pt
http://syzygy.st/javascript-coroutines/

Analyse JavaScript
------------------
https://perf.wiki.kernel.org/index.php/Main_Page
http://www.brendangregg.com/FlameGraphs/cpuflamegraphs.html
http://nerds.airbnb.com
https://httpd.apache.org/docs/2.2/programs/ab.html
https://iperf.fr/

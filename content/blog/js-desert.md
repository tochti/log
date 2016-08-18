+++
title = "Javascript Wüste! 'this' is not the end"
date = 2016-08-17
type = "post"
categories = ["javascript"]
+++

Umgebung
---------

die JavaScript Konsole von mozilla [console][1]. [ES5][3] das heißt Version 5.1 des Sprachstandards

Damit das Sandel auch spaß macht
```bash
git clone https://github.com/rrawrriw/js-desert.git
cd js-desert 
js // Ready to rumble
```

 Prolog
--------

Die Sonne scheint. Crock sitz in seinem Auto. Wind weht durch seine Barthaare und lässt sie im klang der Stille umher tanzen. "Manchmal ist es einfach gut nur zu 'Sein'!". Jedoch ist es nicht so gut, einfach hier zu 'Sein'. Sitzend, in einem kaputten Auto, in der [V8-Wüste][1]! Der Staub verklebt einem die Augen, saugt jeden versuch der Befeuchtung auf wie ein kleines Kind die versteckte Schokolade. Wobei die Nützlichkeit eines Wimbernschlags eines Echos gleichkommt, das erzeugt wurde durch einen einsamen Stein in einem leeren Raum, in dem Wände, Boden und fallender Stein taub sind. "'Sein' möchte ich wo anders sein".


Überlebens-Regel Nr.1 Erkunde die Umgebung
------------------------------------------

Überleben in der Wüste für Fortgeschrittene. Such die Spuren eines verwirrten Kamels. 

Lasst uns in die Wüste gehen und schauen was für große Abenteuer uns erwarten. Führe js auf der Konsole aus und tauche ab in das große Sandelabenteuer das weder feucht noch fröhlich ist.

```javascript
js> load('desert.js');
```
  
Nehmen wir uns ein wenig Zeit um die Umgebung zu erkunden.

```javascript
js> for (o in desert) { print(o) };
desert
sand
```

Sandel, sandel.
```javascript
js>  print("funky!");
```

Die unbeteiligte Stimme aus dem OFF kann sich vorstellen dass, falls ihr noch nicht verdurstet seit, die folgende Frage schwer auf euren Schultern lastet. "Zur Hölle was machen Objekte wie desert und sand in der standard Umgebung?". Diese Objekte haben wir uns eingefangen als wir desert.js geladen hatten. 

```javascript
desert = this;
```

*Solangsam ist das hier alles nicht mehr lustig ich habe durst!*

```javascript
js> beer
typein:8: ReferenceError: beer is not defined
```

*Was ist this, kein Bier?*

```javascript
js> desert.hasOwnProperty('beer')
false
```
*Verdammt, wirklich kein Bier? Jetzt mal mit Ernst!*

Erlöst sollt ihr sein, es soll Bier entstehen.

```javascript
js> var beer = Object.create(Object, {
  status: {
    writeable: false,
    value: 'empty'
  }
});
js> beer.status
"empty"
```
*schluchz*

```javascript
js> beer.status = 'full';
"full"
js> beer.status;
"empty"
```
*so gemein*

Schmollend zeihen wir uns hinter die nächste Düne zurück.

```javascript
js> var dune = {};
js> dune.behind = function () {
  var fn = function () {
    print(this.beer.status)
  };

  fn();
};
js> dune.behind();
empty
```

Mach dir keine Hoffnungen!

Durch Panik getrieben rennen wir mit letzten Kräften richtung Nord. Norden bedeutet kälte, außer man befindet sich auf der untern Hälfte der Südehalbkugel, oder ist zu Fuß in der Wüste unterwegs. Nach unglaublichen langen Sekunden brechen wir zusammen, überlassen uns wimmernt dem Tod, soll er doch kommen! Fieberträume bringen uns Bilder von schönen Frauen.

```javascript
js> dream = function () {
  print('beauty is cuddling', this.me)
}
(function () {print("beauty is cuddling", this.me);})
js> dream()
beauty is cuddling undefined
```
*Halt das ist mein Traum*

```javascript
js> me = {
  me: "Crock"
};
({me:"Crock"})
js> dream.apply(me)
beauty is cuddling Crock
```
*Schon besser*

Doch bekanntlich holt einen die Realität schneller ein als gewünscht.

```javascript
js>  reality = function (person) {
  print(person.me, "is fighting with a camel spider")
};
(function (person) {print(person.me, "is fighting with a camel spider");})
js> reality(me)
Crock is fighting with a camel spider
js> theSpider = {
  toxic: "100%",
  kills: function () {
    print("the spider kills "+ this.me +" to "+ this.toxic)
  }
}
({toxic:"100%", kills:(function () {print("the spider kills " + this.me + " to " + this.toxic);})})
js> Spider = function () {
  return Object.create(theSpider)
}
(function () {return Object.create(theSpider);})
js> spider = new Spider()
({})
js> me.toxic = "0%, instead brings him a beer"
"0%, instead brings him a beer"
js> spider.kills.apply(me)
the spider kills Crock to 0%, instead brings him a beer
```
*Spinne am Arsch! Es lebe die Dynamik.*

Wie von Blitz getroffen fällt es uns ein hatte wir nicht gelesen das Forscher herausgefunden hatten das die Urspinne einen Gendefekt hatte. Dieser Gendefekt führt bei Berührung der Achillesferse dazu das sich die Spinne spontan in eine Flasche Bier verwandelt.
```javascript
js> theSpider.achillesheel = function () {
  this.status = "tasty, full and cool beer!"
}
(function () {this.status = "tasty, full and cool beer!";})
js> spider.__proto__
({toxic:"100%", kills:(function () {print("the spider kills " + this.me + " to " + this.toxic);}), achillesheel:(function () {this.status = "tasty, full and cool beer!";})})
js> spider.achillesheel() 
js> spider.status
"tasty, full and cool beer!"
```
Gelernt ist gelernt! Das ist auch schon das verwirrte Kamel
```javascript
js> desert.lostCamel = {
  ride: function () {
    person = arguments[0];
    print("bring "+ person.me +" to Perlland");
  }
}
({ride:(function () {person = arguments[0];print("bring " + person.me + " to Perlland");})})
js> lostCamel.ride(me)
bring Crock to Perlland
```
Good Bye Partner!

[1]: https://developer.mozilla.org/en-US/docs/Mozilla/Projects/SpiderMonkey/Introduction_to_the_JavaScript_shell "mozilla shell"
[2]: http://en.wikipedia.org/wiki/List_of_ECMAScript_engines "JavaScript Engines"
[3]: http://www.ecma-international.org/publications/standards/Ecma-262.htm "ES5"

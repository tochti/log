+++
date = "2015-11-25T09:07:27+01:00"
draft = false
title = "Backbone"
+++

* Trailing Slash bei URL poroperty.
* events bei view können nur auf Elemente gesetzt werden die sich im $el Element befinden.
* tastypie always_return_data = True damit id zugeordnet werden kann ansonsten wird kein json zurückgegeben.
* wenn tagName nicht gesetzt wird verwendet backbone div. el bekommt ein elemet aus der Webseite zugewiesen. el: "elementName". tagName erzeugt ein Container innerhalb diesem wird der gerenderte html code eingefügt. el referenziert auf das so erzeugt element. wenn kein el angegeben wird erzeugt backbone eines default ist div ohne class und id diese div wird auch nicht mit der aktuellen page verbunden. el: $("#elementInPage")
* this.el gibt html code zurück wenn bereits definiert this.$el.html("html code")
* benutze urlRoot in models. wird model.destroy() aufgerufen kommte es ansonst zu Problemen und es werden alle Element aus der Datenbank gelöscht was wirklich nicht gut ist.
* Wenn man object über eine collection abrufen möchte muss man dazu eine url eigentschaft festlegen kann aber nicht übergeben werden muss aus irgendwelchen gründen object.url = "url" erzeugt werden.
* wenn man object von einer db mittels collection abrufen möchte und die REST Api mittels tastypie bereitgestellt werden muss die parse methode angepasst werden oder der return string von tastypie geändert werden da immer meta daten von tastypie mitgeliefert werden.
* keyup event triggert das event erst nach dem eine Tastegerdückt wurde
* wait = true muss übergeben werden wenn mittel create ein object an eine collection übergeben werden soll und danach das object gleich gerendert werden soll da create sonst einen async call an den server sendet und es sein kann das zur renderzeit noch nicht alle attribute zurverfügung stehen wie zum Beispiel die ID
* tastypie kann so verwenden das man mittels feldname__djangoquerystring nach dingen suche lassen kann bsp. name__contains="something"

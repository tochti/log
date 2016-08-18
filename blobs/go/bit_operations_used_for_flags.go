package main

import "fmt"

// Speicher Flags in einem Integer und überprüfe ob Flag aktiv ist.
func main() {
	// Prinzipelle Idee. Wir nehmen an wir haben 3 Attribute.
	// Diese werden representiert durch die bits
	// 0 0 0 (Dezimal: 0) aktivieren wir nun das
	// erste Attribute sehen die bits wie folgt aus 0 0 1 (Dezimal: 1) würde
	// man nun noch das Zweite Attribute aktivieren würde das ganze wie
	// folgt aussehen 0 1 1 (Dezimal: 3).
	// Möchte man nun wissen ob das erste Attribut gesetzt ist könnte man
	// Mittels dem UND Operator und der richtigen bit folge überprüfen ob ein
	// Bit gestzt ist.
	//   0 1 1
	// & 0 0 1
	// -------
	//   0 0 1 != 0 -> somit aktiv
	// Möchte man wissen ob das zweite Attribut gesetzt ist könnte müsste man folgende Operation vornehmen
	//   0 1 1 (Dezimal: 3)
	// & 0 1 0 (Dezimal: 2)
	// -------
	//   0 1 0 != 0 -> somit aktiv
	// Die benötigte "Vergleichs Bitfolge" kann mittel dem Shift Operator erzeugen werden.
	// Beispiel für Attribut 1:
	// 1 << 0 = 0 0 1
	// Beispiel für Attribute 2
	// 1 << 1 = 0 1 0
	// Beispiel für Attribute 3
	// 1 << 2 = 1 0 0

	// Monster Attribute
	// Hunger: ja/nein
	// Gefährlich für Menschen: ja/nein

	// Zombie
	// Hunger: ja
	// Gefährlich für den Menschen: ja
	fmt.Println("--- Zombie --")
	zombie := 3

	fmt.Printf("Attribute: %b\n", zombie)
	fmt.Printf("Hungrig? %v\n", zombie&(1<<0) != 0)
	fmt.Printf("Gefährlich für den Menschen? %v\n", zombie&(1<<1) != 0)
	fmt.Println("Run away!")

	// Krümmelmonster
	// Hunger: ja
	// Gefährlich für den Menschen: nein
	fmt.Println("--- Krümmelmonster ---")
	kruemmelmonster := 1 << 0
	// alternaive
	kruemmelmonster = 1

	fmt.Printf("Attribute: %b\n", kruemmelmonster)
	fmt.Printf("Hungrig? %v\n", kruemmelmonster&(1<<0) != 0)
	fmt.Printf("Gefährlich für den Menschen? %v\n", kruemmelmonster&(1<<1) != 0)

	// Geist
	// Hunger: nein
	// Gefährlich für den Mensch: ja
	fmt.Println("--- Geist ---")
	ghost := 1 << 1

	fmt.Printf("Attribute: %b\n", ghost)
	fmt.Printf("Hungrig? %v\n", ghost&(1<<0) != 0)
	fmt.Printf("Gefährlich für den Menschen? %v\n", ghost&(1<<1) != 0)

}

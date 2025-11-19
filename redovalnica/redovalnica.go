/*
Package redovalnica

Paket redovalnica vsebuje preproste podatke in funkcije za upravljanje ocen
študentov v majhnem sistemu. Namenjen je kot učni primer: definiramo strukturo
Student in osnovne operacije, kot so dodajanje ocene, izpis vseh ocen in
izpis končnega uspeha študenta.

Eksportirane spremenljivke:
	- StOcen (int): najmanjše število ocen, potrebnih, da se izračuna povprečje
		(privzeto 6). Aplikacija (npr. CLI) lahko to nastavi pred klici funkcij.
	- MinOcena, MaxOcena (float64): najmanjša in največja dovoljena vrednost
		ocene; DodajOceno vrne napako (izpis) kadar je ocena zunaj tega intervala.

Glavne izbrane funkcije:
	- DodajOceno(studenti, vpisnaStevilka, ocena): doda oceno študentu, če
		študent obstaja in je ocena znotraj dovoljenega intervala.
	- IzpisVsehOcen(studenti): izpiše vse študente in njihove ocene.
	- IzpisiKoncniUspeh(studenti): za vsakega študenta izračuna povprečje
		(če ima vsaj StOcen ocen) in izpiše oceno (Odličen/Povprečen/Neuspešen).
	- NovStudent(ime, priimek): konstruktor za ustvarjanje vrednosti Student iz
		drugih paketov (polja strukture niso izvožena).

Uporaba tipično vključuje inicializacijo konfiguracije (nastavitev StOcen,
MinOcena, MaxOcena) in nato klice DodajOceno/IzpisVsehOcen/IzpisiKoncniUspeh.
*/

package redovalnica

import (
	"fmt"
)

var StOcen = 6
var MinOcena float64 = 0
var MaxOcena float64 = 10

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		fmt.Println("Napaka: študent z vpisno številko", vpisnaStevilka, "ne obstaja.")
		return
	}

	if float64(ocena) < MinOcena || float64(ocena) > MaxOcena {
		fmt.Printf("Napaka: ocena mora biti med %.2f in %.2f.\n", MinOcena, MaxOcena)
		return
	}

	st.ocene = append(st.ocene, ocena)
	studenti[vpisnaStevilka] = st
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		return -1.0
	}

	if len(st.ocene) < StOcen {
		return 0
	}

	var vsota int
	for _, ocena := range st.ocene {
		vsota += ocena
	}
	return float64(vsota) / float64(len(st.ocene))
}

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, st := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, st.ime, st.priimek, st.ocene)
	}
}
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {
		p := povprecje(studenti, vpisna)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.ime, student.priimek, p)

		if p >= 9 {
			fmt.Println("Odličen študent!")
		} else if p >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}

func NovStudent(ime, priimek string) Student {
	return Student{ime: ime, priimek: priimek, ocene: []int{}}
}

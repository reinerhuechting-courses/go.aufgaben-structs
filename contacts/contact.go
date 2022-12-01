package contacts

import "fmt"

// Datentyp für Kontakte (z.B. Adressbucheinträge)
// mit Vorname, Name, Telefon und einer Möglichkeit,
// Tags anzugeben. Tags können z.B. so etwas wie "Freunde",
// "Familie", "Wichtig", "Privat", "Geschäftlich" etc. sein.
type Contact struct {
	Name, Surname string
	Phone         string
	Tags          []string
}

// Liefert eine String-Repräsentation eines Kontakts.
func (contact Contact) String() string {
	return fmt.Sprintf("%s %s\n  Telefon: %s\n  Tags: %s\n",
		contact.Name,
		contact.Surname,
		contact.Phone,
		contact.Tags,
	)
}

// Gibt an, ob ein Kontakt einen bestimmten Tag besitzt.
func (contact Contact) HasTag(tag string) bool {
	for _, t := range contact.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// Fügt einen Tag zu einem Kontakt hinzu.
// Soll keine doppelten Einträge erzeugen.
func (contact *Contact) AddTag(tag string) {
	if !contact.HasTag(tag) {
		contact.Tags = append(contact.Tags, tag)
	}
}

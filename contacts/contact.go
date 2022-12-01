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

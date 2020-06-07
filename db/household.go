package db

// Household is the struct representation of a Household record
type Household struct {
	ID   int
	Type string
}

// ValidateHousehold will validate fields of household
func (h *Household) ValidateHousehold() error {
	if h.Type != "Landed" && h.Type != "Condominium" && h.Type != "HDB" {
		return ErrHouseholdTypeInvalid
	}
	return nil
}

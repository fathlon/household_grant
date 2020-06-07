package model

// Household is the struct representation of a Household record
type Household struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// Validate will validate fields of household
func (h *Household) Validate() error {
	if h.Type != "Landed" && h.Type != "Condominium" && h.Type != "HDB" {
		return ErrHouseholdTypeInvalid
	}
	return nil
}

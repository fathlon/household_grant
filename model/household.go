package model

// Household is the struct representation of a Household record
type Household struct {
	ID      int            `json:"id"`
	Type    string         `json:"type"`
	Members []FamilyMember `json:"members,omitempty"`
}

// Validate will validate fields of household
func (h *Household) Validate() error {
	if h.Type != "Landed" && h.Type != "Condominium" && h.Type != "HDB" {
		return ErrHouseholdTypeInvalid
	}
	return nil
}

// AddMember adds a family member into the household
func (h *Household) AddMember(f FamilyMember) {
	if f.ID != 0 && h.MemberExists(f.ID) {
		return
	}
	h.Members = append(h.Members, f)
}

// MemberExists checks if member of the given id exists
func (h *Household) MemberExists(id int) bool {
	for _, f := range h.Members {
		if f.ID == id {
			return true
		}
	}
	return false
}

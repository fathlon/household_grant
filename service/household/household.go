package household

import (
	"github.com/fathlon/household_grant/db"
	"github.com/fathlon/household_grant/model"
	"github.com/fathlon/household_grant/service"
)

// Create is the service method for creating household
func Create(ds *db.Datastore, h model.Household) (model.Household, error) {
	if err := h.Validate(); err != nil {
		return model.Household{}, service.NewValidationError(err)
	}

	return ds.CreateHousehold(h)
}

// AddMember adds a family member to the household of the given ID
func AddMember(ds *db.Datastore, householdID int, f model.FamilyMember) (model.FamilyMember, error) {
	if err := f.Validate(); err != nil {
		return model.FamilyMember{}, service.NewValidationError(err)
	}

	h, err := ds.RetrieveHousehold(householdID)
	if err != nil {
		return model.FamilyMember{}, service.NewValidationError(err)
	}

	if f.ID == 0 {
		f, err = ds.CreateFamilyMember(f)
		if err != nil {
			return model.FamilyMember{}, service.NewValidationError(err)
		}
	}

	h.AddMember(f)
	// save updated household back into DB
	ds.UpdateHousehold(h)

	return f, nil
}

// RetrieveAll returns all households in the db
func RetrieveAll(ds *db.Datastore) []model.Household {
	return ds.RetrieveHouseholds()
}

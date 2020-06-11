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

	// clear out family member to prevent saving of members through creation
	h.Members = []model.FamilyMember{}

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

// Retrieve returns the household of the given id
func Retrieve(ds *db.Datastore, householdID int) (model.Household, error) {
	result, err := ds.RetrieveHousehold(householdID)
	if err != nil {
		return model.Household{}, service.NewValidationError(err)
	}

	return result, nil
}

// Delete deletes the given household by id and returns the deleted household
func Delete(ds *db.Datastore, householdID int) (model.Household, error) {
	result, err := ds.DeleteHousehold(householdID)
	if err != nil {
		return model.Household{}, service.NewValidationError(err)
	}

	return result, nil
}

// DeleteMember deletes the given family member from household by ids and returns the deleted family member
func DeleteMember(ds *db.Datastore, householdID, memberID int) (model.FamilyMember, error) {
	h, err := ds.RetrieveHousehold(householdID)
	if err != nil {
		return model.FamilyMember{}, service.NewValidationError(err)
	}

	if !h.MemberExists(memberID) {
		return model.FamilyMember{}, service.NewValidationError(model.ErrHouseholdFamilyMemberNotExists)
	}

	h.DeleteMember(memberID)
	// save updated household back into DB
	ds.UpdateHousehold(h)

	result, err := ds.DeleteFamilyMember(memberID)
	if err != nil {
		return model.FamilyMember{}, service.NewValidationError(err)
	}

	return result, nil
}

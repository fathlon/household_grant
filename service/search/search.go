package search

import (
	"github.com/fathlon/household_grant/db"
	"github.com/fathlon/household_grant/model"
	"github.com/fathlon/household_grant/service"
)

// Retrieves returns all results which matches the given search operation
func Retrieves(ds *db.Datastore, so model.SearchOperation) ([]model.Household, error) {
	if err := so.Validate(); err != nil {
		return []model.Household{}, service.NewValidationError(err)
	}

	result := ds.Search(so)

	// remove member slice and return only household
	if so.WholeHousehold {
		for i, r := range result {
			r.Members = []model.FamilyMember{}
			result[i] = r
		}
	}

	return result, nil
}

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

	return ds.AddHousehold(h)
}

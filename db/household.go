package db

import "github.com/fathlon/household_grant/model"

// AddHousehold takes a Housedhold object and saves it into the datastore
func (d *Datastore) AddHousehold(h model.Household) (model.Household, error) {
	currentIdx := nextHseIndex()
	if _, exist := d.households[currentIdx]; exist {
		return model.Household{}, ErrHouseholdDuplicateID
	}

	h.ID = currentIdx
	d.households[currentIdx] = h

	return h, nil
}

// RetrieveHousehold retrieves the household of the given ID
func (d *Datastore) RetrieveHousehold(id int) (model.Household, error) {
	result, ok := d.households[id]
	if !ok {
		return model.Household{}, ErrHouseholdNotFound
	}

	return result, nil
}

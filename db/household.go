package db

import "github.com/fathlon/household_grant/model"

// CreateHousehold takes a Housedhold object and saves it into the datastore
func (d *Datastore) CreateHousehold(h model.Household) (model.Household, error) {
	currentIdx := nextHseIndex()
	if _, exist := d.Households[currentIdx]; exist {
		return model.Household{}, ErrHouseholdDuplicateID
	}

	h.ID = currentIdx
	d.Households[currentIdx] = h

	return h, nil
}

// RetrieveHousehold retrieves the household of the given ID
func (d *Datastore) RetrieveHousehold(id int) (model.Household, error) {
	result, ok := d.Households[id]
	if !ok {
		return model.Household{}, ErrHouseholdNotFound
	}

	return result, nil
}

// UpdateHousehold updates the given household into the datastore
func (d *Datastore) UpdateHousehold(h model.Household) error {
	if h.ID == 0 {
		return ErrHouseholdInvalid
	}
	d.Households[h.ID] = h

	return nil
}

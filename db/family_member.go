package db

import "github.com/fathlon/household_grant/model"

// CreateFamilyMember takes a FamilyMember object and saves it into the datastore
func (d *Datastore) CreateFamilyMember(f model.FamilyMember) (model.FamilyMember, error) {
	currentIdx := nextMemIndex()
	if _, exist := d.Members[currentIdx]; exist {
		return model.FamilyMember{}, ErrFamilyMemberDuplicateID
	}

	f.ID = currentIdx
	d.Members[currentIdx] = f

	return f, nil
}

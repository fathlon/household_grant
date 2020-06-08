package db

import "github.com/fathlon/household_grant/model"

// AddFamilyMember takes a FamilyMember object and saves it into the datastore
func (d *Datastore) AddFamilyMember(f model.FamilyMember) (model.FamilyMember, error) {
	currentIdx := nextMemIndex()
	if _, exist := d.members[currentIdx]; exist {
		return model.FamilyMember{}, ErrFamilyMemberDuplicateID
	}

	f.ID = currentIdx
	d.members[currentIdx] = f

	return f, nil
}

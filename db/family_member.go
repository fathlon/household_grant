package db

import "github.com/fathlon/household_grant/model"

// CreateFamilyMember takes a FamilyMember object and saves it into the datastore
func (d *Datastore) CreateFamilyMember(f model.FamilyMember) (model.FamilyMember, error) {
	if nameExists(f.Name, d.Members) {
		return model.FamilyMember{}, ErrFamilyMemberDuplicateName
	}

	currentIdx := nextMemIndex()
	if _, exist := d.Members[currentIdx]; exist {
		return model.FamilyMember{}, ErrFamilyMemberDuplicateID
	}

	f.ID = currentIdx
	d.Members[currentIdx] = f

	return f, nil
}

func nameExists(name string, members map[int]model.FamilyMember) bool {
	for _, f := range members {
		if f.Name == name {
			return true
		}
	}
	return false
}

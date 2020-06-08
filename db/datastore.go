package db

import "github.com/fathlon/household_grant/model"

var (
	hseIndex int
	memIndex int
)

// Datastore is the struct represetation of a persistent datastore
type Datastore struct {
	households map[int]model.Household
	members    map[int]model.FamilyMember
}

// NewDatastore is the constructor for initializing and creating a new datastore
func NewDatastore() *Datastore {
	return &Datastore{
		households: make(map[int]model.Household),
		members:    make(map[int]model.FamilyMember),
	}
}

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

// nextHseIndex returns the next index value
func nextHseIndex() int {
	hseIndex++
	return hseIndex
}

// nextMemIndex returns the next index value
func nextMemIndex() int {
	memIndex++
	return memIndex
}

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

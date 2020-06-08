package db

import "github.com/fathlon/household_grant/model"

var (
	HseIndex int
	MemIndex int
)

// Datastore is the struct represetation of a persistent datastore
type Datastore struct {
	Households map[int]model.Household
	Members    map[int]model.FamilyMember
}

// NewDatastore is the constructor for initializing and creating a new datastore
func NewDatastore() *Datastore {
	return &Datastore{
		Households: make(map[int]model.Household),
		Members:    make(map[int]model.FamilyMember),
	}
}

// nextHseIndex returns the next index value
func nextHseIndex() int {
	HseIndex++
	return HseIndex
}

// nextMemIndex returns the next index value
func nextMemIndex() int {
	MemIndex++
	return MemIndex
}

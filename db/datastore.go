package db

var index int

// Datastore is the struct represetation of a persistent datastore
type Datastore struct {
	store map[int]Household
}

// NewDatastore is the constructor for initializing and creating a new datastore
func NewDatastore() *Datastore {
	return &Datastore{
		store: make(map[int]Household),
	}
}

// AddHousehold takes a Housedhold object and saves it into the datastore
func (d *Datastore) AddHousehold(h Household) (Household, error) {
	currentIdx := nextIndex()
	if _, exist := d.store[currentIdx]; exist {
		return Household{}, ErrDuplicateID
	}

	h.ID = currentIdx
	d.store[currentIdx] = h

	return h, nil
}

// nextIndex returns the next index value
func nextIndex() int {
	index++
	return index
}

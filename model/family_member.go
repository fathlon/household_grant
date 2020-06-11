package model

import "time"

// FamilyMember is the struct representation of a FamilyMember record
type FamilyMember struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Gender         string    `json:"gender"`
	MaritalStatus  string    `json:"marital_status"`
	Spouse         string    `json:"spouse"`
	OccupationType string    `json:"occupation_type"`
	AnnualIncome   int       `json:"annual_income"`
	DOB            time.Time `json:"dob"`
}

// Validate will validate fields of family member
func (f *FamilyMember) Validate() error {
	if f.Name == "" {
		return ErrFamilyMemberNameInvalid
	}

	if f.Gender != "M" && f.Gender != "F" {
		return ErrFamilyMemberGenderInvalid
	}

	if f.OccupationType != "Unemployed" && f.OccupationType != "Student" && f.OccupationType != "Employed" {
		return ErrFamilyMemberOccupationTypeInvalid
	}

	if f.MaritalStatus != "Single" && f.MaritalStatus != "Married" {
		return ErrFamilyMemberMaritalStatusInvalid
	}

	now := time.Now()
	if f.DOB.IsZero() || f.DOB.Local().After(now) {
		return ErrFamilyMemberDOBInvalid
	}

	return nil
}

// Age returns current age based on DOB
func (f *FamilyMember) Age() int {
	now := time.Now()
	dobYear := f.DOB.Local().Year()

	currentAge := now.Year() - dobYear
	monthChecker := time.Date(dobYear, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)

	if monthChecker.Before(f.DOB.Local()) {
		currentAge--
	}

	return currentAge
}

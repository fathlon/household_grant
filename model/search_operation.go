package model

// SearchOperation is the struct representation containing all possible search functions
type SearchOperation struct {
	// Household conditions
	HouseholdIncomeGT int `json:"household_income_gt"`
	HouseholdIncomeLT int `json:"household_income_lt"`
	HouseholdIncome   int `json:"household_income"`

	HouseholdSizeGT int `json:"household_size_gt"`
	HouseholdSizeLT int `json:"household_size_lt"`
	HouseholdSize   int `json:"household_size"`

	HasCouple   bool `json:"has_couple"`
	HasChildren bool `json:"has_children"`

	// FamilyMember conditions
	AnnualIncomeGT int `json:"annual_income_gt"`
	AnnualIncomeLT int `json:"annual_income_lt"`
	AnnualIncome   int `json:"annual_income"`

	AgeGT int `json:"age_gt"`
	AgeLT int `json:"age_lt"`
	Age   int `json:"age"`

	// Response
	WholeHousehold bool `json:"whole_household"`

	// Flags
	CompareHouseholdIncome bool
	CompareHouseholdSize   bool
	CompareAnnualIncome    bool
	CompareAge             bool
}

// Validate checks if there are conflicting comparison specified
func (so *SearchOperation) Validate() error {
	if so.HouseholdIncome != 0 && (so.HouseholdIncomeGT != 0 || so.HouseholdIncomeLT != 0) {
		return ErrSearchConflictingComparison
	}

	if so.HouseholdSize != 0 && (so.HouseholdSizeGT != 0 || so.HouseholdSizeLT != 0) {
		return ErrSearchConflictingComparison
	}

	if so.Age != 0 && (so.AgeGT != 0 || so.AgeLT != 0) {
		return ErrSearchConflictingComparison
	}

	if so.AnnualIncome != 0 && (so.AnnualIncomeGT != 0 || so.AnnualIncomeLT != 0) {
		return ErrSearchConflictingComparison
	}

	return nil
}

// SetFlags update SearchOperation flags for retrieval check
func (so *SearchOperation) SetFlags() {
	if so.HouseholdIncome != 0 || so.HouseholdIncomeGT != 0 || so.HouseholdIncomeLT != 0 {
		so.CompareHouseholdIncome = true
	}

	if so.HouseholdSize != 0 || so.HouseholdSizeGT != 0 || so.HouseholdSizeLT != 0 {
		so.CompareHouseholdSize = true
	}

	if so.Age != 0 || so.AgeGT != 0 || so.AgeLT != 0 {
		so.CompareAge = true
	}

	if so.AnnualIncome != 0 || so.AnnualIncomeGT != 0 || so.AnnualIncomeLT != 0 {
		so.CompareAnnualIncome = true
	}
}

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
}

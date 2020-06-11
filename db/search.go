package db

import (
	"github.com/fathlon/household_grant/model"
)

// Search retrieves household and family members which matches the search criteria
func (d *Datastore) Search(so model.SearchOperation) []model.Household {
	// return empty if no search criteria specified
	if (so == model.SearchOperation{}) {
		return []model.Household{}
	}

	so.SetFlags()
	result := []model.Household{}

	totalHouseholdIncomeMap := d.sumHouseholdIncome()

	for hID, h := range d.Households {
		// skip household with no members
		if len(h.Members) == 0 {
			continue
		}

		matchedHousehold := h
		matchedHousehold.Members = []model.FamilyMember{}

		// skip if criteria specified but do not match
		if so.CompareHouseholdIncome && !matchHouseholdIncome(totalHouseholdIncomeMap[hID], so) {
			continue
		}

		if so.CompareHouseholdSize && !matchHouseholdSize(len(h.Members), so) {
			continue
		}

		if so.HasCouple != nil && containsCouple(h.Members) != *so.HasCouple {
			continue
		}

		if so.HasChildrenByAge != 0 && !containsChildrenByAge(so.HasChildrenByAge, h.Members) {
			continue
		}

		for _, f := range h.Members {
			if so.CompareAnnualIncome && !matchAnnualIncome(f.AnnualIncome, so) {
				continue
			}

			if so.CompareAge && !matchAge(f.Age(), so) {
				continue
			}
			matchedHousehold.Members = append(matchedHousehold.Members, f)
		}

		// skip household if no member in household matched criteria
		if len(matchedHousehold.Members) == 0 {
			continue
		}

		result = append(result, matchedHousehold)
	}

	return result
}

// containsCouple loops the members slice and checks if both couple lives in same household
func containsCouple(members []model.FamilyMember) bool {
	spouseMap := make(map[string]string)
	for _, f := range members {
		if f.Spouse != "" {
			if _, exists := spouseMap[f.Name]; exists {
				return true
			}
			// value is only placeholder value
			spouseMap[f.Name] = ""

			if _, exists := spouseMap[f.Spouse]; exists {
				return true
			}
			// add both member and spouse to map
			spouseMap[f.Spouse] = ""
		}
	}

	return false
}

// containsChildrenByAge loops the members slice and checks if any member is below the given age
func containsChildrenByAge(val int, members []model.FamilyMember) bool {
	for _, f := range members {
		if f.Age() < val {
			return true
		}
	}
	return false
}

func matchHouseholdIncome(val int, so model.SearchOperation) bool {
	if so.HouseholdIncome != 0 && so.HouseholdIncome != val {
		return false
	}
	if so.HouseholdIncomeGT != 0 && val <= so.HouseholdIncomeGT {
		return false
	}
	if so.HouseholdIncomeLT != 0 && val >= so.HouseholdIncomeLT {
		return false
	}

	return true
}

func matchHouseholdSize(val int, so model.SearchOperation) bool {
	if so.HouseholdSize != 0 && so.HouseholdSize != val {
		return false
	}
	if so.HouseholdSizeGT != 0 && val <= so.HouseholdSizeGT {
		return false
	}
	if so.HouseholdSizeLT != 0 && val >= so.HouseholdSizeLT {
		return false
	}

	return true
}

func matchAnnualIncome(val int, so model.SearchOperation) bool {
	if so.AnnualIncome != 0 && so.AnnualIncome != val {
		return false
	}
	if so.AnnualIncomeGT != 0 && val <= so.AnnualIncomeGT {
		return false
	}
	if so.AnnualIncomeLT != 0 && val >= so.AnnualIncomeLT {
		return false
	}

	return true
}

func matchAge(val int, so model.SearchOperation) bool {
	if so.Age != 0 && so.Age != val {
		return false
	}
	if so.AgeGT != 0 && val <= so.AgeGT {
		return false
	}
	if so.AgeLT != 0 && val >= so.AgeLT {
		return false
	}

	return true
}

// sumHouseholdIncome calculates the total household income and returns as map
func (d *Datastore) sumHouseholdIncome() map[int]int {
	result := make(map[int]int)
	for hID, h := range d.Households {
		sum := 0
		for _, fm := range h.Members {
			sum += fm.AnnualIncome
		}
		result[hID] = sum
	}
	return result
}

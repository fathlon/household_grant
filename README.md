# Household Grant

## Assumptions

The below documented are the various assumptions made while implementing, which directly affects the behaviour and error handling of the system.

- Each HTTP request is fired sequentially
- Saving of family members during household creation is **NOT** allowed.
- Each family member has to be added individually to the household, including married couple with spouse. Having `Spouse` field filled does **NOT** naturally mean they live in the same household
- When adding family member to household,
  - every member without a non-zero value in `ID` field (primary key) will be treated as a new member regardless.
  - if the member with a non-zero value in `ID` already exists in the household, the current member data will **NOT** overwrite the existing member.
  - this will not be used as a way to update existing member's details
- Format of `DOB` field in family memeber follows RFC3339 standard.
- `Spouse` field in family member stores the **Name** of the spouse
- All comparison are **case-sensitive**.
- Search without any valid key and value will return empty result.
- Search with duplicate key, only the first value will be taken.
- For Search, conflicting comparison is **NOT** allowed. E.g HouseholdIncome and HouseholdIncomeLT or HouseholdIncomeGT
- Int value used does not exceed int32 MAX_INT value
- Zero value of type **NOT** allowed for search, except for `bool` type fields
- `DOB` given will **NOT** be later than current time
- For Search, `has_children_by_age` takes in an age value which retrieves all members below the given age
- `Name` field in family member should be unique.
- For Search, specify `whole_household`=`true` to get only qualifying household response, family members list will be hidden.

## Accepted values for Household and FamilyMember

The accepted values are **case-sensitive**

### Household

- Type: `Landed`, `Condominium`, `HDB`

### FamilyMember

- Gender: `M`, `F`
- OccupationType: `Unemployed`, `Student`, `Employed`
- MaritalStatus: `Single`, `Married`

## Local Setup

- In an environment with golang v1.13 installed, execute `go run cmd/api/main.go` from the root folder
- Access the server at `http://localhost:8080`

### Routes

#### Create Household

- `POST` to `/households`

  Sample Request

  ```json
  {
    "type": "HDB"
  }
  ```

#### Add FamilyMember to Household

- `POST` to `/households/{id}/familymember`

  Sample Request

  ```json
  {
    "name": "WOMAN",
    "gender": "M",
    "spouse": "MAN",
    "occupation_type": "Employed",
    "marital_status": "Married",
    "dob": "1990-06-09T09:05:18+08:00"
  }
  ```

#### List all Households

- `GET` to `/households`

#### Show Household

- `GET` to `/households/{id}`

#### Search

- `GET` to `/search`

  List of available search key(s)

  - household_income_gt (int)
  - household_income_lt (int)
  - household_income (int)
  - household_size_gt (int)
  - household_size_lt (int)
  - household_size (int)
  - has_couple (bool)
  - has_children_by_age (int)
  - annual_income_gt (int)
  - annual_income_lt (int)
  - annual_income (int)
  - age_gt (int)
  - age_lt (int)
  - age (int)
  - whole_household (bool)
  - ...

  Refer to `model/search_operation.go` for a list of all available search key(s)

  Sample Request

  ```http
  /search?annual_income_lt=100000&whole_household=true

  /search?has_couple=true&has_children_by_age=18

  /search?age_gt=50
  ```

#### Delete Household

- `DELETE` to `/households/{id}`

#### Delete Family Member from Household

- `DELETE` to `/households/{id}/familymember/{fid}`

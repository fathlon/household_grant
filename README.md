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

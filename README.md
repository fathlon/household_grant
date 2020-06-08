# Household Grant

## Assumptions

The below documented are the various assumptions made while implementing, which directly affects the behaviour and error handling of the system.

- Each HTTP request is fired sequentially
- Each family member has to be added individually to the household, including married couple with spouse. Having `Spouse` field filled does **NOT** naturally mean they live in the same household
- When adding family member to household, every member without a non-zero value in `ID` field (primary key) will be treated as a new member regardless.

## Accepted values for Household and FamilyMember

The accepted values are **case-sensitive**

### Household

- Type: `Landed`, `Condominium`, `HDB`

### FamilyMember

- Gender: `M`, `F`
- OccupationType: `Unemployed`, `Student`, `Employed`
- MaritalStatus: `Single`, `Married`

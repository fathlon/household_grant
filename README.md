# Household Grant

## Assumptions

The below documented are the various assumptions made while implementing, which directly affects the behaviour and error handling of the system.

- Each HTTP request is fired sequentially

## Accepted values for Household and FamilyMember

The accepted values are **case-sensitive**

### Household

- Type: `Landed`, `Condominium`, `HDB`

### FamilyMember

- Gender: `M`, `F`
- OccupationType: `Unemployed`, `Student`, `Employed`
- MaritalStatus: `Single`, `Married`

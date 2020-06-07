package household

import (
	"strconv"
	"testing"

	"github.com/fathlon/household_grant/db"
	"github.com/fathlon/household_grant/model"
	"github.com/fathlon/household_grant/service"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		givenDatastore *db.Datastore
		givenHousehold model.Household
		expectedError  error
	}{
		{
			givenDatastore: db.NewDatastore(),
			givenHousehold: model.Household{},
			expectedError:  service.NewValidationError(model.ErrHouseholdTypeInvalid),
		},
		{
			givenDatastore: db.NewDatastore(),
			givenHousehold: model.Household{Type: "Landed"},
			expectedError:  nil,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			// When:
			result, err := Create(tc.givenDatastore, tc.givenHousehold)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.NotZero(t, result.ID)
				require.Equal(t, tc.givenHousehold.Type, result.Type)
			}
		})
	}
}

package person

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePerson(t *testing.T) {
	testCases := map[string]struct {
		Name            string
		Surname         string
		Age             int
		expectedName    string
		expectedSurname string
		expectedAge     int
	}{
		"person 1": {
			Name:            "Asd",
			Surname:         "Xyz",
			Age:             20,
			expectedName:    "Asd",
			expectedSurname: "Xyz",
			expectedAge:     20,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			Person := CreatePerson(tc.Name, tc.Surname, tc.Age)
			require.Equal(t, Person.Name, tc.expectedName, "unexpected name")
			require.Equal(t, Person.Surname, tc.expectedSurname, "unexpected surname")
			require.Equal(t, Person.Age, tc.expectedAge, "unexpected age")
		})
	}

}

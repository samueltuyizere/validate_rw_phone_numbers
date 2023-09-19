package validate_rw_phone_numbers_test

import (
	"testing"

	"github.com/samueltuyizere/validate_rw_phone_numbers"
)

type TestCase struct {
	number          string
	expectedResults bool
}

func TestMTN(t *testing.T) {
	t.Parallel()
	mtnTestCases := []TestCase{
		{number: "0788000000", expectedResults: true},
		{number: "0798000000", expectedResults: true},
		{number: "+250788000000", expectedResults: true},
		{number: "+250798000000", expectedResults: true},
		{number: "250788000000", expectedResults: true},
		{number: "250798000000", expectedResults: true},
		{number: "0738000000", expectedResults: false},
		{number: "0728000000", expectedResults: false},
		{number: "+250738000000", expectedResults: false},
		{number: "+250728000000", expectedResults: false},
		{number: "250738000000", expectedResults: false},
		{number: "250728000000", expectedResults: false},
		{number: "0778000000", expectedResults: false},
		{number: "+250778000000", expectedResults: false},
		{number: "250778000000", expectedResults: false},
	}

	for _, c := range mtnTestCases {
		actualResults := validate_rw_phone_numbers.ValidateMtn(c.number)
		if actualResults != c.expectedResults {
			t.Errorf("for phone number %v, we expected %v and gotten %v", c.number, c.expectedResults, actualResults)
		}
	}
}

func TestAirtelTigo(t *testing.T) {
	t.Parallel()
	airtelTigoTestCases := []TestCase{
		{number: "0788000000", expectedResults: false},
		{number: "0798000000", expectedResults: false},
		{number: "+250788000000", expectedResults: false},
		{number: "+250798000000", expectedResults: false},
		{number: "250788000000", expectedResults: false},
		{number: "250798000000", expectedResults: false},
		{number: "0738000000", expectedResults: true},
		{number: "0728000000", expectedResults: true},
		{number: "+250738000000", expectedResults: true},
		{number: "+250728000000", expectedResults: true},
		{number: "250738000000", expectedResults: true},
		{number: "250728000000", expectedResults: true},
		{number: "0778000000", expectedResults: false},
		{number: "+250778000000", expectedResults: false},
		{number: "250778000000", expectedResults: false},
	}

	for _, c := range airtelTigoTestCases {
		actualResults := validate_rw_phone_numbers.ValidateAirtelTigo(c.number)
		if actualResults != c.expectedResults {
			t.Errorf("for phone number %v, we expected %v and gotten %v", c.number, c.expectedResults, actualResults)
		}
	}
}
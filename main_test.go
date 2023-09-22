package validate_rw_phone_numbers_test

import (
	"testing"

	"github.com/samueltuyizere/validate_rw_phone_numbers"
)

type TestCase struct {
	number          string
	expectedResults bool
}

type InfoTestCase struct {
	number          string
	expectedResults validate_rw_phone_numbers.PhoneNbrInfo
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

func TestGettingPhoneNbrDetails(t *testing.T) {
	t.Parallel()
	airtelResults := &validate_rw_phone_numbers.PhoneNbrInfo{
		true,
		"Airtel/Tigo",
		"Airtel Money",
	}
	mtnResults := &validate_rw_phone_numbers.PhoneNbrInfo{
		true,
		"MTN",
		"MTN Mobile Money",
	}
	invalidResults := &validate_rw_phone_numbers.PhoneNbrInfo{
		false,
		"unkown",
		"unkown",
	}
	testCases := []InfoTestCase{
		{
			number:          "0788000000",
			expectedResults: *mtnResults,
		},
		{
			number:          "0798000000",
			expectedResults: *mtnResults,
		},
		{
			number:          "+250788000000",
			expectedResults: *mtnResults,
		},
		{
			number:          "+250798000000",
			expectedResults: *mtnResults,
		},
		{
			number:          "250788000000",
			expectedResults: *mtnResults,
		},
		{
			number:          "250798000000",
			expectedResults: *mtnResults,
		},
		{
			number:          "0738000000",
			expectedResults: *airtelResults,
		},
		{
			number:          "0728000000",
			expectedResults: *airtelResults,
		},
		{
			number:          "+250738000000",
			expectedResults: *airtelResults,
		},
		{
			number:          "+250728000000",
			expectedResults: *airtelResults,
		},
		{
			number:          "250738000000",
			expectedResults: *airtelResults,
		},
		{
			number:          "250728000000",
			expectedResults: *airtelResults,
		},
		{
			number:          "0778000000",
			expectedResults: *invalidResults,
		},
		{
			number:          "+250778000000",
			expectedResults: *invalidResults,
		},
		{
			number:          "250778000000",
			expectedResults: *invalidResults},
	}
	for _, c := range testCases {
		results := validate_rw_phone_numbers.GetPhoneNumberInfo(c.number)
		if results != c.expectedResults {
			t.Errorf("for phone number %v, we expected %v, but gotten %v", c.number, c.expectedResults, results)
		}
	}
}

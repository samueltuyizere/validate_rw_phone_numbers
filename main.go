package validate_rw_phone_numbers

import "strings"

func ValidateMtn(phoneNumber string) bool {
	if strings.HasPrefix(phoneNumber, "078") || strings.HasPrefix(phoneNumber, "079") {
		if len(phoneNumber) == 10 {
			return true
		}
		return false
	}
	if strings.HasPrefix(phoneNumber, "+25078") || strings.HasPrefix(phoneNumber, "+25079") {
		if len(phoneNumber) == 13 {
			return true
		}
		return false
	}
	if strings.HasPrefix(phoneNumber, "25078") || strings.HasPrefix(phoneNumber, "25079") {
		if len(phoneNumber) == 12 {
			return true
		}
		return false
	}

	return false
}

func ValidateAirtelTigo(phoneNumber string) bool {
	if strings.HasPrefix(phoneNumber, "072") || strings.HasPrefix(phoneNumber, "073") {
		if len(phoneNumber) == 10 {
			return true
		}
		return false
	}
	if strings.HasPrefix(phoneNumber, "+25072") || strings.HasPrefix(phoneNumber, "+25073") {
		if len(phoneNumber) == 13 {
			return true
		}
		return false
	}
	if strings.HasPrefix(phoneNumber, "25072") || strings.HasPrefix(phoneNumber, "25073") {
		if len(phoneNumber) == 12 {
			return true
		}
		return false
	}

	return false
}

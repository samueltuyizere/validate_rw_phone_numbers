package validate_rw_phone_numbers

import "strings"

type PhoneNbrInfo struct {
	IsValid    bool
	Telco      string
	WalletName string
}

func GetPhoneNumberInfo(phoneNumber string) PhoneNbrInfo {
	if ValidateMtn(phoneNumber) {
		return PhoneNbrInfo{
			true,
			"MTN",
			"MTN Mobile Money",
		}
	}
	if ValidateAirtelTigo(phoneNumber) {
		return PhoneNbrInfo{
			true,
			"Airtel/Tigo",
			"Airtel Money",
		}
	}
	return PhoneNbrInfo{
		false,
		"unkown",
		"unkown",
	}
}

func NormalizePhoneNumber(phoneNumber string) string {
	phoneNumber = strings.Replace(phoneNumber, " ", "", -1)
	phoneNumber = strings.Replace(phoneNumber, "-", "", -1)
	phoneNumber = strings.Replace(phoneNumber, "+", "", -1)
	phoneNumber = strings.Replace(phoneNumber, "(", "", -1)
	phoneNumber = strings.Replace(phoneNumber, ")", "", -1)
	phoneNumber = strings.Replace(phoneNumber, ".", "", -1)
	phoneNumber = strings.Replace(phoneNumber, ",", "", -1)
	return phoneNumber
}

func ValidateMtn(phoneNumber string) bool {
	phoneNumber = NormalizePhoneNumber(phoneNumber)
	if strings.HasPrefix(phoneNumber, "078") ||
		strings.HasPrefix(phoneNumber, "079") {
		return len(phoneNumber) == 10
	}
	if strings.HasPrefix(phoneNumber, "25078") ||
		strings.HasPrefix(phoneNumber, "25079") {
		return len(phoneNumber) == 12
	}
	return false
}

func ValidateAirtelTigo(phoneNumber string) bool {
	phoneNumber = NormalizePhoneNumber(phoneNumber)
	if strings.HasPrefix(phoneNumber, "072") ||
		strings.HasPrefix(phoneNumber, "073") {
		return len(phoneNumber) == 10
	}
	if strings.HasPrefix(phoneNumber, "25072") ||
		strings.HasPrefix(phoneNumber, "25073") {
		return len(phoneNumber) == 12
	}
	return false
}

func GetLocalFormat(phoneNumber string) string {
	phoneNumber = NormalizePhoneNumber(phoneNumber)
	if ValidateMtn(phoneNumber) {
		if len(phoneNumber) == 10 {
			return phoneNumber
		}
		return phoneNumber[2:]
	}
	if ValidateAirtelTigo(phoneNumber) {
		if len(phoneNumber) == 10 {
			return phoneNumber
		}
		return phoneNumber[2:]
	}
	return "unkown"
}

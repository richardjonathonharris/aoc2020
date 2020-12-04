package passporter

import (
	"fmt"
	"testing"
)

func TestCanValidateKeysOneLine(t *testing.T) {
	testString := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"
	resp := ValidatePassport(testString, false)
	if resp != true {
		t.Errorf("ValidatePassport should return true if all keys are found (one line)")
	}
}

func TestCanValidateKeysOneLineNoCID(t *testing.T) {
	testString := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 hgt:183cm"
	resp := ValidatePassport(testString, false)
	if resp != true {
		t.Errorf("ValidatePassport should return true if all keys but cid are found (one line)")
	}
}

func TestCanValidateMissingKeysOneLine(t *testing.T) {
	testString := "ecl:gry"
	resp := ValidatePassport(testString, false)
	if resp != false {
		t.Errorf("Validate passport should return false if some keys are missing (one line).")
	}
}

func TestCanValidateKeysMultipleLines(t *testing.T) {
	testString := `hcl:#ae17e1 iyr:2013
		eyr:2024
		ecl:brn pid:760753108 byr:1931
		hgt:179cm
		cid:abc123`
	resp := ValidatePassport(testString, false)
	if resp != true {
		t.Errorf("ValidatePassport should return true if all keys are found (multiple lines)")
	}
}

func TestCanValidateKeysMultipleLinesNoCID(t *testing.T) {
	testString := `hcl:#ae17e1 iyr:2013
		eyr:2024
		ecl:brn pid:760753108 byr:1931
		hgt:179cm`
	resp := ValidatePassport(testString, false)
	if resp != true {
		t.Errorf("ValidatePassport should return true if all keys are found but CID (multiple lines)")
	}
}

func TestCanValidateMissingKeysMultipleLines(t *testing.T) {
	testString := `hcl:#ae17e1 iyr:2013
		eyr:2024
		hgt:179cm`
	resp := ValidatePassport(testString, false)
	if resp != false {
		t.Errorf("ValidatePassport should return false if there are missing keys (multiple lines)")
	}
}

func testValidators(expected bool, actual bool) func(*testing.T) {
	return func(t *testing.T) {
		if actual != expected {
			t.Errorf(fmt.Sprintf("Expected %t but got %t", expected, actual))
		}
	}
}

func TestAllValidations(t *testing.T) {
	t.Run("ValidateYear:not int", testValidators(false, ValidateYear("abc", 1920, 2002)))
	t.Run("ValidateYear:too early year", testValidators(false, ValidateYear("1919", 1920, 2002)))
	t.Run("ValidateYear:too late year", testValidators(false, ValidateYear("2003", 1920, 2002)))
	t.Run("ValidateYear:acceptable year", testValidators(true, ValidateYear("1999", 1920, 2002)))

	t.Run("ValidateHeight:no cm or inches", testValidators(false, ValidateHeight("abc123")))
	t.Run("ValidateHeight:cm too small", testValidators(false, ValidateHeight("149cm")))
	t.Run("ValidateHeight:cm too big", testValidators(false, ValidateHeight("194cm")))
	t.Run("ValidateHeight:in too small", testValidators(false, ValidateHeight("58in")))
	t.Run("ValidateHeight:in too big", testValidators(false, ValidateHeight("77in")))
	t.Run("ValidateHeight:cm acceptable", testValidators(true, ValidateHeight("155cm")))
	t.Run("ValidateHeight:in acceptable", testValidators(true, ValidateHeight("65in")))

	t.Run("ValidateHairColor:no #", testValidators(false, ValidateHairColor("abc123")))
	t.Run("ValidateHairColor:not 6 characters after #", testValidators(false, ValidateHairColor("#1")))
	t.Run("ValidateHairColor:string contains non 0-9 or a-f characters", testValidators(false, ValidateHairColor("#zxy123")))
	t.Run("ValidateHairColor:acceptable", testValidators(true, ValidateHairColor("#012abc")))

	t.Run("ValidateEyeColor:is not acceptable", testValidators(false, ValidateEyeColor("abc123")))
	t.Run("ValidateEyeColor:is acceptable", testValidators(true, ValidateEyeColor("gry")))

	t.Run("ValidatePassportId:not a number", testValidators(false, ValidatePassportId("abc123")))
	t.Run("ValidatePassportId:not a nine-digit number", testValidators(false, ValidatePassportId("1")))
	t.Run("ValidatePassportId:is nine-digit number", testValidators(true, ValidatePassportId("000000000")))
}

func TestFullyInvalidPassport(t *testing.T) {
	passport := &Passport{
		byr: "abc123",
		iyr: "abc123",
		eyr: "abc123",
		hgt: "abc123",
		hcl: "abc123",
		ecl: "abc123",
		pid: "abc123",
	}
	if passport.AllFieldsValid() != false {
		t.Errorf("AllFieldsValid should be false when at least one field is not valid")
	}
}

func TestFullyValidPassport(t *testing.T) {
	passport := &Passport{
		byr: "2002",
		iyr: "2011",
		eyr: "2021",
		hgt: "60in",
		hcl: "#123abc",
		ecl: "brn",
		pid: "000000001",
	}
	if passport.AllFieldsValid() != true {
		t.Errorf("AllFieldsValid should be true when all useful fields are valid")
	}
}

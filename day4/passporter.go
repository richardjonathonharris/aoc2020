package passporter

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func matchToKey(passport *Passport, keyVal string) {
	// This is gross, but oh well!
	key := keyVal[0:3]
	value := keyVal[4:]
	if key == "byr" {
		passport.byr = value
	} else if key == "iyr" {
		passport.iyr = value
	} else if key == "eyr" {
		passport.eyr = value
	} else if key == "hgt" {
		passport.hgt = value
	} else if key == "hcl" {
		passport.hcl = value
	} else if key == "ecl" {
		passport.ecl = value
	} else if key == "pid" {
		passport.pid = value
	} else if key == "cid" {
		passport.cid = value
	}
}

func validateKeys(passport *Passport) bool {
	return passport.byr != "" && passport.iyr != "" && passport.eyr != "" && passport.hgt != "" && passport.hcl != "" && passport.ecl != "" && passport.pid != ""
}

func ValidateYear(propYear string, minYear int, maxYear int) bool {
	year, err := strconv.Atoi(propYear)
	if err != nil {
		return false
	}
	if year < minYear || year > maxYear {
		return false
	}
	return true
}

func ValidateHeight(height string) bool {
	if len(height) == 0 {
		return false
	}
	units := height[len(height)-2:]
	if units != "cm" && units != "in" {
		return false
	}
	value, err := strconv.Atoi(height[0 : len(height)-2])
	if err != nil {
		return false
	}
	if units == "cm" && (value < 150 || value > 193) {
		return false
	}
	if units == "in" && (value < 59 || value > 76) {
		return false
	}
	return true
}

func ValidateHairColor(hc string) bool {
	if len(hc) == 0 {
		return false
	}
	if string(hc[0]) != "#" {
		return false
	}
	hairColor := hc[1:]
	if len(hairColor) != 6 {
		return false
	}
	goodChars := regexp.MustCompile(`[g-z]`)
	return len(goodChars.ReplaceAllString(strings.ToLower(hairColor), "")) == 6
}

func ValidateEyeColor(ec string) bool {
	acceptableEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, match := range acceptableEyeColors {
		if ec == match {
			return true
		}
	}
	return false
}

func ValidatePassportId(pid string) bool {
	_, err := strconv.Atoi(pid)
	if err != nil {
		return false
	}
	return len(pid) == 9
}

func (p *Passport) AllFieldsValid() bool {
	return ValidateYear(p.byr, 1920, 2002) && ValidateYear(p.iyr, 2010, 2020) && ValidateYear(p.eyr, 2020, 2030) && ValidateHeight(p.hgt) && ValidateHairColor(p.hcl) && ValidateEyeColor(p.ecl) && ValidatePassportId(p.pid)
}

func ValidatePassport(passString string, validateFields bool) bool {
	passport := &Passport{}
	spaces := regexp.MustCompile(`\s+`)
	fixedString := spaces.ReplaceAllString(strings.ReplaceAll(passString, "\n", " "), " ")
	for _, keyVal := range strings.Split(fixedString, " ") {
		matchToKey(passport, keyVal)
	}
	if !validateFields {
		return validateKeys(passport)
	} else {
		return passport.AllFieldsValid()
	}
}

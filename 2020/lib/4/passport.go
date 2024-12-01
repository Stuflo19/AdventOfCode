package day4

import (
	"regexp"
	"strconv"
	"strings"
)

//byr (Birth Year)
//iyr (Issue Year)
//eyr (Expiration Year)
//hgt (Height)
//hcl (Hair Color)
//ecl (Eye Color)
//pid (Passport ID)
//cid (Country ID)

const (
	BIRTH_YEAR      = "byr"
	ISSUE_YEAR      = "iyr"
	EXPIRATION_YEAR = "eyr"
	HEIGHT          = "hgt"
	HAIR_COLOR      = "hcl"
	EYE_COLOR       = "ecl"
	PASSPORT_ID     = "pid"
	COUNTRY_ID      = "cid"
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

func NewPassport(data string) Passport {
	byr := ""
	iyr := ""
	eyr := ""
	hgt := ""
	hcl := ""
	ecl := ""
	pid := ""
	cid := ""

	keyValueSplit := strings.Split(data, " ")

	for _, value := range keyValueSplit {
		keyValue := strings.Split(value, ":")
		key := keyValue[0]
		value := keyValue[len(keyValue)-1]

		if key == BIRTH_YEAR {
			byr = value
		} else if key == ISSUE_YEAR {
			iyr = value
		} else if key == EXPIRATION_YEAR {
			eyr = value
		} else if key == HEIGHT {
			hgt = value
		} else if key == HAIR_COLOR {
			hcl = value
		} else if key == EYE_COLOR {
			ecl = value
		} else if key == PASSPORT_ID {
			pid = value
		} else if key == COUNTRY_ID {
			cid = value
		}
	}

	return Passport{
		byr,
		iyr,
		eyr,
		hgt,
		hcl,
		ecl,
		pid,
		cid,
	}
}

func (p Passport) validatePassport() bool {
	if p.byr == "" ||
		p.iyr == "" ||
		p.eyr == "" ||
		p.hgt == "" ||
		p.hcl == "" ||
		p.ecl == "" ||
		p.pid == "" {
		return false
	}

	return true
}

func (p Passport) validatePassport2() bool {
	if !p.validBirthYear() ||
		!p.validIssueYear() ||
		!p.validExpirationYear() ||
		!p.validHeight() ||
		!p.validHairColor() ||
		!p.validEyeColor() ||
		!p.validPassportID() {
		return false
	}

	return true
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func (p Passport) validBirthYear() bool {
	if len(p.byr) < 4 || len(p.byr) > 4 {
		return false
	}

	intYear, err := strconv.Atoi(p.byr)

	if err != nil {
		return false
	}

	if intYear < 1920 || intYear > 2002 {
		return false
	}

	return true
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func (p Passport) validIssueYear() bool {
	if len(p.iyr) < 4 || len(p.iyr) > 4 {
		return false
	}

	intYear, err := strconv.Atoi(p.iyr)

	if err != nil {
		return false
	}

	if intYear < 2010 || intYear > 2020 {
		return false
	}

	return true
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func (p Passport) validExpirationYear() bool {
	if len(p.eyr) < 4 || len(p.eyr) > 4 {
		return false
	}

	intYear, err := strconv.Atoi(p.eyr)

	if err != nil {
		return false
	}

	if intYear < 2020 || intYear > 2030 {
		return false
	}

	return true
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func (p Passport) validHeight() bool {
	trimmedCm, inCM := strings.CutSuffix(p.hgt, "cm")
	trimmedIn, inIn := strings.CutSuffix(p.hgt, "in")

	if inCM {
		integerHeight, err := strconv.Atoi(trimmedCm)

		if err != nil {
			return false
		}

		if integerHeight < 150 || integerHeight > 193 {
			return false
		}
	} else if inIn {
		integerHeight, err := strconv.Atoi(trimmedIn)

		if err != nil {
			return false
		}

		if integerHeight < 59 || integerHeight > 76 {
			return false
		}
	} else {
		return false
	}

	return true
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func (p Passport) validHairColor() bool {
	r, err := regexp.Compile(`^#(([0-9a-f]{2}){3}|([0-9a-f]){3})$`)

	if err != nil {
		panic("Invalid Regex")
	}

	result := r.MatchString(p.hcl)

	return result
}

const (
	AMBER = "amb"
	BLUE  = "blu"
	BROWN = "brn"
	GREY  = "gry"
	GREEN = "grn"
	HAZEL = "hzl"
	OTHER = "oth"
)

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func (p Passport) validEyeColor() bool {
	if p.ecl == AMBER ||
		p.ecl == BLUE ||
		p.ecl == BROWN ||
		p.ecl == GREY ||
		p.ecl == GREEN ||
		p.ecl == HAZEL ||
		p.ecl == OTHER {
		return true
	}
	return false
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func (p Passport) validPassportID() bool {
	if len(p.pid) != 9 {
		return false
	}

	_, err := strconv.Atoi(p.pid)

	if err != nil {
		return false
	}

	return true
}

//cid (Country ID) - ignored, missing or not.

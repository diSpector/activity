package validators

import (
	"regexp"
)

const (
	MAX_DESC_LENGTH = 50
	MIN_DESC_LENGTH = 5
	MAX_PERSONS     = 10
)

func ValidateAdd(desc string, cnt int, args []string) (bool, error) {
	var err error

	err = validateEmptyArgs(args)
	if err != nil {
		return false, err
	}

	err = validateDescription(desc)
	if err != nil {
		return false, err
	}

	err = validateCnt(cnt)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ValidateSearch(desc string, args []string) (bool, error) {
	var err error

	err = validateEmptyArgs(args)
	if err != nil {
		return false, err
	}

	err = validateDescription(desc)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ValidateList(args []string) (bool, error) {
	err := validateEmptyArgs(args)
	if err != nil {
		return false, err
	}

	return true, nil
}

func validateEmptyArgs(args []string) error {
	if len(args) != 0 {
		return ErrArgsNotNil
	}

	return nil
}

func validateDescription(desc string) error {
	if desc == `` {
		return ErrDescEmpty
	}

	if len(desc) > MAX_DESC_LENGTH {
		return ErrDescTooBig
	}

	if len(desc) < MIN_DESC_LENGTH {
		return ErrDescTooSmall
	}

	r := regexp.MustCompile(`[a-zA-z]+`)
	if !r.MatchString(desc) {
		return ErrDescNoLetters
	}

	return nil
}

func validateCnt(cnt int) error {
	if cnt == 0 {
		return ErrCntZero
	}

	if cnt < 0 {
		return ErrCntNeg
	}

	if cnt > MAX_PERSONS {
		return ErrCntTooBig
	}

	return nil
}

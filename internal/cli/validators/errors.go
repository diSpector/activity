package validators

import "github.com/pkg/errors"

var (
	ErrArgsNotNil    = errors.New(`arguments count should be equals to zero`)
	ErrDescEmpty     = errors.New(`description length should not be empty`)
	ErrDescTooBig    = errors.Errorf("description length should be less than %d", MAX_DESC_LENGTH)
	ErrDescTooSmall  = errors.Errorf("description length should be more than %d", MIN_DESC_LENGTH)
	ErrDescNoLetters = errors.New("description should contain letters")
	ErrCntZero       = errors.New("persons count should be more than zero")
	ErrCntNeg        = errors.New("persons count should be positive")
	ErrCntTooBig     = errors.Errorf("persons count should be less than %d", MAX_PERSONS)
)

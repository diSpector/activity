package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// самый простой вариант
func TestValidateEmptyArgs(t *testing.T) {
	var args []string
	var err error

	args = []string{`some`, `arg`}

	err = validateEmptyArgs(args)
	if err == nil {
		t.Error(`expected err, got nil`)
	}

	args = []string{}
	err = validateEmptyArgs(args)
	if err != nil {
		t.Errorf(`expected err = nil, got = %s`, err)
	}
}

// табличный тест
func TestValidateEmptyArgs_Table(t *testing.T) {
	var testTable = []struct {
		Name  string
		Args  []string
		IsErr bool
	}{
		{
			Name:  `non-empty args`,
			Args:  []string{`some`, `arg`},
			IsErr: true,
		},
		{
			Name:  `empty args`,
			Args:  []string{},
			IsErr: false,
		},
		{
			Name:  `empty args`,
			Args:  nil,
			IsErr: false,
		},
	}

	for _, test := range testTable {
		t.Log(`test name:`, test.Name)
		err := validateEmptyArgs(test.Args)
		if (err == nil && test.IsErr) || (err != nil && !test.IsErr) {
			t.Errorf(`expected err %v, got %s`, test.IsErr, err)
		}
	}
}

func TestValidateEmptyArgs_Table_Run(t *testing.T) {
	var testTable = []struct {
		Name  string
		Args  []string
		IsErr bool
	}{
		{
			Name:  `non-empty args`,
			Args:  []string{`some`, `arg`},
			IsErr: true,
		},
		{
			Name:  `empty args`,
			Args:  []string{},
			IsErr: false,
		},
		{
			Name:  `empty args`,
			Args:  nil,
			IsErr: false,
		},
	}

	for _, test := range testTable {
		t.Run(test.Name, func(t *testing.T) {
			err := validateEmptyArgs(test.Args)
			if (err == nil && test.IsErr) || (err != nil && !test.IsErr) {
				t.Errorf(`expected err %v, got %s`, test.IsErr, err)
			}
		})
	}
}

func TestValidateEmptyArgs_Table_Run_Parallel(t *testing.T) {
	var testTable = []struct {
		Name  string
		Args  []string
		IsErr bool
	}{
		{
			Name:  `non-empty args`,
			Args:  []string{`some`, `arg`},
			IsErr: true,
		},
		{
			Name:  `empty args`,
			Args:  []string{},
			IsErr: false,
		},
		{
			Name:  `empty args`,
			Args:  nil,
			IsErr: false,
		},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			err := validateEmptyArgs(test.Args)
			if (err == nil && test.IsErr) || (err != nil && !test.IsErr) {
				t.Errorf(`expected err %v, got %s`, test.IsErr, err)
			}
		})
	}
}

func TestValidateEmptyArgs_Table_Run_Asserts(t *testing.T) {
	var testTable = []struct {
		Name  string
		Args  []string
		IsErr bool
	}{
		{
			Name:  `non-empty args`,
			Args:  []string{`some`, `arg`},
			IsErr: true,
		},
		{
			Name:  `empty args`,
			Args:  []string{},
			IsErr: false,
		},
		{
			Name:  `empty args`,
			Args:  nil,
			IsErr: false,
		},
	}

	for _, test := range testTable {
		t.Run(test.Name, func(t *testing.T) {
			err := validateEmptyArgs(test.Args)
			if test.IsErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateEmptyArgs_Table_Run_Require(t *testing.T) {
	var testTable = []struct {
		Name  string
		Args  []string
		IsErr bool
	}{
		{
			Name:  `non-empty args`,
			Args:  []string{`some`, `arg`},
			IsErr: true,
		},
		{
			Name:  `empty args`,
			Args:  []string{},
			IsErr: false,
		},
		{
			Name:  `empty args`,
			Args:  nil,
			IsErr: false,
		},
	}

	for _, test := range testTable {
		t.Run(test.Name, func(t *testing.T) {
			err := validateEmptyArgs(test.Args)
			if test.IsErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestValidateDescription(t *testing.T) {
	var testTable = []struct {
		Name        string
		Description string
		Err         error
	}{
		{
			Name:        `empty description`,
			Description: ``,
			Err:         ErrDescEmpty,
		},
		{
			Name:        `description is too long`,
			Description: `sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss`,
			Err:         ErrDescTooBig,
		},
		{
			Name:        `description is too short`,
			Description: `kkk`,
			Err:         ErrDescTooSmall,
		},
		{
			Name:        `description without letters`,
			Description: `11111111111111`,
			Err:         ErrDescNoLetters,
		},
		{
			Name:        `description is ok`,
			Description: `Some desc`,
			Err:         nil,
		},
	}

	for _, test := range testTable {
		t.Run(test.Name, func(t *testing.T) {
			err := validateDescription(test.Description)
			assert.Equal(t, test.Err, err)
		})
	}
}

func BenchmarkValidateDescription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validateDescription(`test`)
	}
}

func FuzzValidateDescription(f *testing.F) {
	// fuzzy
}
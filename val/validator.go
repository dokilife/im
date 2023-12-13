package val

import "fmt"

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}

	return nil
}

func ValidateMessage(value string) error {
	if err := ValidateString(value, 1, 1000); err != nil {
		return err
	}
	return nil
}

package helpers

import (
	"fmt"
)

func CheckConnectionArgs(first *int, after *string, last *int, before *string) error {
	if first == nil && last == nil {
		return fmt.Errorf("first or last must be used")
	}
	if first == nil && after != nil {
		return fmt.Errorf("after must be with first")
	}
	if (last != nil && before == nil) || (last == nil && before != nil) {
		return fmt.Errorf("last and before must be used together")
	}
	if first != nil && after != nil && last != nil && before != nil {
		return fmt.Errorf("incorrect arguments usage")
	}
	return nil
}

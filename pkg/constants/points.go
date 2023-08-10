//nolint:gomnd
package constants

import "fmt"

type Points uint64 // Баллы

func (p Points) Readable() string {
	lastTwoDigits := p % 100
	lastDigit := p % 10

	if lastTwoDigits > 10 && lastTwoDigits < 20 {
		return fmt.Sprintf("%d баллов", p)
	}

	if lastDigit > 1 && lastDigit < 5 {
		return fmt.Sprintf("%d балла", p)
	}

	if lastDigit == 1 {
		return fmt.Sprintf("%d балл", p)
	}

	return fmt.Sprintf("%d баллов", p)
}

package declension

import "fmt"

type (
	Words  [3]string // Именительный падеж | Родительный падеж | Множественное число
	Number interface {
		int | int8 | int16 | int32 | int64 |
			uint | uint8 | uint16 | uint32 | uint64
	}
)

func Numeral[N Number](n N, w Words) string {
	var (
		lastTwoDigits = n % 100
		lastDigit     = n % 10

		nominative = w[0] // Именительный падеж
		genitive   = w[1] // Родительный падеж
		plural     = w[2] // Множественное число
	)

	if lastTwoDigits > 10 && lastTwoDigits < 20 {
		return fmt.Sprintf("%d %s", n, plural)
	}

	if lastDigit > 1 && lastDigit < 5 {
		return fmt.Sprintf("%d %s", n, genitive)
	}

	if lastDigit == 1 {
		return fmt.Sprintf("%d %s", n, nominative)
	}

	return fmt.Sprintf("%d %s", n, plural)
}

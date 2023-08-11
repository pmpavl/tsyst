package constants

import "github.com/pmpavl/tsyst/pkg/declension"

type Points uint64 // Баллы

const PointsZero Points = 0

func (p Points) Readable() string {
	return declension.Numeral(uint64(p), [3]string{"балл", "балла", "баллов"})
}

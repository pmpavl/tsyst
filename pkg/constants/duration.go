package constants

import (
	"fmt"
	"time"

	"github.com/pmpavl/tsyst/pkg/declension"
)

type Duration time.Duration // Длительность

const DurationsZero Duration = 0

func (d Duration) Time() time.Duration   { return time.Duration(d) }
func (d Duration) End() time.Time        { return time.Now().Add(d.Time()) }
func (d Duration) RoundMinute() Duration { return Duration(d.Time().Round(time.Minute)) }
func (d Duration) RoundSecond() Duration { return Duration(d.Time().Round(time.Second)) }

func (d Duration) Readable() string {
	if d == DurationsZero {
		return ""
	}

	var (
		minutes = int64(d.Time().Minutes())
		seconds = int64(d.Time().Seconds()) - minutes*60

		minutesDeclension = declension.Numeral(minutes, [3]string{"минута", "минуты", "минут"})
		secondsDeclension = declension.Numeral(seconds, [3]string{"секунда", "секунды", "секунд"})
	)

	if seconds == 0 {
		return minutesDeclension
	}

	if minutes == 0 {
		return secondsDeclension
	}

	return fmt.Sprintf("%s %s", minutesDeclension, secondsDeclension)
}

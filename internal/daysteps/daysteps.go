package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dayStepsData := strings.Split(datastring, ",")
	if len(dayStepsData) != 2 {
		return fmt.Errorf("wrong datastring")
	}

	steps, err := strconv.Atoi(dayStepsData[0])
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(dayStepsData[1])
	if err != nil {
		return err
	}

	err = spentenergy.ValidateParseData(steps, duration)
	if err != nil {
		return err
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, spentCalories), nil
}

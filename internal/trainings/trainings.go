package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	RunnigActivity = "Бег"
	WalkingActiviy = "Ходьба"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	trainingData := strings.Split(datastring, ",")
	if len(trainingData) != 3 {
		return fmt.Errorf("wrong datastring")
	}

	steps, err := strconv.Atoi(trainingData[0])
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(trainingData[2])
	if err != nil {
		return err
	}

	err = spentenergy.ValidateParseData(steps, duration)
	if err != nil {
		return err
	}

	t.Steps = steps
	t.TrainingType = trainingData[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var spentCalories float64
	var err error
	switch t.TrainingType {
	case RunnigActivity:
		spentCalories, err = spentenergy.RunningSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case WalkingActiviy:
		spentCalories, err = spentenergy.WalkingSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType,
		t.Duration.Hours(), distance, speed, spentCalories), nil
}

package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var ErrLessThanZero = errors.New("value should be positive")

func ValidateParseData(steps int, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("%w: steps = %d", ErrLessThanZero, steps)
	}

	if duration <= 0 {
		return fmt.Errorf("%w: duration = %d", ErrLessThanZero, duration)
	}

	return nil
}

func validateSpentCaloriesData(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("%w: steps = %d", ErrLessThanZero, steps)
	}

	if weight <= 0 {
		return fmt.Errorf("%w: weight = %d", ErrLessThanZero, steps)
	}

	if height <= 0 {
		return fmt.Errorf("%w: height = %d", ErrLessThanZero, steps)
	}

	if duration <= 0 {
		return fmt.Errorf("%w: duration = %d", ErrLessThanZero, steps)
	}

	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	spentCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validateSpentCaloriesData(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	return (weight * meanSpeed * duration.Minutes()) / minInH, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return (stepLength * float64(steps)) / mInKm
}

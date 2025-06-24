package spentenergy

import (
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

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("function WalkingSpentCalories must receive a positive argument")
	}
	avSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	result := (weight * avSpeed * minutes / minInH) * walkingCaloriesCoefficient
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("function RunningSpentCalories must receive a positive argument")
	}
	avSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	result := weight * avSpeed * minutes / minInH
	return result, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0
	}
	hours := duration.Hours()
	return Distance(steps, height) / hours
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 3 {
		return fmt.Errorf("method Parse received corrupted string argument")
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("method Parse received number of steps <= 0")
	}
	t.Steps = steps

	t.TrainingType = splitData[1]

	duration, err := time.ParseDuration(splitData[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("method Parse received duration <= 0")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var calories float64
	var err error
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	avSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	hours := t.Duration.Hours()

	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight,
			t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight,
			t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	}

	result := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", t.TrainingType, hours, distance, avSpeed, calories)
	return result, nil
}

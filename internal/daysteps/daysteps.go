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
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 2 {
		return fmt.Errorf("method Parse received corrupted string argument")
	}

	steps, err := strconv.Atoi(splitData[0])
	if steps <= 0 {
		return fmt.Errorf("method Parse received number of steps <= 0")
	}
	if err != nil {
		return err
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(splitData[1])
	if duration <= 0 {
		return fmt.Errorf("method Parse received duration <= 0")
	}
	if err != nil {
		return err
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return result, nil
}

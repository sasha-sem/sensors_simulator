package sensors

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Sensor interface {
	GenerateValue()
	fmt.Stringer
}
type DefaultSensor struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	MinValue  float64   `json:"min_value"`
	MaxValue  float64   `json:"max_value"`
	Value     float64   `json:"value"`
	TimeStamp time.Time `json:"timestamp"`
}

func (s *DefaultSensor) GenerateValue() {
	var val float64
	val = rand.Float64()*(s.MaxValue-s.MinValue) + s.MinValue
	s.Value = val
	s.TimeStamp = time.Now()
}

func (s DefaultSensor) String() string {
	sMar, _ := json.Marshal(s)
	return string(sMar)
}

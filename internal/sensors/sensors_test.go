package sensors

import (
	"log"
	"testing"
	"time"
)

func TestSensor_GenerateValue(t *testing.T) {
	testSensor := Sensor{ID: 1, Type: "test", Name: "Тестовый датчик", MinValue: 25, MaxValue: 30, Value: 25}
	testStartTime := time.Now()
	testSensor.GenerateValue()
	testStopTime := time.Now()
	if testSensor.Value > testSensor.MaxValue || testSensor.Value < testSensor.MinValue {
		t.Errorf("TestSensor_GenerateValue().Value = %f; want between than %f and %f", testSensor.Value, testSensor.MinValue, testSensor.MaxValue)
	}
	if testSensor.TimeStamp.Unix() < testStartTime.Unix() || testSensor.TimeStamp.Unix() > testStopTime.Unix() {
		t.Errorf("TestSensor_GenerateValue().TimeStamp = %s; want between than %s and %s", testSensor.TimeStamp.String(), testStartTime.String(), testStopTime.String())
	}
}

func TestSensor_String(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalf("%s: %s", "Bad timezone", err)
	}
	testSensor := Sensor{ID: 1, Type: "test", Name: "Тестовый датчик", MinValue: 25, MaxValue: 30, Value: 25, TimeStamp: time.Date(2021, time.June, 23, 10, 30, 0, 0, loc)}
	testString := testSensor.String()
	exampleString := "{\"id\":1,\"type\":\"test\",\"name\":\"Тестовый датчик\",\"min_value\":25,\"max_value\":30,\"value\":25,\"timestamp\":\"2021-06-23T10:30:00+03:00\"}"
	if testString != exampleString {
		t.Errorf("String() = %s; want %s", testString, exampleString)
	}
}

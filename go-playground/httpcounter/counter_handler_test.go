package main

import (
	"net/http/httptest"
	"testing"
	"time"
)

func TestMakeCounterHandler(t *testing.T) {
	timestamps := make(chan time.Time)
	go func() { _ = <-timestamps }()
	counts := make(chan int)
	go func() { counts <- 100 }()

	rec := httptest.NewRecorder()

	handler := &CounterHandler{timestamps: timestamps, counts: counts}
	handler.ServeHTTP(rec, nil)

	response := rec.Body.String()
	expected := "100"
	if response != expected {
		t.Errorf("Wrong response %s, expected %s", response, expected)
	}
}

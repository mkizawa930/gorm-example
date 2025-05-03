package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	db, err := NewTestDB()
	if err != nil {
		panic(err)
	}

	if err := Migrate(db); err != nil {
		panic(err)
	}

	code := m.Run()

	if code > 0 {
		os.Exit(code)
	}

	os.Exit(0)
}

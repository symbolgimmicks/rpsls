package main

import (
	"github.com/DATA-DOG/godog"
)

func theApplicationStarted() error {
	return nil
}

func iRunTheTest() error {
	return nil
}

func theTestRuns() error {
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the application started$`, theApplicationStarted)
	s.Step(`^I run the test$`, iRunTheTest)
	s.Step(`^the test runs$`, theTestRuns)

	s.BeforeScenario(func(interface{}) {
		// clean the state before every scenario
	})
}

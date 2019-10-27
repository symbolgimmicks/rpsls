package randomnumber_test

import (
	"errors"
	"fmt"

	"github.com/DATA-DOG/godog"
	RNG "github.com/symbolgimmicks/rpsls/randomnumber"
)

type validateIsValidFeature struct {
	num               RNG.RandomNumber
	lastIsValidResult bool
	lastGetError      error
}

func (a *validateIsValidFeature) minIsDefinedAs(arg1 int) (err error) {
	err = nil
	if RNG.Min != arg1 {
		err = errors.New(fmt.Sprintf("Expected RNG.Min = %d.  Actual Result %d = %d", arg1, RNG.Min, arg1))
	}
	return
}

func (a *validateIsValidFeature) maxIsDefinedAs(value int) (err error) {
	err = nil
	if RNG.Max != value {
		err = errors.New(fmt.Sprintf("Expected RNG.Max = %d.  Actual Result %d = %d", value, RNG.Max, value))
	}
	return
}

func (a *validateIsValidFeature) initializeARandomNumberTo(value int) (err error) {
	err = nil
	a.num = RNG.RandomNumber{Value: value}
	return
}

func (a *validateIsValidFeature) initializeARandomNumberToURL(url string) (err error) {
	err = nil
	a.num = RNG.RandomNumber{Value: 0}
	err = a.num.GenerateFromService(url)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to generate number from [%s]: %v", url, err))
	}
	a.lastGetError = err
	return
}

func (a *validateIsValidFeature) initializeARandomNumberToDefaultURL() (err error) {
	err = a.initializeARandomNumberToURL(RNG.DefaultRNGServiceURL)
	return
}

func (a *validateIsValidFeature) isValidSucceeds() (err error) {
	err = nil
	if actual := a.num.IsValid(); actual != true {
		err = errors.New(fmt.Sprintf("isValid failed. Actual Value [%v], GET result [%v]", a.num.Value, a.lastGetError))
	}
	return
}

func (a *validateIsValidFeature) isValidFails() (err error) {
	err = nil
	if actual := a.num.IsValid(); actual != false {
		err = errors.New("isValid succeeded")
	}
	return
}

func (a *validateIsValidFeature) reset() {
	a.num = RNG.RandomNumber{Value: 0}
}

func FeatureContext(s *godog.Suite) {
	validateApi := &validateIsValidFeature{}

	s.BeforeScenario(func(interface{}) {
		validateApi.reset()
	})

	s.Step(`^Min is defined as (\d+)$`, validateApi.minIsDefinedAs)
	s.Step(`^Max is defined as (\d+)$`, validateApi.maxIsDefinedAs)
	s.Step(`^I initialize a RandomNumber to (\d+)$`, validateApi.initializeARandomNumberTo)
	s.Step(`^I send a GET request to initialize a RandomNumber using the Default RNG endpoint$`, validateApi.initializeARandomNumberToDefaultURL)
	s.Step(`^I send a GET request to initialize a RandomNumber using the endpoint ([^"]*)$`, validateApi.initializeARandomNumberToURL)
	s.Step(`^isValid succeeds$`, validateApi.isValidSucceeds)
	s.Step(`^isValid fails$`, validateApi.isValidFails)
}

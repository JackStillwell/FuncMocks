package models

type OrderedMocks struct{}

type MockDecision struct {
	DecisiveMock interface{}
	Paths        []interface{}
}

type MockLoop struct {
	Mocks        OrderedMocks
	EndCondition func() bool
}

package utils

type TestCase struct {
	Name       string
	SetupMocks func() (expectedValue interface{})
	WantErr    bool
}

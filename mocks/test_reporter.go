package mocks

type testReporter interface {
	Fatalf(format string, args ...any)
	Fatal(args ...any)
	Cleanup(func())
}
package mock

// Foo .
//go:generate mockgen -destination=mocks/mock_foo.go -package=mocks github.com/archerwq/go-lab/concept/test/mock Foo
type Foo interface {
	Bar(x int) int
}

// Check .
func Check(f Foo, x int) bool {
	return f.Bar(x) > 1024
}

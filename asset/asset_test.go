package asset

type testKind string

func (k testKind) Kind() string {
	return string(k)
}

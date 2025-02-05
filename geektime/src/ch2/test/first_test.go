package try_test

import "testing"

/*
Run in terminal:
go test -v first_test.go
*/
// *testing.T表示testing.T的指针
func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}

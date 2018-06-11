package asheets

import (
	"fmt"
	"testing"
)

func TestCol(t *testing.T) {
	a := ColNameToNumber("A")
	b := NumberToColName(a)
	fmt.Println("A", a, b)
	t.Log("A", a, b)
}

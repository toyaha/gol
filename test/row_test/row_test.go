package row_test

import (
	"fmt"
	"testing"

	"github.com/toyaha/gol"
)

func TestRow_NewRow(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		row := gol.NewRow()

		var str string
		err := row.Scan(&str)
		target := fmt.Sprintf("%v", err)
		check := "row not found"
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	})
}

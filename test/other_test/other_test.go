package null_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/toyaha/gol"
)

func TestChangeQueryForMssql(t *testing.T) {
	fn := func(t *testing.T, query string, check string) {
		target := gol.ChangeQueryForMssql(query)
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	}

	t.Run("base", func(t *testing.T) {
		var query = "abcABC123-^\\@[;:],./!\"#$%&'()=~|`{+*}<>?_???"
		var check = "abcABC123-^\\@[;:],./!\"#$%&'()=~|`{+*}<>@v1_@v2@v3@v4"
		fn(t, query, check)
	})
}

func TestChangeValueListForMssql(t *testing.T) {
	fn := func(t *testing.T, dataList []interface{}, checkList []interface{}) {
		valueList := gol.ChangeValueListForMssql(dataList...)
		target := fmt.Sprintf("%+v", valueList)
		check := fmt.Sprintf("%+v", checkList)
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	}

	t.Run("base", func(t *testing.T) {
		var dataList = []interface{}{1, sql.NamedArg{Name: "a", Value: "b"}, true}
		var checkList = []interface{}{
			sql.NamedArg{Name: "v1", Value: 1},
			sql.NamedArg{Name: "a", Value: "b"},
			sql.NamedArg{Name: "v2", Value: true},
		}
		fn(t, dataList, checkList)
	})
}

func TestChangeQueryForPostgresql(t *testing.T) {
	fn := func(t *testing.T, query string, check string) {
		target := gol.ChangeQueryForPostgresql(query)
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	}

	t.Run("base", func(t *testing.T) {
		var query = "abcABC123-^\\@[;:],./!\"#$%&'()=~|`{+*}<>?_???"
		var check = "abcABC123-^\\@[;:],./!\"#$%&'()=~|`{+*}<>$1_$2$3$4"
		fn(t, query, check)
	})
}

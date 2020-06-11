package null_test

import (
	"testing"
	"time"

	"github.com/toyaha/gol"
)

func TestNullBool(t *testing.T) {
	fn := func(t *testing.T, data *gol.NullBool, checkMap map[string]interface{}) {
		{
			target := data.Get()
			check, ok := checkMap["Get"]
			if !ok {
				t.Error("Get is not exist")
			}
			if check == true {
				if target == nil {
					t.Errorf("\ntarget: not exist")
				}
			} else {
				if target != nil {
					t.Errorf("\ntarget: exist")
				}
			}
		}

		{
			target := data.GetValue()
			check, ok := checkMap["GetValue"]
			if !ok {
				t.Error("GetValue is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetValueWithDefault(true)
			check, ok := checkMap["GetValueWithDefault_true"]
			if !ok {
				t.Error("GetValueWithDefault_true is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetValueWithDefault(false)
			check, ok := checkMap["GetValueWithDefault_false"]
			if !ok {
				t.Error("GetValueWithDefault_false is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetString()
			check, ok := checkMap["GetString"]
			if !ok {
				t.Error("GetString is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetStringWithDefault("default")
			check, ok := checkMap["GetStringWithDefault_default"]
			if !ok {
				t.Error("GetStringWithDefault_default is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	}

	t.Run("none", func(t *testing.T) {
		data := &gol.NullBool{}
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     false,
			"GetValueWithDefault_true":     true,
			"GetValueWithDefault_false":    false,
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})

	t.Run("true", func(t *testing.T) {
		data := &gol.NullBool{}
		data.Set(true)
		checkMap := map[string]interface{}{
			"Get":                          true,
			"GetValue":                     true,
			"GetValueWithDefault_true":     true,
			"GetValueWithDefault_false":    true,
			"GetString":                    "true",
			"GetStringWithDefault_default": "true",
		}
		fn(t, data, checkMap)
	})

	t.Run("delete", func(t *testing.T) {
		data := &gol.NullBool{}
		data.Set(true)
		data.Delete()
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     false,
			"GetValueWithDefault_true":     true,
			"GetValueWithDefault_false":    false,
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})
}

func TestNullFloat64(t *testing.T) {
	fn := func(t *testing.T, data *gol.NullFloat64, checkMap map[string]interface{}) {
		{
			target := data.Get()
			check, ok := checkMap["Get"]
			if !ok {
				t.Error("Get is not exist")
			}
			if check == true {
				if target == nil {
					t.Errorf("\ntarget: not exist")
				}
			} else {
				if target != nil {
					t.Errorf("\ntarget: exist")
				}
			}
		}

		{
			target := data.GetValue()
			check, ok := checkMap["GetValue"]
			if !ok {
				t.Error("GetValue is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetValueWithDefault(1.23)
			check, ok := checkMap["GetValueWithDefault_1.23"]
			if !ok {
				t.Error("GetValueWithDefault_1.23 is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetString()
			check, ok := checkMap["GetString"]
			if !ok {
				t.Error("GetString is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetStringWithDefault("default")
			check, ok := checkMap["GetStringWithDefault_default"]
			if !ok {
				t.Error("GetStringWithDefault_default is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	}

	t.Run("none", func(t *testing.T) {
		data := &gol.NullFloat64{}
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     float64(0),
			"GetValueWithDefault_1.23":     float64(1.23),
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})

	t.Run("3.14", func(t *testing.T) {
		data := &gol.NullFloat64{}
		data.Set(3.14)
		checkMap := map[string]interface{}{
			"Get":                          true,
			"GetValue":                     float64(3.14),
			"GetValueWithDefault_1.23":     float64(3.14),
			"GetString":                    "3.14",
			"GetStringWithDefault_default": "3.14",
		}
		fn(t, data, checkMap)
	})

	t.Run("delete", func(t *testing.T) {
		data := &gol.NullFloat64{}
		data.Set(3.14)
		data.Delete()
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     float64(0),
			"GetValueWithDefault_1.23":     float64(1.23),
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})
}

func TestNullInt32(t *testing.T) {
	fn := func(t *testing.T, data *gol.NullInt32, checkMap map[string]interface{}) {
		{
			target := data.Get()
			check, ok := checkMap["Get"]
			if !ok {
				t.Error("Get is not exist")
			}
			if check == true {
				if target == nil {
					t.Errorf("\ntarget: not exist")
				}
			} else {
				if target != nil {
					t.Errorf("\ntarget: exist")
				}
			}
		}

		{
			target := data.GetValue()
			check, ok := checkMap["GetValue"]
			if !ok {
				t.Error("GetValue is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetValueWithDefault(123)
			check, ok := checkMap["GetValueWithDefault_123"]
			if !ok {
				t.Error("GetValueWithDefault_123 is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetString()
			check, ok := checkMap["GetString"]
			if !ok {
				t.Error("GetString is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetStringWithDefault("default")
			check, ok := checkMap["GetStringWithDefault_default"]
			if !ok {
				t.Error("GetStringWithDefault_default is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	}

	t.Run("none", func(t *testing.T) {
		data := &gol.NullInt32{}
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     int(0),
			"GetValueWithDefault_123":      int(123),
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})

	t.Run("100", func(t *testing.T) {
		data := &gol.NullInt32{}
		data.Set(100)
		checkMap := map[string]interface{}{
			"Get":                          true,
			"GetValue":                     int(100),
			"GetValueWithDefault_123":      int(100),
			"GetString":                    "100",
			"GetStringWithDefault_default": "100",
		}
		fn(t, data, checkMap)
	})

	t.Run("delete", func(t *testing.T) {
		data := &gol.NullInt32{}
		data.Set(1000)
		data.Delete()
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     int(0),
			"GetValueWithDefault_123":      int(123),
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})
}

func TestNullInt64(t *testing.T) {
	fn := func(t *testing.T, data *gol.NullInt64, checkMap map[string]interface{}) {
		{
			target := data.Get()
			check, ok := checkMap["Get"]
			if !ok {
				t.Error("Get is not exist")
			}
			if check == true {
				if target == nil {
					t.Errorf("\ntarget: not exist")
				}
			} else {
				if target != nil {
					t.Errorf("\ntarget: exist")
				}
			}
		}

		{
			target := data.GetValue()
			check, ok := checkMap["GetValue"]
			if !ok {
				t.Error("GetValue is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetValueWithDefault(123)
			check, ok := checkMap["GetValueWithDefault_123"]
			if !ok {
				t.Error("GetValueWithDefault_123 is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetString()
			check, ok := checkMap["GetString"]
			if !ok {
				t.Error("GetString is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetStringWithDefault("default")
			check, ok := checkMap["GetStringWithDefault_default"]
			if !ok {
				t.Error("GetStringWithDefault_default is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	}

	t.Run("none", func(t *testing.T) {
		data := &gol.NullInt64{}
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     int(0),
			"GetValueWithDefault_123":      int(123),
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})

	t.Run("100", func(t *testing.T) {
		data := &gol.NullInt64{}
		data.Set(100)
		checkMap := map[string]interface{}{
			"Get":                          true,
			"GetValue":                     int(100),
			"GetValueWithDefault_123":      int(100),
			"GetString":                    "100",
			"GetStringWithDefault_default": "100",
		}
		fn(t, data, checkMap)
	})

	t.Run("delete", func(t *testing.T) {
		data := &gol.NullInt64{}
		data.Set(1000)
		data.Delete()
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     int(0),
			"GetValueWithDefault_123":      int(123),
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})
}

func TestNullString(t *testing.T) {
	fn := func(t *testing.T, data *gol.NullString, checkMap map[string]interface{}) {
		{
			target := data.Get()
			check, ok := checkMap["Get"]
			if !ok {
				t.Error("Get is not exist")
			}
			if check == true {
				if target == nil {
					t.Errorf("\ntarget: not exist")
				}
			} else {
				if target != nil {
					t.Errorf("\ntarget: exist")
				}
			}
		}

		{
			target := data.GetValue()
			check, ok := checkMap["GetValue"]
			if !ok {
				t.Error("GetValue is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetValueWithDefault("abc")
			check, ok := checkMap["GetValueWithDefault_abc"]
			if !ok {
				t.Error("GetValueWithDefault_abc is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetString()
			check, ok := checkMap["GetString"]
			if !ok {
				t.Error("GetString is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetStringWithDefault("default")
			check, ok := checkMap["GetStringWithDefault_default"]
			if !ok {
				t.Error("GetStringWithDefault_default is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	}

	t.Run("none", func(t *testing.T) {
		data := &gol.NullString{}
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     "",
			"GetValueWithDefault_abc":      "abc",
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})

	t.Run("100", func(t *testing.T) {
		data := &gol.NullString{}
		data.Set("add")
		checkMap := map[string]interface{}{
			"Get":                          true,
			"GetValue":                     "add",
			"GetValueWithDefault_abc":      "add",
			"GetString":                    "add",
			"GetStringWithDefault_default": "add",
		}
		fn(t, data, checkMap)
	})

	t.Run("delete", func(t *testing.T) {
		data := &gol.NullString{}
		data.Set("add")
		data.Delete()
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     "",
			"GetValueWithDefault_abc":      "abc",
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})
}

func TestNullTime(t *testing.T) {
	timeDefault := time.Time{}
	timeFormat := "2006/01/02 15:04:05"
	time2000Str := "2000/01/01 00:00:00"
	time2000, err := time.Parse(timeFormat, time2000Str)
	if err != nil {
		t.Error(err)
		return
	}
	time1999Str := "1999/01/01 00:00:00"
	time1999, err := time.Parse(timeFormat, time1999Str)
	if err != nil {
		t.Error(err)
		return
	}

	fn := func(t *testing.T, data *gol.NullTime, checkMap map[string]interface{}) {
		{
			target := data.Get()
			check, ok := checkMap["Get"]
			if !ok {
				t.Error("Get is not exist")
			}
			if check == true {
				if target == nil {
					t.Errorf("\ntarget: not exist")
				}
			} else {
				if target != nil {
					t.Errorf("\ntarget: exist")
				}
			}
		}

		{
			target := data.GetValue()
			check, ok := checkMap["GetValue"]
			if !ok {
				t.Error("GetValue is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			time1999, err := time.Parse(timeFormat, time1999Str)
			if err != nil {
				t.Error(err)
				return
			}
			target := data.GetValueWithDefault(time1999)
			check, ok := checkMap["GetValueWithDefault_1999"]
			if !ok {
				t.Error("GetValueWithDefault_1999 is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetString(timeFormat)
			check, ok := checkMap["GetString"]
			if !ok {
				t.Error("GetString is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := data.GetStringWithDefault(timeFormat, "default")
			check, ok := checkMap["GetStringWithDefault_default"]
			if !ok {
				t.Error("GetStringWithDefault_default is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	}

	t.Run("none", func(t *testing.T) {
		data := &gol.NullTime{}
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     timeDefault,
			"GetValueWithDefault_1999":     time1999,
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})

	t.Run("add", func(t *testing.T) {
		data := &gol.NullTime{}
		ti, err := time.Parse(timeFormat, time2000Str)
		if err != nil {
			t.Error(err)
			return
		}
		data.Set(ti)
		checkMap := map[string]interface{}{
			"Get":                          true,
			"GetValue":                     time2000,
			"GetValueWithDefault_1999":     time2000,
			"GetString":                    time2000Str,
			"GetStringWithDefault_default": time2000Str,
		}
		fn(t, data, checkMap)
	})

	t.Run("delete", func(t *testing.T) {
		data := &gol.NullTime{}
		ti, err := time.Parse(timeFormat, time2000Str)
		if err != nil {
			t.Error(err)
			return
		}
		data.Set(ti)
		data.Delete()
		checkMap := map[string]interface{}{
			"Get":                          false,
			"GetValue":                     timeDefault,
			"GetValueWithDefault_1999":     time1999,
			"GetString":                    "",
			"GetStringWithDefault_default": "default",
		}
		fn(t, data, checkMap)
	})
}

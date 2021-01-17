package gol

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"time"
)

type NullInterface interface {
	IsNull() bool
}

var _ NullInterface = &NullBool{}

type NullBool struct {
	sql.NullBool
}

func (rec *NullBool) Get() *bool {
	if !rec.NullBool.Valid {
		return nil
	}

	val := rec.NullBool.Bool

	return &val
}

func (rec *NullBool) GetValue() bool {
	return rec.NullBool.Bool
}

func (rec *NullBool) GetValueWithDefault(value bool) bool {
	if rec.Valid {
		return rec.NullBool.Bool
	}

	return value
}

func (rec *NullBool) GetString() string {
	return rec.GetStringWithDefault("")
}

func (rec *NullBool) GetStringWithDefault(value string) string {
	if rec.Valid {
		return strconv.FormatBool(rec.Bool)
	}

	return value
}

func (rec *NullBool) IsNull() bool {
	return !rec.Valid
}

func (rec *NullBool) Set(value bool) {
	rec.NullBool.Bool = value
	rec.NullBool.Valid = true
}

func (rec *NullBool) Delete() {
	var val bool
	rec.NullBool.Bool = val
	rec.NullBool.Valid = false
}

func (rec NullBool) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Bool)
}

func (rec *NullBool) UnmarshalJSON(value []byte) error {
	var val *bool

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val != nil {
		rec.Bool = *val
		rec.Valid = true
	} else {
		rec.Valid = false
	}

	return nil
}

var _ NullInterface = &NullFloat64{}

type NullFloat64 struct {
	sql.NullFloat64
}

func (rec *NullFloat64) Get() *float64 {
	if !rec.NullFloat64.Valid {
		return nil
	}

	val := rec.NullFloat64.Float64

	return &val
}

func (rec *NullFloat64) GetValue() float64 {
	return rec.NullFloat64.Float64
}

func (rec *NullFloat64) GetValueWithDefault(value float64) float64 {
	if rec.Valid {
		return rec.NullFloat64.Float64
	}

	return value
}

func (rec *NullFloat64) GetString() string {
	return rec.GetStringWithDefault("")
}

func (rec *NullFloat64) GetStringWithDefault(value string) string {
	if rec.NullFloat64.Valid {
		return strconv.FormatFloat(rec.Float64, 'f', -1, 64)
	}

	return value
}

func (rec *NullFloat64) IsNull() bool {
	return !rec.Valid
}

func (rec *NullFloat64) Set(value float64) {
	rec.NullFloat64.Float64 = value
	rec.NullFloat64.Valid = true
}

func (rec *NullFloat64) Delete() {
	var val float64
	rec.NullFloat64.Float64 = val
	rec.NullFloat64.Valid = false
}

func (rec NullFloat64) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Float64)
}

func (rec *NullFloat64) UnmarshalJSON(value []byte) error {
	var val *float64

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val != nil {
		rec.Float64 = *val
		rec.Valid = true
	} else {
		rec.Valid = false
	}

	return nil
}

var _ NullInterface = &NullInt32{}

type NullInt32 struct {
	sql.NullInt32
}

func (rec *NullInt32) Get() *int {
	if !rec.NullInt32.Valid {
		return nil
	}

	val := int(rec.NullInt32.Int32)

	return &val
}

func (rec *NullInt32) GetValue() int {
	return int(rec.NullInt32.Int32)
}

func (rec *NullInt32) GetValueWithDefault(value int) int {
	if rec.Valid {
		return int(rec.NullInt32.Int32)
	}

	return value
}

func (rec *NullInt32) GetString() string {
	return rec.GetStringWithDefault("")
}

func (rec *NullInt32) GetStringWithDefault(value string) string {
	if rec.NullInt32.Valid {
		return strconv.Itoa(int(rec.Int32))
	}

	return value
}

func (rec *NullInt32) IsNull() bool {
	return !rec.Valid
}

func (rec *NullInt32) Set(value int) {
	rec.NullInt32.Int32 = int32(value)
	rec.NullInt32.Valid = true
}

func (rec *NullInt32) Delete() {
	var val int32
	rec.NullInt32.Int32 = val
	rec.NullInt32.Valid = false
}

func (rec NullInt32) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Int32)
}

func (rec *NullInt32) UnmarshalJSON(value []byte) error {
	var val *int32

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val != nil {
		rec.Int32 = *val
		rec.Valid = true
	} else {
		rec.Valid = false
	}

	return nil
}

var _ NullInterface = &NullInt64{}

type NullInt64 struct {
	sql.NullInt64
}

func (rec *NullInt64) Get() *int {
	if !rec.NullInt64.Valid {
		return nil
	}

	val := int(rec.NullInt64.Int64)

	return &val
}

func (rec *NullInt64) GetValue() int {
	return int(rec.NullInt64.Int64)
}

func (rec *NullInt64) GetValueWithDefault(value int) int {
	if rec.Valid {
		return int(rec.NullInt64.Int64)
	}

	return value
}

func (rec *NullInt64) GetString() string {
	return rec.GetStringWithDefault("")
}

func (rec *NullInt64) GetStringWithDefault(value string) string {
	if rec.NullInt64.Valid {
		return strconv.Itoa(int(rec.Int64))
	}

	return value
}

func (rec *NullInt64) IsNull() bool {
	return !rec.Valid
}

func (rec *NullInt64) Set(value int) {
	rec.NullInt64.Int64 = int64(value)
	rec.NullInt64.Valid = true
}

func (rec *NullInt64) Delete() {
	var val int64
	rec.NullInt64.Int64 = val
	rec.NullInt64.Valid = false
}

func (rec NullInt64) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Int64)
}

func (rec *NullInt64) UnmarshalJSON(value []byte) error {
	var val *int64

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val != nil {
		rec.Int64 = *val
		rec.Valid = true
	} else {
		rec.Valid = false
	}

	return nil
}

var _ NullInterface = &NullString{}

type NullString struct {
	sql.NullString
}

func (rec *NullString) Get() *string {
	if !rec.NullString.Valid {
		return nil
	}

	val := rec.NullString.String

	return &val
}

func (rec *NullString) GetValue() string {
	return rec.NullString.String
}

func (rec *NullString) GetValueWithDefault(value string) string {
	if rec.Valid {
		return rec.NullString.String
	}

	return value
}

func (rec *NullString) GetString() string {
	return rec.GetStringWithDefault("")
}

func (rec *NullString) GetStringWithDefault(value string) string {
	if rec.NullString.Valid {
		return rec.NullString.String
	}

	return value
}

func (rec *NullString) IsNull() bool {
	return !rec.Valid
}

func (rec *NullString) Set(value string) {
	rec.NullString.String = value
	rec.NullString.Valid = true
}

func (rec *NullString) Delete() {
	var val string
	rec.NullString.String = val
	rec.NullString.Valid = false
}

func (rec NullString) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.String)
}

func (rec *NullString) UnmarshalJSON(value []byte) error {
	var val *string

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val != nil {
		rec.String = *val
		rec.Valid = true
	} else {
		rec.Valid = false
	}

	return nil
}

var _ NullInterface = &NullTime{}

type NullTime struct {
	sql.NullTime
}

func (rec *NullTime) Get() *time.Time {
	if !rec.NullTime.Valid {
		return nil
	}

	val := rec.NullTime.Time

	return &val
}

func (rec *NullTime) GetValue() time.Time {
	return rec.NullTime.Time
}

func (rec *NullTime) GetValueWithDefault(value time.Time) time.Time {
	if rec.Valid {
		return rec.NullTime.Time
	}

	return value
}

func (rec *NullTime) GetString(format string) string {
	return rec.GetStringWithDefault(format, "")
}

func (rec *NullTime) GetStringWithDefault(format string, value string) string {
	if rec.NullTime.Valid {
		return rec.NullTime.Time.Format(format)
	}

	return value
}

func (rec *NullTime) IsNull() bool {
	return !rec.Valid
}

func (rec *NullTime) Set(value time.Time) {
	rec.NullTime.Time = value
	rec.NullTime.Valid = true
}

func (rec *NullTime) Delete() {
	var val time.Time
	rec.NullTime.Time = val
	rec.NullTime.Valid = false
}

func (rec NullTime) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Time)
}

func (rec *NullTime) UnmarshalJSON(value []byte) error {
	var val *time.Time

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val != nil {
		rec.Time = *val
		rec.Valid = true
	} else {
		rec.Valid = false
	}

	return nil
}

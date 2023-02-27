package wsa_lib_utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"regexp"
	"strings"
)

type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = err == nil
	return err
}

// make string to null
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func IsContains(str, word string) bool {
	return strings.Contains(str, word)
}

func FormatStringTime(str string) string {
	str1 := strings.ReplaceAll(str, "T", " ")
	return strings.Split(str1, ".")[0]
}

func FormatStringTimePostgres(str string) string {
	str1 := strings.ReplaceAll(str, "T", " ")
	return strings.Split(str1, "+")[0]
}

func FormatStringForex(s, t string) string {
	if t == "FRONT" {
		return FormatCurrency(strings.Split(s, ".")[0])
	} else if t == "BACK" {
		return FormatCurrency(strings.Split(s, ".")[0]) + "," + strings.Split(s, ".")[1][:2]
	} else {
		return FormatCurrency(strings.Split(s, ".")[0]) + "," + strings.Split(s, ".")[1]
	}
}

func FormatCurrency(s string) string {
	var buffer bytes.Buffer
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if count%3 == 0 && count != 0 {
			buffer.WriteString(".")
		}
		buffer.WriteByte(s[i])
		count++
	}
	return ReverseString(buffer.String())
}

func ReverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func ListContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ClearString(str string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

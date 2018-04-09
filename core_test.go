package go_case

import (
	"fmt"
	"testing"
)

var testCases = []string{
	"stringService",
	"StringService",
	"Stringservice",
	"string_service",
	"String_service",
	"string_Service",
	"JSONService",
	"jsonService",
	"JSONServiceV2",
	"JSONService V2",
	"String.Service",
	"String--Service",
	"ТестКириллицы",
	"Ещё-Тест Кириллицы",
	"и-ещё-тест.Кириллицы",
}

func Test_ToSnakeCase(t *testing.T) {
	answers := []string{
		"string_service",
		"string_service",
		"stringservice",
		"string_service",
		"string_service",
		"string_service",
		"json_service",
		"json_service",
		"json_service_v2",
		"json_service_v2",
		"string_service",
		"string__service",
		"тест_кириллицы",
		"ещё_тест_кириллицы",
		"и_ещё_тест_кириллицы",
	}
	if len(testCases) != len(answers) {
		t.Fatal("different amount of test cases and expected answers")
	}
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("%d input %s", i+1, tt), func(t *testing.T) {
			if got := ToSnakeCase(tt); got != answers[i] {
				t.Errorf("ToSnakeCase(%s) = %v, want %v", tt, got, answers[i])
			}
		})
	}
}

func Test_ToUrlSnakeCase(t *testing.T) {
	answers := []string{
		"string-service",
		"string-service",
		"stringservice",
		"string-service",
		"string-service",
		"string-service",
		"json-service",
		"json-service",
		"json-service-v2",
		"json-service-v2",
		"string-service",
		"string--service",
		"тест-кириллицы",
		"ещё-тест-кириллицы",
		"и-ещё-тест-кириллицы",
	}
	if len(testCases) != len(answers) {
		t.Fatal("different amount of test cases and expected answers")
	}
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("%d input %s", i+1, tt), func(t *testing.T) {
			if got := ToUrlSnakeCase(tt); got != answers[i] {
				t.Errorf("ToUrlSnakeCase(%s) = %v, want %v", tt, got, answers[i])
			}
		})
	}
}

func Test_ToDotSnakeCase(t *testing.T) {
	answers := []string{
		"string.service",
		"string.service",
		"stringservice",
		"string.service",
		"string.service",
		"string.service",
		"json.service",
		"json.service",
		"json.service.v2",
		"json.service.v2",
		"string.service",
		"string..service",
		"тест.кириллицы",
		"ещё.тест.кириллицы",
		"и.ещё.тест.кириллицы",
	}
	if len(testCases) != len(answers) {
		t.Fatal("different amount of test cases and expected answers")
	}
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("%d input %s", i+1, tt), func(t *testing.T) {
			if got := ToDotSnakeCase(tt); got != answers[i] {
				t.Errorf("ToDotSnakeCase(%s) = %v, want %v", tt, got, answers[i])
			}
		})
	}
}

func Test_ToCamelCase(t *testing.T) {
	answers := []string{
		"StringService",
		"StringService",
		"Stringservice",
		"StringService",
		"StringService",
		"StringService",
		"JSONService",
		"JsonService",
		"JSONServiceV2",
		"JSONServiceV2",
		"StringService",
		"StringService",
		"ТестКириллицы",
		"ЕщёТестКириллицы",
		"ИЕщёТестКириллицы",
	}
	if len(testCases) != len(answers) {
		t.Fatal("different amount of test cases and expected answers")
	}
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("%d input %s", i+1, tt), func(t *testing.T) {
			if got := ToCamelCase(tt); got != answers[i] {
				t.Errorf("ToCamelCase(%s) = %v, want %v", tt, got, answers[i])
			}
		})
	}
}

func Test_ToCamelCaseLowerFirst(t *testing.T) {
	answers := []string{
		"stringService",
		"StringService",
		"Stringservice",
		"stringService",
		"StringService",
		"stringService",
		"JSONService",
		"jsonService",
		"JSONServiceV2",
		"JSONServiceV2",
		"StringService",
		"StringService",
		"ТестКириллицы",
		"ЕщёТестКириллицы",
		"иЕщёТестКириллицы",
	}
	if len(testCases) != len(answers) {
		t.Fatal("different amount of test cases and expected answers")
	}
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("%d input %s", i+1, tt), func(t *testing.T) {
			if got := ToCamelCaseLowerFirst(tt); got != answers[i] {
				t.Errorf("ToCamelCaseLowerFirst(%s) = %v, want %v", tt, got, answers[i])
			}
		})
	}
}

func Test_ToNoCase(t *testing.T) {
	answers := []string{
		"stringService",
		"StringService",
		"Stringservice",
		"string_service",
		"String_service",
		"string_Service",
		"JSONService",
		"jsonService",
		"JSONServiceV2",
		"JSONService V2",
		"String.Service",
		"String--Service",
		"ТестКириллицы",
		"Ещё-Тест Кириллицы",
		"и-ещё-тест.Кириллицы",
	}
	if len(testCases) != len(answers) {
		t.Fatal("different amount of test cases and expected answers")
	}
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("%d input %s", i+1, tt), func(t *testing.T) {
			if got := ToNoCase(tt); got != answers[i] {
				t.Errorf("ToNoCase(%s) = %v, want %v", tt, got, answers[i])
			}
		})
	}
}

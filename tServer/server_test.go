package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestDoubleFunc(t *testing.T) {

	testCases := []struct {
		Name       string
		Input      string
		StatusCode int
		Error      string
		Result     int
	}{
		{Name: "double test 3", Input: "3", StatusCode: http.StatusOK, Error: "", Result: 6},
		{Name: "double test 8", Input: "8", StatusCode: http.StatusOK, Error: "", Result: 16},
		{Name: "double test not a number", Input: "n", StatusCode: http.StatusBadRequest, Error: "not a numbern", Result: -2},
		{Name: "double test miss value", Input: "", StatusCode: http.StatusBadRequest, Error: "missing error", Result: 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodGet, "localhost:8080/double?v="+testCase.Input, nil)
			if err != nil {
				t.Fatalf("crate request fail %v", err)
				return
			}

			rec := httptest.NewRecorder()

			doubleFunc(rec, req)
			res := rec.Result()
			if res.StatusCode != testCase.StatusCode {
				t.Errorf("response code err %d", testCase.StatusCode)
				return
			}
			rBytes, err := io.ReadAll(res.Body)
			defer res.Body.Close()

			if err != nil {
				t.Fatalf("read body fail %v", err)
			}

			trimRes := strings.TrimSpace(string(rBytes))
			if res.StatusCode != http.StatusOK {
				if trimRes != testCase.Error {
					t.Errorf("response messge not equal got %s want %s", trimRes, testCase.Error)
				}
				return
			}

			result, err := strconv.Atoi(trimRes)
			if err != nil {
				t.Fatalf("convery error %v", err)
				return
			}

			if result != testCase.Result {
				t.Errorf("result error get %d, want %d", result, 12)
			}
		})
	}
}

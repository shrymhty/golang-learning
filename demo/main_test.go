package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
)

func TestHandleRoot(t *testing.T) {
	w := httptest.NewRecorder()
	
	handleRoot(w, nil)

	desireCode := http.StatusOK

	if w.Code != desireCode {
		t.Errorf("Expected status code %d, but got %d\nBody: %s", desireCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Welcome to our homepage")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("Bad return. Got %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleGoodbye(t *testing.T) {
	w := httptest.NewRecorder()

	handleGoodbye(w, nil)

	desiredCode := http.StatusOK

	if w.Code != desiredCode {
		t.Errorf("Expected status code %d, but got %d\nBody: %s", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Goodbye!")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("Bad return. Got %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleParameterized(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?user=TestMan", nil)

	w := httptest.NewRecorder()

	handleParameterized(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad response code, expected: %v but got %v\nBody: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, TestMan!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("Bad return, got %q, expected %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleHelloParameterizedNoParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	handleParameterized(w, req)

	desiredCode := http.StatusOK

	if desiredCode != w.Code {
		t.Errorf("Bad response code, expected %v got %v\nBody: %s", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, User!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("Bad return. Got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleParameterizedWrongParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?foo=bar", nil)
	w := httptest.NewRecorder()

	handleParameterized(w, req)

	desiredCode := http.StatusOK

	if desiredCode != w.Code {
		t.Errorf("Bad response code, expected %v got %v\nBody: %s", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, User!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("Bad return. Got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleUserResponsesHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/responses/TestMan/hello/", nil)
	req.SetPathValue("user", "TestMan")

	w := httptest.NewRecorder()

	handleUserResponsesHello(w, req)

	desiredCode := http.StatusOK

	if desiredCode != w.Code {
		t.Errorf("Bad response code, expected %v got %v\nBody: %s", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, User!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("Bad return. Got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}
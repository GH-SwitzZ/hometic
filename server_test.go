package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func TestCreatePairDevice(t *testing.T) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(Pair{DeviceID: 1234,UserID: 4321})
	req := httptest.NewRequest(http.MethodPost, "/pair-device", payload)
	rec := httptest.NewRecorder()

	PairDeviceHandler(rec, req)
	if http.StatusOK != rec.Code {
		t.Error("expect 200 OK but get", rec.Code)
	}

	expected := `{"status":"active"}`
	if rec.Body.String() != expected {
		t.Errorf("expect %q but get %q\n", expected, rec.Body.String())
	}
}

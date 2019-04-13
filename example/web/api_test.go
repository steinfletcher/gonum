package web

import (
	"github.com/steinfletcher/apitest"
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	apitest.New().
		Debug().
		Handler(newApp().Router).
		Post("/user/search").
		Body(`{"name": "jan", "verification_status": "VERIFIED"}`).
		Expect(t).
		Status(http.StatusOK).
		Body(`{"id":"1234","name":"jan","verification_status":"VERIFIED"}`).
		End()
}

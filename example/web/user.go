package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//go:generate gonum -types=VerificationStatusEnum

type VerificationStatusEnum struct {
	Verified    string `enum:"VERIFIED"`
	NotVerified string `enum:"NOT_VERIFIED"`
	Expired     string `enum:"EXPIRED"`
}

type User struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	VerificationStatus VerificationStatus `json:"verification_status"`
}

func getUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &user)
		if err != nil {
			panic(err)
		}

		if user.VerificationStatus == Verified && user.Name == "jan" {
			u := User{Name: user.Name, ID: "1234", VerificationStatus: Verified}
			jsonData, err := json.Marshal(u)
			if err != nil {
				panic(err)
			}
			_, err = w.Write(jsonData)
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}
}


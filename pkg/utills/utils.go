package utills

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)
func ParseBody(r *http.Request, x interface{}) { // interface can hold value of any type
	if body, err := ioutil.ReadAll(r.Body); err == nil { // read the body of request
		if err := json.Unmarshal([]byte(body), x); err != nil { // unmarshal the request body to the interface
			return
		}
	}
}
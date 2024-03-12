package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts an apikey from
// the headers of an HTTP request
//eg :=
//Authorization : Apikey {insert apikey here}

func GetAPIKey(headers http.Header) (string, error)  {
	val := headers.Get("Authorization")	
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth headers")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	 
	return vals[1], nil
}

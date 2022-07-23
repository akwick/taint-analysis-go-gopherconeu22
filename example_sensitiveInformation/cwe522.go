// CWE-522 Insufficently Protected Credentials
package cwe522

import (
	"fmt"
	"log"
)

type Authentication struct {
	Username string
	Password string `datapolicy:"password"`
}

func authenticate(a Authentication) (*AuthenticationResponse, error) {
	resp, err := makeAuthenticationRequest(a)
	if err != nil {
		log.Printf("unable to make authenticated request: %v\n", a)
		return nil, err
	}
	return resp, nil
}

func (a Authentication) String() string {
	return fmt.Sprintf("%s:%s", a.Username, a.Password)
}

func (a Authentication) StringLog() string {
	return fmt.Sprintf("%s:***", a.Username)
}

// just a stub, to allow the code to compile
type AuthenticationResponse struct{}

// just a stub, to allow the code to compile
func makeAuthenticationRequest(Authentication) (*AuthenticationResponse, error) { return nil, nil }

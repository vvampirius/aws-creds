// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html

package main

import (
	"encoding/json"
	"io"
	"log"
)

type Credentials struct {
	Version int
	AccessKeyId string
	SecretAccessKey string
	//SessionToken string
	//Expiration string
}

func (credentials *Credentials) Fprint(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(credentials)
	if err != nil { log.Println(err.Error()) }
	return err
}

func NewCredentials(accessKeyId, secretAccessKey string) *Credentials {
	credentials := Credentials{
		Version: 1,
		AccessKeyId: accessKeyId,
		SecretAccessKey: secretAccessKey,
	}
	return &credentials
}
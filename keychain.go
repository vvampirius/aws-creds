package main

import (
	"errors"
	"fmt"
	Keychain "github.com/keybase/go-keychain"
	"log"
)

func GetPasswordFromKeychain(name string, keychainPath string) ([]byte, error) {
	keychain := Keychain.NewWithPath(keychainPath)
	if err := keychain.Status(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	queryItem := Keychain.NewItem()
	queryItem.SetSecClass(Keychain.SecClassGenericPassword)
	queryItem.SetMatchLimit(Keychain.MatchLimitOne)
	queryItem.SetReturnData(true)
	queryItem.SetLabel(name)

	items, err := Keychain.QueryItem(queryItem)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if len(items) != 1 {
		err := errors.New(fmt.Sprintf("Query return %d results!\n", len(items)))
		log.Println(err.Error())
		return nil, err
	}

	return items[0].Data, nil
}

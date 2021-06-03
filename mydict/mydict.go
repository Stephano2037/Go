package mydict

import "errors"

//alias (별명) Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
//found by "key" value
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

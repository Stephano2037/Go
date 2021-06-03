package mydict

import "errors"

//alias (별명) Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word  already exists")
var errCantUpdate = errors.New("Can't update non-existing word")
var errCantDelete = errors.New("Can't delete word because there is no word")

// Search for a word
//found by "key" value
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
// found by key
func (d Dictionary) Add(word, def string) error {
	/*
		1.found redundunt word
		2. if not, then add new word
	*/

	_, err := d.Search(word)

	//switch
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	return nil
}

//Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil

}

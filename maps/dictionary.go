package maps

// Dictionary wraps a map of string k & v
type Dictionary map[string]string

// DictionaryErr wrapper for errors
type DictionaryErr string

const (
	// ErrorNotFound message to return when word not found
	ErrorNotFound = DictionaryErr("could not find the word you are looking for")

	// ErrorWordExists message to return when word already exists
	ErrorWordExists = DictionaryErr("the word alread exists in dictionary")

	// ErrorWordDoesNotExist message to return when word already exists
	ErrorWordDoesNotExist = DictionaryErr("the word does not alread exists in dictionary")
)

// Search finds the given element in a map
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

// Add adds words to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

// Update changes an existing value
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete removes a given word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}

package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func HandleCustomErrors(w http.ResponseWriter, err error) error {
	var sErr *json.SyntaxError
	var unmarshallError *json.UnmarshalTypeError

	switch {

	case errors.As(err, &sErr):
		msg := fmt.Sprintf("request body contains badly-formed JSON at position %d", sErr.Offset)
		http.Error(w, msg, http.StatusBadRequest)
		return nil

	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := "request body contains badly-formed JSON"
		http.Error(w, msg, http.StatusBadRequest)
		return nil

	case err.Error() == "http: request body too large ":
		msg := "request body must not be larger than 1MB"
		http.Error(w, msg, http.StatusRequestEntityTooLarge)
		return nil

	case errors.As(err, &unmarshallError):
		msg := fmt.Sprintf("request body contains invalid value for %q field (at position %d)", unmarshallError.Field, unmarshallError.Offset)
		http.Error(w, msg, http.StatusBadRequest)
		return nil

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg := fmt.Sprintf("request body contains unknown field %s", fieldName)
		http.Error(w, msg, http.StatusBadRequest)
		return nil

	case errors.Is(err, io.EOF):
		msg := "request body cannot be empty"
		http.Error(w, msg, http.StatusBadRequest)
		return nil

	default:
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	}
}

func CreateRandomToken(tokenLength int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	characterList := "abcdefighijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rChars := []rune(characterList)
	var token []rune
	if tokenLength > 10 {
		for i := 0; i <= tokenLength; i++ {
			randomInteger := rand.Intn(len(rChars))
			randomChar := rChars[randomInteger]
			token = append(token, randomChar)
		}
		return string(token), nil
	} else {
		return "", errors.New("token length cannot be lower than 10")
	}
}

func ConvertToJson(targetData interface{}) ([]byte, error) {
	convertedData, err := json.Marshal(targetData)
	if err != nil {
		msg := fmt.Sprintf("error found: %s", err)
		return nil, errors.New(msg)
	}
	return convertedData, nil
}

func ConvertFromJson(marshaledData []byte) (interface{}, error) {
	var placeholder interface{}
	err := json.Unmarshal(marshaledData, &placeholder)
	if err != nil {
		msg := fmt.Sprintf("error found: %s", err)
		return nil, errors.New(msg)
	}
	return placeholder, nil
}

func CheckValidHttpMethod(r *http.Request, acceptableMethod string) (bool, error) {
	if r.Method == acceptableMethod {
		return true, nil
	} else {
		return false, errors.New("wrong method provided")
	}
}

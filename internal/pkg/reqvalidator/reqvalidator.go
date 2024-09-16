package reqvalidator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate            *validator.Validate
	phoneRegex          = `^(\+[1-9]\d{9,14}|\d{10,15})$`
	slugRegex           = `^[a-z0-9]+(?:-[a-z0-9]+)*$`
	compiledPhoneRegexp *regexp.Regexp
)

func init() {
	validate = validator.New()

	compiledPhoneRegexp = regexp.MustCompile(phoneRegex)

	err := validate.RegisterValidation("phone", validatePhoneNumber)
	if err != nil {
		log.Printf("[reqvalidator][init] Unable to put validator for phonuNumber %v", err)
	}
}

// ReadRequest is
func ReadRequest(r *http.Request, request interface{}) (string, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	if err = json.Unmarshal(bodyBytes, &request); err != nil {
		return "", fmt.Errorf("json.Unmarshal: %w", err)
	}

	if err = validate.Struct(request); err != nil {
		return "", fmt.Errorf("validate.Struct: %w", err)
	}

	return string(bodyBytes), nil
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	return compiledPhoneRegexp.MatchString(phoneNumber)
}

func validateSlugString(slugString string) bool {
	re := regexp.MustCompile(slugRegex)
	return re.MatchString(slugString)
}

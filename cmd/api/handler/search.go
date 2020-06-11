package handler

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/fathlon/household_grant/model"
)

// Search executes search based on the given request params
func Search(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// mapSearchOperation takes params (map[string][]string) and maps to SearchOperation struct
func mapSearchOperation(params url.Values) (model.SearchOperation, error) {
	rt := reflect.TypeOf(model.SearchOperation{})
	newStruct := reflect.New(rt).Elem()

	for k, v := range params {
		if len(v) == 0 {
			continue
		}
		// only accept the first value if multiple exists
		currentValStr := v[0]

		// loop each fields in SearchOperation struct
		for i := 0; i < rt.NumField(); i++ {
			structField := rt.Field(i)
			jsonFullTag := structField.Tag.Get("json")
			// skip fields with no json tag
			if jsonFullTag == "" {
				continue
			}

			// extract only the tag name value
			jsonTag := strings.Split(jsonFullTag, ",")[0]

			if k == jsonTag {
				currentField := newStruct.Field(i)
				if !currentField.CanInterface() || !currentField.CanSet() {
					continue
				}

				// checks type of Field, and set value by type appropriately
				switch currentField.Interface().(type) {
				case int:
					val, err := strconv.ParseInt(currentValStr, 10, 64)
					if err != nil {
						return model.SearchOperation{}, errors.New("invalid type of value")
					}
					currentField.SetInt(val)
				case *bool:
					val, err := strconv.ParseBool(currentValStr)
					if err != nil {
						return model.SearchOperation{}, errors.New("invalid type of value")
					}
					currentField.Set(reflect.ValueOf(&val))
				default:
					return model.SearchOperation{}, errors.New("not implemented, type not supported")
				}
			}
		}
	}

	return newStruct.Interface().(model.SearchOperation), nil
}

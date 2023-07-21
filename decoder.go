package urlvaluescanner

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/vingarcia/structscanner"
	"gopkg.in/yaml.v3"
)

func Unmarshal(uv url.Values, obj any) error {
	return structscanner.Decode(obj, newDecoder(uv))
}

// decoder can be used to fill a struct with the values of from url.Values.
type decoder struct {
	sourceValues url.Values
}

func newDecoder(sourceValues url.Values) decoder {
	return decoder{
		sourceValues: sourceValues,
	}
}

// DecodeField implements the TagDecoder interface
func (e decoder) DecodeField(info structscanner.Field) (any, error) {
	tag := strings.Split(info.Tags["schema"], ",")
	key := tag[0]

	required := false
	if len(tag) > 1 {
		required = (tag[1] == "required")
	}

	value, ok := e.sourceValues[key]
	if !ok && required {
		return nil, fmt.Errorf("missing required query param: '%s'", key)
	}

	if info.Kind == reflect.String {
		return value[0], nil
	} else if info.Kind == reflect.Slice {
		return value, nil
	}

	v := reflect.New(info.Type)
	err := yaml.Unmarshal([]byte(value[0]), v.Interface())
	if err != nil {
		return nil, fmt.Errorf("unable to convert input value '%s=%s' to type %v: %w", key, value[0], info.Type, err)
	}

	return v.Elem().Interface(), nil
}

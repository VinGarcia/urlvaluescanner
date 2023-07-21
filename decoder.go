package urlvaluescanner

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/vingarcia/structscanner"
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

	v, ok := e.sourceValues[key]
	if !ok && required {
		return nil, fmt.Errorf("missing required query param: '%s'", key)
	}

	if info.Kind == reflect.Slice {
		return v, nil
	}

	return v[0], nil
}

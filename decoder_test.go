package urlvaluescanner_test

import (
	"net/url"
	"testing"

	"github.com/vingarcia/urlvaluescanner"
	tt "github.com/vingarcia/urlvaluescanner/internal/testtools"
)

func TestDecoder(t *testing.T) {
	t.Run("should unmarshal url values correctly", func(t *testing.T) {
		uv := url.Values{
			"name": []string{"fakeName"},
			"type": []string{"fakeType"},
		}

		var dto struct {
			Name string `schema:"name"`
			Type string `schema:"type"`
		}
		err := urlvaluescanner.Unmarshal(uv, &dto)
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, dto.Name, "fakeName")
		tt.AssertEqual(t, dto.Type, "fakeType")
	})

	t.Run("should unmarshal slices of strings", func(t *testing.T) {
		uv := url.Values{
			"name": []string{"fakeName1", "fakeName2"},
			"type": []string{"fakeType"},
		}

		var dto struct {
			Names []string `schema:"name,required"`
			Type  string   `schema:"type,required"`
		}
		err := urlvaluescanner.Unmarshal(uv, &dto)
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, dto.Names, []string{"fakeName1", "fakeName2"})
		tt.AssertEqual(t, dto.Type, "fakeType")
	})

	t.Run("should unmarshal required values correctly", func(t *testing.T) {
		uv := url.Values{
			"name": []string{"fakeName"},
			"type": []string{"fakeType"},
		}

		var dto struct {
			Name string `schema:"name,required"`
			Type string `schema:"type,required"`
		}
		err := urlvaluescanner.Unmarshal(uv, &dto)
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, dto.Name, "fakeName")
		tt.AssertEqual(t, dto.Type, "fakeType")
	})

	t.Run("should return an error if a required field is missing", func(t *testing.T) {
		uv := url.Values{
			"type": []string{"fakeType"},
		}

		var dto struct {
			Name string `schema:"name,required"`
			Type string `schema:"type,required"`
		}
		err := urlvaluescanner.Unmarshal(uv, &dto)
		tt.AssertErrContains(t, err, "missing", "query param", "name")
	})

	t.Run("should unmarshal any types != string using yaml.Unmarshal", func(t *testing.T) {
		uv := url.Values{
			"someInt": []string{"42"},
		}

		var dto struct {
			SomeInt int `schema:"someInt"`
		}
		err := urlvaluescanner.Unmarshal(uv, &dto)
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, dto.SomeInt, 42)
	})
}

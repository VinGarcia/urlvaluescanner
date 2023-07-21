
## How to use it

```golang
uv := url.Values{
	"name": []string{"fakeName"},
	"type": []string{"fakeType"},
}

var dto struct {
	Name string `schema:"name"`
	Type string `schema:"type,required"`
}
err := urlvaluescanner.Unmarshal(uv, &dto)
```

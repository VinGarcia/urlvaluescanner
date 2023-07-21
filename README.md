
## How to use it

```golang
uv := url.Values{
	"name": []string{"some name"},
	"type": []string{"type1", "type2"},
}

var dto struct {
	Name string `schema:"name"`
	Types []string `schema:"type,required"`
}
err := urlvaluescanner.Unmarshal(uv, &dto)
```

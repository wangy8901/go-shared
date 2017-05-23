package utils

// An interface that produces properties matched with a key
type PropertySource interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	// Generic method to get any property
	Get(key string) interface{}
}

// A simple property source that uses map[string]interface{} to store properties
// This implementation only supports top level property keys. Also, it does not
// ensure type compatibility so illegal type assertion will fail.
type mapPropertySource struct {
	p map[string]interface{}
}

func (s mapPropertySource) GetString(key string) string {
	return s.p[key].(string)
}

func (s mapPropertySource) GetInt(key string) int {
	return s.p[key].(int)
}

func (s mapPropertySource) GetBool(key string) bool {
	return s.p[key].(bool)
}

func (s mapPropertySource) Get(key string) interface{} {
	return s.p[key]
}

func NewMapPropertySource(source map[string]interface{}) PropertySource {
	return mapPropertySource{p: source}
}

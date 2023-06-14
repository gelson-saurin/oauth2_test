package dbav

type EndpointConfigPK struct {
	Url    string
	Method string
}

func (ec *EndpointConfigPK) ToStringMap() map[string]string {
	return map[string]string{
		"url":    ec.Url,
		"method": ec.Method,
	}
}

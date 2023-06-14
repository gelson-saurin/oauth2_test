package service

type EndpointConfigModel struct {
	EndpointConfigKeyModel EndpointConfigKeyModel
	Data                   string
}

type EndpointConfigKeyModel struct {
	Url    string
	Method string
}

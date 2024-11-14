package server

type AppKitServer[T any] struct {
	Dependencies map[string][]Dependency
	Controllers  map[string]Controller
	Config       *T
}

type IController = func() error

type Dependency interface {
	Process() error
}

type Controller struct {
}

type Transport interface {
	Run() error
}

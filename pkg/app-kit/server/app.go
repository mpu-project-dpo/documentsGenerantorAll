package server

func WithApp(controllers ...IController) {
	for _, controller := range controllers {
		err := controller()
		if err != nil {
			panic(err)
		}
	}
}

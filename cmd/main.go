package main

import "log"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	configProvider, err := runtime.NewEnv()
	if err != nil {
		return err
	}

	app, err := apirest.BuildApplication(
		configProvider,
		apirest.BuildRepositories(
			configProvider))
	if err != nil {
		return err
	}

	return app.Run()
}

package src

import "fmt"

type Configuration struct {
	EventStore struct {
		Address  string `envconfig:"EVENTSTORE_ADDR"`
		Username string `envconfig:"EVENTSTORE_USER"`
		Password string `envconfig:"EVENTSTORE_PWD"`
	}

	ServiceHost string `envconfig:"CERT"`
	Debug       bool
}

func NewConfiguration() *Configuration { return &Configuration{} }

func (c *Configuration) Loaded() { fmt.Println("configuration was loaded") }

func (c *Configuration) NamePrefix() string { return "" }

func (c *Configuration) Validate() error {

	fmt.Println("validate the configuration")

	return nil
}

package config

import "github.com/google/uuid"

var Version = ""

func init() {
	Version = uuid.New().String()
}

const Port = 8080

const PageTitle = "RepForge"

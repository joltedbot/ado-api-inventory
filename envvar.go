package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
)

type EnvVars struct {
	TenantId     string `validate:"required,len=36,uuid4"`
	ClientId     string `validate:"required,len=36,uuid4"`
	ClientSecret string `validate:"required,max=64,printascii"`
	Organization string `validate:"required,max=50,printascii"`
}

var validate *validator.Validate

func getAndValidateEnvVars() (EnvVars, error) {
	validate = validator.New(validator.WithRequiredStructEnabled())
	environment := EnvVars{
		TenantId:     os.Getenv("ADO_TENANT_ID"),
		ClientId:     os.Getenv("ADO_CLIENT_ID"),
		ClientSecret: os.Getenv("ADO_CLIENT_SECRET"),
		Organization: os.Getenv("ADO_ORGANIZATION"),
	}

	err := validate.Struct(environment)
	if err != nil {
		log.Println(err)
		return EnvVars{}, err
	}

	return environment, nil
}

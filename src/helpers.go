package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readConfigFile(filename string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func readConfigFileWithDefaultValues(filename string) map[string]interface{} {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return map[string]interface{}{}
	}
	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return map[string]interface{}{}
	}
	return config
}

func getEnvironmentVariables() map[string]string {
	var env map[string]string
	err := os.Getenv("GOOS")
	if err != nil {
		return env
	}
	return env
}

func getEnvironmentVariablesWithDefaultValues() map[string]string {
	var env map[string]string
	err := os.Getenv("GOOS")
	if err != nil {
		return map[string]string{}
	}
	return env
}

func getEnvironmentVariablesWithDefaultValuesFromConfig() map[string]string {
	var env map[string]string
	err := readConfigFile("config.json")
	if err != nil {
		return map[string]string{}
	}
	err = json.Unmarshal(data, &env)
	if err != nil {
		return map[string]string{}
	}
	return env
}

func getEnvironmentVariablesFromConfigWithDefaultValues() map[string]string {
	var env map[string]string
	err := readConfigFileWithDefaultValues("config.json")
	if err != nil {
		return map[string]string{}
	}
	err = json.Unmarshal(data, &env)
	if err != nil {
		return map[string]string{}
	}
	return env
}
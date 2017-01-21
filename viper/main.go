package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	// Note: Viper does not require any initialization before using, unless we'll be dealing multiple different configurations.
	// check [working with multiple vipers](https://github.com/spf13/viper#working-with-multiple-vipers)

	// Set config file we want to read. 2 ways to do this.
	// 1. Set config file path including file name and extension
	viper.SetConfigFile("./configs/env.json")

	// OR
	// 2. Register path to loook for config files in. It can accept multiple paths.
	// It will search these paths in given order
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("$HOME/configs")
	// And then register config file name (no extension)
	viper.SetConfigName("env")
	// Optionally we can set specific config type
	//viper.SetConfigType("json")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	// https://godoc.org/github.com/spf13/viper#Get
	// Get can retrieve any value given the key to use.
	// Get is case-insensitive for a key.
	// Get has the behavior of returning the value associated with the first place from where it is set.
	// Viper will check in the following order: override, flag, env, config file, key/value store, default
	// Get returns an interface. For a specific value use one of the Get____ methods.

	port := viper.Get("prod.port") // returns string
	//port := viper.GetInt("prod.port") // returns integer
	fmt.Printf("Value: %v, Type: %T\n", port, port)

	// Check if a particular key is set
	// Notice that we can trverse nested configuration e.g. prod.port
	if !viper.IsSet("prod.port") {
		log.Fatal("missing port number")
	}

	// Extract sub-tree using `Sub`
	prod := viper.Sub("prod")

	// Unmarshal into struct
	type config struct {
		Host    string
		Port    int
		enabled bool
	}

	var C config

	err := prod.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Println(C.Host)

}

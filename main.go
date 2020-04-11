package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	ConfigFileName      = "config_pig"   //configure file name
	ConfigFileExtension = "json"         //configure file extension
	ConfigPath          = "$HOME/Config" //configure path
)

func main() {
	fmt.Println("Starting point:")

	var (
		age     int
		fn, ln  string
		marital bool
	)

	pflag.IntVar(&age, "age", 23, "age of person")
	pflag.StringVar(&fn, "firstName", "atticus", "first name of person")
	pflag.StringVar(&ln, "lastName", "li", "last name of person")
	pflag.BoolVar(&marital, "married", true, "marital status")
	pflag.Parse()

	viper.SetConfigName(ConfigFileName)      // name of config file (without extension)
	viper.SetConfigType(ConfigFileExtension) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(ConfigPath)          // path to look for the config file in
	viper.AddConfigPath(".")                 // call multiple times to add many search paths
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetConfigFile(ConfigFileName + "." + ConfigFileExtension)
		} else {
			panic(fmt.Errorf("Read config err %s", err))
		}
	}

	if err = viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(fmt.Errorf("BindPFlags err %s", err))
	}

	age = viper.GetInt("age")
	fn = viper.GetString("firstName")
	ln = viper.GetString("lastName")
	marital = viper.GetBool("married")

	fmt.Println("Parms:")
	fmt.Println("--age ", age, "--firstName ", fn, "--lastName ", ln, "--married ", marital)

	if err = viper.WriteConfig(); err != nil {
		panic(fmt.Errorf("Write config %s", err))
	}

}

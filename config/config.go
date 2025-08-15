package config

import "os"
import "github.com/spf13/viper"

var Config *Conf

type Conf struct {
	System *System `yaml:"system"`
	Redis  *Redis  `yaml:"redis"`
	Jwt    *Jwt    `yaml:"jwt"`
}

type Redis struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisPassword string `yaml:"redisPassword"`
	RedisDbName   int    `yaml:"redisDbName"`
	RedisNetwork  string `yaml:"redisNetwork"`
}

type System struct {
	AppEnv   string `yaml:"appEnv"`
	Domain   string `yaml:"domain"`
	Version  string `yaml:"version"`
	HttpPort string `yaml:"httpPort"`
	Host     string `yaml:"host"`
}

type Jwt struct {
	Secret   string `yaml:"secret"`
	LifeSpan int    `yaml:"lifespan"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/local")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}

package config

type Configuration struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Grpc Grpc `yaml:"grpc"`
	Http Http `yaml:"http"`
}

type Grpc struct {
	Port int `yaml:"port"`
}

type Http struct {
	Port int `yaml:"port"`
}
type Database struct {
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Port     int    `yaml:"port"`
}

package config

type Config struct {
	DataBase string `yaml:"dataBase"`
}

type ConfigDB struct {
	UserDB     string `yaml:"dbUser"`
	PasswordDB string `yaml:"dbPassword"`
	NameDB     string `yaml:"dbName"`
	HostDB     string `yaml:"dbHost"`
	PortDB     int    `yaml:"dbPort"`
}

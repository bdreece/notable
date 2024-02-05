package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrParse = errors.New("failed to parse config")

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Logging struct {
	File  string `yaml:"file"`
	Level int    `yaml:"level"`
}

type Storage struct {
	RootDirectory  string `yaml:"root_dir"`
	TasksDirectory string `yaml:"templates_dir"`
	NotesDirectory string `yaml:"users_dir"`
}

type TLS struct {
	ServerCertPath   string `yaml:"server_cert_path"`
	ServerKeyPath    string `yaml:"server_key_path"`
	ClientCARootPath string `yaml:"client_ca_root_path"`
}

type Token struct {
	Issuer   string `yaml:"issuer"`
	Audience string `yaml:"audience"`
	Secret   string `yaml:"secret"`
}

type Config struct {
	Mode     string   `yaml:"mode"`
	Database Database `yaml:"database"`
	Logging  Logging  `yaml:"logging"`
	Storage  Storage  `yaml:"storage"`
	TLS      TLS      `yaml:"tls"`
	Token    Token    `yaml:"token"`
}

func Parse(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Join(err, ErrParse)
	}

	d := yaml.NewDecoder(f)
	c := new(Config)
	if err = d.Decode(c); err != nil {
		return nil, errors.Join(err, ErrParse)
	}

	return c, nil
}

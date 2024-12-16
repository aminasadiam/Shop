package config

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DatabaseConfig struct {
	Driver           string `json:"driver"`
	ConnectionString string `json:"connection_string"`
}

type ShopConfig struct {
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	ItemsPerPage     int    `json:"items_per_page"`
	EnableUserReview bool   `json:"enable_user_reviews"`
	Locale           string `json:"locale"`
	Timezone         string `json:"timezone"`
}

type SecurityConfig struct {
	JWTSecret      string `json:"jwt_secret"`
	SessionTimeout int    `json:"session_timeout"`
}

type EmailConfig struct {
	SMTPHost     string `json:"smtp_host"`
	SMTPPort     int    `json:"smtp_port"`
	SMTPUser     string `json:"smtp_user"`
	SMTPPassword string `json:"smtp_password"`
}

type PaymentConfig struct {
	DefaultGateway string `json:"default_gateway"`
	MerchantID     string `json:"merchant_id"`
	CallbackURL    string `json:"callback_url"`
}

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Shop     ShopConfig     `json:"shop"`
	Security SecurityConfig `json:"security"`
	Email    EmailConfig    `json:"email"`
	Payment  PaymentConfig  `json:"payment"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

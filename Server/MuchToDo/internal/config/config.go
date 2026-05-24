package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config stores all application configuration.
type Config struct {
	ServerPort         string   `mapstructure:"PORT"`
	MongoURI           string   `mapstructure:"MONGO_URI"`
	DBName             string   `mapstructure:"DB_NAME"`
	JWTSecretKey       string   `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationHours int      `mapstructure:"JWT_EXPIRATION_HOURS"`
	EnableCache        bool     `mapstructure:"ENABLE_CACHE"`
	RedisAddr          string   `mapstructure:"REDIS_ADDR"`
	RedisPassword      string   `mapstructure:"REDIS_PASSWORD"`
	LogLevel           string   `mapstructure:"LOG_LEVEL"`
	LogFormat          string   `mapstructure:"LOG_FORMAT"`
	CookieDomains      []string `mapstructure:"COOKIE_DOMAINS"`
	SecureCookie       bool     `mapstructure:"SECURE_COOKIE"`
	AllowedOrigins     []string `mapstructure:"ALLOWED_ORIGINS"`
}

// LoadConfig reads configuration ONLY from environment variables (Docker-safe)
func LoadConfig(path string) (Config, error) {
	var config Config

	// 🔥 CRITICAL: Make env vars take priority
	viper.AutomaticEnv()

	// Explicit binding (prevents weird parsing issues)
	_ = viper.BindEnv("MONGO_URI")
	_ = viper.BindEnv("PORT")
	_ = viper.BindEnv("DB_NAME")
	_ = viper.BindEnv("JWT_SECRET_KEY")
	_ = viper.BindEnv("REDIS_ADDR")

	// Defaults
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENABLE_CACHE", false)
	viper.SetDefault("JWT_EXPIRATION_HOURS", 72)
	viper.SetDefault("ALLOWED_ORIGINS", "http://localhost:5173")

	// ❌ DO NOT use ReadInConfig in Docker (this is causing your issue)
	_ = os.Setenv("CONFIG_MODE", "docker")

	// Load into struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	// Clean CSV fields
	config.AllowedOrigins = cleanCSV(viper.GetString("ALLOWED_ORIGINS"))
	config.CookieDomains = cleanCSV(viper.GetString("COOKIE_DOMAINS"))

	return config, nil
}

// cleanCSV safely parses comma-separated values
func cleanCSV(input string) []string {
	parts := strings.Split(input, ",")
	var result []string

	for _, p := range parts {
		v := strings.TrimSpace(p)
		v = strings.Trim(v, "\"'")
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}

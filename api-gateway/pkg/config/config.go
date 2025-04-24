package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Port               string `mapstructure:"PORT"`
	SupabaseURL        string `mapstructure:"SUPABASE_URL"`
	SupabaseKey        string `mapstructure:"SUPABASE_KEY"`
	HelpdeskServiceURL string `mapstructure:"HELPDESK_SERVICE_URL"`
	TicketServiceURL   string `mapstructure:"TICKET_SERVICE_URL"`
	IssuesServiceURL   string `mapstructure:"ISSUES_SERVICE_URL"`
	TrackServiceURL    string `mapstructure:"TRACK_SERVICE_URL"`
	NetworkServiceURL  string `mapstructure:"NETWORK_SERVICE_URL"`
	FAQServiceURL      string `mapstructure:"FAQ_SERVICE_URL"`
}

// LoadConfig loads configuration from environment variables and config file
func LoadConfig() (*Config, error) {
	// Set default values
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("HELPDESK_SERVICE_URL", "localhost:50051")
	viper.SetDefault("TICKET_SERVICE_URL", "localhost:50052")
	viper.SetDefault("ISSUES_SERVICE_URL", "localhost:50053")
	viper.SetDefault("TRACK_SERVICE_URL", "localhost:50054")
	viper.SetDefault("NETWORK_SERVICE_URL", "localhost:50055")
	viper.SetDefault("FAQ_SERVICE_URL", "localhost:50056")

	// Read environment variables
	viper.AutomaticEnv()

	// Check for required env vars
	requiredEnvVars := []string{
		"SUPABASE_URL",
		"SUPABASE_KEY",
	}

	for _, env := range requiredEnvVars {
		if viper.GetString(env) == "" {
			if os.Getenv("GO_ENV") != "development" {
				return nil, fmt.Errorf("required environment variable %s is not set", env)
			} else {
				// In development mode, use placeholders for required variables
				viper.Set(env, fmt.Sprintf("dev-%s-value", env))
			}
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

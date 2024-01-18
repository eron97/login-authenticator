package connect_db

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func Test_defaultConnectionProvider_NewConnectionDB(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	successConfig := Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	failedConfig := Config{
		User:     "adminmock",
		Password: "senhamock",
		Host:     "hostmock",
		Port:     "portamock",
		Database: "databasemock",
	}

	tests := []struct {
		name    string
		p       *defaultConnectionProvider
		wantErr bool
	}{
		{
			name: "sucesso na conexão",
			p: &defaultConnectionProvider{
				config: successConfig,
			},
			wantErr: false,
		},
		{
			name: "falha na conexão",
			p: &defaultConnectionProvider{
				config: failedConfig,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.p.NewConnectionDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultConnectionProvider.NewConnectionDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

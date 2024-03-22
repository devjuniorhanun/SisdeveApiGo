package env

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var Env *config

type config struct {
	GoEnv        string `mapstructure:"GO_ENV"`       // GO_ENV: Determina o ambiente, pode ser production, development, stage.
	GoPort       string `mapstructure:"GO_PORT"`      // GO_PORT: Determina a porta que vamos usar para receber requisições, vamos usar a porta :8080
	DatabaseURL  string `mapstructure:"DATABASE_URL"` // DATABASE_URL: Aqui fica a url de conexão com o banco de dados
	ViaCepURL    string `mapstructure:"VIA_CEP_URL"`
	JwtSecret    string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    *jwtauth.JWTAuth
}

// Função LoadingConfig()
// Responsável para carregar os arquivos de configurações com o viper
func LoadingConfig(path string) (*config, error) {
	viper.SetConfigFile("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&Env)
	if err != nil {
		return nil, err
	}
	Env.TokenAuth = jwtauth.New("HS256", []byte(Env.JwtSecret), nil)
	return Env, nil
}

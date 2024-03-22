// Programas executáveis iniciam pelo pacote logger
package logger

/*
Os códigos em Go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports
*/
import (
	"log/slog" // Pacote responsável por gerar logs
	"os"       // Pacote responsável por interagir com o Sistema Hospedeiro
)

// / Função que inicializa o slog
func InitLogger() {
	// Setando os Logs no formado Json
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

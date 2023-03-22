package app

type App struct {
	srv *server.Server
	lg  logger.Logger
}

func New(lg logger.Logger, cfg config.Config) (*App, error) {

}

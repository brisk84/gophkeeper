package logger

type Logger interface {
	Errorln(args ...interface{})
	Infoln(args ...interface{})
	Fatalln(args ...interface{})
}

func New(debug bool) (*zap.SugaredLogger, error) {

}

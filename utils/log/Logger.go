package log

type LoggerInterface interface {
	Error() error
	Warn() error
	Info() error
}

func Error(logger LoggerInterface) (err error) {
	err = logger.Error()
	return
}

func Warn(logger LoggerInterface) (err error) {
	err = logger.Warn()
	return
}

func Info(logger LoggerInterface) (err error) {
	err = logger.Info()
	return
}

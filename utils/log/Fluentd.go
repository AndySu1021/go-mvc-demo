package log

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/spf13/viper"
)

var FluentdClient *fluent.Fluent

type Fluentd struct {
	Tag  string
	Data map[string]string
}

func (fluentd Fluentd) Error() (err error) {
	fluentd.Data["log_level"] = "ERROR"
	err = fluentd.log(fluentd.Tag, fluentd.Data)
	return
}

func (fluentd Fluentd) Warn() (err error) {
	fluentd.Data["log_level"] = "WARN"
	err = fluentd.log(fluentd.Tag, fluentd.Data)
	return
}

func (fluentd Fluentd) Info() (err error) {
	fluentd.Data["log_level"] = "INFO"
	err = fluentd.log(fluentd.Tag, fluentd.Data)
	return
}

func (fluentd Fluentd) log(tag string, data map[string]string) (err error) {
	err = FluentdClient.Post(tag, data)
	return
}

func InitFluentd() (err error) {
	FluentdClient, err = fluent.New(fluent.Config{
		FluentPort: viper.GetInt("logger.fluentd.port"),
		FluentHost: viper.GetString("logger.fluentd.host"),
	})
	return
}

func Close() (err error) {
	err = FluentdClient.Close()
	return
}
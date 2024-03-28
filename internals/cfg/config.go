package cfg

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cfg struct {
	Port       string
	DBname     string
	DBuser     string
	DBpass string
	DBhost     string
	DBport     string
}

func LoadConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("GGKIT_SERV")
	v.Set("PORT", "8080")
	v.Set("DBNAME", "learn_ggkit_db")
	v.Set("DBUSER", "postgres")
	v.Set("DBPASS", "root")
	v.Set("DBHOST", "")
	v.Set("DBPORT", "5432")
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}
	return cfg

}

func (cfg *Cfg) GetDBConnetcUrl() string { //маленький метод для сборки строки соединения с БД
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBuser,
		cfg.DBpass,
		cfg.DBhost,
		cfg.DBport,
		cfg.DBname,
	)
}
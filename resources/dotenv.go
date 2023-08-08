package resources

import (
	"fmt"
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/pkg/log"
)

const dotEnvName string = ".env"

func (r *Resources) initDotEnv() error {
	var dotEnv map[string]string

	if _, err := os.Stat(dotEnvName); err == nil || os.IsExist(err) {
		if err := godotenv.Load(dotEnvName); err != nil {
			return errors.Wrapf(err, "godotenv load")
		}

		if dotEnv, err = godotenv.Read(dotEnvName); err != nil {
			return errors.Wrapf(err, "godotenv read")
		}
	}

	r.dotEnvLog(dotEnv)

	return nil
}

// Логирование прочитанной мапы dotEnv.
func (r *Resources) dotEnvLog(dotEnv map[string]string) {
	if len(dotEnv) == 0 {
		log.Logger.Info().Msg("dotenv empty")

		return
	}

	dotEnvSlice := make([]string, 0, len(dotEnv))

	for env := range dotEnv {
		dotEnvSlice = append(dotEnvSlice, env)
	}

	sort.Strings(dotEnvSlice)

	dotEnvLog := "init dotenv success"
	for _, env := range dotEnvSlice {
		dotEnvLog += fmt.Sprintf("\n\t%s: %s", env, dotEnv[env])
	}

	log.Logger.Info().Msg(dotEnvLog)
}

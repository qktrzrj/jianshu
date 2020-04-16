package setting

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/yaml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/config/source/flag"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	defaultRootPath   = "app"
	defaultETCDPrefix = "/hello/user/conf"
	etcdConfig        defaultETCDConfig
	storageConfig     defaultStorageConfig
	loggerConfig      defaultLoggerConfig
	mailConfig        defaultEmailConfig
	redisConfig       defaultRedisConfig
	sp                = string(filepath.Separator)
	once              sync.Once

	HTTP_PORT  int
	JWT_SECRET string
)

func InitConfig() {
	once.Do(func() {
		// base config from env and override with flags
		config.Load(env.NewSource(), flag.NewSource())

		// switch to work space
		appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))
		os.Chdir(appPath)
		// get file source path
		cf := filepath.Join(appPath, "conf")
		fileSource := file.NewSource(file.WithPath(cf+sp+"config.yml"), source.WithEncoder(yaml.NewEncoder()))

		// override with file if exist
		config.Load(fileSource)

		log.Info().Str("Init", "loading config ...").Send()

		// first load etcd config, if etcd is work, then put etcdsource into config
		if err := config.Get(defaultRootPath, "etcd").Scan(&etcdConfig); err != nil {
			etcdConfig.Enable = false
		}
		// if etcdConfig enable,override with etcd
		if etcdConfig.Enable {
			log.Info().Str("Init", "check etcd dial ...").Send()
			etcdCfg := clientv3.Config{
				Endpoints:   []string{fmt.Sprintf("%s:%d", etcdConfig.Host, etcdConfig.Port)},
				DialTimeout: time.Second * 15,
			}
			client, err := clientv3.New(etcdCfg)
			if err != nil {
				panic(err)
			}
			defer client.Close()
			timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
			response, err := client.Get(timeout, defaultETCDPrefix)
			if err != nil {
				panic(err)
			}
			if response.Count > 0 {
				log.Info().Str("Init", "switch to etcd source ...").Send()
				config.Load(etcd.NewSource(
					etcd.WithAddress(fmt.Sprintf("%s:%d", etcdConfig.Host, etcdConfig.Port)),
					etcd.WithPrefix(defaultETCDPrefix),
					etcd.StripPrefix(true),
					etcd.WithDialTimeout(time.Second*5),
				))
			}
		}

		HTTP_PORT = config.Get(defaultRootPath, "http", "port").Int(8008)
		JWT_SECRET = config.Get(defaultRootPath, "jwt", "secret").String("shyptr")

		config.Get(defaultRootPath, "storage").Scan(&storageConfig)
		config.Get(defaultRootPath, "logger").Scan(&loggerConfig)
		config.Get(defaultRootPath, "mail").Scan(&mailConfig)
		config.Get(defaultRootPath, "redis").Scan(&redisConfig)
		// add config path watcher
		go watchEtcd()
	})
}

func watchEtcd() {
	watch, err := config.Watch(defaultRootPath, "etcd")
	if err != nil {
		// do nothing, but log
		log.Warn().Str("watch", "etcd watch start failed").AnErr("message", err).Send()
	}
	for v, err := watch.Next(); err == nil; v, err = watch.Next() {
		v.Scan(&etcdConfig)
	}
	watch.Stop()
}

func GetStorageConfig() StorageConfig {
	return storageConfig
}

func GetETCDConfig() ETCDConfig {
	return etcdConfig
}

func GetMailConfig() MailConfig {
	return mailConfig
}

func GetLoggerConfig() LoggerConfig {
	return loggerConfig
}

func GetRedisConfig() RedisConfig {
	return redisConfig
}

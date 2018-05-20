package anagram

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
)

// ConfigVals are configuration values for redis
type ConfigVals struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// RedisConfig is a wrapper for ConfigVals to support yaml structure
type RedisConfig struct {
	Redis ConfigVals
}

func (config *RedisConfig) getConfig() *RedisConfig {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
}

// RedisClient connects to redis and provides a redis client for cache
func RedisClient() *redis.Client {
	var config RedisConfig
	config.getConfig()
	address := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	time, _ := time.ParseDuration("2s")
	client := redis.NewClient(&redis.Options{
		Addr:        address,
		Password:    config.Redis.Password,
		DB:          config.Redis.Db,
		DialTimeout: time,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

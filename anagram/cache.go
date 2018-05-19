package anagram

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/go-redis/redis"
)

// RedisConfig in an ideal word should be externally configurable
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
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
func RedisClient() (*redis.Client){
	var config RedisConfig
	config.getConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:", "%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Db,
	})
	return client
	// _, _ := client.Ping().Result()
	// fmt.Println(pong, err)
	// // Output: PONG <nil>
}


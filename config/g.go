package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	projectName := "go-mega"
	getConfig(projectName)
}

func getConfig(projectName string) {
	// name of config file (without extension)
	viper.SetConfigName("config")
	// optionally look for config in the working directory
	viper.AddConfigPath(".")
	// call multiple times to add many search paths
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))
	// path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// GetMysqlConnectingString func
// Read mysql config
func GetMysqlConnectingString() string {
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")
	//tcp connection
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%-s?charset=%s&parseTime=true", user, pwd, host, db, charset)
}

// GetSMTPConfig func
func GetSMTPConfig() (server string, port int, user, pwd string) {
	server = viper.GetString("mail.smtp")
	port = viper.GetInt("mail.smtp-port")
	user = viper.GetString("mail.user")
	pwd = viper.GetString("mail.password")
	return
}

// GetServerURL func
func GetServerURL() (url string) {
	url = viper.GetString("server.url")
	return
}

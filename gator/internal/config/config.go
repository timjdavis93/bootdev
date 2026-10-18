package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"

func getConfigPath() (string, error) {

	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	configFilePath := home_dir + configFileName
	return configFilePath, nil

}

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func ReadConfig() (*Config, error) {

	configFilePath, err := getConfigPath()
	if err != nil {
		return &Config{}, err
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return &Config{}, err
	}

	config := Config{}
	if err := json.Unmarshal(configFile, &config); err != nil {
		return &Config{}, err
	}

	return &config, nil
}

func (c *Config) SetUser(currentUserName string) error {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configFilePath := home_dir + configFileName

	c.Current_user_name = currentUserName

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if err := os.WriteFile(configFilePath, data, 0666); err != nil {
		return err
	}

	return nil
}

/*func createFileIfNotExist(path string) error {
	file, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_EXCL|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil

} */

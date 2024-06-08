package mysql

import "eCommerce/config"

type MySql struct {
	config *config.Config
}

func NewMySql(config *config.Config) (*MySql, error) {
	m := &MySql{
		config: config,
	}

	return m, nil
}

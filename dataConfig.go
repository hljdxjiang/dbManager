package dataManager

import "os"
import "fmt"

type config struct {
	file string
}

func InitConfig(sfile string) (*config, error) {
	o := new(config)
	o.file = sfile
	_, err := os.Stat(sfile)
	if err == nil {
		return o, nil
	} else {
		return nil, fmt.Errorf(sfile + "dataManager config file is not existed")
	}
}

func (cf *config) GetFile() string {
	return cf.file
}

package unpacker

import (
	"github.com/mholt/archiver/v3"
	"log"
	"regexp"
)

type Unpacker struct {
	Logger        *log.Logger
}

func New(config *Unpacker) *Unpacker {
	return config
}

func (u *Unpacker) Extract(filename string, destinatin string) error {
	err := archiver.Unarchive(filename, destinatin)
	if err != nil {
		return err
	}
	return nil
}

func (u *Unpacker) GetAliasesDistrib(fileName string) string {
	regexp := regexp.MustCompile(`^1c-enterprise[\d]*-[\d\.\d\.\d*\.\d*]*-*([a-z-]*)_([\d\.\d\.\d*\.\d*-]*)_(amd64)\.([a-z]*)$`)
	resultFileName := regexp.ReplaceAllString(fileName, `$1-$2.$4`)
	return resultFileName
}
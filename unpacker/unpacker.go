package unpacker

import (
	"github.com/mholt/archiver/v3"
	"regexp"
)

var re = regexp.MustCompile(`^1c-enterprise[\d]*-[\d\.\d\.\d*\.\d*]*-*([a-z-]*)_([\d\.\d\.\d*\.\d*-]*)_(amd64)\.([a-z]*)$`)

func Extract(filename string, destinatin string) error {
	err := archiver.Unarchive(filename, destinatin)
	if err != nil {
		return err
	}
	return nil
}

func GetAliasesDistrib(fileName string) string {
	resultFileName := re.ReplaceAllString(fileName, `$1-$2.$4`)
	return resultFileName
}

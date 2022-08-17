package GoSystemD

import (
	"os"
	"path/filepath"
)

func GetExecutablePath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	return dir + "/" + filepath.Base(os.Args[0]), err
}

func NewService(name string, prms Parameters) Service {
	srv := Service{
		Name:   name,
		Params: prms,
	}

	return srv
}

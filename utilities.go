package GoSystemD

import (
	"bufio"
	"os"
	"os/user"
)

func isRoot() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, ErrUnableToGetUser
	}
	return currentUser.Username == "root", nil
}

func saveToFile(fileName string, content string) error {
	var err error

	f, err := os.Create(fileName)
	if err == nil {
		w := bufio.NewWriter(f)
		_, err = w.WriteString(content)
		w.Flush()
	}
	defer f.Close()

	return err
}

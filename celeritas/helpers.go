package celeritas

import "os"

func (c *Celeritas) CreateDirIfNotExist(path string) error {
	// permissions for folder
	const mode = 0755

	// creates a folder if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)

		// couldn't create folder
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Celeritas) CreateFileIfNotExists(path string) error {

	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if err != nil {
			return err
		}

		// close the file when we are done
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	return nil
}
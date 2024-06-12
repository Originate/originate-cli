package utils

import (
	"fmt"
	"os"
)

func CopyDir(src string, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	file, err := f.Stat()
	if err != nil {
		return err
	}

	if !file.IsDir() {
		return fmt.Errorf("Source " + file.Name() + " is not a directory!")
	}

	if err := os.Mkdir(dest, 0755); err != nil {
		return err
	}

	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			if err := CopyDir(fmt.Sprintf("%s/%s", src, f.Name()), fmt.Sprintf("%s/%s", dest, f.Name())); err != nil {
				return err
			}
		}

		if !f.IsDir() {
			content, err := os.ReadFile(src + "/" + f.Name())
			if err != nil {
				return err
			}

			err = os.WriteFile(dest+"/"+f.Name(), content, 0755)
			if err != nil {
				return err

			}
		}

	}

	return nil
}

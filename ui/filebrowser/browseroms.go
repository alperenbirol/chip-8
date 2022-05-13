package filebrowser

import "github.com/sqweek/dialog"

func BrowseRom() (string, error) {
	filename, err := dialog.File().Filter("ch8 CHIP-8 ROM", "ch8").Load()
	if err != nil {
		return "", err
	}
	return filename, nil
}

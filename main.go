package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"

	fileUtil "github.com/danielnetop/notesnook-to-standardnotes/internal/file"
	"github.com/danielnetop/notesnook-to-standardnotes/internal/notesnook"
	"github.com/danielnetop/notesnook-to-standardnotes/internal/sn"
)

const tagsFileName = "0_tags.txt"

func main() {
	backupFile, err := fileUtil.GetExportDataFile()
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	zf, err := fileUtil.ReadBackup(backupFile)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	defer func(zf *zip.ReadCloser) {
		err := zf.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(zf)

	files, err := notesnook.ValidateBackupFiles(zf)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	for _, file := range files {
		contentNotes, err := sn.ProcessConversion(file)
		if err != nil {
			fmt.Println(err.Error())

			return
		}

		err = fileUtil.CreateFileFromContent(contentNotes, file.FileName)
		if err != nil {
			fmt.Println(err.Error())

			return
		}
	}

	if tags := sn.ConvertNotebooksToTags(); len(tags.Items) > 0 {
		contentTags, err := json.Marshal(tags)
		if err != nil {
			fmt.Println(err.Error())

			return
		}

		err = fileUtil.CreateFileFromContent(contentTags, tagsFileName)
		if err != nil {
			fmt.Println(err.Error())

			return
		}
	}
}

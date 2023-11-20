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

	var nooks []notesnook.Nook

	for _, file := range files {
		processedNooks, err := notesnook.ProcessNotesnookExportData(file)
		if err != nil {
			fmt.Println(err.Error())

			return
		}

		// This needs to be done this way because the note title and content might not be
		// in the same notesnook backup file, so we need to fetch all data from backup
		// and only after that we can process them into a sn file
		nooks = append(nooks, processedNooks...)
	}

	err = sn.ProcessConversionAndSaveToFile(nooks)
	if err != nil {
		fmt.Println(err.Error())

		return
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

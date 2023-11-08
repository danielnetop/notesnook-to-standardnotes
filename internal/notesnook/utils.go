//nolint:gosec //Used only to check hash
package notesnook

import (
	"archive/zip"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	fileUtil "github.com/danielnetop/notesnook-to-standardnotes/internal/file"
)

func ValidateBackupFiles(zf *zip.ReadCloser) ([]ExportData, error) {
	var files []ExportData

	for _, file := range zf.File {
		if file.Name == ".nnbackup" {
			continue
		}

		content, err := readAll(file)
		if err != nil {
			return nil, err
		}

		if content.Encrypted {
			return nil, errNotDecryptedBackup
		}

		content.FileName = fmt.Sprintf("%s_converted.txt", file.Name)

		files = append(files, content)
	}

	return files, nil
}

func readAll(file *zip.File) (ExportData, error) {
	content, err := fileUtil.GetFileContent(file)
	if err != nil {
		return ExportData{}, errGetFileContentData
	}

	var exportData ExportData

	err = json.Unmarshal(content, &exportData)
	if err != nil {
		return ExportData{}, errUnmarshallingInputFile
	}

	// validate data with it's md5 hash
	if exportData.Hash != getMD5Hash(exportData.Data) {
		return ExportData{}, errDataDoesntMatchWithHashSum
	}

	return exportData, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ProcessNotesnookExportData(file ExportData) ([]Nook, error) {
	decode, err := base64.StdEncoding.DecodeString(file.Data)
	if err != nil {
		return []Nook{}, err
	}

	data, err := fileUtil.GUnzipData(decode)
	if err != nil {
		return []Nook{}, err
	}

	var nooks []Nook

	err = json.Unmarshal(data, &nooks)
	if err != nil {
		return []Nook{}, err
	}

	return nooks, nil
}

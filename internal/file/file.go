package file

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cespare/xxhash/v2"
	"github.com/gabriel-vasile/mimetype"
)

func GetFileContent(file *zip.File) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return []byte{}, err
	}

	defer func(f io.ReadCloser) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	return io.ReadAll(f)
}

func GUnzipData(data []byte) ([]byte, error) {
	var r io.Reader

	r, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var res bytes.Buffer

	_, err = res.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}

func CreateFileFromContent(content []byte, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return errExportingNotes
	}

	_, err = f.Write(content)
	if err != nil {
		return errWritingToFile
	}

	err = f.Close()
	if err != nil {
		return errSavingFile
	}

	return nil
}

func GetExportDataFile() (string, error) {
	// Validate arguments
	if len(os.Args) < 2 {
		return "", errNoFileProvided
	}

	flag.Parse()
	file := flag.Arg(0)

	if file == "" {
		return "", errNoFileProvided
	}

	return file, nil
}

func ReadBackup(backupFile string) (*zip.ReadCloser, error) {
	zf, err := zip.OpenReader(backupFile)
	if err != nil {
		return nil, err
	}

	return zf, nil
}

const (
	markdownImageFormat = "![%s](%s)"
	markdownFileFormat  = "[%s](%s)"
)

func ConvertFileToBase64(hash string) string {
	zipFile := "attachments.zip"

	zf, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func(zf *zip.ReadCloser) {
		err := zf.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(zf)

	for _, file := range zf.File {
		// this is because notesnook keeps same names in some cases
		// for example screenshots from mac are uploaded as image.png, and this will lead to multiple
		// image.png across the backup file, this way we compare the hash of the file and check if it's
		// indeed the correct file
		content, err := GetFileContent(file)
		if err != nil {
			fmt.Println(err.Error())
		}

		if hash == fmt.Sprintf("%016x", xxhash.Sum64(content)) {
			mimeType := mimetype.Detect(content)

			base64Encoding := fmt.Sprintf("data:%s;base64,%s", mimeType, base64.StdEncoding.EncodeToString(content))

			if strings.Contains(mimeType.String(), "image/") {
				return fmt.Sprintf(markdownImageFormat, file.Name, base64Encoding)
			}

			return fmt.Sprintf(markdownFileFormat, file.Name, base64Encoding)
		}
	}

	return ""
}

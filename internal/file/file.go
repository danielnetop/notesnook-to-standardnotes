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

func ConvertFileToBase64(filename string) string {
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
		// this is because notesnook doesn't keep original names in some cases
		// I uploaded one image with name `IMG_0001.jpg`, and then uploaded a different file
		// with same name `IMG_0001.jpg` and after the second upload the attachment became `img_0001.jpg`,
		// and it only kept the first attachment
		if strings.EqualFold(file.Name, filename) {
			content, err := GetFileContent(file)
			if err != nil {
				fmt.Println(err.Error())
			}

			mimeType := mimetype.Detect(content)

			base64Encoding := fmt.Sprintf("data:%s;base64,%s", mimeType, base64.StdEncoding.EncodeToString(content))

			if strings.Contains(mimeType.String(), "image/") {
				return fmt.Sprintf(markdownImageFormat, filename, base64Encoding)
			}

			return fmt.Sprintf(markdownFileFormat, filename, base64Encoding)
		}
	}

	return ""
}

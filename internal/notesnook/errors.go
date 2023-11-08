package notesnook

import "errors"

var errNotDecryptedBackup = errors.New("backup file is encrypted, please use a non encrypted backup")

var errUnmarshallingInputFile = errors.New("error extracting data from input file")

var errGetFileContentData = errors.New("error getting content data from file")

var errDataDoesntMatchWithHashSum = errors.New("data doesn't match with md5 hash")

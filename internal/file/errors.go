package file

import "errors"

var errExportingNotes = errors.New("error exporting notes")

var errWritingToFile = errors.New("error writing export to file")

var errSavingFile = errors.New("error saving file")

var errNoFileProvided = errors.New("path for `Notesnook` backup file is required")

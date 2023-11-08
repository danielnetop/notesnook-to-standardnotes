# Notesnook to Standard Notes converter

This is a CLI tool used to convert the full backup from [Notesnook](https://notesnook.com) into a [Standard Notes](https://standardnotes.com) decrypted import with tags and notes.
After the tool runs the Notesnook information will be ready to import into Standard Notes.

## Usage

- Download the latest release for your platform [here](https://github.com/danielnetop/notesnook-to-standardnotes/releases).
- Extract the file you've just downloaded.
- Run the CLI tool
  - Windows
    - `notesnook-to-standardnotes.exe path/to/notesnook-backup-timestamp.nnbackupz`
  - Linux/MacOS
    - `./notesnook-to-standardnotes path/to/notesnook-backup-timestamp.nnbackupz`
- Download a decrypted backup of your Notesnook account on `Settings > Backup & export > Create backup`
  - Also download your attachments on `Settings > Profile > Open manager > Select all > Download`
    - Move the backup file (`nnbackupz` file) to the folder where you extracted the `notesnook-to-standardnotes` tool
    - Move the attachments file (`attachments.zip` file) to the folder where you extracted the `notesnook-to-standardnotes` tool
- Find the converted Standard Notes import files in the same folder
  - You'll find multiple new files
    - `0-plain-id_converted.txt`
    - `0_tags.txt`

- Steps to import into Standard Notes
  - Go to `Preferences > Backups > Import backup`
    - Select files (you can start from 0) and select import
    - After all files are imported and if the file `0_tags.txt` is present you should import it 

### Example
```
# If you have the backup file on the same folder as the notesnook-to-standardnotes tool
## Windows
notesnook-to-standardnotes.exe notesnook-backup-timestamp.nnbackupz
## Linux/MacOS
./notesnook-to-standardnotes notesnook-backup-timestamp.nnbackupz
```

### MacOS

When executing the above command you might get the following message:

`"notesnook-to-standardnotes" can't be opened because Apple cannot check it for malicious software`

Press `Show in Finder` and right-click on the tool and click `Open` after that you might get another message:

`macOS cannot verify the developer of "notesnook-to-standardnotes". Are you sure you want to open it?`

Press `Open` and a new terminal window should open with the following:
```
notesnook-to-standardnotes-1.0.6-darwin-arm64/notesnook-to-standardnotes ; exit;
path for `Notesnook` backup file is required
```

Close the terminal window and follow the steps on [Usage](#usage).


## Badges

![Build Status](https://github.com/danielnetop/notesnook-to-standardnotes/workflows/Test/badge.svg)
[![License](https://img.shields.io/github/license/danielnetop/notesnook-to-standardnotes)](/LICENSE)
[![Release](https://img.shields.io/github/release/danielnetop/notesnook-to-standardnotes.svg)](https://github.com/danielnetop/notesnook-to-standardnotes/releases/latest)

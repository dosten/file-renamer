# File Renamer

_NOTE: This project is intended for my personal use as it do not offer a general solution for the main problem
but you are free to use it if it helps you._

Rename photos and videos in batch to a standarized name.

## How it Works

It recursively traverses all the files in the given directory and then:

1. Try to guess the creation date of the file:

    - Using the filename. Ex: `IMG-20220115-WA001.jpeg -> 2022/01/15`, if not possible,
    - Using the EXIF metadata, if not possible,
    - Keep old name for that file

2. Rename the file with a standard format: `YYYY-MM-DD-XXX.ext` where `XXX` is an incremental counter starting in 000.

## How to Build

```
make build
```

## How to Use

```
./bin/file-renamer [-dry-run] path/to/root/folder
```

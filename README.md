# CSCE4600

![Coverage](https://img.shields.io/badge/Coverage-92.6%25-brightgreen)

## To run project 2:

1. `cd Project2` if not already in the directory
2. `go run main.go`

----

### Commands I implemented

1. echo
   1. supports printing environment variables in a linux fashion, i.e `echo $VAR`
2. ls
   1. Supports the following flags
      1. -R    Recursively list files and directories
        -S    Sort files by size
        -a    Show hidden files
        -d    List directories themselves
        -h    Print file sizes in human-readable format
        -l    Use long listing format
        -r    Reverse the order of the listing
        -t    Sort files by modification time
   2. Can also combine flags like `ls -l -h`
3. mkdir
4. rm
5. cat

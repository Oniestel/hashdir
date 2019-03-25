# hashdir
A utility that creates a list of file hashes and tracks changes for a directory. 

## Usage

### hash
Creates a list of file hashes.
```
>hashdir hash directory
Start hashing. Directory: dir / Storage: (not set)
Hashing completed. Result saved: dir.csv
```

### check
Checks directory for changes and displays them.
```
>hashdir check -s dir.csv dir
Check started. Directory: dir / Storage: dir.csv
Check completed.
Added files:
dir/file2.txt
```

### copy
Copies changes to dest.
```
>hashdir copy -s dir.csv -d dir2 dir
Copy started. Directory: dir / Storage: dir.csv / Destination: dir2
dir/file2.txt
Copy completed.
```


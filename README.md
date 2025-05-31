## Command `gat` or `cat`

`gat` is a command line like `cat` but written in Go.

### Usage

```bash
gat [file1] [file2] ...
```

### Example

```bash
gat file1.txt file2.txt
```
This will concatenate the contents of `file1.txt` and `file2.txt` and print them to the standard output.

### Installation
Download the binary version for your platform from the [releases page](https://github.com/gabrieldebem/gat/releases) and place it in your `PATH`.

### Building from source

- git clone `git clone github.com/gabrieldebem/gat.git`
- cd `gat`
- run `go build -o gat .` 

rmdot
=====

Utility for removing all dotfiles such as `.git`.

## Install

Download the latest [golang](https://golang.org) and check the version at first.

```console
$ go version
```

Then run `go get`.

```console
go get -u github.com/moutend/rmdot/cmd/rmdot
```

## Quick start

```console
# List files before removing dotfiles.
$ ls -a
.
..
.git
.gitconfig
.gitignore
file1
file2
file3

# Remove all dotfiles.
$ rmdot -quiet ./

# That's it!
$ ls -a
.
..
file1
file2
file3
```

## Usage

```console
rmdot [options] path1 path2 ...
```

- Options:
  - `-d, -dry-run`: do not remove files
  - `-q, -quiet`: do not print anything

## Shell script vs `rmdot`

You can do the same thing with shell script, but the `rmdot` have the advantage of cross platform.

## LICENSE

MIT

## Author

`Yoshiyuki Koyanagi <moutend@gmail.com>`

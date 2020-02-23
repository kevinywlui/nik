# nik

`nik` - cd using subsequence matching and frecency.
Powered by golang and sqlite.

This program is essentially the same as
[autojump](https://github.com/wting/autojump),
[z](https://github.com/rupa/z),
and
[fasd](https://github.com/clvv/fasd). It was created to help me learn golang.

## Warning

The testing and error handling is severely lacking.

## Installation

Installation is done in 2 steps, installing the golang binary and then shell helper. 

### golang binary 

First install `golang` following the instructions here: <https://golang.org/doc/install>. Next install `nik`:
```
go get github.com/kevinywlui/nik
```

### Shell helper

Now we need to source the `nik.zsh` file. Add the following somewhere in your `~/.zshrc`
```
source `go env GOPATH`/src/github.com/kevinywlui/nik/nik.zsh
```


## Usage

The basic usage is executing
```
j bar
```
which will try to `cd` into a path, `/pre/.../fix/dir`, so that `dir` is a
subsequence of `bar`. Alternatively, executing
```
j foo bar
```
will try to `cd` into a path, `/pre/.../fix/dir`, so that `bar` is a subsequence
of `dir` and `foo` is a subsequence of `/pre/.../fix/`.

Alternatively, run `nik help` to see the possibilities with the `nik` command.

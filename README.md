# nik

`nik` provides an easy way to navigate between distant locations in the filesystem.

## Usage

The basic usage is executing
```
nik bar
```
which will try to `cd` into a path, `/pre/.../fix/dir`, so that `dir` is a
subsequence of `bar`. Alternatively, executing
```
nik foo bar
```
will try to `cd` into a path, `/pre/.../fix/dir`, so that `bar` is a subsequence
of `dir` and `foo` is a subsequence of `/pre/.../fix/`.

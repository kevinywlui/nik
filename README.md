# nik

`nik` - cd around the filesystem using subsequence matching and frecency.
Powered by `golang` and `sqlite`.

This program is essentially the same as
[autojump](https://github.com/wting/autojump),
[z](https://github.com/rupa/z),
or
[fasd](https://github.com/clvv/fasd). It was created to help me learn `golang`.

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

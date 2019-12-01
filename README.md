# git-get

Download a single file or directory from a git repository. Inspired by `svn export`.

Currently, we depend on the fact that most version control hosts provide a way to access and download the raw files of a repository. On Github we can download a file by executing `wget https://raw.githubusercontent.com/user/project/branch/file`, on Bitbucket it's `wget https://bitbucket.org/user/project/raw/commit-hash/file` and on Gitlab `wget https://gitlab.com/user/project/raw/branch/file`. As you can see, there's no standard way to download a file, let alone a whole folder.

Right now this is done by making an in-memory clone and copying the requested files. This is not viable for very large repos. The objective is to be able to download files without having to do a full clone.

## Installing

Just execute:

```
go get github.com/GuillemCastro/git-get
```

## Usage

```
git get <remote> <file or directory>

Arguments:
    remote  a Git repository
    file    the file or directory to be downloaded
```

For example to download the `main.go` file from this repository, we would execute:

`git get github.com/GuillemCastro/git-get main.go`

This will download the file onto the current folder. Note that it will override any local file.

To download a folder, the procedure is the same. If we want to downlaod the `src` folder from this repository https://github.com/GuillemCastro/rt-data, we would execute:

`git get github.com/GuillemCastro/rt-data src`

This will create a folder `src` in the current folder.

## Status

[x] Download a file or folder from the master/default branch
[ ] Download from a branch, tag or a specific commit
[ ] Don't require a complete clone

## License

```
MIT License

Copyright (c) 2019 Guillem Castro

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
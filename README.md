[![Build Status](https://travis-ci.org/yowcow/ghr.svg?branch=master)](https://travis-ci.org/yowcow/ghr)

GHR
===

A tool to fetch a release version string from GitHub releases.

HOW TO INSTALL
--------------

With Go:

```
go install github.com/yowcow/ghr
```

Or download binary executable from [releases](https://github.com/yowcow/ghr/releases), and place it in `$PATH`.

HOW TO USE
----------

To print the latest Vim version, do:

```
ghr -repo vim/vim HEAD
```

To print the version 3 releases prior to the latest, do:

```
ghr -repo vim/vim HEAD^^^
```

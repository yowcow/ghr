[![Build Status](https://travis-ci.org/yowcow/vimver.svg?branch=master)](https://travis-ci.org/yowcow/vimver)

VimVer
======

A small binary to find Vim version string from CLI.

HOW TO INSTALL
--------------

With Go:

```
go install github.com/yowcow/vimver
```

Or download binary executable from [releases](https://github.com/yowcow/vimver/releases), and place it in `$PATH`.

HOW TO USE
----------

To print the latest Vim version, do:

```
vimver HEAD
```

To print the version 3 releases prior to the latest, do:

```
vimver HEAD^^^
```

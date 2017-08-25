[![Build Status](https://travis-ci.org/yowcow/vim-ver.svg?branch=master)](https://travis-ci.org/yowcow/vim-ver)

Vim-Ver
=======

A small binary to find Vim version string from CLI.

HOW TO INSTALL
--------------

With Go:

```
go install github.com/yowcow/vim-ver
```

Or download binary executable from [releases](https://github.com/yowcow/vim-ver/releases), and place it in `$PATH`.

HOW TO USE
----------

To print the latest Vim version, do:

```
vim-ver HEAD
```

To print the version 3 releases prior to the latest, do:

```
vim-ver HEAD^^^
```

Clarent [![Build Status](https://magnum.travis-ci.com/fortytw2/clarent.svg?token=z6RRRTsAmDVdjJqpsjYs&branch=master)](https://magnum.travis-ci.com/fortytw2/clarent)
============
Clarent is a complete userland akin to the GNU Coreutils.

Not POSIX/SuS compliant by any means, but fully tested and feature complete.

Installation
------------
Clone the repo, build the binary with `CGO_ENABLED=0 go build -a` and do what 
you will with it (if you're cool with a dynamic binary, simply `go build`)


Developing applets
------------------
- Copy `applets/template` and name the copy like your applet
- Rename `template.go` and edit its contents to fit your applet
- Add your applet to `applets.go`

The template provides the basic framework you should stick to, ish.

Bugs
----
Yep. Report them on GitHub :)

Contact
-------
If you have ideas for missing applets, found a bug or have a suggestion, use
this [project’s issues](https://github.com/fortytw2/clarent/issues).
If you want to participate, just fork and code away. For questions contact me:
fortytw2 (at) gmail.com

5DPrint (Go implementation)
=======
[ status - unstable ]

5DPrint / fai·di·print / is tailor-made for the MakiBox A6 and modern 3D printing. 
The UI is designed for simplicity and letting the user get straight to printing. Devices are 
automatically detected and connected. Moving the extruder around has never been easier with 
the interactive print area.

(Go)5DPrint is a [Go](https://golang.org) based web application that generates a webserver and
uses a HTML5/CSS/Javascript based UI to provide the user experience and communication to the attached
3D printer. Building for the various platforms, make sure you're using Go1.2 and run:

```shell
cd <path-to-5DPrint>/
export GOPATH=`pwd`
cd src/github.com/makible/
go build -o 5DPrint *.go
```

Currently supports the MakiBox A6 (LT & HT) running the [MakiBox specific firmware](https://bitbucket.org/makible/makibox-firmware) 
living in bitbucket (__*Will soon port to github as well.*__).

For the *(preferred)* [Chrome App](https://www.google.com/intl/en/chrome/webstore/apps-create.html) implementation,
please refer to [5DPrint](https://github.com/Makible/5DPrint) for more details.

__*MakiBox flavored G/MCode info, coming soon*__

License
=======
Copyright (c) 2013, Makible Ltd. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted 
provided that the following conditions are met:

- Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

- Redistributions in binary form must reproduce the above copyright notice, this
  list of conditions and the following disclaimer in the documentation and/or
  other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR 
IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND 
FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR 
CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL 
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, 
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER 
IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT 
OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

# siv

A pure Go implementation of of the SIV-CMAC AEAD as described in
RFC 5297. SIV-CMAC does not require a nonce, allowing for both
deterministic and resistance to nonce re- or misuse.

## fuzzing

_build and install [dvyukov/go-fuzz](https://github.com/dvyukov/go-fuzz) first_

```bash
go-fuzz-build -o gofuzz-workdir/project-fuzz.zip github.com/joekir/siv-go
cd gofuzz-workdir
go-fuzz -bin=project-fuzz.zip -workdir=.
``` 

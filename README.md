# shamazing

A library to discover amazing things about a SHA1 hash (or MD5 or whatever).
It's **sha**-mazing. Almost as shamazing as that pun.

## Areas of shamazingness

- Longest string (9c2**cddfaedaea**9689a22e37aaa20191041554fe8)
- Longest integer (9c2cddfaedaea9689a22e37aaa**20191041554**fe8)
- Longest repeating sequence (9c2cddfaedaea9689a22e37**aaa**20191041554fe8)

## Install



## Development

If you'd like to hack on shamazing and add something neato or fix something that's lame-o, make sure you have [Go installed](https://golang.org), and then [install libgit2](https://github.com/libgit2/libgit2) — on OS X, this might be as easy for you as `brew install libgit2`.

## Usage

```sh
$ shamazing 9c2cddfaedaea9689a22e37aaa20191041554fe8
Longest string: cddfaedaea
Longest integer: 20191041554
Longest repeating: aaa

$ shamazing 9c2cddfaedaea9689a22e37aaa20191041554fe8 --string
cddfaedaea

$ shamazing 9c2cddfaedaea9689a22e37aaa20191041554fe8 -s
cddfaedaea

$ shamazing 9c2cddfaedaea9689a22e37aaa20191041554fe8 --integer
20191041554

$ shamazing 9c2cddfaedaea9689a22e37aaa20191041554fe8 -i
20191041554

$ shamazing 9c2cddfaedaea9689a22e37aaa20191041554fe8 --repeating
aaa

$ shamazing 9c2cddfaedaea9689a22e37aaa20191041554fe8 -r
aaa

# Search through the shas using `git log --format="%H"` in the current directory.
$ shamazing -s
cddfaedaea

# Return the full SHA with --full
$ shamazing --full
Longest string: 17705a5a37fbd11017f0d5e053b474dabbbd4022
Longest integer: 0c4b61fc2c5e7dd5566d42d0de1c431984899ddf
Longest repeating: 17705a5a37fbd11017f0d5e053b474dabbbd4022
```

## An Holman Project

Just a little diddy for dicking around with all these hashes. Made by
[@holman](https://twitter.com/holman).

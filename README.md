# shamazing

A library to discover amazing things about a SHA1 hash (or MD5 or whatever).
It's **sha**-mazing. Almost as shamazing as that pun.

## Areas of shamazingness

- Longest string (9c2**cddfaedaea**9689a22e376aa20191041554fe8)
- Longest integer (9c2cddfaedaea9689a22e376aa**20191041554**fe8)

## Install

    gem install shamazing

## Usage

### Command Line Interface

```sh
$ shamazing 9c2cddfaedaea9689a22e376aa20191041554fe8
Longest string: cddfaedaea
Longest integer: 20191041554

$ shamazing 9c2cddfaedaea9689a22e376aa20191041554fe8 --string
cddfaedaea

$ shamazing 9c2cddfaedaea9689a22e376aa20191041554fe8 -s
cddfaedaea

$ shamazing 9c2cddfaedaea9689a22e376aa20191041554fe8 --integer
20191041554

$ shamazing 9c2cddfaedaea9689a22e376aa20191041554fe8 -i
20191041554
```

### Ruby Interface

```ruby
require 'shamazing'

sha = '9c2cddfaedaea9689a22e376aa20191041554fe8'

Shamazing.string(sha)
# => cddfaedaea

Shamazing.integer(sha)
# => 20191041554
```

## An Holman Project

Just a little diddy for dicking around with all these hashes. Made by
[@holman](https://twitter.com/holman).
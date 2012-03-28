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

# Search through the shas using `git log --format="%H"` in the current directory.
$ shamazing -s
cddfaedaea

# Return the full SHA with --full
$ shamazing --full
Longest string: 17705a5a37fbd11017f0d5e053b474dabbbd4022
Longest integer: 0c4b61fc2c5e7dd5566d42d0de1c431984899ddf
```

### Ruby Interface

```ruby
require 'shamazing'

sha = '9c2cddfaedaea9689a22e376aa20191041554fe8'

Shamazing.string(sha)
# => cddfaedaea

Shamazing.integer(sha)
# => 20191041554

shas = %w(
  fdb31214c2cca29e4f723ad676cddb043bd73986
  0c4b61fc2c5e7dd5566d42d0de1c431984899ddf
  9c2cddfaedaea9689a22e376aa20191041554fe8
  f1b4c270f6746cbfff99bbf0f5a2388f4e509943
)
Shamazing.string_from_array(shas)
# => cddfaedaea
Shamazing.integer_from_array(shas)
# => 20191041554
```

## An Holman Project

Just a little diddy for dicking around with all these hashes. Made by
[@holman](https://twitter.com/holman).
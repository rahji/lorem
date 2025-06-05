# lorem (ipsum)

This is a simple command-line tool that gives you some [lorem ipsum](https://www.lipsum.com/).

It's really just a wrapper for a [fork](https://github.com/ro-ag/golorem) of a [Go package](chttps://github.com/drhodes/golorem)
that does the hard work. I just wanted to have an actual command that I could use in any operating system,
so I made this.

To install it, download the appropriate binary from the [Releases](releases) link to the right, then place
it somewhere in your PATH. Or just install it using Go, by typing this command in your terminal:
`go install github.com/rahji/lorem@latest`

To run the command, open your terminal application and type one of the following:

```bash
lorem word 4 10       # outputs a single word made up of 4-10 characters
lorem sentence 6 12   # outputs a single sentence made up of 6-12 words
lorem paragraph 4 8   # outputs a single paragraph made up of 4-8 sentences
```

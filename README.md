# go-dice-passwd
## A simple go app for generating dice roll based passphrases using the https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt wordlist.

### Usage

For now the program accepts a single argument which is the number of words to generate.
Please note, since the program is not yet finished, and accepts only a single char which will be converted to an integer. So entering a sinlge letter might will roll the dice based on the unicode value of the char. 
Numbers > 10 will only use the firt char for the passwd generation.

### Example 

```go run genpass.go```

This will prompt for a single char. 
Entering e.g. 8 will generate a passpherase with 8 words from the wordlist.  
to be continued...



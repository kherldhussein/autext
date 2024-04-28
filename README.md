# Autex

This tool is a simple text manipulation program written in Go. It performs various modifications to a given text file based on specified rules, such as converting words to uppercase or lowercase, replacing hexadecimal or binary numbers with their decimal equivalents, handling punctuation, and more.

## Features

- **Word Capitalization**: Converts words to uppercase or capitalizes them based on specified rules.
- **Word Case Conversion**: Converts words to lowercase based on specified rules.
- **Hexadecimal and Binary Conversion**: Replaces hexadecimal and binary numbers with their decimal equivalents.
- **Punctuation Handling**: Ensures proper formatting of punctuation marks.
- **Word Replacement**: Replaces "a" with "an" based on the following word's initial letter.
- **Quotation Handling**: Correctly positions quotation marks around words or phrases.
- **Vowel Handling**: Ensures proper usage of "a" or "an" before words starting with vowels or 'h'.

## Usage

1. **Installation**: Clone this repository to your local machine.

   ```bash
   git clone https://github.com/kherldhussein/autext
   ```

2. **Build**: Build the project using Go.

   ```bash
   go build .
   ```

3. **Run**: Execute the program with input and output file names.

   ```bash
   ./autex input.txt output.txt
   ```

## Rules

1. **(cap)**: Capitalizes the word before "(cap)".
2. **(low)**: Converts the word before "(low)" to lowercase.
3. **(up)**: Converts the word before "(up)" to uppercase.
4. **(hex)**: Replaces the word before "(hex)" with its decimal equivalent (if it's a hexadecimal number).
5. **(bin)**: Replaces the word before "(bin)" with its decimal equivalent (if it's a binary number).
6. **(cap, n)**: Capitalizes the previous n words before "(cap)".
7. **(low, n)**: Converts the previous n words before "(low)" to lowercase.
8. **(up, n)**: Converts the previous n words before "(up)" to uppercase.
9. **Punctuation**: Handles punctuation marks according to specified rules.
10. **Single Quotes**: Handles single quotes placement.
11. **"a" to "an"**: Replaces "a" with "an" if the following word starts with a vowel or 'h'.

## Examples

```bash
$ ./autex sample.txt result.txt
```

Input:
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

Output:
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

## Contributing

Contributions are welcome! If you have suggestions for improvements, open an issue or create a pull request.

## Author

This project was created and maintained by [Kherld Hussein](https://github.com/kherldhussein)

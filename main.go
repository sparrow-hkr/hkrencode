package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func encodeURL(input string, depth int) string {
	for i := 0; i < depth; i++ {
		input = url.QueryEscape(input)
	}
	return input
}

func fullHexEncode(input string) string {
	var encoded strings.Builder
	for _, c := range input {
		encoded.WriteString(fmt.Sprintf("%%%02x", c))
	}
	return encoded.String()
}

func decodeURL(input string) string {
	decoded, err := url.QueryUnescape(input)
	if err != nil {
		return "ERROR decoding: " + err.Error()
	}
	return decoded
}

func readLinesFromFile(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func main() {
	// Flags
	urlsFlag := flag.String("urls", "", "Comma-separated list of URLs to process")
	urlFileFlag := flag.String("urlfile", "", "Path to file containing list of URLs (one per line)")
	depthFlag := flag.Int("depth", 1, "Number of times to encode the URL")
	decodeFlag := flag.Bool("decode", false, "Decode the input instead of encoding")
	stdinFlag := flag.Bool("stdin", false, "Read input from standard input")
	outFlag := flag.String("out", "", "Save output to file")
	fullHexFlag := flag.Bool("fullhex", false, "Encode every character into hexadecimal ASCII")

	flag.Parse()

	var inputs []string

	// Read from stdin
	if *stdinFlag {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				inputs = append(inputs, line)
			}
		}
	} else if *urlsFlag != "" {
		inputs = strings.Split(*urlsFlag, ",")
	} else if *urlFileFlag != "" {
		lines, err := readLinesFromFile(*urlFileFlag)
		if err != nil {
			fmt.Println("Failed to read URL file:", err)
			return
		}
		inputs = lines
	} else {
		fmt.Println("No input provided. Use --urls, --urlfile, or --stdin.")
		return
	}

	var results []string
	for _, input := range inputs {
		if *decodeFlag {
			results = append(results, decodeURL(input))
		} else if *fullHexFlag {
			results = append(results, fullHexEncode(input))
		} else {
			results = append(results, encodeURL(input, *depthFlag))
		}
	}

	// Output
	if *outFlag != "" {
		file, err := os.Create(*outFlag)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}
		defer file.Close()
		for _, line := range results {
			file.WriteString(line + "\n")
		}
		fmt.Println("Output saved to", *outFlag)
	} else {
		for _, line := range results {
			fmt.Println(line)
		}
	}
}

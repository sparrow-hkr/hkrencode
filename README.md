# hkrencode
A url encode and decode easy way with hkrencode

# Instalation
```bash
go install github.com/sparrow-hkr/hkrencode@latest
```

# Usage
Help
```bash
hkrencode --help
```
Encode  
```bash
hkrencode --urls="https://example.com/admin?name=xor"
```
Encode URLs from file  
```bash
hkrencode --urlfile=urls.txt
```
Encode with depth  
```bash
hkrencode --urls="https://example.com/test" --depth=2
```
Decode Url
```bash
hkrencode --urls="https%3A%2F%2Fexample.com%2Ftest" --decode
```
Use Standard Input  
```bash
echo "https://example.com/test" | hkrencode --stdin
```
Save Output  
```bash
hkrencode --urls="https://example.com/test" --depth=3 --out=encoded.txt
```
Encode every character  
```bash
hkrencode --urls="/home/index.html" --fullhex
```

# Greph

Just like grep but for html document.

## Compile
Clone repository :
```bash
git clone https://github.com/ahmadrosid/greph.git
```

Compile :
```bash
go test
go build
```

## Usage

Grep from url :
```bash
./greph https://example.com "p[0].text"
```

Grep from stdin :
```bash
echo "<p>Paragraph</p>" | ./greph "p[0].text"
```

## Query
Format `selector[index]extractor`.

### Selector
Select by tag :
```bash
echo "<p>Paragraph</p>" | ./greph "p[0].text"
```
Select by class :
```bash
echo "<h1 class="title">Title</h1>" | ./greph ".title[0].text"
```

### Index
Get one from index `0`:
```bash
./greph https://example.com "p[0].text"
```

Get all bwtween index `2` and `5`:
```bash
./greph https://example.com "p[2:5].text"
```

Get all until index `5`:
```bash
./greph https://example.com "p[0:5].text"
```

### Extractor
Extract the text from tag :
```bash
echo "<p>Paragraph</p>" | ./greph "p[0].text"
```
Extract the text from attributes :
```bash
echo "<a href="https://example.com">Title</h1>" | ./greph "a[0]:href"
```
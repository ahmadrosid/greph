# Grep for Html document

Just like grep but for html document.

## Usage

Grep from url :
```bash
greph https://example.com "p[0].text"
```

Grep from stdin :
```bash
echo "<p>Paragraph</p>" | greph "p[0].text"
```


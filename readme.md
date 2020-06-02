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


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

| Selector                    | Index                                    | Extractor         |
| ----------------------------| -----------------------------------------| ----------------- |
| Class if started with `.`   | [0] get one at index `0`                 | Text if .text     |
| Tag if not started with `.` | [2:5] get all between index `2` and `5`  | Attr if :href,etc |
|                             | [:5] get all from index `0`              | Attr if :href,etc |

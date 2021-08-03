# Markhub
Parser for markdown using api github.

## Install
Get of module.
```
go get github.com/informeai/markhub
```
## Usage
Create markhub.
```
m := NewMarkHub()
```
Read of string markdown.
```
err := m.ParseString("# this is h1")
t := m.HTML()
fmt.Println(t)
```
Or read of file.
```
err := m.ParseFile(".test.md")
t := m.HTML()
fmt.Println(t)
```
### Serve
Serving of local address.
```
err := m.ParseFile(".test.md")
err = m.Serve(":8080")
```


by wellington gadelha :+1:

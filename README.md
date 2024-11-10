# edgo

A minimal implementation of the classic Unix line editor `ed` written in Go.

## Usage

### Basic Operation

```bash
# Start with empty buffer
edgo

# Edit a file
edgo filename.txt

# Start with custom prompt
edgo -p "* "
```

### Command Line Flags

- `-p string`: Set custom prompt (default "\*")
- `-v`: Display version information and exit

### Example Session

```
$ edgo
*a
hello world
line two
.
*w file.txt
2
*q
```

Please make sure to update tests as appropriate.

## TODO

- [ ] print buffer size on load
- [ ] don't show prompt when appending/inserting
- [ ] support basic error messages
- [ ] enable showing error messages with H similar to P
- [ ] Line ranges (e.g., `1,5p`)
- [ ] Search commands (`/pattern/`)
- [ ] Substitute command (`s/old/new/`)
- [ ] Copy/move lines
- [ ] Undo functionality

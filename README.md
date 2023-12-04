### Calstatus

read/add a text to your `calender.json`

### Usage:

- Read the next text from `calender.json` if available
```
$ calstatus
```
- Add a text to `calender.json`
```
$ calstatus --add "this is my text"
$ calstatus --time 2M --add "add this text to be shown in 2 minutes"
$ calstatus --at "2023-12-03 23:55:00" --add "this text to be shown at 23:55"
```

### Problem:

The `--at` flag:
- Should be not used with `--time`.
- Should be used with `--add`.
- Should be used like `+%Y-%d-%a %Hh%M` because go does not intelligent like `date` command in `unix/linux`.

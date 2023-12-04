# CALSTATUS

`calstatus` is a command-line tool that allows you to read and add texts to your `calender.json` file. You can use it to keep track of your events, tasks, or reminders in a simple and convenient way.

## Installation

To install Calstatus, you need to have Go installed on your system, Then, you can clone this repository and build the executable file:

```bash
$ git clone https://github.com/ayoubelmhamdi/calstatus.git
$ cd calstatus
$ go build
```

## Usage

To use `calstatus`, you need to have a `calender.json` file in your current directory. The file should have an array of objects, each with a `text` and a `time` property. The `text` property is the message that you want to display, and the `time` property is the date and time when you want to display it.

For example, your `calender.json` file could look like this:

```json
[
  {
    "text": "Meeting with John at 10 AM",
    "time": "2023-12-04 10:00:00"
  },
  {
    "text": "Buy groceries on the way home",
    "time": "2023-12-04 18:00:00"
  }
]
```

To read the next text from `calender.json` if available, you can run:

```bash
$ calstatus
```

To add a text to `calender.json`, you can use the `--text` flag with the message that you want to add. For example:

```bash
$ calstatus --text "Call mom tomorrow"
```

This will add the text to the end of the `calender.json` file with the current date and time.

You can also specify the date and time when you want to display the text by using the `--time` or the `--at` flag. The `--time` flag takes a relative time in minutes, hours, days, or weeks. For example:

```bash
$ calstatus --time 2M --text "Take a break"
```

This will add the text to the `calender.json` file with the date and time two minutes from now.

The `--at` flag takes an absolute date and time in the format `+%Y-%d-%a %Hh%M`. For example:

```bash
$ calstatus --at "2023-12-03 23:55:00" --text "Happy birthday!"
```

This will add the text to the `calender.json` file with the specified date and time.

Note: The `--at` flag should not be used with the `--time` flag, and it should be used with the `--text` flag.

### TODO:

This project is still in development, and needs more features. I will be happy to add more features, if you send me a PR.

- User Interface
    - [ ] Create a dmenu to pick time for the `--at` flag.
    - [ ] Create a dmenu to remove text from the `calender.json` file.
    - [ ] Create a dmenu to edit text in the `calender.json` file.
    - [ ] use profile mode(home/work) to add/retrieve text.
    - [ ] Create a frontend for the `calender.json` file with flutter/react/flet/raylib/pure js.

- Data Source
    - [ ] Change the location of the `calender.json` file.
    - [ ] Create some tools to fetch data from mobile calendar.
    - [ ] Create some tools to fetch data from web calendar.
    - [ ] Create some tools to generate data from some website, such as Amazon.

- Synchronization
    - [ ] Create some tools to sync data with mobile calendar.

- Compatibility
    - [ ] Support cross-platform.
    - [ ] Support cron job automatization.


## Contributing

Calstatus is an open source project, and contributions are welcome. If you want to contribute to this project, please follow these steps:

- Fork this repository and clone it to your local machine.
- Create a new branch for your feature or bug fix.
- Make your changes and commit them with a descriptive message.
- Push your branch to your forked repository and create a pull request.
- Wait for the code review and feedback.

## License

Calstatus is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.

## Acknowledgements

Calstatus was inspired by the [calcurse] project, a text-based calendar and scheduling application. I would like to thank the developers and contributors of calcurse for their work and inspiration.

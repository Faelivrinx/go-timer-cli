# Go Timer CLI

Go Timer CLI is a command-line utility that allows you to set a timer for a specific date and display the remaining time in the terminal, updating every second.

## Installation
To install Go Timer CLI, you must have Go installed on your system. If you don't have Go installed, you can download it from the official website: https://golang.org/dl/

Once you have Go installed, you can install Go Timer CLI using the following command:
```bash
go install github.com/Faelivrinx/go-timer-cli
```

Replace <username> and <repository-name> with your GitHub username and the name of your repository, respectively.

## Usage
To use Go Timer CLI, run the following command:

```bash
go-timer-cli --deadline=<deadline>
```

Replace <deadline> with the date and time you want to set the timer for, in the format YYYY-MM-DD HH:MM:SS.

For example, to set a timer for March 30, 2023 at 8:30 AM, you would run the following command:

```bash
go-timer-cli --deadline=2023-03-30 08:30:00
```

The timer will then start, and the remaining time will be displayed in the terminal, updating every second.

To stop the timer before the deadline, press Ctrl-C.

## License
Go Timer CLI is released under the MIT License. See LICENSE for more information.

---
I hope that helps! Let me know if you have any further questions.


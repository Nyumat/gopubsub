# pubsubbin 🚀

## Description

A simple go program which implements the publish-subscribe model as seen in [Go Patterns](https://refactoring.guru/design-patterns/go). I'll be facilitating inter-process communication with several goroutines that talk over a channel. There's an additional templ-based UI to visualize the messages being sent and received. (WIP)

The way it works is that the `Publisher` sends messages to the `Broker` which then sends the messages to the `Subscriber`. The `Broker` is the middleman that facilitates the communication between the `Publisher` and the `Subscriber`.

The `Publisher` and `Subscriber` are both goroutines that communicate with the `Broker` over a channel. The `Broker` is a singleton that manages the communication between the `Publisher` and `Subscriber`.

## Built With

- [Go](https://go.dev)
- [HTMX](https://htmx.org/)
- [Templ](https://templ.guide/)
- [Chi](https://go-chi.io/#/)
- [Task](https://taskfile.dev)
- [TailwindCSS](https://tailwindcss.com/)

## Prerequisites

- [Go](https://go.dev) 1.16 or later
- [Task](https://taskfile.dev)
- [Make](https://www.gnu.org/software/make)
- [npm](https://www.npmjs.com/get-npm)

## Usage

To run the program, both a Makefile and Taskfile is provided to simplify the process. The following commands are available:

---

| Command         | Description                    |
| --------------- | ------------------------------ |
| `make build`    | Builds the program             |
| `make run`      | Runs the program               |
| `make test`     | Runs the tests                 |
| `make clean`    | Cleans the build files         |
| `task run`      | Re-builds and Runs the program |
| `task test`     | Runs the tests                 |
| `task assets`   | Build the static assets        |
| `task generate` | Generate templ and mocks       |
| `task install`  | Install the program            |

---

For the full list of commands, run `make help` or `task --list`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file I provided for more details.

## Acknowledgements

- [Go Patterns](https://refactoring.guru/design-patterns/go)]

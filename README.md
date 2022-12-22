# pg-terminal

pg-terminal is a command-line tool for connecting to and interacting with PostgreSQL databases. With pg-terminal, you can easily run SQL statements, view database schema and data, and manage your PostgreSQL databases from the terminal.

## Features

- Connect to PostgreSQL databases using a user name, password, host, port, and database name.
- Run SQL statements from a file or from the command line.
- View database schema and data.
- Execute SQL statements and view the results in the terminal.

## Installation

To install pg-terminal, you will need to have Go installed on your system.

1. Download the source code for pg-terminal.
2. Navigate to the directory containing the source code.
3. Build the tool using the `go build` command. This will create an executable file called "pg-terminal".
4. Install the tool by moving the executable file to a directory in your system's `PATH` environment variable.

## Usage

To use pg-terminal, you will need to provide the connection details for your PostgreSQL database. These details include the user name, password, host, port, and database name. You can provide these details either as flags or as environment variables.

Here is an example of how to connect to a PostgreSQL database and run a SQL statement from a file:
~~~ bash
pg-terminal --file path/to/sql/file --user my_user --database my_database --password my_password --port 5432 --ip 127.0.0.1
~~~

You can also run SQL statement with overriting Environment Variable **PG_TERMINAL_SQL**
~~~ bash
PG_TERMINAL_SQL='<your sql statement>' pg-terminal --file path/to/sql/file --user my_user --database my_database --password my_password --port 5432 --ip 127.0.0.1
~~~

For a complete list of available flags and options, run `pg-terminal --help`.

## License

pg-terminal is released under the MIT License. See the [LICENSE](LICENSE) file for more details.


## Motivation

During work I searched for a cli postgres client I can just copy into a small docker container and use it without adding depedencies. Every other tool I tried needed to install more or had to have an interpreter or compiler. So I decided to publish a small lightweight version of a postgres CLI tool to interact with an postgres instance. 

## Disclaimer
This code is written in cooperation with chatGPT3. the main part or the code, tests, github actions and readme is created with chatGPT3. 
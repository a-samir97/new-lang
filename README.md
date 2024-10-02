# New Lang

New Lang is a programming language project written in Go.

## Features

- **If-Else Expression Evaluation**: Supports the evaluation of if-else expressions in the evaluator.
- **Infix Expression Evaluation**: Supports the evaluation of infix expressions, including operations for integers and booleans.
- **Postfix Expression Evaluation**: Supports the evaluation of prefix expressions, including handling of the '!' and '-' operators.

## Installation

To install New Lang, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/a-samir97/new-lang.git
   ```
2. Navigate to the project directory:
   ```sh
   cd new-lang
   ```
3. Build the project:
   ```sh
   go build
   ```

## Usage

To use New Lang, run the following command:
```sh
./new-lang [options] <input-file>
```

## Project Structure
### Specific Folder Responsibilities

- **ast/**: Responsible for defining the Abstract Syntax Tree (AST) for the language. This includes the structures and types used to represent the syntax of the language.
- **lexer/**: Responsible for lexical analysis. This folder contains the code for breaking down the input source code into tokens that can be further processed by the parser.
- **parser/**: Responsible for parsing the tokens produced by the lexer into an Abstract Syntax Tree (AST). This folder contains the code for syntactic analysis.
- **evaluator/**: Contains the code for evaluating the AST nodes. This includes the implementation of evaluation rules for different expressions and statements in the language.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries, please contact [a.samir97](mailto:a.samir9710@gmail.com).
```

For more details on PRs, you can visit the [pull requests page](https://github.com/a-samir97/new-lang/pulls?state=all).

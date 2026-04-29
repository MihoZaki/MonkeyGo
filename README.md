# 🐒 Monkey Interpreter (Extended)

> A handcrafted interpreter for the **Monkey** programming language, built from scratch in Go. Originally based on Thorsten Ball's *[Writing an Interpreter in Go](https://interpreterbook.com/)*, now actively extended with additional built-in functions, improved error handling, and plans for language-level expansions.

[![Go Version](https://img.shields.io/badge/Go-1.26%2B-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Tests](https://img.shields.io/badge/Tests-Passing-brightgreen)]()

## 🚀 Quick Start

```bash
# Clone & run
git clone https://github.com/MihoZaki/MonkeyGo.git
cd MonkeyGo
go run main.go

# Run the REPL
>> let x = [1, 2, 3];
>> let y = reverse(x);
>> y;
[3, 2, 1]
>> let z = pop(x);
>> z;
[1, 2]
```

## ✨ Features

### Core (Book Implementation)
- ✅ Lexical analysis with token streaming
- ✅ Pratt parser with full precedence climbing
- ✅ AST generation & structural equality testing
- ✅ Tree-walking evaluator with lexical scoping
- ✅ First-class functions, closures & higher-order functions
- ✅ Arrays, hash maps, string literals, booleans & integers
- ✅ Graceful error recovery & descriptive parser messages

### 🛠️ Extended Builtins
| Function | Description | Example |
|----------|-------------|---------|
| `pop(arr)` | Returns a new array with the last element removed | `pop([1,2,3]) → [1,2]` |
| `reverse(arr)` | Returns a new reversed array | `reverse([1,2,3]) → [3,2,1]` |
| `len(x)` | Length of arrays/strings | `len("hello") → 5` |
| `first(arr)`, `rest(arr)`, `push(arr, x)` | Functional list operations | `rest([1,2]) → [2]` |

## 🏗️ Architecture

```
Source Code
   │
   ▼
[Lexer]  →  Token Stream  →  [Parser]  →  AST  →  [Evaluator]  →  Runtime Objects
   │              │               │           │          │
   └─ ASCII       └─ Precedence  └─ Pratt    └─ Type    └─ Environment
                     Climbing       Descent     Switch     + Scoping
```

- **Lexer**: Converts raw ASCII source into typed tokens (byte-based scanning)
- **Parser**: Recursive descent + Pratt precedence parsing → produces a clean AST
- **Evaluator**: Direct AST traversal → resolves values in lexical environments
- **REPL**: Read-Eval-Print-Loop with live feedback & error reporting

## 📖 Language Syntax Examples

```monkey
// Variables & Functions
let double = fn(x) { x * 2 };
let numbers = [1, 2, 3, 4, 5];

// Higher-order functions
let mapped = map(numbers, double); // [2, 4, 6, 8, 10]

// Extended array operations
let reversed = reverse(mapped);    // [10, 8, 6, 4, 2]
let shortened = pop(reversed);     // [10, 8, 6, 4]
```

## 🗺️ Roadmap

- [ ] Native loop constructs (`for`, `while`)
- [ ] String manipulation builtins (`slice`, `contains`, `trim`)
- [ ] Unicode identifier support (rune-based lexer)
- [x] Single-line comments
- [x] Multi-lines comments
- [ ] Postfix operators(i++, i--)
- [ ] Index Expression for String literals
- [ ] Float type support

## 🧪 Testing

```bash
# Run full test suite
go test ./...

# Run with coverage
go test -cover ./...
```

## 🤝 Contributing

Pull requests are welcome! If you're extending builtins, adding language features, or improving error recovery:
1. Fork & create a feature branch
2. Keep changes scoped & well-tested
3. Update `README.md` & `ROADMAP.md` if adding public APIs
4. Submit a PR with a clear description & test coverage

## 📚 Acknowledgments

- **[Thorsten Ball](https://thorstenball.com)** – Author of *Writing an Interpreter in Go*. This project started as a faithful implementation of his teachings and is now evolving beyond it.

## 📄 License

MIT © [MihoZaki]


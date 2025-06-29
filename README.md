<h1 align="center">
  <img src="md/font.png" alt="Lite Java"> 
</h1>

**LiteJava** is a minimal, beginner-friendly programming language inspired by Java, designed to reduce boilerplate and simplify syntax. This project includes both:

- A **Transpiler** that converts LiteJava source code into valid Java code.
- A **Graphical IDE** built with the Fyne UI toolkit, enabling users to write, transpile, and view Java output in a user-friendly environment.

---

## 🌟 Features

- ✅ Converts LiteJava into valid Java source code
- ✅ Simple syntax for classes, functions, loops, and conditionals
- ✅ Automatic return type inference (`void`, `int`, `String`)
- ✅ Inline printing using `print`
- ✅ Built-in GUI IDE for writing and testing code
- ✅ Automatically saves the output as `Main.java` in the user's home directory

---

## 🧠 About LiteJava

LiteJava is a streamlined, pseudo-Java language designed to help beginners understand the structure of Java without overwhelming syntax. Key characteristics:

- **No semicolons**, braces `{}` can often be omitted
- **Simplified function and variable declarations**
- **Automatic type inference** for basic types
- **Less verbose entry point**: just write `main`, no need for full `public static void main(...)`

### LiteJava Example:

```python
class HelloWorld {
  main
    print "Hello, world!"
    var x = 5
    if (x > 3)
      print "x is greater than 3"
    func greet(name)
      print "Hello " + name
}
```

### Generated Java Code:
```java
public class HelloWorld {
  public static void main(String[] args) {
    System.out.println("Hello, world!");
    int x = 5;
    if (x > 3) {
      System.out.println("x is greater than 3");
    }
  }

  public static void greet(String name) {
    System.out.println("Hello " + name);
  }
}

```

## 📁 Project Structure

```pgsql
litejava/
├── transpiler/
│   └── transpiler.go     # LiteJava-to-Java conversion logic
├── main.go               # GUI application with Fyne
├── go.mod
├── go.sum
├── README.md
├── license.txt
└── md/
    └── font.png
    └── icon.png
```

## 🚀 Getting Started

### Requirements

- Go >= 1.18
- Java JDK (for running the output)
- Fyne Toolkit (go get fyne.io/fyne/v2)

### Build & Run

```bash
go mod tidy
go run main.go
```

### Output File

The transpiled Java code will be saved as:

```bash
~/Main.java
```

You can compile and run it with:

```bash
javac ~/Main.java
java HelloWorld
```

## 🔤 Supported LiteJava Syntax

| LiteJava Syntax   | Java Equivalent                          |
| ----------------- | ---------------------------------------- |
| `class Hello {`   | `public class Hello {`                   |
| `main`            | `public static void main(String[] args)` |
| `print "text"`    | `System.out.println("text");`            |
| `var x = 5`       | `int x = 5;`                             |
| `if (condition)`  | `if (condition) {`                       |
| `for (...)`       | `for (...) {`                            |
| `func name(args)` | `public static <type> name(String args)` |
| `return value`    | `return value;` (type inferred)          |


## 📦 Return Type Inference

Functions automatically determine their return type by scanning for return statements:

- return "Hello" → String
- return 42 → int
- No return → void

## 🖥️ GUI IDE

The included IDE allows you to:

- Write LiteJava code in a text editor
- Click "Translate" to convert it to Java
- View the generated Java code in a separate pane
- Automatically save the output file

## 📄 License

Licensed under the Apache License 2.0.

You may use, modify, and distribute this software under the terms of this license.

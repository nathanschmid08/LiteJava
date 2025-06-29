<h1 align="center">
  <img src="md/font.png" alt="Lite Java"> 
</h1>
<p align="center">
  <img src="https://img.shields.io/badge/Status-Stable-blue" alt="Status">
  <img src="https://img.shields.io/badge/Version-1.0-blue" alt="Version">
  <img src="https://img.shields.io/badge/Language-Go-blue" alt="Language">
</p>
<div align="center">
  
**LiteJava** is a minimal, beginner-friendly programming language inspired by Java, designed to reduce boilerplate and simplify syntax. This project includes both:
A **Transpiler** that converts LiteJava source code into valid Java code.
A **Graphical IDE** built with the Fyne UI toolkit, enabling users to write, transpile, and view Java output in a user-friendly environment.

  <h4>
    <em>Minimal • Beginner-Friendly • Zero Dependencies</em>
  </h4>
</div>
<p align="center">
  <a href="#-features"><b>🌟 Features</b></a> •
  <a href="#-getting-started"><b>🚀 Getting Started</b></a> •
  <a href="#-gui-ide"><b>🖥️ GUI IDE</b></a> •
  <a href="#-project-structure"><b>📁 Project Structure</b></a>
</p>
<hr>
<img align="right" src="md/icon.png" width="150">

## 🧠 About LiteJava
LiteJava is a streamlined, pseudo-Java language designed to help beginners understand the structure of Java without overwhelming syntax. Key characteristics:
- **No semicolons**, braces `{}` can often be omitted
- **Simplified function and variable declarations**
- **Automatic type inference** for basic types
- **Less verbose entry point**: just write `main`, no need for full `public static void main(...)`
> ✨ Modern language design with automatic type inference and simplified syntax
<hr>

## 🌟 Features
<table>
  <tr>
    <td width="200"><h3 align="center">🔄</h3><h3 align="center"><b>Transpiler</b></h3></td>
    <td>✅ Converts LiteJava into valid Java source code</td>
  </tr>
  <tr>
    <td width="200"><h3 align="center">📝</h3><h3 align="center"><b>Simple Syntax</b></h3></td>
    <td>✅ Simple syntax for classes, functions, loops, and conditionals</td>
  </tr>
  <tr>
    <td width="200"><h3 align="center">🧠</h3><h3 align="center"><b>Auto Inference</b></h3></td>
    <td>✅ Automatic return type inference (`void`, `int`, `String`)</td>
  </tr>
  <tr>
    <td width="200"><h3 align="center">🖨️</h3><h3 align="center"><b>Inline Printing</b></h3></td>
    <td>✅ Inline printing using `print`</td>
  </tr>
  <tr>
    <td width="200"><h3 align="center">🖥️</h3><h3 align="center"><b>GUI IDE</b></h3></td>
    <td>✅ Built-in GUI IDE for writing and testing code</td>
  </tr>
  <tr>
    <td width="200"><h3 align="center">💾</h3><h3 align="center"><b>Auto Save</b></h3></td>
    <td>✅ Automatically saves the output as `Main.java` in the user's home directory</td>
  </tr>
</table>
<hr>

## 🖥️ Example
### LiteJava Code
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
### Generated Java Code
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

### 🔤 Supported LiteJava Syntax
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
<hr>
  
## 🚀 Getting Started
### Requirements
- **Go** (>= 1.18)
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
> ⚠️ Ensure Java JDK is installed for compiling generated code
<hr>
  
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
<hr>
  
## 🖥️ GUI IDE
The included IDE allows you to:
- Write LiteJava code in a text editor
- Click "Translate" to convert it to Java
- View the generated Java code in a separate pane
- Automatically save the output file
<hr>
  
## 📦 Return Type Inference
Functions automatically determine their return type by scanning for return statements:
- return "Hello" → String
- return 42 → int
- No return → void
<hr>
  
## 📄 License

Licensed under the Apache License 2.0.
You may use, modify, and distribute this software under the terms of this license.
<div align="center">
  <p><i>© 2025 Simplified Java Programming Made Beautiful</i></p>
</div>

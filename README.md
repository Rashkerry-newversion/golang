# Go Learning Journey

Welcome to my Go (Golang) learning repository! 🚀  
This repo documents my journey of learning Go, starting from installation to writing my first program.

---

## 📥 Installation

### Windows

1. Download Go from the [official site](https://go.dev/dl/)

![Official site for GO](images/image1.png
)
2. Run the installer. By default, Go will be installed in:

   ```pwsh
   C:\Program Files\Go
   ```

3. Ensure Go is added to your PATH:
   - Open **Command Prompt** or **PowerShell**
   - Run:

     ```bash
     go version
     ```

   - If it prints the version, Go is ready to use!

![Go version](images/image2.png)

### Linux / Mac

```bash
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

---

## 📝 First Program

- Create a new folder for your project, for example `hello`.

- Inside, make a file called `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

![Folder Format](images/image3.png)

Run the program:

```bash
go run main.go
```

Output:

```bash
Hello, Go!
```

![Hello World](images/image4.png)

---

## 📚 What's Next?

- Learn about **variables and types**
- Explore **functions**
- Understand **packages and imports**
- Build simple CLI tools

This repo will be updated as I progress in my Go learning journey.  
Stay tuned for more lessons!

---

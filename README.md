# Securelee 🔐

A Open Source CLI tool built using Golang and Appwrite Cloud for Sharing Files Safely.

## 💻 Tech Stack :
- [Golang](https://go.dev/)
- [Appwrite](https://appwrite.io/)

## 🛠️ Installation :

### Requirements : 
- Golang ( version >= 1.19.3 )  [Installation Guide](https://golangdocs.com/install-go-windows)


 Run this command to install Securelee :
 ```
 go install github.com/Siddheshk02/Securelee@latest
 ```
 
 Start with Securelee :
 ```
 Securelee
 ```
 
 ## 📌 Commands :
 <!-- commands -->
 
 * [`Securelee [command] --help`](#Securelee-help)
 * [`Securelee SignUp`](#Securelee-SignUp)
 * [`Securelee login`](#Securelee-login)
 * [`Securelee share`](#Securelee-share)
 * [`Securelee get`](#Securelee-get)
 * [`Securelee view`](#Securelee-view)
 * [`Securelee downloads`](#Securelee-downloads)
 * [`Securelee delete`](#Securelee-delete)
 * [`Securelee whoami`](#Securelee-whoami)
 * [`Securelee logout`](#Securelee-logout)

## `Securelee [command] --help`

```
> Securelee --help

Securelee is a Cloud based CLI tool for Sharing Files securely.

Usage:
  Securelee [flags]
  Securelee [command]

Available Commands:
  SignUp      Sign-up to Securelee using your browser.
  completion  Generate the autocompletion script for the specified shell
  delete      Delete a Shared File.
  downloads   User who downloaded the Files shared by you
  get         Get a File.
  help        Help about any command
  login       Login to Securelee using your browser.
  logout      Command to Logout Securelee.
  share       Share a File
  view        Get list of Files Shared by you.
  whoami      A Command to see the Current logged-in User.

Flags:
  -h, --help     help for Securelee
  -t, --toggle   Help message for toggle

Use "Securelee [command] --help" for more information about a command.
```

## `Securelee SignUp`

```
Sign-up to Securelee using your email.

Usage:
  Securelee SignUp [flags]

Flags:
  -h, --help   help for SignUp
```

## `Securelee login`

```
Login to Securelee using your email.

Usage:
  Securelee login [flags]

Flags:
  -h, --help   help for login
```

## `Securelee share`

```
Upload a File -> Get a code -> Share the code with others.

Usage:
  Securelee share [flags]

Flags:
  -h, --help   help for share
```

## `Securelee get`

```
Enter the Code -> Get the File.

Usage:
  Securelee get [flags]

Flags:
  -h, --help   help for get
```

## `Securelee view`

```
Get list of Files Shared by you.

Usage:
  Securelee view [flags]

Flags:
  -h, --help   help for view
```

## `Securelee downloads`

```
Get the list of Users who downloaded the Files shared by you.

Usage:
  Securelee downloads [flags]

Flags:
  -h, --help   help for downloads

```

## `Securelee delete`

```
Delete a Shared File.

Usage:
  Securelee delete [flags]

Flags:
  -h, --help   help for delete
```

## `Securelee whoami`

```
A Command to see the Current logged-in User.

Usage:
  Securelee whoami [flags]

Flags:
  -h, --help   help for whoami

```

## `Securelee logout`

```
Command to Logout of Securelee.

Usage:
  Securelee logout [flags]

Flags:
  -h, --help   help for logout
```
 

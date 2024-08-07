# Mi CLI

This is a simple CLI to make encrypting/decrypting projects secrets that are safely stored in [Konbini](https://github.com/juancwu/mi)
among other features such as registering to Konbini, viewing all your bentos, etc... For all the available commands read [Command List](#command-list)

## Getting Started

### Method 1: Download from releases

```shell
# Change the variables as needed
curl -L -o mi https://github.com/juancwu/mi/releases/download/$VERSION/mi-$OS-$ARCH

# Make downloaded binary an executable
chmod +x mi

# Move executable to bin
sudo mv mi /usr/bin
```

### Method 2: Build Locally

You will need to have Go installed in your local machine.
 
```shell
# Clone the repo and cd into the repo
git clone git@github.com:juancwu/mi.git

# Build
go build -o mi .

# Move executable to bin
sudo mv mi /usr/bin
```

You can remove the repository if you want to.

## Command List

> Use `mi [command] --help` for more information about a specific command.

| Command | Description |
| ------- | ----------- |
| `mi auth signup` | Use to register and unlock the features of Konbini. |
| `mi bento prepare <bento-name>` | Fills a bento with the key value pairs in the given env file. |
| `mi bento fill <path-to-env-file>` | Fills a bento with the key value pairs in the given env file. |

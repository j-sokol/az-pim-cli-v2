# az-pim-cli

## Description

Command line tool to list and activate Azure PIM Role Assignments.

## Installation

```bash

git clone https://github.com/j-sokol/az-pim-cli-v2
cd az-pim-cli
go build -o az-pim-cli cmd/main.go

mv ./az-pim-cli /usr/local/bin
```

## Usage

This application accepts the following command-line parameters:

### List

- `-l`: List available role assignments

```bash
az-pim-cli -l
```

### Activate

- `-a`: Activate role
- `-r`: Role to activate
- `-s`: Scope of activation (i.e. Subscription / Resource)

Here's an example of how to use this application:

```bash
az-pim-cli -a -r 'key vault administrator' -s 'subscribtion_name' 
```

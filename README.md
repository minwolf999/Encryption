# Encryption

Encryption is a program for encrypting and decrypting an entire folder.

⚠️ This program can only decrypt folders encrypted by this program ⚠️

## Features

- Encrypt an entire folder while keeping an encrypted version and an unencrypted version.
- Decrypt an entire folder while keeping a decrypted version and an encrypted version.

## Prerequisites

Go version 1.22.0 ou supérieure

## Installation

Clone this repository and compile the project:

```bash
git clone https://github.com/minwolf999/encryption.git
cd encryption
(cd code/decrypt && go build -o ../../decrypt)
(cd code/encrypt && go build -o ../../encrypt)
```

## Utilisation

- Copy the folder you want to encrypt (or decrypt) into the `previousFile` folder
- Run the executable that corresponds to what you want to do: decrypt or encrypt
- The encrypted or decrypted result (depending on the executable used) will be added to the `transformFile` folder
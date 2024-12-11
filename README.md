# Zorro
![Awesome](https://awesome.re/badge.svg)

Zorro is a command-line tool for encryption, decryption, hashing, and RSA key management written in Go. It provides a simple interface for users to secure their data and manage cryptographic keys efficiently.

## Introduction

Zorro is a versatile command-line tool designed for various cryptographic operations, including encryption, decryption, hashing, and RSA key management. Developed with simplicity and efficiency in mind, Zorro offers a user-friendly interface for securing sensitive data and managing cryptographic keys effectively.

## Features

- **AES Encryption and Decryption**: Encrypt and decrypt files or strings using the Advanced Encryption Standard (AES) algorithm.
- **Hashing Algorithms**: Compute hashes of values using popular hashing algorithms like SHA-256 and MD5.
- **RSA Key Generation**: Generate RSA key pairs of custom bit lengths and optionally save them to a database for secure key management.
- **Database Integration**: Store and manage RSA keys in a database, providing centralized access control and key retrieval.

## Installation

Getting Zorro up and running is a breeze. Just follow these simple steps:

1. **Clone the Zorro repository:**

   Begin by cloning the Zorro repository to your local machine using Git. Open your terminal and execute the following command:

   ```bash
   git clone https://github.com/Peeentaa/Zorro.git
   ```
   
2.  **Build the project with Go:**

    Once you've cloned the repository, navigate into the zorro directory:

    ```bash
    cd zorro
    ```

    Now, build the project using Go:

    ```bash
    go build
    ```

3. **Optional: Install Zorro globally:**

    If you'd like to use Zorro from any directory on your system, you can install it globally:

    ```bash
    go install
    ```
   By installing Zorro globally, you can run it as a command-line tool from anywhere in your terminal.

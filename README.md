# Vernam Encrypter

Command-line tool which allows encrypting files with
[one-time pad encryption](https://en.wikipedia.org/wiki/One-time_pad). This looks like Gilbert Vernam encrypting device.

## Usage

```shell
vernam_encrypter command <arguments...>
```

Available commands:
* `keygen <length> <filename>` - generate key file. Arguments:
  * `<length>` - key length in bytes;
  * `<filename>` - path to generated key file;
* `encrypt <filename> <keyfile> <target_file>` - encrypt file with key. Arguments:
  * `<filename>` - path to source file to encrypt;
  * `<keyfile>` - path to key file to use;
  * `<target_file>` - path to target encrypted file;
* `decrypt <filename> <keyfile> <target_file>` - decrypt file with key. Arguments:
  * `<filename>` - path to source file to decrypt;
  * `<keyfile>` - path to key file to use;
  * `<target_file>` - path to target decrypted file;

## How to build

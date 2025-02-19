# Stegano

Stegano is a simple and efficient image steganography tool written in Go. It allows you to hide images within other images using least significant bit (LSB) encoding. This provides a basic way to conceal data in plain sight while maintaining the appearance of the host image.

## Features
- Hide any text into an image.
- Extract the hidden text back.
- Hide an image inside another image.
- Extract the hidden image from a stego-image.
- Supports PNG/JPG image formats.
- Uses least significant bit (LSB) encoding for steganography.
- Automatic resizing of the hidden image to fit within the host image.
- Simple CLI interface for easy usage.

## Usage
### Embedding a Text
To embed a hidden text inside a host image:
```sh
./stegano embed "Hello World" into secret.png as result.png
```
This will create `result.png` containing the hidden `Hello World` inside `secret.png`.

### Extracting a Hidden Text
To extract a hidden text from a stego-image:
```sh
./stegano ext from stego.png as extracted.png
```
This will extract the hidden text from `stego.png` and save it as `extracted.png`.


### Embedding an Image
To embed a hidden image inside a host image:
```sh
./stegano embed image input.png into secret.png as result.png
```
This will create `result.png` containing the hidden `input.png` inside `secret.png`.

### Extracting a Hidden Image
To extract a hidden image from a stego-image:
```sh
./stegano ext image from stego.png as extracted.png
```
This will extract the hidden image from `stego.png` and save it as `extracted.png`.

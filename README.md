# OCR GO

gosseract and Imagick web server implementation using gin gonic as framework

# How to install

# Install Imagick

To install Imagick, run the following command:

## Ubuntu/Linux
sudo apt-get install libmagickwand-dev

# Install Tesseract

## Linux

Tesseract is available directly from many Linux distributions. The package is generally called **'tesseract'** or **'tesseract-ocr'** - search your distribution's repositories to find it.
Thus you can install Tesseract 4.x and it's developer tools on Ubuntu 18.x bionic by simply running:
```
sudo apt install tesseract-ocr
sudo apt install libtesseract-dev
```

# Install Dependencies

```
cd project-folder
go get -u github.com/kardianos/govendor
govendor sync
```

# Run the project

Source or execute the `env.sh` then either `go build` and run the executable, or `go run main.go` to get debug log in terminal.

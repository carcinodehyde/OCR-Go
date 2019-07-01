# gosseractcv

gosseract and opencv implementation

# How to install

# Install GoCV

To install GoCV, run the following command:

```
go get -u -d gocv.io/x/gocv
```

To run code that uses the GoCV package, you must also install OpenCV 4.0.0 on your system. Here are instructions for Ubuntu, Raspian, macOS, and Windows.

## Ubuntu/Linux

### Installation

You can use `make` to install OpenCV 4.1.0 with the handy `Makefile` included with this repo. If you already have installed OpenCV, you do not need to do so again. The installation performed by the `Makefile` is minimal, so it may remove OpenCV options such as Python or Java wrappers if you have already installed OpenCV some other way.

#### Quick Install

The following commands should do everything to download and install OpenCV 4.1.0 on Linux:

	cd $GOPATH/src/gocv.io/x/gocv
	make install

If it works correctly, at the end of the entire process, the following message should be displayed:

	gocv version: 0.20.0
	opencv lib version: 4.1.0

That's it, now you are ready to use GoCV.

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
# aesnet256

## Overview:

This project is designed to provide an offline alternative to aesencryption.net. This is mainly to provide a secondary means to access the crypto puzzles by Master Boot Record and Keygen Church in the event that aesencryption.net were to go down or offline permanently. The goal is to provide as close as possible interoperability between encrypted data created on that site. See the section on [Reporting Bugs](https://github.com/irq-conflict/aesnet256?tab=readme-ov-file#reporting-issues) for information on creating bug reports.


## Live site:

You can visit a live working instance here: [AESNET256 Live Site](https://aesnet256.dragoncircle.co/)

## Installation:
The repository provides compiled binaries for Arm64 linux, AMD64 linux and AMD64 windows. These are located under the releases section. You can also [find them directly here](https://github.com/irq-conflict/aesnet256/releases). 

## Usage:
This application can run in two modes. First, there is a web mode which starts a locally hosted web service on port 7654 (which can be changed using the port flag). Simply run the application specific to your platform/hardware like this:

Linux/Intel or AMD64:

```
aesnet256.amd64 -server
```

Linux/ARM64:

```
aesnet256.arm64 -server
```

Windows/Intel or AMD64:

```
aesnet256.exe -server
```

You should be greeted with some small output about usage, after that open a web browser and browse to:

http://localhost:7654

and the page that allows you to choose encryption or decryption should show up.

The second way to use the utility is to specify command line parameters using -mode -key -text 

This will run the program using the provided input and provide any output assuming the program can decrypt the text using the provided key.

To run the application you provide the following parameters on the command line:

<pre>
   -server ( Enable server mode)
        Enable the http server to provide the service as a REST API and a HTMX web front end.

   -port (default port 7654) 
        Specify the default port to run the REST service and associated HTML pages on.

   -key (This is the encryption key) 
        Note: This can be less than 16 bytes, per the OpenSSL spec and the AESencryption website, these are zero padded to the proper length. 
        Keys that are longer than alloted will be truncated as needed. 
        Same as the aesencryption.net website.

   -text (This is the text you want to either encrypt or decrypt) 

   -iv [optional] 
        You can specify your own Initilization Vector. 
        If it is not supplied, the default one from the aesencryption.net website will be used.

   -aes_type [optional] You can specify 128, 192 or 256. The default is 256.

   -mode use encrypt or decrypt as options for the program.

   -v will display the current version number and exit.

</pre>

## Reporting issues:

 Open an Issue in the repository. Again, this program is mainly provided as a backup so these puzzles can be solved in the event of an issue with the Aesencryption.net website.

## Building:

This is designed to be built with GO111MODULE=auto

Clone the repository and then place it in your GOPATH
Usually this is $HOME/go/src

This should allow the build script contained in the cmd subfolder to be run.
It will create executable binaries for amd64/linux, arm64/linux and amd64/windows.

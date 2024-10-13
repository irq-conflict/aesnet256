# aesnet256

## Overview:

This project is designed to provide an offline alternative to aesencryption.net. This is mainly to provide a secondary means to access the crypto puzzles by Master Boot Record and Keygen Church in the event that aesencryption.net were to go down or offline permanently. 

## Installation:
The repository provides compiled binaries for Arm64 linux, AMD64 linux and AMD64 windows. 

## Usage:

To run the application you provide the following parameters on the command line:

<pre>
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

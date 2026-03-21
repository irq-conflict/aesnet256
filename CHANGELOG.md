```
ARG/Crypto Puzzle Compatibility List:
            * Hardwarez Bandcamp Cipher
            * PC Cipher [1/2]
            * PC Cipher [2/2] (OLD & NEW)
            * PC Bonus Page Cipher
            * Interrupt Request Bandcamp/Server Cipher
            * Interrupt Request Bonus Page Cipher
            * Floppy Disk Overdrive Bandcamp/Server Cipher
            * Floppy Disk Overdrive Bonus Page Cipher
            * Defrag Bandcamp/Server Cipher
            * Defrag Bonus Page Cipher
            * Internet Protocol Bonus Page Cipher
            * Internet Protocol [1/2]
            * Internet Protocol [2/2]

=================
    Version 0.31
        = BUGFIX: Fix uncaught error when IV length isn't equal to block size.
        = Migrate to Go Modules
        = Move compatibility list to top of CHANGELOG.md Additional confirmed support should
          be added as a single line in the CHANGELOG and added to the universal list up top.

=================
    Version 0.3
        = BUGFIX: Added functionality to remove errant line breaks in 
                  Decryption Text Input via Web UI
                  which facilitates direct copy/paste from mbrserver.com/bonus.
        = Properly formatted CSS stylesheet
        = Ensured compatibilty with the following puzzle components:
            * Hardwarez Bandcamp Cipher
            * PC Cipher [1/2]
            * PC Cipher [2/2] (OLD & NEW)
            * PC Bonus Page Cipher
            * Interrupt Request Bandcamp/Server Cipher
            * Interrupt Request Bonus Page Cipher
            * Floppy Disk Overdrive Bandcamp/Server Cipher
            * Floppy Disk Overdrive Bonus Page Cipher
            * Defrag Bandcamp/Server Cipher
            * Defrag Bonus Page Cipher
            * Internet Protocol Bonus Page Cipher
            * Internet Protocol [1/2]
            * Internet Protocol [2/2]

=================
    Version 0.21
        = Changed keysize from text box to drop down for ease of use.
        = Added missing error notification edge cases to AES decyrption / encryption
        = Added enhanced error messages to HTMX pages
        = BUGFIX: malformed GET requests and requests with 
                  recursive paths were not functioning correctly.
                  Added redirect to index.html from the default handler.
        = BUGFIX: StaticHandler was returning 400 instead of 404 when
                  page is not found.
        = Ensured compatibility with the following puzzles:
            * Hardwarez
            * Internet Protocol
            * Interrupt Request
            * Floppy Disk Overdrive
            * Defrag

=================
    Version 0.2
        = Added CHANGELOG.md
        = Implemented rudementary REST API style interface
		= Created basic HTMX page to use system
        = Ensured compatibility with the following puzzles:
            * Hardwarez
            * Internet Protocol
            * Interrupt Request
            * Floppy Disk Overdrive

=================
    Version 0.1
        = Initial AES Decryption and Encryption
        = Ensured compatibility with the following puzzles:
            * Hardwarez
            * Internet Protocol
            * Interrupt Request
        = Command line version completed
```

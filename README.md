### Vera

Allows applications to easily mount encrypted volumes using the VeraCrypt cli

#### Notes

1. A program using Vera needs extended permissions to mount volumes, mounted volumes may not be visible in
   the VeraCrypt GUI for this reason
2. When running the tests, depending on the operating system, error messages may appear stating that the folder could
   not be found. This is because some operating systems (e.g. Ubuntu) try to open newly mounted devices in file
   explorer. Since the tests dismount the mounted volumes after the tests are done, they are then no longer to be found. 
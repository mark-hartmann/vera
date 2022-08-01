## Overview

This package relies on the exported `ExecCommand` function and `Param` structs which are passed. Internally, the corresponding 
VeraCrypt command is assembled and then executed.  

`vera` provides some predefined functions that try to simplify common commands, such as `List()`, `MountSlot()` and 
`DismountSlot()`. For a complete list of all exported functions, see the list below:

* List
* DismountAll
* DismountSlot
* DismountVolume
* PropertiesSlot
* PropertiesVolume
* Installed
* MountPath
* MountSlot

Each of the above uses the `ExecCommand`, so the user may choose to implement some of these functions by 
themselves. However, this can quickly become a mess: 
```go
stdout, err := vera.ExecCommand(vera.Param{Name: "list", IsFlag: true})

if err == nil {
    // code to actually parse the stdout buffer and turn it into an actual list
}
```
The shorter version for the above code would be:
```go
mounts, err := vera.List() // the stdout buffer is already parsed and "mounts" contains a list of MountProperties
```

### MountProperties

The MountProperties struct is returned when calling `List`, `PropertiesSlot`, `PropertiesVolume`, `MountPath`, and `MountSlot`. It 
provides some details about the mounted volume. The structure is as follows:

```go
type MountProperties struct {
  Slot       uint8  // the slot the volume is mounted at
  Volume     string // the volume path
  MountPoint string // the directory the volume is mounted to
}
```

###### Slot

VeraCrypt uses "slots" to mount volumes to. These slots are limited and VeraCrypt only allow up to 64 mounted volumes. 
The slots range from **1** to **64**

###### Volume

Volume contains the absolute path of the mounted volume, even if a relative path was used to mount it

###### MountPoint

The mount directory of this volume. If not provided by the user, VeraCrypt / the operating system will choose one. On 
Linux Mint for example, the default mount directories are named "veracrypt" + the used slot, e.g. `/media/veracrypt1` or 
`/media/veracrypt3`

---

### Param

Param is a simple struct used for all kinds of arguments one might want to use:

```go
type Param struct {
   Name   string
   Value  string // Leaving the Value empty doesn't indicate the param is a flag, use IsFlag instead
   IsFlag bool   // IsFlag must be set to true for flags, e.g. --truecrypt or --version
}
```
###### Usage
```go
// create a flag
tc := vera.Param{Name: "truecrypt", IsFlag: true} // --truecrypt

// create a normal parameter
pwd := vera.Param{Name: "password", Value: "<your password>"} // --password="<your password>"

// create an argument
volume := vera.Param{Value: "./stuff.vc"} // "./stuff.vc"

// mount volume
_, err := vera.ExecCommand(tc, volume, pwd)
```

Depending on the application requirements, it may be helpful to set up frequently used parameters as variables:
```go
var NoFileSystem := vera.Param{Name: "filesystem", Value: "none"}

func mountNoFs() {
    _, err := vera.MountSlot("./container.vc", 1, "<password>", NoFileSystem)
    if err != nil { 
        // do something
    }
}
```

### Running Unit Tests
For the tests to work properly, a `mount directory must be created in the testdata directory to be used as a mount point.

---
#### Notes

1. A program using this package may need extended permissions to mount volumes
2. Mounted volumes may not be visible in the VeraCrypt GUI if the volumes have been mounted with different permissions
3. When running the tests, depending on the settings, error messages may appear stating that the directory could
   not be found. This is because some operating systems (e.g. Ubuntu) try to open newly mounted devices in file
   explorer. Since the tests dismount the mounted volumes after the tests are done, they are then no longer to be found.

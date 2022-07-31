package vera

import "fmt"

// trueCrypt tells VeraCrypt to mount the given volume in trueCrypt mode for backward compatability
var trueCrypt = Param{Name: "truecrypt", IsFlag: true}

// nonInteractive represents the "--non-interactive" flag. This is only for internal use, as direct communication is not
// supported (yet?)
var nonInteractive = Param{Name: "non-interactive", IsFlag: true}

// textOnly represents the "-t" / "--text" flag and is internally used to prevent the VeraCrypt GUI from taking over
var textOnly = Param{Name: "t", IsFlag: true}

// version is used to return the version of the currently installed VeraCrypt
var version = Param{Name: "version", IsFlag: true}

// dismount is used to unmount one or all mounted volumes
var dismount = Param{Name: "dismount", IsFlag: true}

// list is used to return a list of all currently mounted volumes
var list = Param{Name: "list", IsFlag: true}

// arg returns a configured Param instance to use as argument
func arg(value string) Param {
	return Param{
		Value: value,
	}
}

// Param is used for various functions to add additional arguments for the VeraCrypt CLI
type Param struct {
	Name   string
	Value  string // Leaving the Value empty doesn't indicate the param is a flag, use IsFlag instead
	IsFlag bool   // IsFlag must be set to true for flags, e.g. --truecrypt or --version
}

// String returns the string representation of a param
func (p Param) String() string {
	var param string
	if len(p.Name) < 3 {
		param = "-" + p.Name
	} else {
		param = "--" + p.Name
	}

	// a flag does not have a value, so we can return
	if p.IsFlag {
		return param
	}

	if len(p.Name) < 3 {
		return fmt.Sprintf("%s %s", param, p.Value)
	} else {
		return fmt.Sprintf("%s=%s", param, p.Value)
	}
}

package os

import (
	"fmt"
	"path/filepath"
)

// CmdPkgInstallBinary installs a binary in a system
func (os OS) CmdPkgInstallBinary(c Client, usr, name string, b []byte) error {
	switch os {
	case Linux:
		ssh, ok := (c).(ClientSSH)
		if !ok {
			return ErrUnsupportedClient
		}

		home, err := os.CmdFSHome(c, usr)
		if err != nil {
			return fmt.Errorf("error installing the binary: %v", err)
		}

		binDir := filepath.Join(home, ".bin")

		exists, err := os.CmdFSCheckDir(c, binDir)
		if err != nil {
			return fmt.Errorf("error installing the binary: %v", err)
		}
		if !exists {
			if err = os.CmdFSMkdir(c, binDir); err != nil {
				return fmt.Errorf("error creating the bin directory: %v", err)
			}

			if err = os.CmdFSAppendToFile(c, filepath.Join(home, ".profile"), []byte(`\n#Generated by DRLM\nexport PATH="$PATH:$HOME/.bin"\n`)); err != nil {
				return fmt.Errorf("error adding the bin dir to the PATH: %v", err)
			}
		}

		if err := ssh.WriteFile(filepath.Join(binDir, name), b); err != nil {
			return fmt.Errorf("error writting the binary: %v", err)
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}

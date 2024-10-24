package agent

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/mutagen-io/mutagen/pkg/filesystem"
	"github.com/mutagen-io/mutagen/pkg/mutagen"
	"github.com/mutagen-io/mutagen/pkg/platform"
)

const (
	// BaseName is the base name for agent executables (sans any
	// platform-specific suffix like ".exe").
	BaseName = "mutagen-agent"
)

// installPath computes and creates the parent directories of the path where the
// current executable should be installed if it is an agent binary with the
// current Mutagen version.
func installPath() (string, error) {
	// Compute (and create) the path to the agent parent directory.
	parent, err := filesystem.Mutagen(true, filesystem.MutagenAgentsDirectoryName, mutagen.Version)
	if err != nil {
		return "", fmt.Errorf("unable to compute parent directory: %w", err)
	}

	// Compute the target executable name.
	executableName := platform.ExecutableName(BaseName, runtime.GOOS)

	// Compute the installation path.
	return filepath.Join(parent, executableName), nil
}

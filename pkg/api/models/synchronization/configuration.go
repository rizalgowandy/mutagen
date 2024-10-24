package synchronization

import (
	"github.com/mutagen-io/mutagen/pkg/api/models/types"
	"github.com/mutagen-io/mutagen/pkg/filesystem"
	"github.com/mutagen-io/mutagen/pkg/filesystem/behavior"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	"github.com/mutagen-io/mutagen/pkg/synchronization/compression"
	"github.com/mutagen-io/mutagen/pkg/synchronization/core"
	"github.com/mutagen-io/mutagen/pkg/synchronization/core/ignore"
	"github.com/mutagen-io/mutagen/pkg/synchronization/hashing"
)

// Configuration represents synchronization session configuration.
type Configuration struct {
	// Mode specifies the default synchronization mode.
	Mode core.SynchronizationMode `json:"mode,omitempty" yaml:"mode" mapstructure:"mode"`
	// Hash specifies the hashing algorithm to use for content.
	Hash hashing.Algorithm `json:"hash,omitempty" yaml:"hash" mapstructure:"hash"`
	// MaximumEntryCount specifies the maximum number of filesystem entries
	// that endpoints will tolerate managing.
	MaximumEntryCount uint64 `json:"maxEntryCount,omitempty" yaml:"maxEntryCount" mapstructure:"maxEntryCount"`
	// MaximumStagingFileSize is the maximum (individual) file size that
	// endpoints will stage. It can be specified in human-friendly units.
	MaximumStagingFileSize types.ByteSize `json:"maxStagingFileSize,omitempty" yaml:"maxStagingFileSize" mapstructure:"maxStagingFileSize"`
	// ProbeMode specifies the filesystem probing mode.
	ProbeMode behavior.ProbeMode `json:"probeMode,omitempty" yaml:"probeMode" mapstructure:"probeMode"`
	// ScanMode specifies the filesystem scanning mode.
	ScanMode synchronization.ScanMode `json:"scanMode,omitempty" yaml:"scanMode" mapstructure:"scanMode"`
	// StageMode specifies the filesystem staging mode.
	StageMode synchronization.StageMode `json:"stageMode,omitempty" yaml:"stageMode" mapstructure:"stageMode"`
	// Ignore contains parameters related to synchronization ignore
	// specifications.
	Ignore struct {
		// Syntax specifies the ignore syntax and semantics.
		Syntax ignore.Syntax `json:"syntax,omitempty" yaml:"syntax" mapstructure:"syntax"`
		// Paths specifies the default list of ignore specifications.
		Paths []string `json:"paths,omitempty" yaml:"paths" mapstructure:"paths"`
		// VCS specifies the VCS ignore mode.
		VCS ignore.IgnoreVCSMode `json:"vcs,omitempty" yaml:"vcs" mapstructure:"vcs"`
	} `json:"ignore" yaml:"ignore" mapstructure:"ignore"`
	// Symlink contains parameters related to symbolic link handling.
	Symlink struct {
		// Mode specifies the symbolic link mode.
		Mode core.SymbolicLinkMode `json:"mode,omitempty" yaml:"mode" mapstructure:"mode"`
	} `json:"symlink" yaml:"symlink" mapstructure:"symlink"`
	// Watch contains parameters related to filesystem monitoring.
	Watch struct {
		// Mode specifies the file watching mode.
		Mode synchronization.WatchMode `json:"mode,omitempty" yaml:"mode" mapstructure:"mode"`
		// PollingInterval specifies the interval (in seconds) for poll-based
		// file monitoring. A value of 0 specifies that Mutagen's internal
		// default interval should be used.
		PollingInterval uint32 `json:"pollingInterval,omitempty" yaml:"pollingInterval" mapstructure:"pollingInterval"`
	} `json:"watch" yaml:"watch" mapstructure:"watch"`
	// Permissions contains parameters related to permission handling.
	Permissions struct {
		// Mode specifies the permissions mode.
		Mode core.PermissionsMode `json:"mode,omitempty" yaml:"mode" mapstructure:"mode"`
		// DefaultFileMode specifies the default permission mode to use for new
		// files in "portable" permission propagation mode.
		DefaultFileMode filesystem.Mode `json:"defaultFileMode,omitempty" yaml:"defaultFileMode" mapstructure:"defaultFileMode"`
		// DefaultDirectoryMode specifies the default permission mode to use for
		// new files in "portable" permission propagation mode.
		DefaultDirectoryMode filesystem.Mode `json:"defaultDirectoryMode,omitempty" yaml:"defaultDirectoryMode" mapstructure:"defaultDirectoryMode"`
		// DefaultOwner specifies the default owner identifier to use when
		// setting ownership of new files and directories in "portable"
		// permission propagation mode.
		DefaultOwner string `json:"defaultOwner,omitempty" yaml:"defaultOwner" mapstructure:"defaultOwner"`
		// DefaultGroup specifies the default group identifier to use when
		// setting ownership of new files and directories in "portable"
		// permission propagation mode.
		DefaultGroup string `json:"defaultGroup,omitempty" yaml:"defaultGroup" mapstructure:"defaultGroup"`
	} `json:"permissions" yaml:"permissions" mapstructure:"permissions"`
	// Compression contains parameters related to compression.
	Compression struct {
		// Algorithm specifies the compression algorithm.
		Algorithm compression.Algorithm `json:"algorithm,omitempty" yaml:"algorithm" mapstructure:"algorithm"`
	} `json:"compression" yaml:"compression" mapstructure:"compression"`
}

// loadFromInternal sets a configuration to match an internal
// Protocol Buffers representation. The configuration must be valid.
func (c *Configuration) loadFromInternal(configuration *synchronization.Configuration) {
	// Propagate top-level configuration.
	c.Mode = configuration.SynchronizationMode
	c.Hash = configuration.HashingAlgorithm
	c.MaximumEntryCount = configuration.MaximumEntryCount
	c.MaximumStagingFileSize = types.ByteSize(configuration.MaximumStagingFileSize)
	c.ProbeMode = configuration.ProbeMode
	c.ScanMode = configuration.ScanMode
	c.StageMode = configuration.StageMode

	// Propagate ignore configuration.
	c.Ignore.Syntax = configuration.IgnoreSyntax
	c.Ignore.Paths = make([]string, 0, len(configuration.DefaultIgnores)+len(configuration.Ignores))
	c.Ignore.Paths = append(c.Ignore.Paths, configuration.DefaultIgnores...)
	c.Ignore.Paths = append(c.Ignore.Paths, configuration.Ignores...)
	c.Ignore.VCS = configuration.IgnoreVCSMode

	// Propagate symbolic link configuration.
	c.Symlink.Mode = configuration.SymbolicLinkMode

	// Propagate watch configuration.
	c.Watch.Mode = configuration.WatchMode
	c.Watch.PollingInterval = configuration.WatchPollingInterval

	// Propagate permission configuration.
	c.Permissions.Mode = configuration.PermissionsMode
	c.Permissions.DefaultFileMode = filesystem.Mode(configuration.DefaultFileMode)
	c.Permissions.DefaultDirectoryMode = filesystem.Mode(configuration.DefaultDirectoryMode)
	c.Permissions.DefaultOwner = configuration.DefaultOwner
	c.Permissions.DefaultGroup = configuration.DefaultGroup

	// Propagate compression configuration.
	c.Compression.Algorithm = configuration.CompressionAlgorithm
}

// ToInternal converts a public configuration representation to an internal
// Protocol Buffers session configuration. It does not validate the resulting
// configuration.
func (c *Configuration) ToInternal() *synchronization.Configuration {
	return &synchronization.Configuration{
		SynchronizationMode:    c.Mode,
		HashingAlgorithm:       c.Hash,
		MaximumEntryCount:      c.MaximumEntryCount,
		MaximumStagingFileSize: uint64(c.MaximumStagingFileSize),
		ProbeMode:              c.ProbeMode,
		ScanMode:               c.ScanMode,
		StageMode:              c.StageMode,
		SymbolicLinkMode:       c.Symlink.Mode,
		WatchMode:              c.Watch.Mode,
		WatchPollingInterval:   c.Watch.PollingInterval,
		IgnoreSyntax:           c.Ignore.Syntax,
		Ignores:                c.Ignore.Paths,
		IgnoreVCSMode:          c.Ignore.VCS,
		PermissionsMode:        c.Permissions.Mode,
		DefaultFileMode:        uint32(c.Permissions.DefaultFileMode),
		DefaultDirectoryMode:   uint32(c.Permissions.DefaultDirectoryMode),
		DefaultOwner:           c.Permissions.DefaultOwner,
		DefaultGroup:           c.Permissions.DefaultGroup,
		CompressionAlgorithm:   c.Compression.Algorithm,
	}
}

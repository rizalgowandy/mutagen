package url

import (
	"errors"
	"fmt"
	"math"
	"path/filepath"

	"github.com/mutagen-io/mutagen/pkg/comparison"
	"github.com/mutagen-io/mutagen/pkg/extension"
	"github.com/mutagen-io/mutagen/pkg/url/forwarding"
)

// Supported returns whether or not a URL kind is supported.
func (k Kind) Supported() bool {
	switch k {
	case Kind_Synchronization:
		return true
	case Kind_Forwarding:
		return true
	default:
		return false
	}
}

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (p Protocol) MarshalText() ([]byte, error) {
	var result string
	switch p {
	case Protocol_Local:
		result = "local"
	case Protocol_SSH:
		result = "ssh"
	case Protocol_Docker:
		result = "docker"
	default:
		result = "unknown"
	}
	return []byte(result), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (p *Protocol) UnmarshalText(textBytes []byte) error {
	// Convert the bytes to a string.
	text := string(textBytes)

	// Convert to a protocol.
	switch text {
	case "local":
		*p = Protocol_Local
	case "ssh":
		*p = Protocol_SSH
	case "docker":
		*p = Protocol_Docker
	default:
		return fmt.Errorf("unknown protocol specification: %s", text)
	}

	// Success.
	return nil
}

// EnsureValid ensures that URL's invariants are respected.
func (u *URL) EnsureValid() error {
	// Ensure that the URL is non-nil.
	if u == nil {
		return errors.New("nil URL")
	}

	// Ensure that the kind is supported.
	if !u.Kind.Supported() {
		return errors.New("unsupported URL kind")
	}

	// Validate the User, Host, Port, and Environment components based on
	// protocol.
	if u.Protocol == Protocol_Local {
		if u.User != "" {
			return errors.New("local URL with non-empty username")
		} else if u.Host != "" {
			return errors.New("local URL with non-empty hostname")
		} else if u.Port != 0 {
			return errors.New("local URL with non-zero port")
		} else if len(u.Environment) != 0 {
			return errors.New("local URL with environment variables")
		} else if len(u.Parameters) != 0 {
			return errors.New("local URL with parameters")
		}
	} else if u.Protocol == Protocol_SSH {
		if u.Host == "" {
			return errors.New("SSH URL with empty hostname")
		} else if u.Port > math.MaxUint16 {
			return errors.New("SSH URL with invalid port")
		} else if len(u.Environment) != 0 {
			return errors.New("SSH URL with environment variables")
		}
	} else if u.Protocol == Protocol_Docker {
		// In the case of Docker, we intentionally avoid validating environment
		// variables since the values used could change over time. Since we
		// default to empty values for unspecified environment variables, this
		// works out fine, at least so long as Docker continues to treat empty
		// environment variables the same as unspecified ones.
		if u.Host == "" {
			return errors.New("Docker URL with empty container identifier")
		} else if u.Port != 0 {
			return errors.New("Docker URL with non-zero port")
		}
	} else {
		return errors.New("unknown or unsupported protocol")
	}

	// Validate the path component depending on the URL kind.
	if u.Kind == Kind_Synchronization {
		// Ensure the path is non-empty.
		if u.Path == "" {
			return errors.New("empty path")
		}

		// If this is a local URL, then ensure that the path is absolute.
		//
		// HACK: The Mutagen Extension for Docker Desktop needs to avoid this
		// check because of its internal faux-local URLs. In particular, Windows
		// absolute paths will appear as non-absolute in the extension's
		// Linux-based backend container. Technically we only need to avoid this
		// check for alpha URLs, but since we control all of the extension's
		// URLs, and since we don't know which endpoint this URL is targeting,
		// we just disable it entirely when running in the extension.
		if !extension.EnvironmentIsExtension() {
			if u.Protocol == Protocol_Local && !filepath.IsAbs(u.Path) {
				return errors.New("local URL with relative path")
			}
		}

		// If this is a Docker URL, we can actually do a bit of additional
		// validation.
		if u.Protocol == Protocol_Docker {
			if !(u.Path[0] == '/' || u.Path[0] == '~' || isWindowsPath(u.Path)) {
				return errors.New("incorrect first path character")
			}
		}
	} else if u.Kind == Kind_Forwarding {
		// Parse the forwarding endpoint URL to ensure that it's valid.
		protocol, address, err := forwarding.Parse(u.Path)
		if err != nil {
			return fmt.Errorf("invalid forwarding endpoint URL: %w", err)
		}

		// If this is a local URL and represents a Unix domain socket endpoint,
		// then ensure that the socket path is absolute.
		if u.Protocol == Protocol_Local && protocol == "unix" && !filepath.IsAbs(address) {
			return errors.New("local Unix domain socket URL with relative path")
		}

		// TODO: It would be nice to perform some sort of validation on Windows
		// named pipe addresses, but there's not much we can do because the
		// allowed formats vary between source and destination endpoints (so
		// we'd have to weave that information through this function). The only
		// difference is that the ServerName component (see the link below) must
		// be "." for source endpoints but can also name a remote server in the
		// case of destination endpoints. But that's not really the biggest
		// issue. The problem is that the name specification is kind of vague.
		// It says that the PipeName component (again, see the link below) "can
		// include any character other than a backslash, including numbers and
		// special characters", but it doesn't mention whitespace characters
		// (for example a newline character), which, as far as I'm aware, are
		// not allowed. It also limits the "entire pipe name string" to 256
		// characters, but it's not clear if this refers to the PipeName
		// component or the entire address. Finding an appropriate matcher for
		// possible server names is also an uphill battle. This might be
		// specified in the UNC specification. In the end though, we're probably
		// just better off letting the OS decide what to accept and simply
		// returning its errors directly. For further reading, see:
		// https://docs.microsoft.com/en-us/windows/win32/ipc/pipe-names
	}

	// Success.
	return nil
}

// Equal returns whether or not the URL is equivalent to another. The result of
// this method is only valid if both URLs are valid.
func (u *URL) Equal(other *URL) bool {
	// Ensure that both are non-nil.
	if u == nil || other == nil {
		return false
	}

	// Perform an equivalence check.
	return u.Kind == other.Kind &&
		u.Protocol == other.Protocol &&
		u.User == other.User &&
		u.Host == other.Host &&
		u.Port == other.Port &&
		u.Path == other.Path &&
		comparison.StringMapsEqual(u.Environment, other.Environment) &&
		comparison.StringMapsEqual(u.Parameters, other.Parameters)
}

package datatypes

type Directory struct {
	Title, Desc, BasePath string
	Platform              Platform
}

// enum type creation w/ companion pieces
type Platform int

const (
	GCP Platform = iota // iota keyword instantiates all the contants w/ 0 to N values automatically, i.e. 0,1,2.  This is a unique integer, it prevents conflicts with other constants from other types .
	WINDOWS
	UNIX
	UNKNOWN
)

// companion function to Platform type, now if we try to print our platform, it will use this function to print out the result instead of printing the integer.
func (s Platform) String() string {
	switch s {
	case GCP:
		return "GCP"
	case WINDOWS:
		return "WINDOWS"
	case UNIX:
		return "UNIX"
	default:
		return "Unknown"
	}
}

func GetPlatformType(platformString string) Platform {
	var platform Platform
	switch platformString {
	case "WINDOWS":
		platform = WINDOWS
	case "UNIX":
		platform = UNIX
	case "GCP":
		platform = GCP
	default:
		platform = UNKNOWN
	}

	return platform
}

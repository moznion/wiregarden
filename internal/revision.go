package internal

var (
	Version  string
	Revision string
)

func ShortRevision() string {
	if len(Revision) < 7 {
		return ""
	}
	return Revision[:7]
}

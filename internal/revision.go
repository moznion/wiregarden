package internal

var (
	Revision string
)

func ShortRevision() string {
	if len(Revision) < 7 {
		return ""
	}
	return Revision[:7]
}

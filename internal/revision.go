package internal

var (
	Revision string
)

func ShortRevision() string {
	return Revision[:7]
}

package version

var (
	mainVersion = "1.0"
	timeVersion = "none"
)

func Version() string {
	return mainVersion + "." + timeVersion
}
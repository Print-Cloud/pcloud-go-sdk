package api

type Api interface {
	GetPostMsg(version string, key string) (msg string, err error)

	GetGetMsg(version string, key string) (msg string, err error)
}

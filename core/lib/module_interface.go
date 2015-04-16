package lib

type IModule interface {
	SetupApiEndpoints()
	SetupMenuItems()
	SetupAdminAssets()
	SetupSections()
}

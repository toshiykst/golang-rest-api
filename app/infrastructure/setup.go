package infrastructure

// Setup infrastructure
func Setup() {
	_ = NewDB()
	SetupRouter()
}

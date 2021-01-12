package infrastructure

// Setup infrastructure
func Setup() {
	_ = NewRepository()
	SetupRouter()
}

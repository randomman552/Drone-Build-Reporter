package reporters

type Reporter interface {
	Report(DroneContext)
}

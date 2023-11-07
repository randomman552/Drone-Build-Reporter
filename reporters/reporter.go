package reporters

type Reporter interface {
	Report(ReporterContext)
}

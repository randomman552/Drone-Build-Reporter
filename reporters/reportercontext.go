package reporters

type ReporterContext struct {
	Message string
}

type DroneContext struct {
	Build        BuildContext
	Commit       CommitContext
	Repo         RepoContext
	FailedStages []string
	FailedSteps  []string
}

type BuildContext struct {
	Action   string
	Created  string
	Event    string
	Finished string
	Link     string
	Number   string
	Parent   string
	Started  string
	Status   string
	Trigger  string
}

type CommitContext struct {
	Commit       string
	After        string
	Author       string
	AuthorAvatar string
	AuthorEmail  string
	Before       string
	Link         string
	Message      string
	Ref          string
	SHA          string
}

type RepoContext struct {
	Branch     string
	Link       string
	Name       string
	Namespace  string
	Owner      string
	Private    bool
	Type       string
	Visibility string
}

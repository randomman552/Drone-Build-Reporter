package types

type DroneContext struct {
	Build        BuildContext
	Commit       CommitContext
	Repo         RepoContext
	FailedStages []string
	FailedSteps  []string
}

type BuildContext struct {
	Action   string
	Created  uint
	Started  uint
	Finished uint
	Event    string
	Link     string
	Number   int
	Parent   int
	Status   string
	Trigger  string
}

type CommitContext struct {
	Hash         string
	Before       string
	After        string
	Author       string
	AuthorAvatar string
	AuthorEmail  string
	AuthorName   string
	Link         string
	Message      string
	Ref          string
}

type RepoContext struct {
	Branch     string
	Link       string
	Name       string
	Namespace  string
	Owner      string
	Private    bool
	Visibility string
}
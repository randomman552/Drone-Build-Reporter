package main

import (
	"log"
	"os"
	"reporter/types"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Build Reporter Plugin"
	app.Usage = "Build Reporter Plugin"
	app.Action = run
	app.Flags = []cli.Flag{
		// Config parameters
		cli.StringFlag{
			Name:   "config.template_directory",
			EnvVar: "PLUGIN_TEMPLATE_DIRECTORY",
			Value:  "./templates",
		},
		cli.StringFlag{
			Name:   "config.gotify.token",
			EnvVar: "PLUGIN_GOTIFY_TOKEN",
		},
		cli.StringFlag{
			Name:   "config.gotify.url",
			EnvVar: "PLUGIN_GOTIFY_URL",
		},
		// Build context parameters
		cli.StringFlag{
			Name:   "build.action",
			EnvVar: "DRONE_BUILD_ACTION",
		},
		cli.Int64Flag{
			Name:   "build.created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.Int64Flag{
			Name:   "build.started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.StringFlag{
			Name:   "build.event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.StringFlag{
			Name:   "build.link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.IntFlag{
			Name:   "build.number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.IntFlag{
			Name:   "build.parent",
			EnvVar: "DRONE_BUILD_PARENT",
		},
		cli.StringFlag{
			Name:   "build.status",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.trigger",
			EnvVar: "DRONE_BUILD_TRIGGER",
		},
		// Commit context parameters
		cli.StringFlag{
			Name:   "commit.hash",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.before",
			EnvVar: "DRONE_COMMIT_BEFORE",
		},
		cli.StringFlag{
			Name:   "commit.after",
			EnvVar: "DRONE_COMMIT_AFTER",
		},
		cli.StringFlag{
			Name:   "commit.author",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.author.name",
			EnvVar: "DRONE_COMMIT_AUTHOR_NAME",
		},
		cli.StringFlag{
			Name:   "commit.link",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		// Repo context parameters
		cli.StringFlag{
			Name:   "repo.branch",
			EnvVar: "DRONE_REPO_BRANCH",
		},
		cli.StringFlag{
			Name:   "repo.link",
			EnvVar: "DRONE_REPO_LINK",
		},
		cli.StringFlag{
			Name:   "repo.name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "repo.namespace",
			EnvVar: "DRONE_REPO_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.BoolFlag{
			Name:   "repo.private",
			EnvVar: "DRONE_REPO_PRIVATE",
		},
		cli.StringFlag{
			Name:   "repo.visibility",
			EnvVar: "DRONE_REPO_VISIBILITY",
		},
		// Failed steps and stages
		cli.StringSliceFlag{
			Name:   "failed_stages",
			EnvVar: "DRONE_FAILED_STAGES",
		},
		cli.StringSliceFlag{
			Name:   "failed_steps",
			EnvVar: "DRONE_FAILED_STEPS",
		},
	}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) {
	buildCreated := time.Unix(c.Int64("build.created"), 0)
	buildStarted := time.Unix(c.Int64("build.started"), 0)
	buildFinished := time.Unix(c.Int64("build.finished"), 0)
	buildDuration := buildFinished.Sub(buildCreated)

	plugin := Plugin{
		Config: types.Config{
			TemplateDirectory: c.String("config.template_directory"),
			GotifyToken:       c.String("config.gotify.token"),
			GotifyUrl:         c.String("config.gotify.url"),
		},
		Context: types.DroneContext{
			Build: types.BuildContext{
				Action:   c.String("build.action"),
				Created:  &buildCreated,
				Started:  &buildStarted,
				Finished: &buildFinished,
				Duration: &buildDuration,
				Event:    c.String("build.event"),
				Link:     c.String("build.link"),
				Number:   c.Int("build.number"),
				Parent:   c.Int("build.parent"),
				Status:   c.String("build.status"),
				Trigger:  c.String("build.trigger"),
			},
			Commit: types.CommitContext{
				Hash:         c.String("commit.hash"),
				Before:       c.String("commit.before"),
				After:        c.String("commit.after"),
				Author:       c.String("commit.author"),
				AuthorAvatar: c.String("commit.author.avatar"),
				AuthorEmail:  c.String("commit.author.email"),
				AuthorName:   c.String("commit.author.name"),
				Link:         c.String("commit.link"),
				Message:      c.String("commit.message"),
				Ref:          c.String("commit.ref"),
			},
			Repo: types.RepoContext{
				Branch:     c.String("repo.branch"),
				Link:       c.String("repo.link"),
				Name:       c.String("repo.name"),
				Namespace:  c.String("repo.namespace"),
				Owner:      c.String("repo.owner"),
				Private:    c.Bool("repo.private"),
				Visibility: c.String("repo.visibility"),
			},
			FailedStages: c.StringSlice("failed_stages"),
			FailedSteps:  c.StringSlice("failed_steps"),
		},
	}

	plugin.Run()
}

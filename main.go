package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/google/go-github/github"
	"github.com/joeshaw/envdecode"
	"golang.org/x/oauth2"
)

type meta struct {
	team   string
	org    string
	client *github.Client
	ctx    context.Context
	opt    *github.RepositoryListByOrgOptions
}

type config struct {
	GithubToken string `env:"GITHUB_TOKEN,required"`
	Team        string `env:"TEAM,required"`
	Org         string `env:"ORG,required"`
}

func main() {
	var cfg config
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatalf("%s\n", err.Error())
		os.Exit(1)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	pageOption := &github.ListOptions{
		PerPage: 100,
	}
	opt := &github.RepositoryListByOrgOptions{
		Type:        "private",
		ListOptions: *pageOption,
	}

	m := meta{
		team:   cfg.Team,
		org:    cfg.Org,
		client: client,
		ctx:    ctx,
		opt:    opt,
	}

	id, _ := getTeamID(m)
	if id == 0 {
		log.Errorf("Team %s does not exist", m.team)
		os.Exit(1)
	}
	addTeamRepo(m, id, listAllRepos()(m))
}

func filterByTeam(teams []*github.Team, name string) int64 {
	if len(teams) >= 0 {
		for _, t := range teams {
			if *t.Name == name {
				return t.GetID()
			}
		}
	}
	return 0
}

func getTeamID(m meta) (int64, error) {
	for {
		teams, _, err := m.client.Organizations.ListTeams(m.ctx, m.org, &m.opt.ListOptions)
		if err != nil {
			return 0, fmt.Errorf("Get team id %s error: %v", m.team, err.Error())
		}
		id := filterByTeam(teams, m.team)
		if id == 0 {
			continue
		}
		return id, nil
	}
}

func listAllRepos() func(meta) []*github.Repository {
	return func(m meta) []*github.Repository {
		var allRepos []*github.Repository
		for {

			repos, resp, _ := m.client.Repositories.ListByOrg(m.ctx, m.org, m.opt)
			allRepos = append(allRepos, repos...)
			if resp.NextPage == 0 {
				break
			}
			m.opt.Page = resp.NextPage
		}
		return allRepos

	}
}

func addTeamRepo(m meta, id int64, repos []*github.Repository) {
	teamoption := &github.OrganizationAddTeamRepoOptions{
		Permission: "pull",
	}
	for _, repo := range repos {
		resp, err := m.client.Organizations.AddTeamRepo(m.ctx, id, m.org, *repo.Name, teamoption)
		if err != nil {
			log.Errorf("add repo to team %s error: %v", *repo.Name, err.Error())
			os.Exit(1)
		}
		if resp.StatusCode == 204 {
			log.Infof("successfully add %s to team id %d", *repo.Name, id)
		}
	}
}

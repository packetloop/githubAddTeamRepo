[![CircleCI](https://circleci.com/gh/packetloop/githubAddTeamRepo.svg?style=svg)](https://circleci.com/gh/packetloop/githubAddTeamRepo)
[![GitHub release](https://img.shields.io/github/release/packetloop/githubAddTeamRepo.svg)](https://github.com/packetloop/githubAddTeamRepo/releases/)
[![All Contributors](https://img.shields.io/github/contributors/packetloop/githubAddTeamRepo.svg?longCache=true&style=flat-square&colorB=orange&label=all%20contributors)](#contributors)
[![Github All Releases](https://img.shields.io/github/downloads/packetloop/githubAddTeamRepo/total.svg)]()


# addTeamRepo

A script to add Github Repos READ-ONLY access to a team.

## Usage:

Download this provider, pick a version you'd like from releases from
[Binary Releases](https://github.com/packetloop/githubAddTeamRepo/releases)

```bash
curl -L \$
   https://github.com/packetloop/githubAddTeamRepo/releases/download/v0.1.0/githubAddTeamRepo_v0.1.0_Darwin_x86_64 \
   -o /tmp/githubAddTeamRepo && \
   chmod +x ~/.terraform.d/plugins/githubAddTeamRepo

ORG=<GITHUB_ORG> TEAM=<GITHUB_TEAM> GITHUB_TOKEN=<GITHUB_TOKEN> /tmp/githubAddTeamRepo
```
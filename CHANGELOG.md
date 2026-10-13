# Changelog

Version v1.5.0
--------------------
Released: 2026-10-13

🚀 Features:
- Determine version from tag
- Use goreleaser to build rpm package
- Add framework for localisation
- Add german translation

🐛 Bug Fixes:
- Ensure version information fits into dialog
- *(deps)* Update gomod

⚙️ Miscellaneous Tasks:
- Enforce conventional commits
- Add new changelog workflow
- *(builder)* Update image ghcr.io/heathcliff26/go-fyne-ci to v202610090432

Version v1.4.1
--------------------
Released: 2026-09-27

🚀 Features:
- *(container)* [**breaking**] Update image ghcr.io/heathcliff26/go-fyne-ci to v202609060502
- *(container)* [**breaking**] Update image ghcr.io/heathcliff26/go-fyne-ci to v202609241228

💼 Other:
- Add support for alternative steam library location
- Fix coverprofiles permissions
- Adapt border color to new fyne design

Version v1.4.0
--------------------
Released: 2026-08-27

🚀 Features:
- *(container)* [**breaking**] Update image ghcr.io/heathcliff26/go-fyne-ci to v202608201646
- *(container)* [**breaking**] Update image ghcr.io/heathcliff26/go-fyne-ci to v202608261637
- *(container)* [**breaking**] Update image ghcr.io/heathcliff26/go-fyne-ci to v202608261947

🐛 Bug Fixes:
- *(deps)* Update gomod
- *(deps)* Update gomod

💼 Other:
- Move renovate and Link Check to tekton
- Remove obsolete Renovate status badge
- Switch to containerized script
- *(deps)* Bump golang.org/x/net from 0.50.0 to 0.55.0
- *(deps)* Bump golang.org/x/image from 0.38.0 to 0.41.0

Version v1.3.9
--------------------
Released: 2026-06-01

🐛 Bug Fixes:
- *(deps)* Update gomod

Version v1.3.8
--------------------
Released: 2026-05-01

🐛 Bug Fixes:
- *(deps)* Update gomod

💼 Other:
- *(deps)* Bump golang.org/x/image from 0.24.0 to 0.38.0

Version v1.3.7
--------------------
Released: 2026-03-01

🐛 Bug Fixes:
- *(deps)* Update gomod

Version v1.3.6
--------------------
Released: 2026-02-01

🐛 Bug Fixes:
- *(deps)* Update gomod

💼 Other:
- Remove unneeded prepare step in renovate workflow

Version v1.3.5
--------------------
Released: 2026-01-01

🐛 Bug Fixes:
- *(deps)* Update module github.com/heathcliff26/godialog to v0.1.8

Version v1.3.4
--------------------
Released: 2025-12-01

🐛 Bug Fixes:
- *(deps)* Update gomod

💼 Other:
- Enable signing artifacts in release workflow
- Consolidate targets after F41 moved to go1.24

Version v1.3.3
--------------------
Released: 2025-11-01

🐛 Bug Fixes:
- *(deps)* Update gomod

💼 Other:
- Add Fedora 43 and Rawhide targets
- Allow customising container image via env variable

Version v1.3.2
--------------------
Released: 2025-09-30

🐛 Bug Fixes:
- *(deps)* Update gomod

💼 Other:
- Add commit when archiving git

Version v1.3.1
--------------------
Released: 2025-08-01

🐛 Bug Fixes:
- *(deps)* Update gomod

💼 Other:
- Additionally force-reload icons
- Fully install app on system

Version v1.3.0
--------------------
Released: 2025-05-12

🐛 Bug Fixes:
- *(deps)* Update module fyne.io/fyne/v2 to v2.6.1

💼 Other:
- Add help target
- Do not wait for unit-tests for running build jobs
- Add gosec to improve code quality
- Move target description to above them
- Fix realease job permissions

Version v1.2.2
--------------------
Released: 2025-04-12

💼 Other:
- Update fyne to v2.6

Version v1.2.1
--------------------
Released: 2025-04-01

🐛 Bug Fixes:
- *(deps)* Update module fyne.io/fyne/v2 to v2.5.5

Version v1.2.0
--------------------
Released: 2025-02-16

💼 Other:
- Use windows style newlines when saving
- Increase infra-space support to version 1.50.448
- Simplify call to build script
- Fix process not exiting when called from cli
- Implement light mode
- Update screenshots with latest version

Version v1.1.4
--------------------
Released: 2025-02-01

🐛 Bug Fixes:
- *(deps)* Update module fyne.io/fyne/v2 to v2.5.4

💼 Other:
- Upload coverprofile to coveralls
- Fix missing explicit permissions for workflow runs
- Add coverpkg flag

Version v1.1.3
--------------------
Released: 2025-01-01

🐛 Bug Fixes:
- *(deps)* Update module github.com/stretchr/testify to v1.10.0
- *(deps)* Update module fyne.io/fyne/v2 to v2.5.3

⚙️ Miscellaneous Tasks:
- Run unit-tests on windows and linux

💼 Other:
- Bump go version to 1.23
- Update go version to 1.N.P notation
- Add fmt and validate commands
- Add validate step to ci job
- Delete intermediary artifact
- Add script for cleaning working directory
- Fix unit-test on windows
- Improve screenshots
- Use build tags for deciding default path during compile
- Add multiple location on linux for save folder

Version v1.1.2
--------------------
Released: 2024-11-01

🐛 Bug Fixes:
- *(deps)* Update module fyne.io/fyne/v2 to v2.5.1
- *(deps)* Update module fyne.io/fyne/v2 to v2.5.2

💼 Other:
- Fix max lvl habitats endlessly updating
- Fix breaking electricity by adding workers to habitats

Version v1.1.1
--------------------
Released: 2024-08-01

🚀 Features:
- *(github-action)* Update actions/create-github-app-token action to v1.7.0
- *(github-action)* Update actions/upload-artifact action to v4.3.0
- *(github-action)* Update actions/create-github-app-token action to v1.8.0
- *(github-action)* Update ncipollo/release-action action to v1.14.0
- *(github-action)* Update actions/create-github-app-token action to v1.9.0
- *(github-action)* Update actions/create-github-app-token action to v1.10.0

🐛 Bug Fixes:
- *(github-action)* Update actions/download-artifact action to v4.1.2
- *(github-action)* Update actions/upload-artifact action to v4.3.1
- *(deps)* Update module fyne.io/fyne/v2 to v2.4.4
- *(github-action)* Update actions/create-github-app-token action to v1.8.1
- *(github-action)* Update actions/cache action to v4.0.1
- *(github-action)* Update actions/download-artifact action to v4.1.3
- *(deps)* Update module github.com/stretchr/testify to v1.9.0
- *(github-action)* Update actions/create-github-app-token action to v1.8.2
- *(github-action)* Update actions/download-artifact action to v4.1.4
- *(github-action)* Update actions/cache action to v4.0.2
- *(github-action)* Update actions/create-github-app-token action to v1.9.1
- *(github-action)* Update actions/create-github-app-token action to v1.9.3
- *(github-action)* Update stefanzweifel/git-auto-commit-action action to v5.0.1
- *(github-action)* Update actions/upload-artifact action to v4.3.2
- *(github-action)* Update actions/download-artifact action to v4.1.5
- *(deps)* Update module fyne.io/fyne/v2 to v2.4.5
- *(github-action)* Update actions/checkout action to v4.1.4
- *(github-action)* Update actions/upload-artifact action to v4.3.3
- *(github-action)* Update actions/download-artifact action to v4.1.7
- *(github-action)* Update actions/setup-go action to v5.0.1
- *(github-action)* Update actions/checkout action to v4.1.5
- *(github-action)* Update actions/checkout action to v4.1.6
- *(github-action)* Update actions/create-github-app-token action to v1.10.1
- *(github-action)* Update actions/checkout action to v4.1.7
- *(github-action)* Update actions/create-github-app-token action to v1.10.2
- *(github-action)* Update actions/create-github-app-token action to v1.10.3
- *(github-action)* Update actions/upload-artifact action to v4.3.4
- *(github-action)* Update actions/download-artifact action to v4.1.8
- *(github-action)* Update actions/setup-go action to v5.0.2
- *(deps)* Update module fyne.io/fyne/v2 to v2.5.0

💼 Other:
- Use new ci repo for workflows where possible
- Fix go coverprofile report workflow
- Switch gh actions automerge type to pr
- Mount cache into go-fyne-ci container to speed up operations
- Use cache action to decrease test and build time
- Use custom fyne-cross image
- Run Renovate less frequently
- Move all workflows to CI repository

Version v1.0.1
--------------------
Released: 2024-01-13

💼 Other:
- Use prebuilt container image for lint and unit-tests
- Set Renovate automerge type to PR
- Add button next to resources for setting value to 1 mio.

Version v0.2.0
--------------------
Released: 2024-01-07

💼 Other:
- Fix trigger for testcover report
- Fix build job
- Fix release workflow

Version v0.1.2
--------------------
Released: 2024-01-04

💼 Other:
- Maintain Workflows
- Enable Renovate
- Always run on changes to go files
- Compress release artifacts

Version v0.1.1
--------------------
Released: 2024-01-02

💼 Other:
- Add initial ci and helper scripts

Version v0.1.0
--------------------
Released: 2023-05-02


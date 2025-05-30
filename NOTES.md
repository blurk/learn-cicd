By default, a step "succeeds" if it exits with a status code of 0 and "fails" if it exits with a status code other than 0.

Every (good) CLI tool that I'm aware of follows the convention of exit code 0 = pass, anything else = fail. For example, if a test case fails, go test will exit with a status code of 1.

### Workflow
A workflow is triggered when an event occurs in your GitHub repository. For example, we'll trigger our "tests" workflow when we open a pull request into the main branch.

In our case, the ci.yml file contains a single workflow called "ci", but we could have named it anything.

### Jobs
A workflow is made up of one or more jobs. A job is a set of steps that run on the same runner, which is just a virtual machine that runs your job on GitHub's servers.

### Steps
A job is made up of one or more steps. A step is a single task that can run commands, a script, or an action. For example, the steps of a job might include:
- Checking out the code
- Installing dependencies
- Running tests

In our case, the "Tests" job contains 3 steps:
- Check out the code
- Set up Go
- Force failure of the CI job

## Notables keyword
- The `uses` key specifies the action to use
- The `with` key specifies the inputs to the action.
- The 'run' key refers to a custom command

### A good CI pipeline typically includes:

- Unit tests
- Integration tests
- Styling checks
- Linting checks
- Security checks
- Any other kind of automated test

## Code Coverage

Code coverage is a measure of how much of your code is being tested. It's a controversial metric, but I'll try to provide a balanced take... granted I'm not without my own biases.

```python
code_coverage = (lines_covered / total_lines) * 100
```

If you have 1000 lines of code in your project, and you have tests that cover the logic in 500 of those lines, then you have 50% code coverage.

Regardless, the principle is the same everywhere. At its core, CI/CD enables us so that:

- When PRs are opened, run static analysis and tests
- When PRs are merged, build and deploy the app automatically

## What Makes a “Good” CI/CD Pipeline?
- Deterministic builds. The same code should always produce the same build.
- Fast builds. The faster the better. This makes getting bug fixes and new features out to users faster.
- Portable. This is why I love when the majority of a CI/CD pipeline is just bash scripts. It's easy to run locally, and it's easy to run on any CI/CD platform.
- Fully automated. The fewer manual steps, the better. It's really annoying to manually run database migrations and click buttons. It's also error-prone.

Recommend running DB migration after the Docker image is built, but before the deployment of the image. That accomplishes two things:

1. We won't run the migration if there is a problem building the image.
2. The migration will be live before the new code is deployed.

## Recap
- Set up a continuous integration pipeline with GitHub Actions that ensures new PRs pass certain checks before they are merged to main:
	- Unit tests pass
	- Formatting checks pass
	- Linting checks pass
	- Security checks pass
	- You configured a cloud-based SQLite database hosted on Turso
- Set up a continuous deployment pipeline with GitHub Actions that does the following whenever changes are merged into main:
	- Builds a new server binary
	- Builds a new Docker image for the server
	- Pushes the Docker image to the Google Artifact Registry
	- Deploys a new Cloud Run revision with the new image and serves the app to the public internet
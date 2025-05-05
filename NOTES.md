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
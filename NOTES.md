By default, a step "succeeds" if it exits with a status code of 0 and "fails" if it exits with a status code other than 0.

Every (good) CLI tool that I'm aware of follows the convention of exit code 0 = pass, anything else = fail. For example, if a test case fails, go test will exit with a status code of 1.
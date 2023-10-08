# Goledger Chaincode 

## Directory Structure

- `/fabric`: Fabric network v2.2 used as a test environment
- `/chaincode`: chaincode-related files
- `/ccapi`: chaincode REST API in Golang project

## Development

The `cc-tools` library has been tested in Fabric v1.4, v2.2 and v2.4 networks.

Dependencies for chaincode and chaincode API:

- Go 1.14 or higher

Dependencies for test environment:

- Docker 20.10.5 or higher
- Docker Compose 1.28.5 or higher

Intallation of the Chaincode API Go:

```bash
$ cd chaincode; go mod vendor; cd ..
$ cd ccapi; go mod vendor; cd ..
```


## Deploying test env in v2.2

After installing, use the script `./startDev.sh` in the root folder to start the development environment. It will
start all components of the project with 3 organizations.

If you want to deploy with 1 organization, run the command `./startDev.sh -n 1`.

To apply chaincode changes, run `$ ./upgradeCC2.sh <version> <sequence>` with a version higher than the current one (starts with 0.1). Append `-n 1` to the command if running with 1 organization.

To apply CC API changes, run `$ ./reloadCCAPI.sh`.

## Automated tryout and test

To test transactions after starting all components, run `$ ./tryout.sh`. 

To test transactions using the godog tool, run `$ ./godog.sh`.

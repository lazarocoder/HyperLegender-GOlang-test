# Chaincode 

# Overview

"The interface is made available in a Docker image. We have a transaction in which we can start three organizations (orgs) that communicate with each other within a blockchain network. The purpose of the application is to perform a CRUD operation, which involves simulating a book, a secret, and a library, with secure contracts that enable the execution of transactions and data updates in each organization that is part of the same network."

## Directory Structure:

- `/fabric`: Fabric network v2.2 used as a test environment
- `/chaincode`: chaincode-related files
- `/ccapi`: chaincode REST API in Golang project

## Development:

Dependencies for chaincode and chaincode API:

- Go 1.14/higher

Dependencies for test environment:

- Docker 20.10.5/higher
- Docker Compose 1.28.5/higher

Intallation of the Chaincode API Go:

```bash
$ cd chaincode; go mod vendor; cd ..
$ cd ccapi; go mod vendor; cd ..
```


## Deploying test env: 

After installing, you must use the script `./startDev.sh` in the root folder to start the development environment. 
It will start all components of the project with (3) organizations.

To deploy with (1) organization, run the command `./startDev.sh -n 1`.

## Automated tryout and test

To test transactions after starting all components, run `$ ./tryout.sh`. 

To test transactions using the godog tool, run `$ ./godog.sh`.

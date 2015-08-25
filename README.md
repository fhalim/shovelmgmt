# Introduction
Tool for building [dynamic shovels](https://www.rabbitmq.com/shovel-dynamic.html) to move data from one cluster to the other.
Currently requires the shovel plugin to be installed on the downstream cluster and for the shovel HTTP admin API to be available.

# Building
`go install github.com/fhalim/shovelmgmt/cmd/{autoshovel,deleteallautoshovels}`

# Executing
To get commandline options

`autoshovel -h`

## Example usage
`autoshovel --upstreamamqpport 5673 --upstreamadminport 15673 --downstreamamqpport 5676 --downstreamadminport 15676`

To delete auto created federations (matching prefix for autocreated shovels)

`deleteallautoshovels -adminport 15676`

# amm
**amm** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://starport.com).

## Scaffold commands
starport s chain github.com/lostak/amm --no-module
starport s module ocean --ibc --ordering ordered --dep account,bank,capability,distr,mint
starport s map pool tokenA:coin tokenB:coin shares:coin swapFee --module ocean --no-message
starport s message create-pool tokenA:coin tokenB:coin shares:coin swapFee --module ocean

## Logic added
types/keys.go:
	- PoolIndex function 
		- Return the index of a pool given two tokens
		- Return error if coins are of same denom
keeper/msg_server_create_pool.go
	- CreatePool function
		- get pool index using ^ 
		- check if pool exists
		- validate tokens
		- create pool
		- set set pool at index using k.SetPool

proto/ocean/pool.proto
	- change generated index to type string
	 
## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.com).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.com/lostak/amm@latest! | sudo bash
```
`lostak/amm` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://starport.com)
- [Tutorials](https://docs.starport.com/guide)
- [Starport docs](https://docs.starport.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/H6wGTY8sxw)

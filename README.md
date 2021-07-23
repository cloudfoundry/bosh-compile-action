# bosh-compile

A CLI tool to allow you to test compile a set of bosh packages without using a bosh director.

The aim of this project is to create a GitHub action that can be used to create fast feedback on opensource bosh releases.

## Usage

Specify a list of packages to build:

```shell
bc compile --file <path-to-bosh-release>.tgz --packages <package-to-compile>
```

Allow `bc` to determine the best packages to build: 

```shell
bc compile --file <path-to-bosh-release>.tgz --guess
```

## Docs

[CLI Documentation](./docs/bc.md) can be found here.

## License

// TODO

## Contributing

// TODO

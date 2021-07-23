# bosh-compile

A CLI tool (and GitHub action) to allow you to test compile a set of bosh packages without using a bosh director.

The aim of this project is provide fast feedback as to whether a bosh release can be compiled or not.

## Usage

Specify a list of packages to build:

```shell
bc compile --file <path-to-bosh-release>.tgz --packages <package-to-compile>
```

Allow `bc` to determine the best packages to build: 

```shell
bc compile --file <path-to-bosh-release>.tgz --guess
```

## Usage in GitHub Actions

This action can be used in combination with the `orange-cloudfoundry/bosh-release-action` to create and test the bosh release. e.g.

```
      - name: Build Dev Release
        id: build
        uses: orange-cloudfoundry/bosh-release-action@v1.1.0

      - name: Compile Dev Release
        uses: cloudfoundry/bosh-compile-action@main
        with:
          file: ${{ steps.build.outputs.file }}
          args: --guess --debug
```

## Docs

[CLI Documentation](./docs/bc.md) can be found here.

## License

// TODO

## Contributing

// TODO

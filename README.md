# Docker Image Migrator


This is a command-line utility written in Go that automates the process of pulling Docker images, retagging them, and pushing them to a new Docker registry.

The utility reads from a configuration file in YAML format to get the list of Docker images to migrate and the target Docker registry.

## Installation

First, make sure you have Go installed on your machine. You also need to have Docker installed and running on your machine.

Clone the repository to your local machine:

```bash
git clone https://github.com/ruslanguns/docker-image-migrator.git
```

Navigate into the directory:

```bash
cd docker-image-migrator
```

To compile the project, run:

```bash
go build -o docker-image-migrator main.go
```

This will create an executable named `docker-image-migrator`.

## Usage

The Docker Image Migrator is invoked from the command line.

You need to provide the path to the YAML configuration file via the --config flag.

Example:

```bash
./docker-image-migrator --config config.yaml
```

## Configuration

The configuration file is in YAML format. It contains the following fields:

```yaml
images:
  - name: alpine
    tag: latest
  - name: ubuntu
    tag: latest
newRegistry: "newregistry"
```

Each item in the images list has a name (the image name) and a tag.

> If the image comes from a private registry, you need to specify the full path to the image, including the registry name. For example, `registry.example.com/alpine`, and login to the registry before running the utility.

newRegistry should be the name of the target Docker registry where the images will be pushed after being pulled and retagged.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

MIT
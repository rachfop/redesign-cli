# Airfoil CLI Documentation

Airfoil is a command-line interface for developing and deploying projects on RunPod's infrastructure.

## Quick Start

1. Create a new project:
   ```
   airfoil create --name my-project
   ```

2. Navigate to the project directory:
   ```
   cd my-project
   ```

3. Start a development session:
   ```
   airfoil dev
   ```

4. Deploy your project:
   ```
   airfoil deploy
   ```


## Commands

### create (new)

Creates a new RunPod project folder on your local machine.

Usage:
```
airfoil create [flags]
airfoil new [flags]
```

Flags:
- `--name`, `-n`: Set the project name, a directory with this name will be created in the current path.
- `--init`, `-i`: Initialize the project in the current directory instead of creating a new one.
- `--model`, `-m`: Specify the Hugging Face model name for the project.
- `--type`, `-t`: Specify the model type for the project.

Example:
```
airfoil create --name my-project --model gpt2 --type LLM
```

### dev (start)

Start a development session for the current project. 
This command establishes a connection between your local development environment and your RunPod project environment, allowing for real-time synchronization of changes.

Usage:
```
airfoil dev [flags]
airfoil start [flags]
```

Flags:
- `--select-volume`: Choose a new default network volume for the project.
- `--prefix-pod-logs`: Include the Pod ID as a prefix in log messages from the project Pod.

Example:

```
airfoil dev --select-volume
```

### deploy

Deploys a serverless endpoint for the RunPod project in the current folder.

Usage:
```
airfoil deploy
```

Example:

```
airfoil deploy
```

### build

Builds a local Dockerfile for the project in the current folder. 
You can use this Dockerfile to build an image and deploy it to any API server.

Usage:
```
airfoil build [flags]
```

Flags:
- `--include-env`: Incorporate environment variables defined in runpod.toml into the generated Dockerfile.

Example:

```
airfoil build --include-env
```

## Global Flags

These flags can be used with any command:

- `--help`, `-h`: Show help for the command

## Configuration

Airfoil uses a `runpod.toml` file in your project directory for configuration. This file is created when you run the `create` command and can be edited manually.

For more detailed information about each command and its options, use the `--help` flag with any command.


## Build the CLI


```
go mod tidy
```

```
go build -o airfoil
```

```
./airfoil
```

# Airfoil

Airfoil is a command-line interface for developing and deploying projects on RunPod's infrastructure.

## Installation

To install Airfoil, run:

```
go install github.com/yourusername/airfoil@latest
```

Ensure that your Go bin directory is in your PATH.


## Documentation

For full documentation on all available commands and options, see the [CLI Documentation](docs/airfoil.md).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).
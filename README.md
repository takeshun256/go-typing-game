# Go Typing Game

This is a simple typing game written in Go.

Play locally or with Docker and test your skills by correctly typing as many randomly selected words as you can. Your score is displayed at the end of the game. Enjoy!

<img width="929" alt="image" src="https://user-images.githubusercontent.com/75155218/226076680-ed54be3c-593a-44b4-bbb4-5a56728e3022.png">


## Requirements

### Docker
- Docker
- Docker Compose

### Local
- Go 1.20

## Installation and Setup

### Docker
1. Clone the repository:

```bash
git clone https://github.com/takeshun256/go-typing-game.git
cd go-typing-game
```

2. Build the Docker image and start the container:

```bash
make docker-build
```

3. Enter the container and Start the game.

```bash
make docker-run
```

5. Stop and remove the Docker container:

```bash
make docker-clean
```

### Local

1. Clone the repository:

```bash
git clone https://github.com/takeshun256/go-typing-game.git
cd go-typing-game
```

2. Run the game:

```bash
make run
```

3. or Build and Run the game:

```bash
make build
./main
```

4. clean the binary executable file
```bash
make clean
```

## How to Play
The game will randomly select a word from a list of words and prompt you to type the word. Your goal is to type as many words as possible within 60 seconds.

If you type the word correctly, you will score one point. If you make a mistake, no points will be awarded. After 60 seconds, the game will end and your score will be displayed.

Good luck!

## Contributing

Thank you for your interest in contributing to this project! As this is a very small project, there may not be any specific guidelines for contributions. However, if you have any suggestions or improvements, please feel free to open an issue or a pull request. We appreciate any and all contributions to make this project better!

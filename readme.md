
# RPSLS

golang implementation of ROCK PAPPER SCISSORS LIZARD AND SPOCK.
This app exposes four endpoint


## Acknowledgements

 - [How to play Rock Paper Scissors Lizard Spock](https://www.youtube.com/watch?v=zjoVuV8EeOU)
 - [ROCK PAPER SCISSORS SPOCK LIZARD THOERY](http://www.samkass.com/theories/RPSSL.html)
![Logo](https://content.instructables.com/ORIG/FIU/AIWE/I7Q0TCUT/FIUAIWEI7Q0TCUT.jpg)

    
## Run Locally

Clone the project

```bash
  git clone https://github.com/acidicyemi/RPSLS_Golang.git
```

Go to the project directory

```bash
  cd RPSLS_Golang
```

Create the dockerfile image

```bash
  docker build -t game-service .
```

Start A Container from the Image where the 8064 is the port you want to expose on your localhost

```bash
  docker run -d -p 8086:8084 game-service
```
## Documentation

[POSTMAN DOCUMENTATION](https://documenter.getpostman.com/view/5487875/TzY6AuM9
)

  
## License

[MIT](https://choosealicense.com/licenses/mit/)
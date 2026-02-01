I notice Game members are initialized in func main itself. Extract thsi code to method on Game itself.

---

main.go is getting a little too long. move the Game class and associated helpers into game.go and keep func main in main.go

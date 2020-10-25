# Warta

Warta is a Golang package for event broadcast/emitters.

## Installation

Use `go get` command.

```bash
go get github.com/raspiantoro/warta
```

## Usage

```golang
func main() {

	var helloTopic warta.TopicName = "hello"

	helloFunc := func(fName string, lName string) {
		fmt.Printf("Hello %s %s\n", fName, lName)
	}

	warta := warta.New()

	l, err := warta.ListenCreate(helloTopic, helloFunc)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	warta.BroadcastCreate(helloTopic, "John", "Doe")

}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
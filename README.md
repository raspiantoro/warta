![warta](https://user-images.githubusercontent.com/11055157/98251386-3ab8f280-1fab-11eb-956f-23e6bdc71924.jpg)


[![Tests Status](https://github.com/raspiantoro/warta/workflows/tests/badge.svg)](https://github.com/raspiantoro/warta/actions)
[![Coverage Status](https://coveralls.io/repos/github/raspiantoro/warta/badge.svg?branch=master&service=github)](https://coveralls.io/github/raspiantoro/warta?branch=master)
[![Go Report Status](https://goreportcard.com/badge/github.com/raspiantoro/warta)](https://goreportcard.com/report/github.com/raspiantoro/warta)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=raspiantoro_warta&metric=alert_status)](https://sonarcloud.io/dashboard?id=raspiantoro_warta)     
[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/dashboard?id=raspiantoro_warta)

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

	warta.BroadcastClose(helloTopic, "John", "Doe")

}
```

## Contributors
1. [Mario Raspiantoro](https://github.com/raspiantoro)
2. [Ishaq Dwiputra (logo design)](https://www.behance.net/ishaq192933b63)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

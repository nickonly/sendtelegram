package main

import (
    "os"
    "fmt"
)



func Usage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " --token|-t YOUR_TOKEN --message|-m YOUR_MESSAGE [--id|-i chat_or_user_id]")
	fmt.Println("Example:")
	fmt.Println(os.Args[0] + " -t 123456788:qwertyuiopasdfghjklzxcvbnmqwertyuio -m 'Hello, World!' -i 101 -i 202 -i 303")
	os.Exit(1)
}



func main() {
	var message, token string
	var chat_ids []string
	for i := 1; i < len(os.Args); i += 2 {
		switch os.Args[i] {
			case "-t": token = os.Args[i + 1]
			case "--token": token = os.Args[i + 1]
			case "-m": message = os.Args[i + 1]
			case "--message": message = os.Args[i + 1]
			case "-i": chat_ids = append(chat_ids, os.Args[i + 1])
			case "--id": chat_ids = append(chat_ids, os.Args[i + 1])
			default: Usage()
		}
	}
	if len(message) < 1 || len(token) < 1 || len(chat_ids) < 1 {
		Usage()
	}
	for _, chat_id := range chat_ids {
		response, err := SendMessage(token, chat_id, message)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if response.Ok == false {
			fmt.Println(response.ErrorCode)
			fmt.Println(response.Description)
			continue
		}
	}
}

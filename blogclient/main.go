package main

import (
	"blogg/x/blogg/types"
	"context"
	"fmt"
	"log"

	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

// Importing the general purpose Cosmos blockchain client

// Importing the types package of your blog blockchain

func main() {
	// обозначаем имя блокчайна при создании
	addressPrefix := "blogg"

	//получаем обьект клиента cosmosSDK
	cosmos, err := cosmosclient.New(
		context.Background(),
		cosmosclient.WithAddressPrefix(addressPrefix),
	)
	if err != nil {
		log.Fatal(err)
	}
	// получаем имя созданных аккаунтов при прослушивании блокчейна
	accountName := "alice"
	// получаем учетную запись аккаунта alice
	account, err := cosmos.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}
	// получаем адрес нашего блокчейна
	addr := account.Address(addressPrefix)
	// создаём сообщение для отправки в блокчейн
	msg := &types.MsgCreatePost{
		Creator: addr,
		Title: "Hello",
		Body: "This is the first post Maksata kotoryi rabotaet v mydatacoin",
	}
	// транслируем в блокчейн сообщение от аккаунта alice
	txResp, err := cosmos.BroadcastTx(account.Name, msg)
	if err != nil {
		log.Fatal(err)
	}
	// принтим ответ от трансляции транзакции
	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)
	// создаём экземпляр клиента для запроса в блокчейн блогг
	queryClient := types.NewQueryClient(cosmos.Context())
	// делаем запрос в блокчейн используя метод клиента POSt
	queryResp, err := queryClient.Posts(context.Background(), &types.QueryPostsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)




}


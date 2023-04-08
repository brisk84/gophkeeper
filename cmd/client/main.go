package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/internal/gophclient"
	"github.com/manifoldco/promptui"
	"github.com/nexidian/gocliselect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}

func intTests(ctx context.Context, c api.GophKeeperClient) int {
	resp, err := c.Register(ctx, &api.RegisterLoginReq{
		Login:    "Vasya",
		Password: "123",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Token)
	return 0
}

func menu(m map[string]func(gc gophclient.GophClient), gc gophclient.GophClient) {
	menu := gocliselect.NewMenu("Welcome to GophKeeper! Choose a command")

	menu.AddItem("Register", "register")
	menu.AddItem("Login", "login")
	menu.AddItem("Save data", "save")
	menu.AddItem("Save login", "save_login")
	menu.AddItem("Save text", "save_text")
	menu.AddItem("Save card", "save_card")
	menu.AddItem("List data", "list_data")
	menu.AddItem("Get data", "get_data")
	menu.AddItem("Get login", "get_login")
	menu.AddItem("Get text", "get_text")
	menu.AddItem("Exit", "exit")

	choice := menu.Display()
	if f, ok := m[choice]; ok {
		f(gc)
	}
}

func Register(gc api.GophKeeperClient) {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("username must have more than 3 characters")
		}
		return nil
	}

	var username string
	u, err := user.Current()
	if err == nil {
		username = u.Username
	}

	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validate,
		Default:  username,
	}

	login, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// fmt.Printf("Your username is %q\n", result)

	validate = func(input string) error {
		if len(input) < 6 {
			return errors.New("password must have more than 6 characters")
		}
		return nil
	}

	prompt = promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}

	password, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// fmt.Printf("Your password is %q\n", result)

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.Register(ctx, &api.RegisterLoginReq{
		Login:    login,
		Password: password,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Token)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func Login() {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("Username must have more than 3 characters")
		}
		return nil
	}

	var username string
	u, err := user.Current()
	if err == nil {
		username = u.Username
	}

	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validate,
		Default:  username,
	}

	login, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// fmt.Printf("Your username is %q\n", result)

	validate = func(input string) error {
		if len(input) < 6 {
			return errors.New("Password must have more than 6 characters")
		}
		return nil
	}

	prompt = promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}

	password, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// fmt.Printf("Your password is %q\n", result)

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.Login(ctx, &api.RegisterLoginReq{
		Login:    login,
		Password: password,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Success, resp.Token)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func SaveData() {
	// validate := func(input string) error {
	// 	if len(input) < 3 {
	// 		return errors.New("Username must have more than 3 characters")
	// 	}
	// 	return nil
	// }

	// var username string
	// u, err := user.Current()
	// if err == nil {
	// 	username = u.Username
	// }

	// prompt := promptui.Prompt{
	// 	Label:    "Username",
	// 	Validate: validate,
	// 	Default:  username,
	// }

	// login, err := prompt.Run()

	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// // fmt.Printf("Your username is %q\n", result)

	// validate = func(input string) error {
	// 	if len(input) < 6 {
	// 		return errors.New("Password must have more than 6 characters")
	// 	}
	// 	return nil
	// }

	// prompt = promptui.Prompt{
	// 	Label:    "Password",
	// 	Validate: validate,
	// 	Mask:     '*',
	// }

	// password, err := prompt.Run()

	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// fmt.Printf("Your password is %q\n", result)

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.SaveData(ctx, &api.SaveDataReq{
		Token: "68c0c3fe6241aae8d03fbcae65dd7bbcaac15a8ec34dbeb72e9765bed876f3a9",
		Title: "Test binary data",
		Data:  []byte("123"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Id)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func SaveLogin() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.SaveLogin(ctx, &api.SaveLoginReq{
		Token: "68c0c3fe6241aae8d03fbcae65dd7bbcaac15a8ec34dbeb72e9765bed876f3a9",
		Title: "Test login/pass data",
		Login: "Login12",
		Pass:  "Pass32",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Id)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func SaveText() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.SaveText(ctx, &api.SaveTextReq{
		Token: "68c0c3fe6241aae8d03fbcae65dd7bbcaac15a8ec34dbeb72e9765bed876f3a9",
		Title: "Test text data",
		Text:  "Test string of text",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Id)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func SaveCard() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.SaveBankCard(ctx, &api.SaveBankCardReq{
		Token:  "68c0c3fe6241aae8d03fbcae65dd7bbcaac15a8ec34dbeb72e9765bed876f3a9",
		Title:  "Test bank card data",
		CardNo: "1234 5678 9012 3456",
		Valid:  "06/25",
		Cvv:    "123",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Id)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func ListData() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.ListData(ctx, &api.ListDataReq{
		Token: "68c0c3fe6241aae8d03fbcae65dd7bbcaac15a8ec34dbeb72e9765bed876f3a9",
	})
	if err != nil {
		panic(err)
	}
	for _, item := range resp.Items {
		fmt.Println(item)
	}
	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func GetData() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.GetData(ctx, &api.GetDataReq{
		Token:  "68c0c3fe6241aae8d03fbcae65dd7bbcaac15a8ec34dbeb72e9765bed876f3a9",
		DataId: 3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func GetLogin() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.GetLogin(ctx, &api.GetDataReq{
		Token:  "68c0c3fe6241aae8d03fbcae65dd7bbcaac15a8ec34dbeb72e9765bed876f3a9",
		DataId: 5,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func GetText() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// exitCode := intTests(ctx, c)
	// fmt.Println("ExitCode:", exitCode)

	resp, err := c.GetText(ctx, &api.GetDataReq{
		Token:  "68c0c3fe6241aae8d03fbcae65dd7bbcaac15a8ec34dbeb72e9765bed876f3a9",
		DataId: 9,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	cancel()
	cc1.Close()
	// os.Exit(exitCode)
}

func main() {

	gc, err := gophclient.New("0.0.0.0:4343")
	if err != nil {
		panic(err)
	}

	m := make(map[string]func(gc *gophclient.GophClient))
	m["register"] = Register
	m["login"] = Login
	m["save"] = SaveData
	m["save_login"] = SaveLogin
	m["save_text"] = SaveText
	m["save_card"] = SaveCard
	m["list_data"] = ListData
	m["get_data"] = GetData
	m["get_login"] = GetLogin
	m["get_text"] = GetText

	menu(m, gc)
	// Register()
	return

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	exitCode := intTests(ctx, c)
	fmt.Println("ExitCode:", exitCode)

	cancel()
	cc1.Close()
	os.Exit(exitCode)
}

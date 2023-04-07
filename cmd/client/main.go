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

func menu() {
	// tm.Clear() // Clear current screen
	// // for {
	// // 	// By moving cursor to top-left position we ensure that console output
	// // 	// will be overwritten each time, instead of adding new.
	// // 	tm.MoveCursor(1, 1)
	// // 	tm.Println("Current Time:", time.Now().Format(time.RFC1123))
	// // 	tm.Flush() // Call it every time at the end of rendering
	// // 	time.Sleep(time.Second)
	// // }
	// tm.Println("GophKeeper client")
	// tm.Println("Enter command")
	// tm.Println("r - register")
	// tm.Println("l - login")
	// tm.Screen.Read()
	// fmt.Println("GophKeeper client")
	// fmt.Println("Enter command")
	// fmt.Println("r - register")
	// fmt.Println("l - login")
	// var comm string
	// fmt.Scan(&comm)
	// fmt.Println(comm)
	menu := gocliselect.NewMenu("Welcome to GophKeeper! Choose a command")

	menu.AddItem("Register", "register")
	menu.AddItem("Login", "login")
	menu.AddItem("Exit", "exit")

	choice := menu.Display()

	fmt.Printf("Choice: %s\n", choice)
}

func Register() {
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

func main() {
	Register()
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

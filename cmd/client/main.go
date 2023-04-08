package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"time"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/domain"
	"github.com/brisk84/gophkeeper/internal/gophclient"
	"github.com/manifoldco/promptui"
	"github.com/nexidian/gocliselect"
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

type callFunc = func(ctx context.Context, gc *gophclient.GophClient) error

func menu(m map[string]callFunc, gc *gophclient.GophClient) {
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := f(ctx, gc)
		if err != nil {
			fmt.Println(err)
		}
		cancel()
	}
}

func getInput(label string, minLen int, defVal string, mask rune) (string, error) {
	validate := func(input string) error {
		if len(input) < minLen {
			return fmt.Errorf("%s must have more than %d characters", label, minLen)
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:   label,
		Default: defVal,
		Mask:    mask,
	}
	if minLen > 0 {
		prompt.Validate = validate
	}
	title, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return title, nil
}

func Register(ctx context.Context, gc *gophclient.GophClient) error {
	var username string
	if u, err := user.Current(); err == nil {
		username = u.Username
	}
	login, err := getInput("Username", 3, username, 0)
	if err != nil {
		return err
	}
	password, err := getInput("Password", 6, "", '*')
	if err != nil {
		return err
	}

	token, success, err := gc.Register(ctx, domain.User{Login: login, Password: password})
	if err != nil {
		panic(err)
	}
	fmt.Println(token, success)
	return nil
}

func Login(ctx context.Context, gc *gophclient.GophClient) error {
	var username string
	if u, err := user.Current(); err == nil {
		username = u.Username
	}
	login, err := getInput("Username", 3, username, 0)
	if err != nil {
		return err
	}
	password, err := getInput("Password", 6, "", '*')
	if err != nil {
		return err
	}

	token, success, err := gc.Login(ctx, domain.User{Login: login, Password: password})
	if err != nil {
		return err
	}
	fmt.Println(token, success)
	return nil
}

func ListData(ctx context.Context, gc *gophclient.GophClient) error {
	data, err := gc.ListData(ctx)
	if err != nil {
		return err
	}
	for _, item := range data {
		fmt.Println(item)
	}
	return nil
}

func SaveText(ctx context.Context, gc *gophclient.GophClient) error {
	title, err := getInput("Title", 0, "", 0)
	if err != nil {
		return err
	}
	text, err := getInput("Text", 0, "", 0)
	if err != nil {
		return err
	}

	id, err := gc.SaveText(ctx, title, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Saved. Rec number:", id)
	return nil
}

func GetText(ctx context.Context, gc *gophclient.GophClient) error {
	sId, err := getInput("DataId", 0, "", 0)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		return err
	}
	resp, err := gc.GetText(ctx, id)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

func SaveLogin(ctx context.Context, gc *gophclient.GophClient) error {
	title, err := getInput("Title", 0, "", 0)
	if err != nil {
		return err
	}
	login, err := getInput("Login", 0, "", 0)
	if err != nil {
		return err
	}
	password, err := getInput("Password", 0, "", 0)
	if err != nil {
		return err
	}

	id, err := gc.SaveLogin(ctx, title, login, password)
	if err != nil {
		return err
	}
	fmt.Println("Saved. Rec number:", id)
	return nil
}

func GetLogin(ctx context.Context, gc *gophclient.GophClient) error {
	sId, err := getInput("DataId", 0, "", 0)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		return err
	}
	resp, err := gc.GetLogin(ctx, id)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

func SaveData(ctx context.Context, gc *gophclient.GophClient) error {
	title, err := getInput("Title", 0, "", 0)
	if err != nil {
		return err
	}
	filename, err := getInput("Filename", 0, "", 0)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	id, err := gc.SaveData(ctx, title, data)
	if err != nil {
		return err
	}
	fmt.Println("Saved. Rec number:", id)
	return nil
}

func GetData(ctx context.Context, gc *gophclient.GophClient) error {
	sId, err := getInput("DataId", 0, "", 0)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		return err
	}
	filename, err := getInput("Filename", 0, "", 0)
	if err != nil {
		return err
	}

	data, err := gc.GetData(ctx, id)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, fs.FileMode(os.O_CREATE))
	if err != nil {
		return err
	}
	fmt.Println(filename + " created")
	return nil
}

func SaveCard(ctx context.Context, gc *gophclient.GophClient) error {
	title, err := getInput("Title", 0, "", 0)
	if err != nil {
		return err
	}
	cardNo, err := getInput("CardNo", 0, "", 0)
	if err != nil {
		return err
	}
	valid, err := getInput("Valid", 0, "", 0)
	if err != nil {
		return err
	}
	cvv, err := getInput("CVV", 0, "", 0)
	if err != nil {
		return err
	}
	cardInfo := domain.CardInfo{
		CardNo: cardNo,
		Valid:  valid,
		CVV:    cvv,
	}

	id, err := gc.SaveBankCard(ctx, title, cardInfo)
	if err != nil {
		return err
	}
	fmt.Println("Saved. Rec number:", id)
	return nil
}

func Exit(ctx context.Context, gc *gophclient.GophClient) error {
	panic(1)
}

func main() {
	gc, err := gophclient.New("0.0.0.0:4343")
	if err != nil {
		panic(err)
	}

	m := make(map[string]callFunc)
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
	m["exit"] = Exit

	for {
		menu(m, gc)
	}
}

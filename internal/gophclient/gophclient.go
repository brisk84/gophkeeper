package gophclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GophClient struct {
	creds   credentials.TransportCredentials
	srvAddr string
	token   string
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := os.ReadFile("cert/ca-cert.pem")
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

func New(srvAddr string) (*GophClient, error) {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		return nil, err
	}
	return &GophClient{creds: tlsCredentials, srvAddr: srvAddr}, nil
}

func (g *GophClient) Register(ctx context.Context, user domain.User) (string, bool, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return "", false, err
	}
	defer conn.Close()
	cl := api.NewGophKeeperClient(conn)

	resp, err := cl.Register(ctx, &api.RegisterLoginReq{
		Login:    user.Login,
		Password: user.Password,
	})
	if err != nil {
		return "", false, err
	}
	if resp.Success {
		g.token = resp.Token
	}
	return resp.Token, resp.Success, nil
}

func (g *GophClient) Login(ctx context.Context, user domain.User) (string, bool, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return "", false, err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.Login(ctx, &api.RegisterLoginReq{
		Login:    user.Login,
		Password: user.Password,
	})
	if err != nil {
		return "", false, err
	}
	if resp.Success {
		g.token = resp.Token
	}
	return resp.Token, resp.Success, nil
}

func (g *GophClient) SaveText(ctx context.Context, title, text string) (int, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.SaveText(ctx, &api.SaveTextReq{
		Token: g.token,
		Title: title,
		Text:  text,
	})
	if err != nil {
		return 0, err
	}
	return int(resp.Id), nil
}

func (g *GophClient) GetText(ctx context.Context, id int) (string, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.GetText(ctx, &api.GetDataReq{
		Token:  g.token,
		DataId: int32(id),
	})
	if err != nil {
		return "", err
	}
	return resp.Text, nil
}

func (g *GophClient) SaveLogin(ctx context.Context, title, login, pass string) (int, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.SaveLogin(ctx, &api.SaveLoginReq{
		Token: g.token,
		Title: title,
		Login: login,
		Pass:  pass,
	})
	if err != nil {
		return 0, err
	}
	return int(resp.Id), nil
}

func (g *GophClient) GetLogin(ctx context.Context, id int) (string, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.GetText(ctx, &api.GetDataReq{
		Token:  g.token,
		DataId: int32(id),
	})
	if err != nil {
		return "", err
	}
	return resp.Text, nil
}

func (g *GophClient) SaveData(ctx context.Context, title string, data []byte) (int, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.SaveData(ctx, &api.SaveDataReq{
		Token: g.token,
		Title: title,
		Data:  data,
	})
	if err != nil {
		return 0, err
	}
	return int(resp.Id), nil
}

func (g *GophClient) GetData(ctx context.Context, id int) ([]byte, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.GetData(ctx, &api.GetDataReq{
		Token:  g.token,
		DataId: int32(id),
	})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (g *GophClient) SaveBankCard(ctx context.Context, title string, cardInfo domain.CardInfo) (int, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.SaveBankCard(ctx, &api.SaveBankCardReq{
		Token:  g.token,
		Title:  title,
		CardNo: cardInfo.CardNo,
		Valid:  cardInfo.Valid,
		Cvv:    cardInfo.CVV,
	})
	if err != nil {
		return 0, err
	}
	return int(resp.Id), nil
}

func (g *GophClient) ListData(ctx context.Context) ([]domain.Data, error) {
	conn, err := grpc.DialContext(ctx, g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	resp, err := cl.ListData(ctx, &api.ListDataReq{
		Token: g.token,
	})
	if err != nil {
		return nil, err
	}
	var ret []domain.Data
	for _, item := range resp.Items {
		ret = append(ret, domain.Data{
			Id:    int(item.Id),
			Title: item.Title,
			Type:  item.Type,
		})
	}
	return ret, nil
}

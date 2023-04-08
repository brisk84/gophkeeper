package handler

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/domain"
)

func (h *Handler) SaveData(ctx context.Context, req *api.SaveDataReq) (*api.SaveDataResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user, err := h.useCase.Auth(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	id, err := h.useCase.SaveData(ctx, user.Id, req.Title, req.Data)
	if err != nil {
		return nil, err
	}
	return &api.SaveDataResp{
		Id: int32(id),
	}, nil
}

func (h *Handler) SaveLogin(ctx context.Context, req *api.SaveLoginReq) (*api.SaveDataResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user, err := h.useCase.Auth(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	id, err := h.useCase.SaveLogin(ctx, user.Id, req.Title, req.Login, req.Pass)
	if err != nil {
		return nil, err
	}
	return &api.SaveDataResp{
		Id: int32(id),
	}, nil
}

func (h *Handler) SaveText(ctx context.Context, req *api.SaveTextReq) (*api.SaveDataResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user, err := h.useCase.Auth(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	id, err := h.useCase.SaveText(ctx, user.Id, req.Title, req.Text)
	if err != nil {
		return nil, err
	}
	return &api.SaveDataResp{
		Id: int32(id),
	}, nil
}

func (h *Handler) SaveBankCard(ctx context.Context, req *api.SaveBankCardReq) (*api.SaveDataResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user, err := h.useCase.Auth(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	card := domain.CardInfo{
		CardNo: req.CardNo,
		Valid:  req.Valid,
		CVV:    req.Cvv,
	}
	id, err := h.useCase.SaveBankCard(ctx, user.Id, req.Title, card)
	if err != nil {
		return nil, err
	}
	return &api.SaveDataResp{
		Id: int32(id),
	}, nil
}

func (h *Handler) ListData(ctx context.Context, req *api.ListDataReq) (*api.ListDataResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user, err := h.useCase.Auth(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	ld, err := h.useCase.ListData(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	var ret api.ListDataResp
	for _, v := range ld {
		ret.Items = append(ret.Items, &api.ListData{
			Id:    int32(v.Id),
			Title: v.Title,
			Type:  v.Type,
		})
	}
	return &ret, nil
}

func (h *Handler) GetData(ctx context.Context, req *api.GetDataReq) (*api.GetDataResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user, err := h.useCase.Auth(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	ret, err := h.useCase.GetData(ctx, user.Id, int(req.DataId))
	if err != nil {
		return nil, err
	}
	return &api.GetDataResp{Data: ret}, nil
}

func (h *Handler) GetLogin(ctx context.Context, req *api.GetDataReq) (*api.GetLoginResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user, err := h.useCase.Auth(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	li, err := h.useCase.GetData(ctx, user.Id, int(req.DataId))
	if err != nil {
		return nil, err
	}
	var ret domain.LoginInfo
	err = json.Unmarshal(li, &ret)
	if err != nil {
		return nil, err
	}
	return &api.GetLoginResp{Login: ret.Login, Pass: ret.Pass}, nil
}

func (h *Handler) GetText(ctx context.Context, req *api.GetDataReq) (*api.GetTextResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user, err := h.useCase.Auth(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	ret, err := h.useCase.GetData(ctx, user.Id, int(req.DataId))
	if err != nil {
		return nil, err
	}
	return &api.GetTextResp{Text: string(ret)}, nil
}

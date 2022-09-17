package main

import (
	"context"
	"fmt"
	"time"
)

const (
	waitDur    = 1 * time.Second
	cancelDur  = 250 * time.Millisecond
	timeoutDur = 500 * time.Millisecond
)

type Config struct {
	SelectTimeOut time.Duration
}

type DB struct {
	cfg Config
}

type User struct {
	Name string
}

func (d *DB) SelectUser(ctx context.Context, email Request) (User, error) {
	ctx2, cancel := context.WithTimeout(ctx, d.cfg.SelectTimeOut)
	defer cancel()

	timer := time.NewTimer(waitDur)
	select {
	case <-timer.C:
		return User{Name: "Gosha"}, nil
	case <-ctx2.Done():
		return User{}, ctx2.Err()
	}
}

type Handler struct {
	db *DB
}

type Response struct {
	User User
}

type Request struct {
	Email string
}

func (h *Handler) HandleApi(ctx context.Context, request Request) (Response, error) {
	u, err := h.db.SelectUser(ctx, request)
	if err != nil {
		return Response{}, err
	}
	return Response{User: u}, nil
}

func main() {
	cfg := Config{SelectTimeOut: timeoutDur}
	db := DB{cfg: cfg}
	handler := Handler{db: &db}
	ctx, cancel := context.WithCancel(context.Background())

	time.AfterFunc(cancelDur, cancel)

	req := Request{Email: "test@yandex.ru"}
	resp, err := handler.HandleApi(ctx, req)
	fmt.Println(resp, err)
}

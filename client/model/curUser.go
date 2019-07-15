package model
import (
	"net"
	"go_ChatSw/public"
)

type CurUser struct {
	Conn net.Conn
	public.User
}
package ping

import (
	"fmt"
	"golang.org/x/net/icmp"
	"net"
	"time"
)

type Ping struct {
	DelayTime int64
	CtlTime chan struct {}
	Rescyle chan struct {}
}

var (
	P *Ping
)

func init(){
	P = new(Ping)
	P.CtlTime = make(chan struct{})
	P.Rescyle = make(chan struct{},1)
}

func (p *Ping) Listen(){
	netaddr, _ := net.ResolveIPAddr("ip4:icmp", "0.0.0.0")
	conn, _ := net.ListenIP("ip4:icmp", netaddr)
	for {
		select {
		case <-P.CtlTime:
			go func(){
				for conn !=nil{
					if len(p.Rescyle) == 1{
						<-p.Rescyle
						fmt.Println("协程退出")
						return
					}
					buf := make([]byte, 1024)
					n, addr, _ := conn.ReadFromIP(buf)
					msg,_:=icmp.ParseMessage(1,buf[0:n])
					fmt.Println(p.DelayTime)
					time.Sleep(time.Duration(p.DelayTime)*time.Millisecond)
					conn.WriteToIP([]byte("111"),addr)
					fmt.Println(n, addr, msg.Type,msg.Code,msg.Checksum)
				}
			}()
		default:
		}
	}
}
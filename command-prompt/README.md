# encoding/gobを用いる
以下のようなメッセージ仕様を定義

- 最初の4バイト (2桁) にデータ長を持つ
- データ長フィールドの後ろに実際のデータ (文字列) が続く
- エンディアンはビッグエンディアン

データフォーマットはホームディレクトリのpkgディレクトリを参照

加えて，一度の接続で複数回送受信できるような仕様にした．

さらに，エンコードとデコードに，golangのencoding/gobパッケージを用いる．

## サーバ側
ソケットを開いてAccept待ちをし，Connectしてきたらデータを受信して大文字にして返すだけのサーバ．

一回処理したら，コネクションを切断する．（本来は良くない）

データの送受信部分でデータフォーマットを使用している．

クライアントが切断してくるまで処理を続ける．現実では，相手がいつまでも切ってこないとか，色々考慮しないといけないが割愛．

## クライアント側
接続したら，`hello world`というデータを送信し，応答を受信して表示する．1回ごとにコネクションを切断して毎回接続しに行く．

データの送受信部分でデータフォーマットを使用している．

毎回切断せずに，1つのコネクション上で複数回データを流す．

コマンド入力ができるようになっている．

## encoding/gob　を用いることで
クライアントサーバ間の通信のはじめに，データ型の情報を伝える．
これは，gobによる仕様であり，対向先で正しくエンコードとデコードをするために型情報を伝えることが必要になる．
gobにより各送受信の通信データサイズが少し増えるが，デコードとエンコードを考慮しなくてもよくなる．

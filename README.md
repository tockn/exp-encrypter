# 情報科学実験最終レポート

## 実行方法

このプロジェクトには`go`(1.10以上推奨), `npm`, `vue-cli2`以上が必要です。  
以下にそれぞれの環境構築方法を示します。

### Go

まずは[こちらのページ](https://golang.org/dl/)から、実行環境のOSおよびCPUアーキテクチャに合ったファイルをダウンロードしてください。  
ダウンロードしたファイルは`tar`コマンドなどで`/usr/local`へ展開して下さい。  
example:  

```
$ wget https://dl.google.com/go/go1.11.1.linux-amd64.tar.gz
$ tar -C /usr/local -xzf go11.1.linux-amd64.tar.gz
```

PATHも通します。`.bashrc` `.zshrc`などに書いておくと良いです。  

```
$ echo export 'PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
$ source ~/.bashrc
```

また、Goには`$GOPATH`という環境変数が必要です。Goは`go get`コマンドなどでpackageをcloneすることができるのですが、cloneする際のディレクトリを`$GOPATH`で指定する必要があります。プログラム中のpackageのimportもここを探します。  
適当に`$HOME`あたりにディレクトリを作成し、そこを`$GOPATH`とします。また、`$GOPATH/bin`にPATHを通します。  

```
$ mkdir ~/go
$ echo 'export GOPATH=$HOME/go' >> ~/.bashrc
$ echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
$ source ~/.bashrc
```

正しくインストールされているか確認します。

```
$ go version
> go version go1.11.1 darwin/amd64
```

以上でgoの環境構築は終了です。

### npmとvue-cli

CentOSの場合、yumでnpmを入れます。

```
$ sudo yum install -y epel-release
$ sudo yum install -y nodejs
$ sudo yum install -y npm
```

次に、npmでvue-cliを入れます。

```
$ sudo npm install -g vue-cli
```

以下のコマンドが通れば完了です。

```
$ vue --version
> 2.9.6
```

### プロジェクトのビルド

`go`と`vue-cli`を入れたら、プロジェクトのclone、ビルドを行います。  
`go get`コマンドにて、githubよりプロジェクトをクローンします。

```
go get -u github.com/tockn/exp-encrypter
```

これで`$GOPATH/src/github.com/tockn/exp-encrypter`へプロジェクトがcloneされました。  
`exp-encrypter`へ移動して、各依存関係を解決します。

```
$ cd $GOPATH/src/github.com/tockn/exp-encrypter
$ make deps
$ npm install --prefix view
```

ビルドします。

```
$ make build
```

このコマンドによって、WebpackによるVueのコンパイル結果が`/static`に展開され、Goのソースファイルがlinux, macOS, windows用(いずれもamd64)にそれぞれバイナリとしてコンパイルされます。

### プロジェクトの実行

実行環境に合ったバイナリファイルを実行することで、サーバーが立ち上がります。  

```
$ ./linux.out
> 2018/11/02 23:37:06 start listening on :8080
```

(一応、`-port`オプションでlistenするポートを指定することができますが、フロント側の叩くAPIのエンドポイントは`http://localhost:8080`で固定されているため意味がないです...)

## プログラムの動作説明

### プレイ方法

![Imgur](https://i.imgur.com/arVIVWX.png)

画面上部が暗号文と復号結果の表になっています。暗号結果の文字を変更する度にAPIサーバーを叩き、復号結果をリアルタイムで反映させます。  
その下にあるのは正解数です。全問正解すると次の問題へのボタンが現れます。  
また、画面左側にあるのが、今回一番工夫した点である暗号文に出現する単語のリストです。これらは登場数順に並んでおり、単語ごとに復号することできます。もちろん、こちらも文字を変更する度にAPIを叩き、結果をリアルタイムで反映させます。上部の表とも連動しています。

### 概要

JAMStackの構成になっています。  
サーバーサイドはGo製で、ルーティングには`gorilla/mux`パッケージを使用しています。  
フロントエンドはJavaScriptのフレームワークであるVue.js, Vue-router, Vuexを使用したSPA(Single Page Application)になっています。
サーバーは、`/api/`以下を叩くとAPIとしてJSON形式でレスポンスし、それ以外のルーティングでは`/static`以下にある静的ファイル(index.html, css, js)を返します。  

### 暗号文と鍵

暗号文と鍵は、`/quiz`以下にそれぞれ`[0-9]+.txt`, `[0-9]+.key`として同じ名前でペアで入れることで問題として追加できます。  
暗号文、鍵の読み込み、復号化などを担うのが`/crypter`にあるcrypterパッケージです。  

### API仕様

```
GET /api/quiz/{quizID}
```
quizID.txtを返します。  

```
POST /api/quiz/{quizID}
```

request bodyとして鍵を送り、その鍵による復号結果、正解数を返します。  

```
GET /api/quiz/{quizID}/words
```

quizID.txtにある単語の配列を重複なく返します。  
これが解答時のヒントとなります。

### そのほか工夫した点

入力文字は１文字しか入力できず、小文字は大文字に自動で変換されます。

## できなかったこと等

このプログラムでは空白文字で区切ることで暗号文中の単語を判定している。カンマやダブルクオートなどの記号はハードコーディングで除去している。これらを形態素解析を行うなどして、もっと賢い方法で単語をピックアップしたい。(暗号文の形態素解析...!?)

# このリポジトリは？
firebase authとの連携、herokuへの自動デプロイを実現するtemplate的なリポジトリ

# いろいろ設定方法
## firebase authenticationとの連携
GCPでサービスアカウントキーを取得し、以下コマンドでbase64化する。

```
cat [jsonkeyファイルパス] | base64
```
これを環境変数FIREBASE_KEYに設定すればfirebaseと連携できる

## postmanとの連携
otherfileフォルダに入っているinit.jsを各自のfirebase設定で変更する。  
signup.htmlでサインアップすることでAccess Tokenを取得できるので、postmanのauth設定 bearerにセットすれば認証済みAPIの検証ができる。  

## herokuの設定
herokuの環境変数には以下を設定する。

```
"HOST": "dbのホスト",
"USER": "dbのuser",
"PASSWORD": "dbのpassword",
"DB": "dbの名前",
"FIREBASE_KEY": "先程base64化したfirebase key"
```

また、containerモードで起動するため、以下コマンドを実行しておく必要がある。  

```
heroku stack:set container -a <app_name>
```

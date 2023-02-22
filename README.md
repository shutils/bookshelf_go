# bookshelf_go
シンプルな本棚アプリです。<br>
本のタイトルを追加していくことができます。

# 必要環境
以下のコマンドが実行できる環境にあること
- docker
- docker-compose
- make

# 起動方法
プロジェクトルートで以下のコマンドを実行するとコンテナが立ち上がります。
```
make local-build
make local-up
```
http://localhost:4000/book へアクセスするとアプリを確認できます。
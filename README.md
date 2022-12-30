# demo_goent
1. dockerでmysqlをlocalに構築する
2. go ent でschemaを作成しテーブルを作成する
3. go entのcurdを使ってリソースを操作する

# dockerでmysqlをlocalに構築する
docker-compose.ymlを叩けばOK
docker-compose.ymlファイルを置いたディレクトリをカレントディレクトにして
```
docker-compose up -d
```

# go ent でschemaを作成しテーブルを作成する


# 参考
主に1,3を参考にした
1. [はじめに \| ent](https://entgo.io/ja/docs/schema-def/)
2. [GoのORM「ent」の話](https://zenn.dev/masamiki/articles/83a8db3f132fcb1c48f0)
3. [100%型安全なgolangORM「ent」を使ってみた \| フューチャー技術ブログ](https://future-architect.github.io/articles/20210728a/)
4. [GoのFacebook製ORM"ent"を使ってみた \- SMARTCAMP Engineer Blog](https://tech.smartcamp.co.jp/entry/try-go-ent#%E3%82%B9%E3%82%AD%E3%83%BC%E3%83%9E%E3%81%AE%E5%AE%9A%E7%BE%A9)

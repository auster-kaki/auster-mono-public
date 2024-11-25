# auster-frontend

## requirement

* node.js
  * v20系を推奨

```bash
# setup
$ nodebrew install v20.18.0
$ nodebrew use v20.18.0

$ node -v
# v20.18.0

以下で.envを作成する
BASE_URL=http://localhost:8080
```

## Build Setup

```bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev

# build for production and launch server
$ yarn build
$ yarn start

# generate static project
$ yarn generate
```

## Deploy

gh-pagesブランチにビルドしたソースコードをpushすることでデプロイされる  

```bash
$ git checkout gh-pages
$ git pull origin main # 最新のmainを取得
$ cd ./frontend/
$ yarn generate
$ cd ../
$ cp -r frontend/dist/* .  
$ git add .
$ git commit -m "deploy"
$ git push origin gh-pages
```

## 動作確認

以下のURLを参照することでホーム画面が表示  
http://localhost:3000/c/home


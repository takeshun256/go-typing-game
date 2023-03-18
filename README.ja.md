# Go Typing Game

これはGo言語で書かれたシンプルなタイピングゲームです。

ローカル環境でプレイするか、Dockerを使用してプレイして、ランダムに選択された単語を正確に入力して、自分のスキルをテストしてください。ゲームの終わりにあなたのスコアが表示されます。楽しんで！

<img width="929" alt="image" src="https://user-images.githubusercontent.com/75155218/226076825-1d71fd43-3412-4600-9dd8-9f9d2b31c775.png">


## 必要なもの

### Docker
- Docker
- Docker Compose

### ローカル
- Go 1.20

## インストールとセットアップ

### Docker
1. リポジトリをクローンします:

```bash
git clone https://github.com/takeshun256/go-typing-game.git
cd go-typing-game
```

2. Dockerイメージをビルドし、コンテナを起動します:

```bash
docker-compose -f docker-compose.yml up -d --build
```

3. コンテナに入り、ゲームを開始します:

```bash
docker-compose -f docker-compose.prod.yml exec app go run main.go
```

4. Dockerコンテナを停止し、削除します:

```bash
docker-compose -f docker-compose.prod.yml down
```


### Local

1. リポジトリをクローンします:

```bash
git clone https://github.com/<your-username>/go-typing-game.git
cd go-typing-game
```

2. ゲームを実行します:

```bash
go run main.go
```

3. または、ビルドしてゲームを実行します:

```bash
go build main.go
./main
```


## 遊び方

ゲームは、単語リストからランダムに単語を選択し、単語をタイプするように促します。60秒間でできるだけ多くの単語をタイプすることが目標です。

単語を正しくタイプすると、1ポイントが加算されます。間違えた場合は、ポイントは加算されません。60秒後に、ゲームが終了し、スコアが表示されます。

Good luck!

## 貢献
貢献に興味を持ってくださり、ありがとうございます！このプロジェクトは非常に小さなものであるため、貢献に関する特別なガイドラインはありません。しかし、提案や改善点がある場合は、自由にissueやpull requestを開いてください。このプロジェクトをより良くするためのあらゆる貢献を歓迎します！


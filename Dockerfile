# ベースイメージをGo 1.22に更新
FROM golang:1.22 AS builder

# 作業ディレクトリ設定
WORKDIR /app

# go.mod と go.sum をコピー
COPY api/go.mod api/go.sum ./

# 依存関係をインストール
RUN go mod tidy

# ソースコードをコピー
COPY api/ ./

# ビルド
RUN go build -o main .

# 実行用イメージ
FROM golang:1.22

WORKDIR /root/
COPY --from=builder /app/main .

# 実行コマンド
CMD ["./main"]

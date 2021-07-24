FROM golang:1.15.14-alpine

# alpineでは apt-get → apk
# go get 時に git が必要になる
RUN apk update && apk add git

# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

# ホットリロードツールの導入(以下の理由により GO111MODULE= とする)
# https://github.com/keitakn/golang-grpc-server/issues/23
RUN GO111MODULE=off go get github.com/oxequa/realize
ENTRYPOINT [ "realize", "start" ]

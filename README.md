# いきものAZ エクスポーター

寂しいですが、[いきもの特化型SNSの「いきものAZ」がサービス終了](https://ikimonoaz.ikimonopal.jp/article/56123)することになりました。ぼくは2年前に投稿が止まってしまいましたが (撮った写真の整理が追いつかなくて) 、せっかく書いた記事やその写真を思い出としてとっておきたいなと思いました。

公式にはエクスポート機能は提供されていないようですので、自分用にエクスポートアプリをつくってみています。

おこられたら停止します。

## 使い方

あとで書く

## ビルド方法

### 必要なもの

- Go 1.16.3
- GNU/Linux
    - cc
    - libglx-dev
    - libgx-dev
    - libxxf86vm-dev

### 手順

```shell-session
# ソースコードもってくる
$ git clone https://github.com/t-sin/ikimonoaz-exporter
# リポジトリ内に移動
$ cd ikimonoaz-exporter
# 依存ライブラリをぜんぶもってくる
$ go get -d
# バイナリをビルド
$ go build
```

## TODO

- [x] 記事とそのコメント、およびユーザ情報 (ユーザ名やプロフィール) を取得
- [x] 記事のメディア情報から写真および動画ファイルを取得
    - [x] ハッシュだけだとファイル名が衝突する問題
- [x] 記事本文の`img`や`video`のリンク先をダウンロードしたファイルに置換
- [x] 取得したデータをローカルで閲覧可能なHTMLファイル群として出力
- [ ] GUIで操作できるよういろいろ追加と修正
    - [ ] 画面をつくる
    - [ ] 進行状態を取得できるようにする (メディアダウンロードとか長いので)
    - [ ] 進行状態に合わせてプログレスバーと何やってるか文字列を表示する
- [ ] GitHub Actionで配布用ファイルを生成

## 作者

- t-sin (<shinichi.tanaka45@gmail.com>)

## スペシャルサンクス (開発中ずっと聴いてた)

- [いのかしら りす音頭](https://www.tokyo-zoo.net/topic/topics_detail?kind=news&inst=&link_num=25526)
- [マヌルネコのうた](https://www.youtube.com/watch?v=5YLSP6b6yHg)

## ライセンス

このアプリはMITライセンスでライセンスされています。

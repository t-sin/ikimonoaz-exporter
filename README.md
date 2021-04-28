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
- [x] GUIで操作できるよういろいろ追加と修正
    - [x] 画面をつくる
    - [ ] ~~進行状態を取得できるようにする (メディアダウンロードとか長いので)~~ 実行中なのだけ表示
    - [ ] ~~進行状態に合わせてプログレスバーと何やってるか文字列を表示する~~
- [x] 全機能を結合して動作確認
- [ ] GitHub Actionで配布用ファイルを生成
- [ ] 「使い方」欄を書く

## 作者

- t-sin (<shinichi.tanaka45@gmail.com>)
    - [いきものなほしいものリスト](https://www.amazon.jp/hz/wishlist/ls/9LL3K23SPC1R?ref_=wl_share)

## スペシャルサンクス (開発中ずっと聴いてた)

- [いのかしら りす音頭](https://www.tokyo-zoo.net/topic/topics_detail?kind=news&inst=&link_num=25526)
- [マヌルネコのうた](https://www.youtube.com/watch?v=5YLSP6b6yHg)

## ライセンス

UIの表示にRounded M+フォントを組み込んで利用しています。ライセンス条文は[こちら (日本語)](LICENSE_J)あるいは[こちら (英語)](LICENSE_E)を参照してください。

このアプリはMITライセンスでライセンスされています。

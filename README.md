# いきものAZ エクスポーター

寂しいですが、[いきもの特化型SNSの「いきものAZ」がサービス終了](https://ikimonoaz.ikimonopal.jp/article/56123)することになりました。ぼくは2年前に投稿が止まってしまいましたが (撮った写真の整理が追いつかなくて) 、せっかく書いた記事やその写真を思い出としてとっておきたいなと思いました。

公式にはエクスポート機能は提供されていないようですので、自分用にエクスポートアプリをつくってみています。

おこられたら停止します。

## 使い方

これはいきもの特化型SNS「いきものAZ」のユーザデータの保存を行うツールです。以下のように使ってください:

1. [このページ](https://github.com/t-sin/ikimonoaz-exporter/releases)を開く
2. 一番上側にある**新しいバージョンの**`ikimonoaz-exporter-XXXXX.zip`をダウンロードして解凍する
    - `XXXXX`は自分の使っているオペレーティングシステムを選んでください
    - おそらくたいていの方は`windows`です
    - 英語が多くてすみません
3. 解凍先にある`ikimonoaz-exporter.exe`をダブルクリックする
4. いきものAZのマイページにいき、アドレス欄からURLをコピペする
5. `保存先フォルダを選ぶ`をクリック (いくつかファイルとフォルダができます)
6. データを保存するフォルダを選択して"Open"を押す (英語ですみません)
7. `エクスポート開始`をクリック
8. 待つ
    - 完了したら`エクスポート完了`
9. 保存先のindex.htmlをダブルクリック`
10. 🐿️🥳

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
    - [x] ~~進行状態を取得できるようにする (メディアダウンロードとか長いので)~~ 実行中なのだけ表示
    - [x] ~~進行状態に合わせてプログレスバーと何やってるか文字列を表示する~~ 実行中なのだけ表示
- [x] 全機能を結合して動作確認
- [x] GitHub Actionで配布用ファイルを生成
- [x] 「使い方」欄を書く

## 作者

- t-sin (<shinichi.tanaka45@gmail.com>)
    - [いきものなほしいものリスト](https://www.amazon.jp/hz/wishlist/ls/9LL3K23SPC1R?ref_=wl_share)

## スペシャルサンクス (開発中ずっと聴いてた)

- [いのかしら りす音頭](https://www.tokyo-zoo.net/topic/topics_detail?kind=news&inst=&link_num=25526)
- [マヌルネコのうた](https://www.youtube.com/watch?v=5YLSP6b6yHg)

## ライセンス

UIの表示にRounded M+フォントを組み込んで利用しています。ライセンス条文は[こちら (日本語)](LICENSE_J)あるいは[こちら (英語)](LICENSE_E)を参照してください。

このアプリはMITライセンスでライセンスされています。

#set text(lang: "ja")
#set text(font: "Noto Sans Javanese")

#set heading(numbering: none)

#import "@preview/slydst:0.1.5" : *

#show: slides.with(
  title: "Web API講座",
  authors: "48th irom",
)

#outline()

= HTTPの基礎

== HTTP

== HTTPリクエストとレスポンス

== HTTPヘッダー

== HTTPボディとJSON

= REST

== REST API

== リソース指向アーキテクチャ(ROA)

== RESTfulなシステムの制約条件

- クライアントとサーバーの分離
- ステートレス性
- キャッシュ可能性
- 階層化システム  
- コードオンデマンド
- 統一インターフェース

== クライアントとサーバーの分離

== ステートレス性

== キャッシュ可能性

== 階層化システム

== コードオンデマンド

== 統一インターフェース

= HTTPメソッドとCRUD操作

== GET

== HTTP

== POST

== PUT / PATCH

== DELETE

== 冪等性と安全性

== パスパラメータ

== クエリパラメータ

== ステータスコード

= アーキテクチャと設計

== 責務の分離

== レイヤードアーキテクチャ
- Handler(Controller)
- Service(UseCase)
- Repository

== Handler

== Service

== Repository

== データ構造
- Model(Entity)
- DTO(Data Transfer Object)

== Model

== DTO

= GoとGinによる実装

== Go言語

== Gin フレームワーク

== ルーティング

== バインディング

== バリデーション

== 依存性注入
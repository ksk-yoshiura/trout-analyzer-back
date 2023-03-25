# trout-analyzer-back

## プロジェクト概要
本プロジェクトは管理釣り場でのルアーによるトラウト釣りの記録アプリ作製プロジェクトです。
正式名称をTranazaに決めました。
このリポジトリにはSPAのバックエンド（Golang）とIaC（Terraform）が含まれています。  
実際にアプリとして使用する場合、trout-analyzer-frontを同時に起動させる必要があります。
フロントエンド  
https://github.com/ksk-yoshiura/trout-analyzer-front  

## 背景
管理釣り場でルアーによるトラウト釣りに近年ハマり、ちょうど個人アプリを作りたいと持っていたところでしたので、自分の腕を向上させるアプリを作れないかと思い、本アプリの作製を開始しました。管理釣り場におけるトラウト釣りが他と大きく異なるところは魚のいる場所を探す必要がないという点です。つまり適切なルアーを選択し、いい感じに動かせる技術が釣果への大きなウェイトを占めています。したがって、選択したルアーおよびルアーと相性のいい道具（竿、リール、ライン）、天候、釣果（釣れたのか、反応がないのかなど）を記録できて、各項目の相関がグラフで表示できれば、次の釣行への対策が練りやすいのではと考えました。

## このアプリでできること
- 釣り場、ルアー、ロッド、リール、ラインの情報が画像付きで、表示・登録・更新・削除ができます。
- ロッド、リール、ラインの組み合わせをタックルとして表示・登録・更新・削除できます。
- 釣果記録の連続登録ができ、それぞれのデータに対して表示・更新・削除ができます。
- リアルタイム釣果記録では、使用ルアー、使用タックル、ルアーの速度、深さ、天候と釣果（釣れた、ルアーに噛み付いた、追跡があった、反応なし）が記録できます。
- 釣果記録の分析ができます。2023年3月時点では、①ルアーの色と深さ、②ルアーの色とルアータイプ、③ルアーの色と天候の組み合わせに対して、釣果ごとのグラフが表示されます。


## 使用技術
- Golang 1.16
- Tarraform 1.2.4

## 起動方法
- git cloneします。
- ルート配下に.envを設置して、.envにはDB、minio（S3）、AWS（minio用）を記載してください。
```
DB_USER=root
DB_PASSWORD=pass
DB_HOST=db
DB_PORT=3306
DB_DATABASE_NAME=trout_analyzer

S3_ENDPOINT=http://minio:9000
S3_BUCKET=trout-analyzer-upload
S3_REGION=ap-northeast-1

AWS_ACCESS_KEY_ID=minio
AWS_SECRET_ACCESS_KEY=minio123
AWS_DEFAULT_REGION=ap-northeast-1
AWS_BUCKET=develop
AWS_PATH_STYLE_ENDPOINT=true

```
- ` cd trout-analyzer-back `で移動してください。
- `docker-compose up -d`で起動してください。
- ホットリロードは、`docker-compose exec app bash`後に` air `を打つと動作します。  
こうすると、保存の度にリロードしてくれるので開発にとても便利です。

## CI/CD
- ` git push `の度にユニットテストが走り、テストが通ると自動的にデプロイされるように、go.ymlに記載しています。  
- 自動デプロイはecspressoで対応しています。
単体テストはまだ完備していません

## IaC
- /terroform配下に各モジュールで定義しています。
- 基本的に各モジュールにてterraformコマンドを実行してインフラ構築を行います。
- 例えば、appモジュールの場合、` cd app/tranaza `まで移動して、` terroform init `後に、` terraform plan `を実行して問題がなければ` terraform apply `を実行します。
- 構築したインフラを消したい場合は` terraform destroy `を実行します。

## DB
- AWS上でDBを準備する場合、AWS CLIでEXECコマンド`aws ecs execute-command 以下略`を用いて、db/migration配下のSQLをgolang-migrateコマンドでDBを構築します。
- 000001_init_schema.up.sqlにはテーブル各種と初期データが記載されています。
- 000002_init_schema.up.sqlにはindex各種が定義されています。全SQLをEXPLAIN済みです。

## 今後の改修予定
- ページネーション  
データ数が全項目中でおそらく最大になるルアーも100か200程度なので後回し中でした。
- 単体テスト  
勉強中です。

## アプリ作成の感想
- 何度か釣行テストを行ったところ、釣りの最中にスマホを起動して記録を取るという行為がとても煩わしい上にバッテリーをとても消費するということがわかりました。
- 今回構築したインフラでは月額100ドル近くかかりました。もっと安くできるかも知れませんが、とりあえず及第点なのかも知れません。
- 2023年3月に構築したインフラを解体しアプリを停止しました。
- 釣りアプリとしてリアルタイムで釣果を記録できるという案は他に類を見ないのではと個人的に思っていますが、実際に使用してみてあまり実用的じゃないと思いました。
- ただ、個人開発として得られたものが凄まじいので、次回に活かそうと思います。
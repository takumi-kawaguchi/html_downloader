# HTMLダウンロードツール
CSV形式でWebページを指定することで、そのページにアクセスしHTML形式で保存できるツールです。

# 用途
- Webページ修正前のバックアップとしてページを取得する
- Webページ修正前後でページを取得し、前後差分を比較する

など、自由にお使いください。

# 使い方
1. 下記の形式でCSVを作成します  
※複数のページをダウンロードする場合は、複数行記載してください
```
[0: 保存するWebページのファイル名(HTML推奨)],[1: ダウンロードしたいページ名],[2: 格納フォルダパス]
```

2. 本ソースコードをダウンロードし、ファイル`main.go`を実行します
```
go run main.go
```

3. `HTMLファイルを格納するフォルダを入力してください`というメッセージが表示されて入力を求められるので、フォルダパスを入力してください

4. `読み込むCSVファイルのパスを入力してください`というメッセージが表示されて入力を求められるので、CSVファイルパスを入力してください

5. 入力後、CSVに設定されたWebページにアクセスし、実行時に入力したフォルダパスに`yyyyMMddHHmmss_htmldownloder`の形式でフォルダを作成し、その中にCSVで記載したフォルダパスとファイル名でHTMLを格納します(yyyyMMddHHmmssは実行時刻)

# 使用例
```
CSV記載例
↓
test.html,https://exmaple.com/,example\test

実行時入力例
↓
・HTMLファイルを格納するフォルダを入力してください: C:\Users\TestUser\DownloadDestination
・読み込むCSVファイルのパスを入力してください: C:\Users\TestUser\test.csv

結果
↓
・https://exmaple.com/にアクセスし、WebページをHTMLとしてダウンロードする
・実行時に入力したフォルダC:\Users\TestUser\DownloadDestinationにフォルダyyyyMMddHHmmss_htmldownloderを作成し、さらにCSVで指定したフォルダを作成してその中にtest.htmlの名前でHTMLを格納する
(C:\Users\TestUser\DownloadDestination\example\test\test.html)
```
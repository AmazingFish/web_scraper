## Weblioから単語の定義を取得してファイルに書き出すツール

単語一覧を読み込んで、対応する意味をweblioから取得して別ファイルに書き出します。

```
// mac
./weblio in.txt out.txt

// windows
./weblio.exe in.txt out.txt

// linux
./weblio-linux in.txt out.txt
```

第二引数の読み込みファイルは、次のように単語を改行区切りで記述してください。

```in.txt
algorithm
euphoria
mayhem
oblivion
scientist
```

尚、第二引数と第三引数に指定するファイル名は任意です。
第三引数に、存在しないファイル名を指定すると、自動でファイルを作成します。
第三引数に、既存のファイルを指定すると、上書きするので注意してください。

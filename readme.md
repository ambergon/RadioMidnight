# ukagaka_RadioMidnight 
このプラグインは、伺かユーザーなら常時起動しているであろうSSPをラジオっぽく運用できないかと考えて作成されました。<br>
深夜0時30から音楽とともにラジオ原稿と天気予報をバルーンに表示します。<br>
メリットとしては、きちんと人が集まったなら毎日知らないテキストを読めるのではないか?と思ったためです。<br>
未読テキスト中毒(造語)の疑いがあります。<br><br>

SSPの音声合成機能や[【SSP】伺かをCeVIO AIの力を借りてフルボイス化した。【Windows10】 -- 異風堂々](https://ambergonslibrary.com/ukagaka/8544/)などと組み合わせるとより楽しめるかもしれません。<br>


## 出来ること
- 曜日ごとのパーソナリティのラジオ放送
- 曜日ごとのパーソナリティにたいする葉書の送信
- ゲストパーソナリティとしての寄稿
- 天気予報の地域設定・削除
- ハンドルネーム / ラジオネームの設定


## Menu
ゴーストを右クリックからプラグイン名を選択することでメニューが出ます。<br>


## 曜日ごとのパーソナリティ
各曜日毎に人に振れるように作りました。<br>
GitHub上に簡単なテキストをアップロードしていただけたら、毎週読み込みます。<br>
詳細はReadMe_Personality.mdをご確認ください。<br>


## ゲストパーソナリティ
曜日ごとのパーソナリティがお休みの場合、もしくは更新がない場合に読み上げられるテキストです。<br>
プラグインメニューの 【寄稿する】から扱えます。<br>
曜日ごとのパーソナリティと違い、読み上げられるタイミングは他のユーザと異なるので注意が必要です。<br>

- 単発原稿をテストする。
    一部のサクラスクリプトや記号は使えないように置換されているためテスト環境を用意しました。<br>
    入力ダイアログにテキストを入力することで、実際にどのようにバルーンに表示されるか確認できます。<br>

- 単発原稿を送信する。
    テストに問題がなければ、こちらをクリックしていただければ、私のほうに届きます。<br>
    こちらは人力でこのリポジトリ上に保存されるので、反映まで時間がかかる場合があります。<br>
    また不適切なテキストの場合は載せない場合があります。<br>
    長期間残るものになるのでお気を付けください。<br>


## ラジオネームについて。
葉書の送信や寄稿時に使用されます。<br>


## 放送について
放送中はゴーストの発言を抑制します。<br>
また、BGMを設定しているので、ゴーストやSSP側の設定で、サウンドを鳴らさないのチェックを外してください。<br>


## aタグについて。
ラジオ放送内でaタグによるリンクのクリックが可能です。<br>
クリックしたものは直ぐに移動はできず、最大一つまで保持されてメニューで確認できます。<br>
メニューからそのリンクアドレスを確認したのち開けるようになっています。<br>
最低限のスパム対策ですね。<br>


## お借りしたもの
yaya.dll<br>
[Releases · YAYA-shiori/yaya-shiori · GitHub](https://github.com/YAYA-shiori/yaya-shiori/releases)<br>

proxy_ex.dll<br>
[Release SAORI : proxy_ex v1.0.2 · ukatech/csaori · GitHub](https://github.com/ukatech/csaori/releases/tag/saori_proxy_ex_v1.0.2)<br>

WebClapModoki-GAS<br>
[GitHub - dullNeko/WebClapModoki-GAS: WebClap-like system by using Google Apps Script & Google spreadsheet.](https://github.com/dullNeko/WebClapModoki-GAS)<br>

BGM<br>
[騒音のない世界](https://noiselessworld.net/music)様より[晴れの日の私](https://soundcloud.com/baron1_3/harenohi)<br>


## Other
このプログラムを利用することによるいかなる問題や損害に対して、私は責任を負いません。<br>


## Author
[ambergon](https://twitter.com/Sc_lFoxGon)












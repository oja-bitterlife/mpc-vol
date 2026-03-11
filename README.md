# mpc-vol
mpc volume controller

## Description
ほぼGeminiさんに作ってもらった、カーソルキーでmpcの音量を変えるコマンドラインユーティリティーです。

(BubbleTeaを勉強しようと思ってGeminiさんに聞きながら作ろうとしたら、いきなり完成形がでてきた(´・ω:;.:... )

mpc volumeだと音量の変化がわからないのでTUIで操作できるようにしたものです。

実際に聞きながら音量を変えられるのが便利だと思います。

![](assets/2026-03-10-20-45-36.png)

## install

```bash
go install github.com/oja-bitterlife/mpc-vol@latest
```

goのbinにパスが通っていればこれでmpc-volコマンドが使えるようになります。


## Usage

- ↑ / ↓ : 5% ずつ調整
- ← / → : 1% ずつ調整
- Enter / Esc / q : 終了


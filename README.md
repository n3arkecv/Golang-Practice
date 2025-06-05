思路：雙指標兩端夾

1. 前後兩支指標 l、r
l 從左往右、r 從右往左移動

2. 略過非英數字元
用 unicode.IsLetter 或 unicode.IsDigit 判斷

3. 比較字母
全部轉小寫（或大寫）再比；只要有一對不相等就不是回文

4. 指標相遇仍沒衝突 → 回文


行為說明：
isAlnum	透過 unicode 判斷英文字母或數字
兩段 for	分別讓 l、r 跳過非英數字
toLower	不用 strings.ToLower 可省一次轉字串；純 ASCII 拉丁字母直接運算

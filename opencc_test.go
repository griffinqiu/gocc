package opencc

import (
	"testing"
)

func assertCases(t *testing.T, s2t *OpenCC, cases map[string]string) {
	t.Helper()

	for k, v := range cases {
		str, err := s2t.Convert(k)
		if err != nil {
			t.Error(err)
		}
		if str != v {
			t.Errorf("\nExpected: %s\nActually: %s", v, str)
		}
	}
}

func TestConvert_s2t(t *testing.T) {
	cases := map[string]string{
		`我们是工农子弟兵`: `我們是工農子弟兵`,
		`从正数第 x 行到倒数第 y 行，截取多行输出文本的部分内容`:                 `從正數第 x 行到倒數第 y 行，截取多行輸出文本的部分內容`,
		`2017 年中国住房租赁市场租金规模约为 1.3 万亿元`:                   `2017 年中國住房租賃市場租金規模約爲 1.3 萬億元`,
		`香煙（英語：Cigarette），為煙草製品的一種。滑鼠是一種很常見及常用的電腦輸入設備。`:  `香煙（英語：Cigarette），為煙草製品的一種。滑鼠是一種很常見及常用的電腦輸入設備。`,
		`香菸（英語：Cigarette），為菸草製品的一種。記憶體是一種很常見及常用的電腦輸入裝置。`: `香菸（英語：Cigarette），為菸草製品的一種。記憶體是一種很常見及常用的電腦輸入裝置。`,
		`乾隆爷是谁的干爷爷？乾爷爷吗？`:                                `乾隆爺是誰的幹爺爺？乾爺爺嗎？`,
		`2021 年汽车零部件板块市值涨幅跑输乘用车板块，估值相对滞涨，主要由于市场对零部件行业存两大担忧：大宗商品、运费上涨致利润承压；全球芯片紧缺致下游排产低于预期。`: `2021 年汽車零部件板塊市值漲幅跑輸乘用車板塊，估值相對滯漲，主要由於市場對零部件行業存兩大擔憂：大宗商品、運費上漲致利潤承壓；全球芯片緊缺致下游排產低於預期。`,
		`我干什么不干你事。`: `我幹什麼不干你事。`,
	}

	s2t, _ := New("s2t")

	assertCases(t, s2t, cases)
}

func TestConvert_s2twp(t *testing.T) {
	cases := map[string]string{
		`鼠标里面的硅二极管坏了，导致光标分辨率降低。`:          `滑鼠裡面的矽二極體壞了，導致游標解析度降低。`,
		`我们在老挝的服务器的硬盘需要使用互联网算法软件解决异步的问题。`: `我們在寮國的伺服器的硬碟需要使用網際網路演算法軟體解決非同步的問題。`,
		`为什么你在床里面睡着？`:                     `為什麼你在床裡面睡著？`,
		`海内存知己`:                           `海內存知己`,
	}

	s2t, _ := New("s2twp")

	assertCases(t, s2t, cases)
}

func TestConvert_s2hk_finance(t *testing.T) {
	cases := map[string]string{
		"保证金":      "按金",
		"保證金":      "按金",
		"募集資金":     "籌集資金",
		"套利交易":     "對沖",
		"下周开始公开配售": "下週開始公開招股",
	}

	s2hk, _ := New("s2hk-finance")

	assertCases(t, s2hk, cases)
}

package opencc

import (
	"fmt"
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
		`高峰`:       `高峰`,
		`什麼`:       `什麼`,
		`讲下`:       `講吓`,
		`抬头`:       `抬頭`,
		`回流`:       `回流`,
		`1 厘 2 厘钱`: `1 厘 2 厘錢`,
		`公厘`:       `公厘`,
		`厘米`:       `厘米`,
		`恒大`:       `恒大`,
	}

	s2t, err := New("s2t")
	if err != nil {
		panic(fmt.Sprintf("Init s2t failed %s", err))
	}

	assertCases(t, s2t, cases)
}

func TestConvert_s2hk_finance(t *testing.T) {
	cases := map[string]string{
		"保证金":      "按金",
		"保證金":      "按金",
		`佣金`:       `佣金`,
		"募集資金":     "籌集資金",
		"套利交易":     "對沖",
		"下周开始公开配售": "下週開始公開招股",
		"户":        "戶",
	}

	s2hk, err := New("s2hk-finance")
	if err != nil {
		panic(fmt.Sprintf("Init s2hk-finance failed %s", err))
	}

	assertCases(t, s2hk, cases)
}

// Special hotfix in this project
func TestSelfSpecialHotfix(t *testing.T) {
	cc, _ := New("s2hk")

	cases := map[string]string{
		"来自于汇丰，以及汇丰银行，汇入的款项": "來自於滙豐，以及滙豐銀行，匯入的款項",
		"汇业银行集团": "滙業銀行集團",
	}

	assertCases(t, cc, cases)
}

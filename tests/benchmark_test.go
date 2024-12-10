package opencc

import (
	"testing"

	"github.com/longbridge/opencc"
)

func BenchmarkConvert_s2t_short(b *testing.B) {
	cc, _ := opencc.New("s2t")

	in := `自有渠道的会员客户粘性较强，对品牌的忠诚度高，且大量客源与数据的沉淀有利于公司内部的资源共享和高效运营。同时其建设投资成本较高，会员培养也需要一定的时间，由此可见自有渠道直销更加适合华住这种连锁酒店集团。`

	// 0.04 ms/op
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_s2t(b *testing.B) {
	cc, _ := opencc.New("s2t")

	in := readFile("html-raw.txt")

	// 2.5 ms/op
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_t2s(b *testing.B) {
	cc, _ := opencc.New("t2s")

	in := readFile("html-s2t.txt")

	// 2.3 ms/op
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_t2s_short(b *testing.B) {
	cc, _ := opencc.New("t2s")

	in := `自有渠道的會員客戶粘性較強，對品牌的忠誠度高，且大量客源與數據的沉澱有利於公司內部的資源共享和高效運營。同時其建設投資成本較高，會員培養也需要一定的時間，由此可見自有渠道直銷更加適合華住這種連鎖酒店集團。`

	// 0.035 ms/op
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_s2tw(b *testing.B) {
	cc, _ := opencc.New("s2tw")

	in := readFile("html-raw.txt")

	// 3.8 ms/op
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_s2tw_short(b *testing.B) {
	cc, _ := opencc.New("s2tw")

	in := `自有渠道的会员客户粘性较强，对品牌的忠诚度高，且大量客源与数据的沉淀有利于公司内部的资源共享和高效运营。同时其建设投资成本较高，会员培养也需要一定的时间，由此可见自有渠道直销更加适合华住这种连锁酒店集团。`

	// 0.04 ms/op
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_s2hk_finance_short(b *testing.B) {
	cc, _ := opencc.New("s2hk-finance")

	in := `自有渠道的会员客户粘性较强，对品牌的忠诚度高，且大量客源与数据的沉淀有利于公司内部的资源共享和高效运营。同时其建设投资成本较高，会员培养也需要一定的时间，由此可见自有渠道直销更加适合华住这种连锁酒店集团。`

	// 0.04 ms/op
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_s2hk_finance(b *testing.B) {
	cc, _ := opencc.New("s2hk-finance")

	in := readFile("html-raw.txt")

	// 5 ms/op
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

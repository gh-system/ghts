/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package x32_http

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	xing_http "github.com/ghts/ghts/xing_http/go"
	"math"
	"testing"
	"time"
)

func TestF계좌번호_관련_함수(t *testing.T) {
	t.Parallel()

	계좌번호_모음, 에러 := xing_http.F계좌번호_모음()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(계좌번호_모음) > 0)

	//계좌_상세명, 에러 := xing_http.F계좌_상세명(계좌번호_모음[0])
	//lib.F테스트_에러없음(t, 에러)
	//lib.F테스트_참임(t, len(strings.TrimSpace(계좌_상세명)) != 0)
}

func TestCSPAQ12200_현물계좌_총평가(t *testing.T) {
	t.Parallel()

	계좌번호 := 계좌번호_모음[0]

	값, 에러 := xing_http.TrCSPAQ12200_현물계좌_총평가(계좌번호)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 값.M계좌번호, 계좌번호)
	lib.F테스트_다름(t, 값.M계좌명, "")
	//lib.F테스트_참임(t, 값.M현금주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M출금가능금액 >= 0)
	lib.F테스트_참임(t, 값.M코스피_금액 >= 0)
	lib.F테스트_참임(t, 값.M코스닥_금액 >= 0)
	lib.F테스트_참임(t, 값.M잔고평가금액 >= 0)
	lib.F테스트_참임(t, 값.M미수금액 >= 0)
	lib.F테스트_참임(t, 값.M예탁자산총액 >= 0)
	lib.F테스트_참임(t, 값.M투자원금 >= 0)
	lib.F테스트_참임(t, 값.M손익율*float64(값.M투자손익금액) >= 0)
	lib.F테스트_참임(t, 값.M신용담보주문금액 >= 0)
	lib.F테스트_참임(t, 값.M예수금 >= 0, 값.M예수금)
	lib.F테스트_참임(t, 값.D1예수금 >= 0)
	lib.F테스트_참임(t, 값.D2예수금 >= 0)
	lib.F테스트_참임(t, 값.M대용금액 >= 0)
	lib.F테스트_참임(t, 값.M현금미수금액 >= 0)
	lib.F테스트_참임(t, 값.M증거금_현금 >= 0)
	lib.F테스트_참임(t, 값.M증거금_대용 >= 0)
	lib.F테스트_참임(t, 값.M수표_금액 >= 0)
	lib.F테스트_참임(t, 값.M대용주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M증거금률100퍼센트주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M증거금률50퍼센트주문가능금액 >= 값.M증거금률100퍼센트주문가능금액)
	lib.F테스트_참임(t, 값.M증거금률35퍼센트주문가능금액 >= 값.M증거금률50퍼센트주문가능금액)
	lib.F테스트_참임(t, 값.M전일매도정산금액 >= 0)
	lib.F테스트_참임(t, 값.M전일매수정산금액 >= 0)
	lib.F테스트_참임(t, 값.M금일매도정산금액 >= 0)
	lib.F테스트_참임(t, 값.M금일매수정산금액 >= 0)
	lib.F테스트_참임(t, 값.D1연체변제소요금액 >= 0)
	lib.F테스트_참임(t, 값.D2연체변제소요금액 >= 0)
	lib.F테스트_참임(t, 값.D1추정인출가능금액 >= 0)
	lib.F테스트_참임(t, 값.D2추정인출가능금액 >= 0)
	lib.F테스트_참임(t, 값.M예탁담보대출금액 >= 0)
	lib.F테스트_같음(t, 값.M신용설정보증금, 0)
	lib.F테스트_같음(t, 값.M융자금액, 0)
	lib.F테스트_같음(t, 값.M변경후담보비율, 0.0)
	lib.F테스트_같음(t, 값.M담보부족금액, 0)
	lib.F테스트_참임(t, 값.M원담보금액 >= 0)
	lib.F테스트_참임(t, 값.M부담보금액 >= 0)
	lib.F테스트_같음(t, 값.M소요담보금액, 0)
	lib.F테스트_참임(t, 값.M원담보부족금액 >= 0)
	lib.F테스트_같음(t, 값.M소요담보금액, 0)
	lib.F테스트_참임(t, 값.M추가담보현금 >= 0)
	lib.F테스트_참임(t, 값.D1주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M신용이자미납금액 >= 0)
	lib.F테스트_참임(t, 값.M기타대여금액 >= 0)
	lib.F테스트_참임(t, 값.M익일추정반대매매금액 >= 0)
	lib.F테스트_같음(t, 값.M원담보합계금액, 0)
	lib.F테스트_참임(t, 값.M신용주문가능금액 >= 0)
	lib.F테스트_같음(t, 값.M부담보합계금액, 0)
	lib.F테스트_같음(t, 값.M신용담보금현금, 0)
	lib.F테스트_같음(t, 값.M신용담보대용금액, 0)
	lib.F테스트_참임(t, 값.M추가신용담보현금 >= 0)
	lib.F테스트_같음(t, 값.M신용담보재사용금액, 0)
	lib.F테스트_참임(t, 값.M추가신용담보대용 >= 0)
	lib.F테스트_같음(t, 값.M매도대금담보대출금액, 0)
	lib.F테스트_같음(t, 값.M처분제한금액, 0)
}

func TestCSPAQ12300_현물계좌_잔고내역_조회(t *testing.T) {
	t.Parallel()

	const 수량 = 5
	const 가격_정상주문 = int64(0) // 현재가 주문은 가격이 0
	const 호가_유형 = lib.P호가_시장가

	계좌번호 := 계좌번호_모음[0]

	값_모음, 에러 := xing_http.TrCSPAQ12300_현물계좌_잔고내역_조회(계좌번호, xt.CSPAQ12300_평균_단가, false)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		// TODO : lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드)
		lib.F테스트_다름(t, 값.M종목명, "")
		lib.F테스트_다름(t, 값.M유가증권잔고유형코드, "")
		lib.F테스트_다름(t, 값.M유가증권잔고유형명, "")
		lib.F테스트_참임(t, 값.M잔고수량 >= 0)
		lib.F테스트_참임(t, 값.M매매기준잔고수량 >= 0)
		lib.F테스트_참임(t, 값.M금일매수체결수량 >= 0)
		lib.F테스트_참임(t, 값.M금일매도체결수량 >= 0)
		lib.F테스트_참임(t, 값.M매도가 >= 0)
		lib.F테스트_참임(t, 값.M매수가 >= 0)

		// 모의투자 세금은 0.25%, 수수료는 0.35% 입니다.
		// 0.6% 로 계산시 해당 값이 나옵니다.
		매도손익금액 := (값.M매도가 - 값.M매수가) * float64(값.M매매기준잔고수량)

		lib.F테스트_참임(t, math.Abs(float64(값.M매도손익금액)-매도손익금액) <= 1, 값.M매도손익금액, 매도손익금액)
		lib.F테스트_참임(t, 값.M손익율*float64(값.M평가손익) >= 0)
		lib.F테스트_참임(t, 값.M현재가 > 0)
		lib.F테스트_참임(t, 값.M신용금액 >= 0)
		lib.F테스트_참임(t, 값.M만기일.Equal(time.Time{}) || 값.M만기일.After(xt.F당일()))
		lib.F테스트_참임(t, 값.M전일매도체결가 >= 0)
		lib.F테스트_참임(t, 값.M전일매도수량 >= 0)
		lib.F테스트_참임(t, 값.M전일매수체결가 >= 0)
		lib.F테스트_참임(t, 값.M전일매수수량 >= 0)
		lib.F테스트_참임(t, 값.M대출일.Equal(time.Time{}) || 값.M대출일.Equal(xt.F당일()))
		lib.F테스트_참임(t, 값.M평균단가 >= 0)
		lib.F테스트_참임(t, 값.M매도가능수량 >= 0)
		lib.F테스트_참임(t, 값.M매도주문수량 >= 0)
		lib.F테스트_참임(t, 값.M금일매수체결금액 >= 0)
		lib.F테스트_참임(t, 값.M금일매도체결금액 >= 0)
		lib.F테스트_참임(t, 값.M전일매수체결금액 >= 0)
		lib.F테스트_참임(t, 값.M전일매도체결금액 >= 0)
		lib.F테스트_참임(t, 값.M잔고평가금액 > 0)

		// 평가손익 = 매입금액 - 평가금액 : 어렵다...
		//if ETF_ETN_종목_여부(값.M종목코드) { // 세율 0%, 모의서버 수수료 0.35%, 합계 0.35%
		//	//오차 := 값.M잔고평가금액 - int64(float64(값.M매입금액)*1.0035) - 값.M평가손익
		//	//lib.F테스트_참임(t, lib.F절대값_정수64(오차) <= 1, 오차, 값.M매입금액, 값.M잔고평가금액, 값.M평가손익, int64(float64(값.M매입금액)*1.0035))
		//} else { // 세율 0.25%, 모의서버 수수료 0.35%, 합계 0.6%
		//	오차 := 값.M잔고평가금액 - int64(float64(값.M매입금액)*1.006) - 값.M평가손익
		//	lib.F테스트_참임(t, lib.F절대값_정수64(오차) <= 1, 오차, 값.M평가손익, 값.M매입금액, int(float64(값.M매입금액)*1.006), 값.M잔고평가금액) // 평가손익 = 매입금액 - 평가금액
		//}

		lib.F테스트_참임(t, 값.M현금주문가능금액 >= 0)
		lib.F테스트_참임(t, 값.M주문가능금액 >= 0)
		lib.F테스트_참임(t, 값.M매도미체결수량 >= 0)
		lib.F테스트_참임(t, 값.M매도미결제수량 >= 0)
		lib.F테스트_참임(t, 값.M매수미체결수량 >= 0)
		lib.F테스트_참임(t, 값.M매수미결제수량 >= 0)
		lib.F테스트_참임(t, 값.M미결제수량 >= 0)
		lib.F테스트_참임(t, 값.M미체결수량 >= 0)
		lib.F테스트_참임(t, 값.M전일종가 > 0)
		lib.F테스트_참임(t, 값.M매입금액 > 0)
		lib.F테스트_같음(t, 값.M등록시장코드,
			xt.CSPAQ12300_코스피, xt.CSPAQ12300_코스닥, xt.CSPAQ12300_코넥스,
			xt.CSPAQ12300_K_OTC, xt.CSPAQ12300_채권, xt.CSPAQ12300_비상장)
		lib.F테스트_같음(t, 값.M대출상세분류코드,
			xt.CSPAQ12300_대출없음, xt.CSPAQ12300_유통융자,
			xt.CSPAQ12300_자기융자, xt.CSPAQ12300_예탁주식담보융자)
		lib.F테스트_참임(t, 값.M예탁담보대출수량 >= 0)
	}
}

func TestT0167_시각_조회(t *testing.T) {
	t.Parallel()

	시각, 에러 := xing_http.TrT0167_시각_조회()

	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 시각.Year(), time.Now().Year())
	lib.F테스트_같음(t, 시각.Month(), time.Now().Month())
	lib.F테스트_같음(t, 시각.Day(), time.Now().Day())

	지금 := time.Now()
	차이 := 시각.Sub(지금)
	lib.F테스트_참임(t, 차이 > (-1*lib.P1시간) && 차이 < lib.P1시간, 시각, 지금)
}

func TestT1305_기간별_주가_조회(t *testing.T) {
	t.Parallel()

	종목코드 := "069500"
	일주월_구분 := ([]xt.T일주월_구분{xt.P일주월_일, xt.P일주월_주, xt.P일주월_월})[lib.F임의_범위_이내_정수값(0, 2)]
	var 이전_일자 time.Time

	값_모음, 에러 := xing_http.TrT1305_기간별_주가_조회(종목코드, 일주월_구분, 300)
	lib.F테스트_에러없음(t, 에러)

	lib.F테스트_참임(t, len(값_모음) > 250, len(값_모음))

	for i, 값 := range 값_모음 {
		lib.F테스트_같음(t, 종목코드, 값.M종목코드)
		lib.F테스트_같음(t, 값.M일자.Hour(), 0)
		lib.F테스트_같음(t, 값.M일자.Minute(), 0)
		lib.F테스트_같음(t, 값.M일자.Second(), 0)
		lib.F테스트_같음(t, 값.M일자.Nanosecond(), 0)
		lib.F테스트_참임(t, 값.M일자.After(이전_일자) || 값.M일자.Equal(이전_일자))
		이전_일자 = 값.M일자

		if i > 0 {
			차이 := lib.F절대값_실수(값.M일자.Sub(값_모음[i-1].M일자).Hours() / 24)

			switch 일주월_구분 {
			case xt.P일주월_일:
				lib.F테스트_참임(t, 차이 >= 1 && 차이 < 13, 종목코드, 값_모음[i-1].M일자, 값.M일자, 차이)
			case xt.P일주월_주:
				lib.F테스트_참임(t, 차이 >= 3 && 차이 < 20, 종목코드, 값_모음[i-1].M일자, 값.M일자, 차이)
			case xt.P일주월_월:
				lib.F테스트_참임(t, 차이 >= 20 && 차이 < 45, 종목코드, 값_모음[i-1].M일자, 값.M일자, 차이)
			default:
				panic(lib.New에러("예상하지 못한 일주월 구분값 : '%v'", 일주월_구분))
			}
		}

		if 값.M고가 > 0 {
			lib.F테스트_참임(t, 값.M고가 >= 값.M시가, 값.M종목코드, 값.M고가, 값.M시가)
			lib.F테스트_참임(t, 값.M고가 >= 값.M종가, 값.M종목코드, 값.M고가, 값.M종가)
			lib.F테스트_참임(t, 값.M고가 >= 값.M저가, 값.M종목코드, 값.M고가, 값.M저가)
			lib.F테스트_참임(t, 값.M저가 <= 값.M시가, 값.M종목코드, 값.M저가, 값.M시가)
			lib.F테스트_참임(t, 값.M저가 <= 값.M종가, 값.M종목코드, 값.M저가, 값.M종가)
		}

		// 저가 한계값에 대한 추정에 자신이 없어서 일단 건너뜀.
		//최소_호가단위, 에러 := lib.F최소_호가단위by종목코드(값.M종목코드)
		//lib.F테스트_에러없음(t, 에러)
		//저가_한계 :=  int64(float64(값.M고가 - 최소_호가단위) * 0.7) - 최소_호가단위
		//lib.F테스트_참임(t, 값.M저가 >= 저가_한계, 값.M저가, 저가_한계)

		switch 값.M전일대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M전일대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M전일대비등락율 >= 0, 값.M전일대비구분, 값.M전일대비등락율)
		case xt.P구분_보합:
			lib.F테스트_같음(t, 값.M전일대비등락폭, 0)
			lib.F테스트_같음(t, 값.M전일대비등락율, 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M전일대비등락폭 < 0,
				"종목코드 : '%v', 구분 : '%v', 등락폭 : '%v'",
				종목코드, 값.M전일대비구분, 값.M전일대비등락폭)

			lib.F테스트_참임(t, 값.M전일대비등락율 <= 0,
				"종목코드 : '%v', 구분 : '%v', 등락율 : '%v'",
				종목코드, 값.M전일대비구분, 값.M전일대비등락율)
		default:
			if lib.F2정수64_단순형(값.M전일대비등락폭) == 0 &&
				lib.F2실수_단순형(값.M전일대비등락율) == 0.0 {
				값.M전일대비구분 = xt.P구분_보합
			} else {
				lib.F문자열_출력("일주월 구분 : '%v', 종목코드 : '%v', 일자 : '%v', 전일대비구분 : '%v'",
					일주월_구분, 값.M종목코드, 값.M일자.Format(lib.P일자_형식), 값.M전일대비구분)
				t.FailNow()
			}
		}

		switch 값.M시가대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M시가대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M시가대비등락율 >= 0)
		case xt.P구분_보합:
			lib.F테스트_참임(t, 값.M시가대비등락폭 == 0)
			lib.F테스트_참임(t, 값.M시가대비등락율 == 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M시가대비등락폭 < 0)
			lib.F테스트_참임(t, 값.M시가대비등락율 <= 0)
		default:
			lib.F문자열_출력("일주월 구분 : '%v', 종목코드 : '%v', 일자 : '%v', 시가대비구분 : '%v'",
				일주월_구분, 값.M종목코드, 값.M일자, 값.M시가대비구분)
			t.FailNow()
		}

		switch 값.M고가대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M고가대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M고가대비등락율 > 0)
		case xt.P구분_보합:
			lib.F테스트_참임(t, 값.M고가대비등락폭 == 0)
			lib.F테스트_참임(t, 값.M고가대비등락율 == 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M고가대비등락폭 < 0)
			lib.F테스트_참임(t, 값.M고가대비등락율 < 0)
		default:
			lib.F문자열_출력("일주월 구분 : '%v', 종목코드 : '%v', 일자 : '%v', 고가대비구분 : '%v'",
				일주월_구분, 값.M종목코드, 값.M일자, 값.M고가대비구분)
			t.FailNow()
		}

		switch 값.M저가대비구분 {
		case xt.P구분_상한, xt.P구분_상승:
			lib.F테스트_참임(t, 값.M저가대비등락폭 > 0)
			lib.F테스트_참임(t, 값.M저가대비등락율 > 0)
		case xt.P구분_보합:
			lib.F테스트_참임(t, 값.M저가대비등락폭 == 0)
			lib.F테스트_참임(t, 값.M저가대비등락율 == 0)
		case xt.P구분_하한, xt.P구분_하락:
			lib.F테스트_참임(t, 값.M저가대비등락폭 < 0)
			lib.F테스트_참임(t, 값.M저가대비등락율 < 0)
		default:
			lib.F문자열_출력("일주월 구분 : '%v', 종목코드 : '%v', 일자 : '%v', 저가대비구분 : '%v'",
				일주월_구분, 값.M종목코드, 값.M일자, 값.M저가대비구분)
			t.FailNow()
		}

		lib.F테스트_참임(t, 값.M거래량 >= 0)
		lib.F테스트_참임(t, 값.M거래대금_백만 >= 0)
		//lib.F테스트_참임(t, 값.M거래_증가율)
		lib.F테스트_참임(t, 값.M체결강도 >= 0)
		lib.F테스트_참임(t, 값.M소진율 >= 0)
		lib.F테스트_참임(t, 값.M회전율 >= 0)
		//lib.F테스트_참임(t, 값.M외국인_순매수)
		//lib.F테스트_참임(t, 값.M기관_순매수)
		//lib.F테스트_참임(t, 값.M개인_순매수)
		lib.F테스트_참임(t, 값.M시가총액_백만 > 0)
	}
}

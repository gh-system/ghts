/* Copyright(C) 2015-2019년 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright(C) 2015-2019년 UnHa Kim(< unha.kim.ghts at gmail dot com >)

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

package lib

import (
	"strconv"
	"time"
)

const (
	P한국_시간대 = `Asia/Seoul`

	P일자_형식     = "2006-01-02"
	P시간_형식     = "2006-01-02 15:04:05.999999999 -0700 MST"
	P간략한_시간_형식 = "2006-01-02 15:04:05"

	P마이너스1초 = -1 * time.Second
	P10밀리초  = 10 * time.Millisecond
	P30밀리초  = 30 * time.Millisecond
	P50밀리초  = 50 * time.Millisecond
	P100밀리초 = 100 * time.Millisecond
	P300밀리초 = 300 * time.Millisecond
	P500밀리초 = 500 * time.Millisecond
	P1초     = 1 * time.Second
	P3초     = 3 * time.Second
	P5초     = 5 * time.Second
	P10초    = 10 * time.Second
	P20초    = 20 * time.Second
	P30초    = 30 * time.Second
	P40초    = 40 * time.Second
	P1분     = 1 * time.Minute
	P3분     = 3 * time.Minute
	P5분     = 5 * time.Minute
	P10분    = 10 * time.Minute
	P1시간    = time.Hour
	P1일     = 24 * time.Hour
	P무기한    = 9999 * time.Hour

	P에러_자료형 = "error"

	p프로세스ID_목록_파일명 = "pid_list.dat"

	// 자주 사용되는 정규식 표현
	P정규식_실수 = `[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?`

	// 입력 구조체 바이트 복사에 사용됨.
	P긴_공백문자열 = "         " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           " +
		"                           "

	// 입력 구조체 바이트 복사에 사용됨.
	P긴_0_문자열 = "0000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000" +
		"00000000000000000000000000000000000000000000000000"
)

const (
	P증권사_NH   = T증권사(byte('N'))
	P증권사_Xing = T증권사(byte('X'))
)

type T증권사 byte

func (v T증권사) String() string {
	switch v {
	case P증권사_NH:
		return "NH"
	case P증권사_Xing:
		return "Xing"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", int8(v))
	}
}

const (
	P양수 T부호 = 1
	P영  T부호 = 0
	P음수 T부호 = -1
)

type T부호 int8

func (v T부호) String() string {
	switch v {
	case P양수:
		return "+"
	case P영:
		return "0(zero)"
	case P음수:
		return "-"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", int8(v))
	}
}

const (
	P비교_같음 T비교 = 0
	P비교_큼  T비교 = 1
	P비교_작음 T비교 = -1
	P비교_불가 T비교 = -99
)

type T비교 int8

func (v T비교) String() string {
	switch v {
	case P비교_같음:
		return "같음"
	case P비교_큼:
		return "큼"
	case P비교_작음:
		return "작음"
	case P비교_불가:
		return "비교 불가"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", int8(v))
	}
}

const (
	JSON      = T변환(byte('J'))
	MsgPack   = T변환(byte('M'))
	Raw       = T변환(byte('R'))
	P변환형식_기본값 = MsgPack
)

type T변환 byte

func (t T변환) G검사() error {
	switch t {
	case JSON, MsgPack, Raw:
		return nil
	default:
		return New에러with출력("예상하지 못한 변환 형식. '%v'", t)
	}
}

func (t T변환) String() string {
	switch t {
	case JSON:
		return "JSON"
	case MsgPack:
		return "MsgPack"
	case Raw:
		return "Raw"
	default:
		return f포맷된_문자열("예상하지 못한 변환 형식 : '%v'", string(t))
	}
}

const (
	// 질의 메시지 구분
	P메시지_질의  = "G" // GET.
	P메시지_설정  = "S" // SET. 있으면 갱신. 없으면 생성 후 갱신.
	P메시지_삭제  = "D" // DELETE
	P메시지_종료  = "Q"
	P메시지_초기화 = "I" // 주로 테스트 할 때 사용.

	// 회신 메시지 구분
	P메시지_OK = "O"
	P메시지_에러 = "E"

	// 기타.
	P메시지_생성 = "C" // CREATE
	P메시지_읽기 = "R" // READ. 질의(GET)와 중복된다고 판단되면 삭제될 수 있음.
	P메시지_갱신 = "U" // UPDATE
)

const 포트_번호_최소값 = 3001

const (
	P주소_주소정보 T주소 = iota
	//P주소_테스트
	P주소_종목정보
	P주소_가격정보
	P주소_가격정보_입수
	P주소_가격정보_배포
	P주소_NH_TR
	P주소_NH_실시간
	P주소_NH_C함수_콜백
	P주소_NH_C함수_호출
	P주소_Xing_TR
	P주소_Xing_실시간
	P주소_Xing_C함수_콜백
	P주소_Xing_C함수_호출
	P주소_신한_TR
	P주소_신한_실시간
	P주소_신한_C함수_콜백
	P주소_신한_C함수_호출
)

type T주소 int

func (t T주소) G값() string {
	switch t {
	case P주소_주소정보: // 최소한의 고정값.
		return "tcp://127.0.0.1:3000"
	default:
		포트_번호 := strconv.Itoa(int(t) + 포트_번호_최소값)
		return "tcp://127.0.0.1:" + 포트_번호
	}
}

func (t T주소) String() string {
	return t.G이름() + " : " + t.G값()
}

func (t T주소) G이름() string {
	switch t {
	case P주소_주소정보:
		return "주소 정보"
		//case P주소_테스트:
		//	return "테스트"
	case P주소_종목정보:
		return "종목 정보"
	case P주소_가격정보:
		return "가격 정보"
	case P주소_가격정보_입수:
		return "가격 정보 입수"
	case P주소_가격정보_배포:
		return "가격 정보 배포"
	case P주소_NH_TR:
		return "NH TR"
	case P주소_NH_실시간:
		return "NH 실시간"
	case P주소_Xing_TR:
		return "Xing TR"
	case P주소_Xing_실시간:
		return "Xing 실시간"
	case P주소_Xing_C함수_호출:
		return "Xing C함수 호출"
	case P주소_Xing_C함수_콜백:
		return "Xing C함수 콜백"
	default:
		return "알려지지 않은 주소 (테스트용?)"
	}
}

const (
	P신호 T신호 = iota
	P신호_OK
	P신호_에러
	P신호_타임아웃
	P신호_초기화
	P신호_종료
)

type T신호 uint8

func (t T신호) String() string {
	switch t {
	case P신호:
		return "신호"
	case P신호_OK:
		return "OK"
	case P신호_에러:
		return "에러"
	case P신호_타임아웃:
		return "타임아웃"
	case P신호_초기화:
		return "초기화"
	case P신호_종료:
		return "종료"
	default:
		return F2문자열("예상하지 못한 신호값 : '%v'", t)
	}
}

const (
	KRW = T통화(byte('K'))
	USD = T통화(byte('U'))
	EUR = T통화(byte('E'))
	CNY = T통화(byte('C'))
)

type T통화 byte

func (v T통화) String() string {
	switch v {
	case KRW:
		return "KRW"
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	case CNY:
		return "CNY"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", byte(v))
	}
}

func (v *T통화) Parse(값 string) {
	switch 값 {
	case "KRW":
		*v = KRW
	case "USD":
		*v = USD
	case "EUR":
		*v = EUR
	case "CNY":
		*v = CNY
	default:
		panic(New에러("예상하지 못한 값 : '%v'", 값))
	}
}

//TR 및 응답 종류
const (
	TR조회 TR구분 = iota
	TR주문
	TR실시간_정보_구독
	TR실시간_정보_해지
	TR실시간_정보_일괄_해지
	TR접속
	TR접속됨
	TR접속_해제
	TR초기화
	TR종료
)

type TR구분 uint8

func (v TR구분) String() string {
	return TR구분_String(v)
}

// 증권사 API 패키지에서 오버라이드 될 수 있음.
var TR구분_String = func(v TR구분) string {
	switch v {
	case TR조회:
		return "TR조회"
	case TR주문:
		return "TR주문"
	case TR실시간_정보_구독:
		return "TR실시간 정보 구독"
	case TR실시간_정보_해지:
		return "TR실시간 정보 해지"
	case TR실시간_정보_일괄_해지:
		return "TR실시간 정보 일괄 해지"
	case TR접속:
		return "TR접속"
	case TR접속됨:
		return "TR접속됨"
	case TR접속_해제:
		return "TR접속 해제"
	case TR초기화:
		return "TR초기화"
	case TR종료:
		return "TR종료"
	default:
		return F2문자열("예상하지 못한 M값 : '%v'", uint8(v))
	}
}

const (
	TR응답_데이터 TR응답_구분 = iota
	TR응답_실시간_정보
	TR응답_메시지
	TR응답_완료
)

type TR응답_구분 int8

func (v TR응답_구분) String() string {
	switch v {
	case TR응답_데이터:
		return "TR응답_데이터"
	case TR응답_실시간_정보:
		return "TR응답_실시간_정보"
	case TR응답_메시지:
		return "TR응답_메시지"
	case TR응답_완료:
		return "TR응답_완료"
	default:
		return F2문자열("예상하지 못한 M값. %v", v)
	}
}

const (
	P시장구분_전체 T시장구분 = iota
	P시장구분_코스피
	P시장구분_코스닥
	P시장구분_코넥스
	P시장구분_ETF
	P시장구분_ETN
	P시장구분_선물옵션
	P시장구분_CME야간선물옵션
	P시장구분_EUREX야간선물옵션
)

type T시장구분 int8

func (p T시장구분) String() string {
	switch p {
	case P시장구분_전체:
		return "전체"
	case P시장구분_코스피:
		return "코스피"
	case P시장구분_코스닥:
		return "코스닥"
	case P시장구분_코넥스:
		return "코넥스"
	case P시장구분_ETF:
		return "ETF"
	case P시장구분_ETN:
		return "ETN"
	case P시장구분_선물옵션:
		return "선물옵션"
	case P시장구분_CME야간선물옵션:
		return "CME야간선물옵션"
	case P시장구분_EUREX야간선물옵션:
		return "EUREX야간선물옵션"
	default:
		return F2문자열("예상하지 못한 시장구분값 : '%v'", p)
	}
}

func (v T시장구분) Parse(값 string) error {
	switch 값 {
	case "코스피":
		v = P시장구분_코스피
	case "코스닥":
		v = P시장구분_코스닥
	case "코넥스":
		v = P시장구분_코넥스
	case "ETF":
		v = P시장구분_ETF
	case "ETN":
		v = P시장구분_ETN
	default:
		return New에러("예상하지 못한 M값. %v", 값)
	}

	return nil
}

type T매도_매수_구분 uint8

const (
	P매도매수_전체 T매도_매수_구분 = iota
	P매도
	P매수
)

func (p T매도_매수_구분) String() string {
	switch p {
	case P매도:
		return "매도"
	case P매수:
		return "매수"
	default:
		return F2문자열("예상하지 못한 값 : %v", int(p))
	}
}

func (p T매도_매수_구분) G검사() error {
	switch p {
	case P매도매수_전체, P매수, P매도:
		return nil
	default:
		return New에러("잘못된 매수 매도 구분값 : %v", int(p))
	}
}

type T체결_구분 uint8

const (
	P체결구분_전체 T체결_구분 = iota
	P체결구분_체결
	P체결구분_미체결
)

func (p T체결_구분) String() string {
	switch p {
	case P체결구분_전체:
		return "전체"
	case P체결구분_체결:
		return "체결"
	case P체결구분_미체결:
		return "미체결"
	default:
		panic(New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

const (
	P조정시가 T가격데이터_구분 = iota
	P조정종가
)

type T가격데이터_구분 uint8

func (v T가격데이터_구분) String() string {
	switch v {
	case P조정시가:
		return "조정시가"
	case P조정종가:
		return "조정종가"
	default:
		return F2문자열("잘못된 구분값. %v", v)
	}
}

const (
	P장_개시 T장_개시_종료 = iota
	P장_종료
)

type T장_개시_종료 uint8

func (v T장_개시_종료) String() string {
	switch v {
	case P장_개시:
		return "장_개시"
	case P장_종료:
		return "장_종료"
	default:
		return F2문자열("잘못된 구분값. %v", v)
	}
}

const (
	P장_중 T장_정보 = iota
	P장_후_시간외
	P장_전_시간외
)

type T장_정보 uint8

func (v T장_정보) String() string {
	switch v {
	case P장_중:
		return "장_중"
	case P장_후_시간외:
		return "장_후_시간외"
	case P장_전_시간외:
		return "장_전_시간외"
	default:
		return F2문자열("잘못된 장 정보 M값. %v", v)
	}
}

type T정렬_구분 uint8

const (
	P정렬구분_해당없음 T정렬_구분 = iota
	P정렬_정순
	P정렬_역순
)

func (p T정렬_구분) String() string {
	switch p {
	case P정렬구분_해당없음:
		return "해당 없음"
	case P정렬_정순:
		return "정순"
	case P정렬_역순:
		return "역순"
	default:
		return F2문자열("예상하지 못한 M값. %v", int(p))
	}
}

type T신규_정정_취소 int8

const (
	P신규 T신규_정정_취소 = iota
	P정정
	P취소
)

func (v T신규_정정_취소) String() string {
	switch v {
	case P신규:
		return "신규"
	case P정정:
		return "정정"
	case P취소:
		return "취소"
	default:
		return F2문자열("예상하지 못한 M값. %v", v)
	}
}

func (v T신규_정정_취소) G검사() error {
	switch v {
	case P신규, P정정, P취소:
		return nil
	}

	return New에러with출력("잘못된 신규 정정 취소 구분값. %v", v)
}

const (
	P주문응답_정상 T주문응답_구분 = iota
	P주문응답_정정
	P주문응답_취소
	P주문응답_거부
	P주문응답_체결
	P주문응답_IOC취소
	P주문응답_FOK취소
)

type T주문응답_구분 int8

func (v T주문응답_구분) G검사() error {
	switch v {
	case P주문응답_정상, P주문응답_정정, P주문응답_취소,
		P주문응답_거부, P주문응답_체결,
		P주문응답_IOC취소, P주문응답_FOK취소:
		return nil
	default:
		return New에러("잘못된 주문응답 구분값. %v", v)
	}
}

func (v T주문응답_구분) String() string {
	switch v {
	case P주문응답_정상:
		return "접수"
	case P주문응답_정정:
		return "정정"
	case P주문응답_취소:
		return "취소"
	case P주문응답_거부:
		return "거부"
	case P주문응답_체결:
		return "체결"
	case P주문응답_IOC취소:
		return "IOC 취소"
	case P주문응답_FOK취소:
		return "FOC 취소"
	default:
		return F2문자열("잘못된 주문응답 구분값. %v", v)
	}
}

const (
	P호가_지정가 T호가유형 = iota
	P호가_시장가
	P호가_조건부_지정가
	P호가_최유리_지정가
	P호가_최우선_지정가
	P호가_장전_시간외
	P호가_장후_시간외
	P호가_시간외_단일가
	P호가_해당없음
)

type T호가유형 int8

func (v T호가유형) String() string {
	switch v {
	case P호가_지정가:
		return "지정가"
	case P호가_시장가:
		return "시장가"
	case P호가_조건부_지정가:
		return "조건부 지정가"
	case P호가_최유리_지정가:
		return "최유리 지정가"
	case P호가_최우선_지정가:
		return "최우선 지정가"
	case P호가_장전_시간외:
		return "장전 시간외"
	case P호가_장후_시간외:
		return "장후 시간외"
	case P호가_시간외_단일가:
		return "시간외 단일가"
	case P호가_해당없음:
		return "해당없음"
	default:
		return F2문자열("예상하지 못한 M값. %v", v)
	}
}

func (v T호가유형) G검사() {
	switch v {
	case P호가_지정가, P호가_시장가, P호가_조건부_지정가,
		P호가_최유리_지정가, P호가_최우선_지정가,
		P호가_장전_시간외, P호가_장후_시간외,
		P호가_시간외_단일가, P호가_해당없음:
		return
	default:
		panic(New에러("잘못된 지정가 시장가 구분값. %v", v))
	}
}

const (
	P신용거래_해당없음 T신용거래_구분 = iota
	P신용거래_유통융자신규
	P신용거래_자기융자신규
	P신용거래_유통대주신규
	P신용거래_자기대주신규
	P신용거래_유통융자상환
	P신용거래_자기융자상환
	P신용거래_유통대주상환
	P신용거래_자기대주상환
	P신용거래_예탁담보대출상환
	P신용거래_청약대출상환
	P신용거래_보통대출상환
	P신용거래_매입대출신규
	P신용거래_매입대출상환
)

type T신용거래_구분 int8

func (v T신용거래_구분) String() string {
	switch v {
	case P신용거래_해당없음:
		return "해당없음"
	case P신용거래_유통융자신규:
		return "유통융자신규"
	case P신용거래_자기융자신규:
		return "자기융자신규"
	case P신용거래_유통대주신규:
		return "유통대주신규"
	case P신용거래_자기대주신규:
		return "자기대주신규"
	case P신용거래_유통융자상환:
		return "유통융자상환"
	case P신용거래_자기융자상환:
		return "자기융자상환"
	case P신용거래_유통대주상환:
		return "유통대주상환"
	case P신용거래_자기대주상환:
		return "자기대주상환"
	case P신용거래_예탁담보대출상환:
		return "예탁담보대출상환"
	default:
		return F2문자열("예상하지 못한 M값. %v", v)
	}
}

func (v T신용거래_구분) G검사() {
	switch v {
	case P신용거래_해당없음, P신용거래_유통융자신규, P신용거래_자기융자신규, P신용거래_유통대주신규,
		P신용거래_자기대주신규, P신용거래_유통융자상환, P신용거래_자기융자상환, P신용거래_유통대주상환:
		return
	default:
		panic(New에러("잘못된 신용거래 구분값. %v", v))
	}
}

const (
	P주문조건_없음 T주문조건 = iota
	P주문조건_IOC
	P주문조건_FOK
)

type T주문조건 int8

func (v T주문조건) String() string {
	switch v {
	case P주문조건_없음:
		return "없음"
	case P주문조건_IOC:
		return "IOC"
	case P주문조건_FOK:
		return "FOK"
	default:
		return F2문자열("잘못된 주문조건 구분값. %v", v)
	}
}

func (v T주문조건) G검사() {
	switch v {
	case P주문조건_없음, P주문조건_IOC, P주문조건_FOK:
		return
	default:
		panic(New에러("잘못된 주문 조건 구분값. %v", v))
	}
}

const (
	P포지션_롱 = T포지션(byte('L'))
	P포지션_숏 = T포지션(byte('S'))
)

type T포지션 byte

func (v T포지션) String() string {
	switch v {
	case P포지션_롱:
		return "롱(매수)"
	case P포지션_숏:
		return "숏(매도)"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", byte(v))
	}
}

type T소켓_접속방식 uint

const (
	P소켓_접속_BIND T소켓_접속방식 = iota
	P소켓_접속_CONNECT
)

func (t T소켓_접속방식) String() string {
	switch t {
	case P소켓_접속_BIND:
		return "BIND"
	case P소켓_접속_CONNECT:
		return "CONNECT"
	default:
		return "예상하지 못한 접속방식 : '" + t.String() + "'"
	}
}

type T소켓_종류 uint8

const (
	P소켓_종류_REQ = iota
	P소켓_종류_REP
	P소켓_종류_DEALER
	P소켓_종류_ROUTER
	P소켓_종류_PUB
	P소켓_종류_SUB
	P소켓_종류_PUSH
	P소켓_종류_PULL
	P소켓_종류_PAIR
)

func (t T소켓_종류) String() string {
	switch t {
	case P소켓_종류_REQ:
		return "REQ"
	case P소켓_종류_REP:
		return "REP"
	case P소켓_종류_DEALER:
		return "DEALER"
	case P소켓_종류_ROUTER:
		return "ROUTER"
	case P소켓_종류_PUB:
		return "PUB"
	case P소켓_종류_SUB:
		return "SUB"
	case P소켓_종류_PUSH:
		return "PUSH"
	case P소켓_종류_PULL:
		return "PULL"
	case P소켓_종류_PAIR:
		return "PAIR"
	default:
		return F2문자열("예상하지 못한 소켓 종류 : '%v'", t)
	}
}

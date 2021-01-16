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

package x32

import (
	"encoding/json"
	"fmt"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

func F질의값_추출_RT처리(w http.ResponseWriter, req *http.Request) {
	var 질의값 *lib.S질의값_복수_종목

	if 에러 := F질의값_추출(req, 질의값); 에러 != nil {
		F회신(w, New응답(에러))
	} else {
		F질의_처리(w, 질의값)
	}
}

func F질의값_추출_TR처리(w http.ResponseWriter, req *http.Request, TR구분 lib.TR구분, TR코드 string, 질의값 lib.I질의값) {
	if lib.F종류(질의값) != reflect.Ptr {
		lib.New에러with출력("포인터형이 아님. %T", 질의값)
		return
	} else if 에러 := F질의값_추출(req, 질의값); 에러 != nil {
		F회신(w, New응답(에러))
		return
	} else {
		질의값.S설정(TR구분, TR코드)
		F질의_처리(w, 질의값)
	}
}

func F문자열_추출(req *http.Request) (문자열 string, 에러 error) {
	if 바이트_모음, 에러 := ioutil.ReadAll(req.Body); 에러 != nil {
		return "", 에러
	} else {
		return lib.F앞뒤_따옴표_제거(string(바이트_모음)), nil
	}
}

func F질의값_추출(req *http.Request, ptr질의값 interface{}) (에러 error) {
	if lib.F종류(ptr질의값) != reflect.Ptr {
		return lib.New에러with출력("포인터형이 아님. %T", ptr질의값)
	} else if 바이트_모음, 에러 := ioutil.ReadAll(req.Body); 에러 != nil {
		return 에러
	} else if 에러 = json.Unmarshal(바이트_모음, ptr질의값); 에러 != nil {
		return 에러
	}

	return nil
}

func F질의_처리(w http.ResponseWriter, 질의값 lib.I질의값) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	var ch회신 chan interface{}

	if 응답 := f질의_처리_도우미(w, 질의값); 응답.Error() != nil {
		return F회신(w, 응답)
	} else if 식별번호, ok := 응답.V.(int); !ok {
		return F회신(w, New응답(lib.New에러("%v : 예상하지 못한 자료형. %T", 질의값.TR코드(), 응답.V)))
	} else {
		ch회신 = 콜백_대기소.S추가(식별번호, 질의값.TR코드())
	}

	select {
	case 값 := <-ch회신:
		switch 변환값 := 값.(type) {
		case error:
			if strings.Contains(변환값.Error(), "XingAPI에 접속되어 있지 않습니다.") {
				lib.F문자열_출력("[%v] XingAPI 재접속 시도.", lib.F지금().Format(lib.P간략한_시간_형식))
				panic("TODO : C32_재시작()") // TODO : C32_재시작()
				lib.F대기(lib.P10초)

				return F질의_처리(w, 질의값)
			} else if !strings.Contains(변환값.Error(), "주문이 접수 대기") &&
				!strings.Contains(변환값.Error(), "원주문번호를 잘못 입력") &&
				!strings.Contains(변환값.Error(), "취소 가능한 수량을 초과하였습니다.") &&
				!strings.Contains(변환값.Error(), "주문수량이 매매가능수량을 초과했습니다") {
				fmt.Println("*********************************************************")
				lib.New에러with출력(변환값)
				fmt.Println("*********************************************************")
			}

			return F회신(w, New응답(변환값))
		default:
			//lib.F문자열_출력("%v 회신값 자료형 '%T'", 질의값.TR코드(), 변환값)

			return F회신(w, New응답(변환값))
		}
	case <-time.After(lib.P1분):
		return F회신(w, New응답(lib.New에러("F질의_처리() 타임아웃. '%v' '%v'", 질의값.TR코드(), 질의값)))
	case <-lib.Ch공통_종료():
		return nil
	}
}

func f질의_처리_도우미(w http.ResponseWriter, 질의값 lib.I질의값) (응답 *S응답) {
	defer lib.S예외처리{}.S실행()

	switch 질의값.TR구분() {
	case xt.TR조회, xt.TR주문:
		f전송_권한_획득(질의값.TR코드())

		defer f전송_시각_기록(질의값.TR코드())
	}

	select {
	case 응답 = <-New질의(질의값, Ch질의).Ch응답:
		return 응답
	case <-time.After(lib.P30초):
		return New응답(lib.New에러("f질의_처리_도우미() 타임아웃 %v %v", 질의값.TR구분(), 질의값.TR코드()))
	}
}

func f전송_권한_획득(TR코드 string) {
	switch TR코드 {
	case "", xt.RT현물_주문_접수_SC0, xt.RT현물_주문_체결_SC1, xt.RT현물_주문_정정_SC2, xt.RT현물_주문_취소_SC3, xt.RT현물_주문_거부_SC4,
		xt.RT코스피_호가_잔량_H1, xt.RT코스피_시간외_호가_잔량_H2, xt.RT코스피_체결_S3, xt.RT코스피_예상_체결_YS3,
		xt.RT코스피_ETF_NAV_I5, xt.RT주식_VI발동해제_VI, xt.RT시간외_단일가VI발동해제_DVI, xt.RT장_운영정보_JIF:
		return
	}

	f10분당_전송_제한_확인(TR코드)

	if f1초_1회_미만_전송_제한_확인(TR코드) == nil {
		f초당_전송_제한_확인(TR코드)
	}
}

func f1초_1회_미만_전송_제한_확인(TR코드 string) lib.I전송_권한 {
	전송_권한, 존재함 := tr코드별_전송_제한_초당_1회_미만[TR코드]

	switch {
	case !존재함:
		return nil // 해당 TR코드 관련 제한이 존재하지 않음.
	case 전송_권한.TR코드() != TR코드:
		panic("예상하지 못한 경우.")
	}

	return 전송_권한.G획득()
}

func f초당_전송_제한_확인(TR코드 string) lib.I전송_권한 {
	전송_권한, 존재함 := tr코드별_전송_제한_1초[TR코드]

	switch {
	case !존재함:
		panic(lib.New에러("전송제한을 찾을 수 없음 : '%v'", TR코드))
	case 전송_권한.TR코드() != TR코드:
		panic("예상하지 못한 경우.")
	case 전송_권한.G남은_수량() > 100:
		panic("전송 한도가 너무 큼. 1초당 한도와 10분당 한도를 혼동한 듯함.")
	}

	return 전송_권한.G획득()
}

func f10분당_전송_제한_확인(TR코드 string) lib.I전송_권한 {
	전송_권한, 존재함 := tr코드별_전송_제한_10분[TR코드]

	switch {
	case !존재함:
		return nil // 해당 TR코드 관련 제한이 존재하지 않음.
	case 전송_권한.TR코드() != TR코드:
		panic("예상하지 못한 경우.")
	}

	return 전송_권한.G획득()
}

func f전송_시각_기록(TR코드 string) {
	// 10분당 전송 제한 기록
	if 전송_권한, 존재함 := tr코드별_전송_제한_10분[TR코드]; 존재함 {
		전송_권한.S기록()
	}

	// 초당 전송 제한 기록
	if 전송_권한, 존재함 := tr코드별_전송_제한_1초[TR코드]; 존재함 {
		전송_권한.S기록()
	}
}

func F회신(w http.ResponseWriter, 값 *S응답) (에러 error) {
	바이트_모음, 에러 := json.Marshal(값)
	if 에러 != nil {
		return 에러
	}

	_, 에러 = w.Write(바이트_모음)

	return 에러
}

func XingAPI디렉토리() (string, error) {
	파일경로, 에러 := lib.F실행파일_검색(xing_dll)
	if 에러 == nil {
		return lib.F디렉토리명(파일경로)
	}

	기본_위치 := `C:\eBEST\xingAPI\xingAPI.dll`
	if _, 에러 := os.Stat(기본_위치); 에러 == nil {
		lib.F실행경로_추가(기본_위치)

		if _, 에러 := lib.F실행파일_검색(xing_dll); 에러 != nil {
			return "", lib.New에러("실행경로에 추가시켰으나 여전히 찾을 수 없음.")
		}

		return lib.F디렉토리명(기본_위치)
	}

	파일경로, 에러 = lib.F파일_검색(`C:\`, xing_dll)
	if 에러 == nil {
		lib.F실행경로_추가(파일경로)

		if _, 에러 := lib.F실행파일_검색(xing_dll); 에러 != nil {
			return "", lib.New에러("실행경로에 추가시켰으나 여전히 찾을 수 없음.")
		}

		return lib.F디렉토리명(파일경로)
	}

	return "", lib.New에러("DLL파일을 찾을 수 없습니다.")
}

func f자료형_문자열_해석(g *xt.TR_DATA) (자료형_문자열 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 자료형_문자열 = "" }}.S실행()

	TR코드 := lib.F2문자열_공백제거(g.TrCode)
	길이 := lib.F2정수_단순형(g.DataLength)

	switch TR코드 {
	case xt.TR현물계좌_총평가_CSPAQ12200:
		switch 길이 {
		case xt.SizeCSPAQ12200OutBlock1:
			return xt.P자료형_CSPAQ12200OutBlock1, nil
		case xt.SizeCSPAQ12200OutBlock2:
			return xt.P자료형_CSPAQ12200OutBlock2, nil
		case xt.SizeCSPAQ12200OutBlock1 + xt.SizeCSPAQ12200OutBlock2:
			return xt.P자료형_CSPAQ12200OutBlock, nil
		}
	case xt.TR현물계좌_잔고내역_조회_CSPAQ12300:
		const 헤더_길이 = xt.SizeCSPAQ12300OutBlock1 + xt.SizeCSPAQ12300OutBlock2 + 5
		// Non-block 모드는 Occurs데이터 수량을 나타내는 5바이트 추가됨.
		if 길이 == 0 {
			return xt.P자료형_nil, nil
		} else if 길이 < 헤더_길이 || (길이-헤더_길이)%xt.SizeCSPAQ12300OutBlock3 != 0 {
			break
		}

		return xt.P자료형_CSPAQ12300OutBlock, nil
	case xt.TR현물계좌_주문체결내역_조회_CSPAQ13700:
		const 헤더_길이 = xt.SizeCSPAQ13700OutBlock1 + xt.SizeCSPAQ13700OutBlock2 + 5
		// Non-block 모드는 Occurs데이터 수량을 나타내는 5바이트 추가됨.
		if 길이 == 0 {
			return xt.P자료형_nil, nil
		} else if 길이 < 헤더_길이 || (길이-헤더_길이)%xt.SizeCSPAQ13700OutBlock3 != 0 {
			break
		}

		return xt.P자료형_CSPAQ13700OutBlock, nil
	case xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200:
		switch 길이 {
		case xt.SizeCSPAQ22200OutBlock1:
			return xt.P자료형_CSPAQ22200OutBlock1, nil
		case xt.SizeCSPAQ22200OutBlock2:
			return xt.P자료형_CSPAQ22200OutBlock2, nil
		case xt.SizeCSPAQ22200OutBlock1 + xt.SizeCSPAQ22200OutBlock2:
			return xt.P자료형_CSPAQ22200OutBlock, nil
		}
	case xt.TR현물_정상_주문_CSPAT00600:
		if 길이 == xt.SizeCSPAT00600OutBlock {
			return xt.P자료형_CSPAT00600OutBlock, nil
		}
	case xt.TR현물_정정_주문_CSPAT00700:
		if 길이 == xt.SizeCSPAT00700OutBlock {
			return xt.P자료형_CSPAT00700OutBlock, nil
		}
	case xt.TR현물_취소_주문_CSPAT00800:
		if 길이 == xt.SizeCSPAT00800OutBlock {
			return xt.P자료형_CSPAT00800OutBlock, nil
		}
	case xt.TR현물_당일_매매일지_t0150:
		const 헤더_길이 = xt.SizeT0150OutBlock + 5
		// Non-block 모드는 Occurs데이터 수량을 나타내는 5바이트 추가됨.
		if 길이 == 0 {
			return xt.P자료형_nil, nil
		} else if 길이 < 헤더_길이 || (길이-헤더_길이)%xt.SizeT0150OutBlock1 != 0 {
			break
		}

		return xt.P자료형_T0150_현물_당일_매매일지_응답, nil
	case xt.TR현물_일자별_매매일지_t0151:
		const 헤더_길이 = xt.SizeT0151OutBlock + 5
		// Non-block 모드는 Occurs데이터 수량을 나타내는 5바이트 추가됨.
		if 길이 == 0 {
			return xt.P자료형_nil, nil
		} else if 길이 < 헤더_길이 || (길이-헤더_길이)%xt.SizeT0151OutBlock1 != 0 {
			break
		}

		return xt.P자료형_T0151_현물_일자별_매매일지_응답, nil
	case xt.TR시간_조회_t0167:
		return xt.P자료형_T0167OutBlock, nil
	case xt.TR현물_체결_미체결_조회_t0425:
		// Non-block 모드는 Occurs데이터 수량을 나타내는 5바이트 추가됨.
		if 길이 == 0 {
			return xt.P자료형_nil, nil
		} else if 길이 < (xt.SizeT0425OutBlock+5) ||
			(길이-(xt.SizeT0425OutBlock+5))%xt.SizeT0425OutBlock1 != 0 {
			break
		}

		return xt.P자료형_T0425OutBlock, nil
	case xt.TR현물_호가_조회_t1101:
		return xt.P자료형_T1101OutBlock, nil
	case xt.TR현물_시세_조회_t1102:
		return xt.P자료형_T1102OutBlock, nil
	case xt.TR현물_기간별_조회_t1305:
		switch {
		case 길이 == xt.SizeT1305OutBlock:
			return xt.P자료형_T1305OutBlock, nil
		case 길이%xt.SizeT1305OutBlock1 == 0:
			return xt.P자료형_T1305OutBlock1, nil
		}
	case xt.TR현물_당일_전일_분틱_조회_t1310:
		switch {
		case 길이 == xt.SizeT1310OutBlock:
			return xt.P자료형_T1310OutBlock, nil
		case 길이%xt.SizeT1310OutBlock1 == 0:
			return xt.P자료형_T1310OutBlock1, nil
		}
	case xt.TR관리_불성실_투자유의_조회_t1404:
		switch {
		case 길이 == xt.SizeT1404OutBlock:
			return xt.P자료형_T1404OutBlock, nil
		case 길이%xt.SizeT1404OutBlock1 == 0:
			return xt.P자료형_T1404OutBlock1, nil
		}
	case xt.TR투자경고_매매정지_정리매매_조회_t1405:
		switch {
		case 길이 == xt.SizeT1405OutBlock:
			return xt.P자료형_T1405OutBlock, nil
		case 길이%xt.SizeT1405OutBlock1 == 0:
			return xt.P자료형_T1405OutBlock1, nil
		}
	case xt.TR_ETF_시간별_추이_t1902:
		switch {
		case 길이 == xt.SizeT1902OutBlock:
			return xt.P자료형_T1902OutBlock, nil
		case 길이%xt.SizeT1902OutBlock1 == 0:
			return xt.P자료형_T1902OutBlock1, nil
		}
	case xt.TR_ETF_LP호가_조회_t1906:
		return xt.P자료형_T1906OutBlock, nil
	case xt.TR재무순위_종합_t3341:
		switch {
		case 길이 == xt.SizeT3341OutBlock:
			return xt.P자료형_T3341OutBlock, nil
		case 길이%xt.SizeT3341OutBlock1 == 0:
			return xt.P자료형_T3341OutBlock1, nil
		}
	case xt.TR현물_멀티_현재가_조회_t8407:
		switch {
		case 길이%xt.SizeT8407OutBlock1 == 0:
			return xt.P자료형_T8407OutBlock1, nil
		}
	case xt.TR현물_차트_틱_t8411:
		switch {
		case 길이 == xt.SizeT8411OutBlock:
			return xt.P자료형_T8411OutBlock, nil
		case 길이%xt.SizeT8411OutBlock1 == 0:
			return xt.P자료형_T8411OutBlock1, nil
		}
	case xt.TR현물_차트_분_t8412:
		switch {
		case 길이 == xt.SizeT8412OutBlock:
			return xt.P자료형_T8412OutBlock, nil
		case 길이%xt.SizeT8412OutBlock1 == 0:
			return xt.P자료형_T8412OutBlock1, nil
		}
	case xt.TR현물_차트_일주월_t8413:
		switch {
		case 길이 == xt.SizeT8413OutBlock:
			return xt.P자료형_T8413OutBlock, nil
		case 길이%xt.SizeT8413OutBlock1 == 0:
			return xt.P자료형_T8413OutBlock1, nil
		}
	case xt.TR증시_주변_자금_추이_t8428:
		switch {
		case 길이 == xt.SizeT8428OutBlock:
			return xt.P자료형_T8428OutBlock, nil
		case 길이%xt.SizeT8428OutBlock1 == 0:
			return xt.P자료형_T8428OutBlock1, nil
		}
	case xt.TR현물_종목_조회_t8436:
		if 길이%xt.SizeT8436OutBlock == 0 {
			return xt.P자료형_T8436OutBlock, nil
		}
	}

	panic(lib.New에러("예상하지 못한 TR코드 & 길이 : '%v' '%v'", TR코드, 길이))

	//case xt.TR선물옵션_주문체결내역조회_CFOAQ00600:
	//	// Non-block 모드는 Occurs데이터 수량을 나타내는 5바이트 추가됨.
	//	if 길이 == 0 {
	//		return xt.P자료형_nil, nil
	//	} else if 길이 < (xt.SizeCFOAQ00600OutBlock1+xt.SizeCFOAQ00600OutBlock2+5) ||
	//		(길이-(xt.SizeCFOAQ00600OutBlock1+xt.SizeCFOAQ00600OutBlock2+5))%xt.SizeCFOAQ00600OutBlock3 != 0 {
	//		break
	//	}
	//
	//	return xt.P자료형_CFOAQ00600OutBlock, nil
	//case xt.TR선물옵션_정상주문_CFOAT00100:
	//	if 길이 == xt.SizeCFOAT00100OutBlock {
	//		return xt.P자료형_CFOAT00100OutBlock, nil
	//	}
	//case xt.TR선물옵션_정정주문_CFOAT00200:
	//	if 길이 == xt.SizeCFOAT00200OutBlock {
	//		return xt.P자료형_CFOAT00200OutBlock, nil
	//	}
	//case xt.TR선물옵션_취소주문_CFOAT00300:
	//	if 길이 == xt.SizeCFOAT00300OutBlock {
	//		return xt.P자료형_CFOAT00300OutBlock, nil
	//	}
	//case xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500:
	//	const 헤더_길이 = xt.SizeCFOBQ10500OutBlock1 + xt.SizeCFOBQ10500OutBlock2 + 5
	//
	//	if 길이 == 0 {
	//		return xt.P자료형_nil, nil
	//	} else if 길이 < 헤더_길이 || (길이-헤더_길이)%xt.SizeCFOBQ10500OutBlock3 != 0 {
	//		break
	//	}
	//
	//	return xt.P자료형_CFOBQ10500OutBlock, nil
	//case xt.TR선물옵션_미결제약정_현황_CFOFQ02400:
	//	const 헤더_길이 = xt.SizeCFOFQ02400OutBlock1 + xt.SizeCFOFQ02400OutBlock2 + 5 + 5
	//
	//	if 길이 == 0 {
	//		return xt.P자료형_nil, nil
	//	} else if 길이 < 헤더_길이 { // 각 Occurs OutBlock 앞에 5자리로 Count 가 들어갑니다.
	//		break
	//	}
	//
	//	return xt.P자료형_CFOFQ02400OutBlock, nil
	//case xt.TR선물옵션_체결_미체결_조회_t0434:
	//	// Non-block 모드는 Occurs데이터 수량을 나타내는 5바이트 추가됨.
	//	if 길이 == 0 {
	//		return xt.P자료형_nil, nil
	//	} else if 길이 < (xt.SizeT0434OutBlock+5) ||
	//		(길이-(xt.SizeT0434OutBlock+5))%xt.SizeT0434OutBlock1 != 0 {
	//		break
	//	}
	//
	//	return xt.P자료형_T0434OutBlock, nil
	//case xt.TR기업정보_요약_t3320:
	//	switch 길이 {
	//	case xt.SizeT3320OutBlock:
	//		return xt.P자료형_T3320OutBlock, nil
	//	case xt.SizeT3320OutBlock1:
	//		return xt.P자료형_T3320OutBlock1, nil
	//	}
	//case xt.TR지수선물_마스터_조회_t8432:
	//	switch {
	//	case 길이%xt.SizeT8432OutBlock == 0:
	//		return xt.P자료형_T8432OutBlock, nil
	//	}
}

func f민감정보_삭제(raw값 []byte, 구분_문자열 string) []byte {

	switch 구분_문자열 {
	case xt.P자료형_CSPAQ12300OutBlock,
		xt.P자료형_CSPAQ13700OutBlock,
		xt.P자료형_CSPAT00600OutBlock: //,
		//xt.P자료형_CFOAQ00600OutBlock,
		//xt.P자료형_CFOBQ10500OutBlock,
		//xt.P자료형_CFOFQ02400OutBlock:
		f민감정보_삭제_도우미(raw값, 25, 8)
	//case xt.P자료형_CFOAT00100OutBlock,
	//	xt.P자료형_CFOAT00200OutBlock,
	//	xt.P자료형_CFOAT00300OutBlock:
	//	f민감정보_삭제_도우미(raw값, 27, 8)
	case xt.P자료형_CSPAQ22200OutBlock1,
		xt.P자료형_CSPAQ22200OutBlock,
		xt.P자료형_CSPAQ12200OutBlock1,
		xt.P자료형_CSPAQ12200OutBlock:
		f민감정보_삭제_도우미(raw값, 28, 8)
	case xt.P자료형_CSPAT00700OutBlock,
		xt.P자료형_CSPAT00800OutBlock:
		f민감정보_삭제_도우미(raw값, 35, 8)
	case xt.RT현물_주문_접수_SC0:
		f민감정보_삭제_도우미(raw값, 277, 8)

	}

	return raw값
}

func f민감정보_삭제_도우미(raw값 []byte, 시작_인덱스, 길이 int) {
	for i := 시작_인덱스; i < (시작_인덱스 + 길이); i++ {
		raw값[i] = 0
	}
}

func f모의투자서버_접속_중() bool {
	return xt.F서버_구분() == xt.P서버_모의투자
}
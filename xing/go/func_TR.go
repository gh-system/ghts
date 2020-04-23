/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"strconv"

	"strings"
	"time"
)

func TrCFOAQ00600_선물옵션_주문체결내역(계좌번호 string, 선물옵션구분 xt.CFOAQ00600_선물옵션분류, 상품군 xt.T선옵_상품군,
	체결구분 lib.T체결_구분, 조회_시작일, 조회_종료일 time.Time) (응답값 *xt.CFOAQ00600_선물옵션_주문체결내역_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	응답값 = new(xt.CFOAQ00600_선물옵션_주문체결내역_응답)

	연속조회_여부 := false
	연속키 := ""

	for {
		질의값 := new(xt.CFOAQ00600_선물옵션_주문체결내역_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR선물옵션_주문체결내역조회_CFOAQ00600)
		질의값.M레코드수량 = 1
		질의값.M계좌번호 = 계좌번호
		// 비밀번호	// 소켓 전송 안 함.
		질의값.M조회_시작일 = lib.F2일자(조회_시작일)
		질의값.M조회_종료일 = lib.F2일자(조회_종료일)
		질의값.M선물옵션분류 = 선물옵션구분
		질의값.M상품군 = 상품군
		질의값.M체결구분 = 체결구분
		질의값.M정렬구분 = lib.P정렬_역순
		질의값.M연속조회_여부 = 연속조회_여부
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		수신값, ok := i응답값.(*xt.CFOAQ00600_선물옵션_주문체결내역_응답)
		lib.F조건부_패닉(!ok, "TrCFOAQ00600, 예상하지 못한 자료형 : '%T'", i응답값)

		응답값.M응답1 = 수신값.M응답1
		응답값.M응답2 = 수신값.M응답2
		응답값.M반복값_모음 = append(응답값.M반복값_모음, 수신값.M반복값_모음...)

		if !수신값.M추가_연속조회_필요 {
			break
		}

		연속조회_여부 = 수신값.M추가_연속조회_필요
		연속키 = 수신값.M연속키
	}

	return 응답값, nil
}

func TrCFOAT00100_선물옵션_정상주문(계좌번호 string, 선옵_종목코드 string, 매매구분 lib.T매도_매수_구분,
	호가유형 xt.T호가유형, 주문가격 float64, 주문수량 int64) (응답값 *xt.CFOAT00100_선물옵션_정상주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := new(xt.CFOAT00100_선물옵션_정상주문_질의값)
	질의값.S질의값_단일_종목 = lib.New질의값_단일_종목()
	질의값.S질의값_단일_종목.M구분 = xt.TR조회
	질의값.S질의값_단일_종목.M코드 = xt.TR선물옵션_정상주문_CFOAT00100
	질의값.S질의값_단일_종목.M종목코드 = 선옵_종목코드
	질의값.M계좌번호 = 계좌번호
	질의값.M매매구분 = 매매구분
	질의값.M호가유형 = 호가유형
	질의값.M주문가격 = 주문가격
	질의값.M주문수량 = 주문수량

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값, ok := i응답값.(*xt.CFOAT00100_선물옵션_정상주문_응답)
	lib.F조건부_패닉(!ok, "TrCFOAT00100 예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값, nil
}

func TrCFOAT00200_선물옵션_정정주문(선옵_종목코드 string, 계좌번호 string, 매매구분 lib.T매도_매수_구분,
	호가유형 xt.T호가유형, 원주문번호 int64, 주문가격 float64, 정정수량 int64) (응답값 *xt.CFOAT00200_선물옵션_정정주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := new(xt.CFOAT00200_선물옵션_정정주문_질의값)
	질의값.S질의값_단일_종목 = lib.New질의값_단일_종목()
	질의값.S질의값_단일_종목.M구분 = xt.TR조회
	질의값.S질의값_단일_종목.M코드 = xt.TR선물옵션_정정주문_CFOAT00200
	질의값.S질의값_단일_종목.M종목코드 = 선옵_종목코드
	질의값.M계좌번호 = 계좌번호
	질의값.M매매구분 = 매매구분
	질의값.M호가유형 = 호가유형
	질의값.M원주문번호 = 원주문번호
	질의값.M주문가격 = 주문가격
	질의값.M정정수량 = 정정수량

	for i := 0; i < 10; i++ { // 최대 10번 재시도
		i응답값, 에러 := F질의_단일TR(질의값)

		if 에러 != nil && (strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태")) {
			lib.F체크포인트(에러.Error())
			continue // 재시도
		}

		lib.F확인(에러)

		응답값, ok := i응답값.(*xt.CFOAT00200_선물옵션_정정주문_응답)
		lib.F조건부_패닉(!ok, "TrCFOAT00200() 예상하지 못한 자료형 : '%T'", i응답값)

		if 응답값.M응답2 != nil && 응답값.M응답2.M주문번호 <= 0 {
			continue
		}

		return 응답값, nil
	}

	return 응답값, nil
}

func TrCFOAT00300_선물옵션_취소주문(선옵_종목코드 string, 계좌번호 string, 원주문번호 int64, 취소수량 int64) (응답값 *xt.CFOAT00300_선물옵션_취소주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := new(xt.CFOAT00300_선물옵션_취소주문_질의값)
	질의값.S질의값_단일_종목 = lib.New질의값_단일_종목()
	질의값.S질의값_단일_종목.M구분 = xt.TR조회
	질의값.S질의값_단일_종목.M코드 = xt.TR선물옵션_취소주문_CFOAT00300
	질의값.S질의값_단일_종목.M종목코드 = 선옵_종목코드
	질의값.M계좌번호 = 계좌번호
	질의값.M원주문번호 = 원주문번호
	질의값.M취소수량 = 취소수량

	for i := 0; i < 10; i++ { // 최대 10번 재시도
		i응답값, 에러 := F질의_단일TR(질의값)

		if 에러 != nil && (strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태")) {
			continue // 재시도
		}

		lib.F확인(에러)

		응답값, ok := i응답값.(*xt.CFOAT00300_선물옵션_취소주문_응답)
		lib.F조건부_패닉(!ok, "TrCFOAT00300() 예상하지 못한 자료형 : '%T'", i응답값)

		if 응답값.M응답2 != nil && 응답값.M응답2.M주문번호 <= 0 {
			continue
		}

		return 응답값, nil
	}

	return nil, lib.New에러("선물옵션 취소 주문 TR 실행 실패.")
}

func TrCFOBQ10500_선물옵션_예탁금_증거금_조회(계좌번호 string) (응답값 *xt.CFOBQ10500_선물옵션_예탁금_증거금_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	응답값 = new(xt.CFOBQ10500_선물옵션_예탁금_증거금_조회_응답)

	연속조회_여부 := false
	연속키 := ""

	for {
		질의값 := new(xt.CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(lib.TR조회, xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500)
		질의값.M레코드수량 = 1
		질의값.M계좌번호 = 계좌번호
		질의값.M연속조회_여부 = 연속조회_여부
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		수신값, ok := i응답값.(*xt.CFOBQ10500_선물옵션_예탁금_증거금_조회_응답)
		lib.F조건부_패닉(!ok, "TrCFOBQ10500() 예상하지 못한 자료형 : '%T'", i응답값)

		응답값.M응답1 = 수신값.M응답1
		응답값.M응답2 = 수신값.M응답2
		응답값.M반복값_모음 = append(응답값.M반복값_모음, 수신값.M반복값_모음...)

		if !수신값.M추가_연속조회_필요 {
			break
		}

		연속조회_여부 = 수신값.M추가_연속조회_필요
		연속키 = 수신값.M연속키
	}

	return 응답값, nil
}

func TrCFOFQ02400_선물옵션_미결제약정(계좌번호 string, 등록시장 xt.CFOFQ02400_등록시장, 매수일자 time.Time) (응답값 *xt.CFOFQ02400_선물옵션_미결제약정_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	응답값 = new(xt.CFOFQ02400_선물옵션_미결제약정_응답)

	연속조회_여부 := false
	연속키 := ""

	for {
		질의값 := new(xt.CFOFQ02400_선물옵션_미결제약정_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(lib.TR조회, xt.TR선물옵션_미결제약정_현황_CFOFQ02400)
		질의값.M레코드수량 = 1
		질의값.M계좌번호 = 계좌번호
		// 비밀번호
		질의값.M등록시장코드 = 등록시장
		질의값.M매수일자 = 매수일자.Format("20060102")
		질의값.M연속조회_여부 = 연속조회_여부
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		수신값, ok := i응답값.(*xt.CFOFQ02400_선물옵션_미결제약정_응답)
		lib.F조건부_패닉(!ok, "TrCFOFQ02400() 예상하지 못한 자료형 : '%T'", i응답값)

		응답값.M응답1 = 수신값.M응답1
		응답값.M응답2 = 수신값.M응답2
		응답값.M반복값1_모음 = append(응답값.M반복값1_모음, 수신값.M반복값1_모음...)
		응답값.M반복값2_모음 = append(응답값.M반복값2_모음, 수신값.M반복값2_모음...)

		if !수신값.M추가_연속조회_필요 {
			break
		}

		연속조회_여부 = 수신값.M추가_연속조회_필요
		연속키 = 수신값.M연속키
	}

	return 응답값, nil
}

func TrCSPAT00600_현물_정상주문(질의값 *xt.CSPAT00600_현물_정상_주문_질의값) (응답값 *xt.CSPAT00600_현물_정상_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값, ok := i응답값.(*xt.CSPAT00600_현물_정상_주문_응답)
	lib.F조건부_패닉(!ok, "TrCSPAT00600() 예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값, nil
}

func TrCSPAT00700_현물_정정주문(질의값 *xt.CSPAT00700_현물_정정_주문_질의값) (응답값 *xt.CSPAT00700_현물_정정_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	for i := 0; i < 10; i++ { // 최대 10번 재시도
		i응답값, 에러 := F질의_단일TR(질의값)

		if 에러 != nil && (strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태입니다")) {
			continue
		}

		lib.F확인(에러)

		응답값, ok := i응답값.(*xt.CSPAT00700_현물_정정_주문_응답)
		lib.F조건부_패닉(!ok, "TrCSPAT00700() 예상하지 못한 자료형 : '%T'", i응답값)

		if 응답값.M응답2 != nil && 응답값.M응답2.M주문번호 <= 0 {
			continue
		}

		return 응답값, nil
	}

	return nil, lib.New에러("정정 주문 TR 실행 실패.")
}

func TrCSPAT00800_현물_취소주문(질의값 *lib.S질의값_취소_주문) (응답값 *xt.CSPAT00800_현물_취소_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	for i := 0; i < 10; i++ { // 최대 10번 재시도
		i응답값, 에러 := F질의_단일TR(질의값)

		if 에러 != nil && (strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태")) {
			continue // 재시도
		}

		lib.F확인(에러)

		응답값, ok := i응답값.(*xt.CSPAT00800_현물_취소_주문_응답)
		lib.F조건부_패닉(!ok, "TrCSPAT00800() 예상하지 못한 자료형 : '%T'", i응답값)

		if 응답값.M응답2 != nil && 응답값.M응답2.M주문번호 <= 0 {
			continue
		}

		return 응답값, nil
	}

	return nil, lib.New에러("취소 주문 TR 실행 실패.")
}

func TrCSPAQ12300_현물계좌_잔고내역_조회(계좌번호 string) (값 *xt.CSPAQ12300_현물계좌_잔고내역_응답, 에러 error) {
	계좌번호_모음 := lib.F확인(F계좌번호_모음()).([]string)

	존재함 := false
	for _, 계좌번호_값 := range 계좌번호_모음 {
		if 계좌번호 == 계좌번호_값 {
			존재함 = true
			break
		}
	}

	lib.F조건부_패닉(!존재함, "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	값 = new(xt.CSPAQ12300_현물계좌_잔고내역_응답)
	연속조회_여부 := false
	연속키 := ""

	for {
		질의값 := new(xt.CSPAQ12300_현물계좌_잔고내역_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물계좌_잔고내역_조회_CSPAQ12300)
		질의값.M레코드_수량 = 1 //    TR의 Inblock에서의 RecCnt는 무조건 "1"을 넣으시기 바랍니다.
		질의값.M계좌번호 = 계좌번호
		질의값.M잔고생성_구분 = "1"    // 0:전체, 1:현물
		질의값.M수수료적용_구분 = "0"   // 0:수수료 미적용, 1:수수료 적용
		질의값.D2잔고기준조회_구분 = "1" // 0:전부조회, 1:D2잔고 0이상만 조회
		질의값.M단가_구분 = "1"      // 0:평균단가, 1:BEP단가
		질의값.M연속조회_여부 = 연속조회_여부
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		수신값, ok := i응답값.(*xt.CSPAQ12300_현물계좌_잔고내역_응답)
		lib.F조건부_패닉(!ok, "TrCSPAQ12300() 예상하지 못한 자료형 : '%T'", i응답값)

		값.M헤더1 = 수신값.M헤더1
		값.M헤더2 = 수신값.M헤더2
		값.M반복값_모음 = append(값.M반복값_모음, 수신값.M반복값_모음...)

		if !수신값.M추가_연속조회_필요 {
			break
		}

		연속조회_여부 = 수신값.M추가_연속조회_필요
		연속키 = 수신값.M연속키
	}

	return 값, nil
}

func TrCSPAQ13700_현물계좌_주문체결내역(계좌번호 string, 주문일 time.Time,
	체결_미체결_구분 xt.T주문_체결_미체결_구분_CSPAQ13700) (값_모음 []*xt.CSPAQ13700_현물계좌_주문체결내역_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	계좌번호_모음 := lib.F확인(F계좌번호_모음()).([]string)

	존재함 := false
	for _, 계좌번호_값 := range 계좌번호_모음 {
		if 계좌번호 == 계좌번호_값 {
			존재함 = true
			break
		}
	}

	lib.F조건부_패닉(!존재함, "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	값_모음 = make([]*xt.CSPAQ13700_현물계좌_주문체결내역_반복값, 0)

	연속조회_여부 := false
	연속키 := ""

	const 역순구분 = "1" // "0":역순, "1":정순
	var 시작주문번호 int64

	if 역순구분 == "0" {
		시작주문번호 = 999999999
	} else {
		시작주문번호 = 000000000
	}

	for {
		질의값 := new(xt.CSPAQ13700_현물계좌_주문체결내역_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물계좌_주문체결내역_조회_CSPAQ13700)
		질의값.M레코드_수량 = 1 // //    TR의 Inblock에서의 RecCnt는 무조건 "1"을 넣으시기 바랍니다.
		질의값.M계좌번호 = 계좌번호
		질의값.M주문시장코드 = "00"                       // "00":전체, "10":거래소, "20":코스닥
		질의값.M매매구분 = "0"                          // "0":전체, "1":매도, "2":매수
		질의값.M종목코드 = ""                           // 주식 : A+종목코드, ELW : J+종목코드
		질의값.M체결여부 = strconv.Itoa(int(체결_미체결_구분)) // "0":전체, "1":체결, "3":미체결
		질의값.M주문일 = 주문일.Format("20060102")        // 주문일
		질의값.M시작주문번호 = 시작주문번호                     // 역순구분이 순 : 000000000, 역순구분이 역순 : 999999999
		질의값.M역순구분 = 역순구분
		질의값.M주문유형코드 = "00" // "00":전체, "98":매도, "99":매수
		질의값.M연속조회_여부 = 연속조회_여부
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		수신값, ok := i응답값.(*xt.CSPAQ13700_현물계좌_주문체결내역_응답)
		lib.F조건부_패닉(!ok, "TrCSPAQ13700() 예상하지 못한 자료형 : '%T'", i응답값)

		값_모음 = append(값_모음, 수신값.M반복값_모음...)

		if !수신값.M추가_연속조회_필요 {
			break
		}

		연속조회_여부 = 수신값.M추가_연속조회_필요
		연속키 = 수신값.M연속키
	}

	return 값_모음, nil
}

func TrCSPAQ22200_현물계좌_예수금_주문가능금액(계좌번호 string) (값 *xt.CSPAQ22200_현물계좌_예수금_주문가능금액_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	계좌번호_모음 := lib.F확인(F계좌번호_모음()).([]string)

	존재함 := false
	for _, 계좌번호_값 := range 계좌번호_모음 {
		if 계좌번호 == 계좌번호_값 {
			존재함 = true
			break
		}
	}

	lib.F조건부_패닉(!존재함, "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	질의값 := lib.New질의값_문자열(xt.TR조회, xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200, 계좌번호)
	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값, ok := i응답값.(*xt.CSPAQ22200_현물계좌_예수금_주문가능금액_응답)
	lib.F조건부_패닉(!ok, "TrCSPAQ22200 예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값, nil
}

func TrT0150_현물_당일_매매일지(계좌번호 string) (응답값_모음 []*xt.T0150_현물_당일_매매일지_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	계좌번호 = strings.ReplaceAll(계좌번호, "-", "")
	계좌번호_모음 := lib.F확인(F계좌번호_모음()).([]string)

	존재함 := false
	for _, 계좌번호_값 := range 계좌번호_모음 {
		if 계좌번호 == 계좌번호_값 {
			존재함 = true
			break
		}
	}

	lib.F조건부_패닉(!존재함, "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	var 연속키_매매구분, 연속키_종목코드, 연속키_단가, 연속키_매체 string

	응답값_모음 = make([]*xt.T0150_현물_당일_매매일지_응답_반복값, 0)

	for {
		질의값 := new(xt.T0150_현물_당일_매매일지_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물_당일_매매일지_t0150)
		질의값.M계좌번호 = 계좌번호
		질의값.M연속키_매매구분 = 연속키_매매구분
		질의값.M연속키_종목코드 = 연속키_종목코드
		질의값.M연속키_단가 = 연속키_단가
		질의값.M연속키_매체 = 연속키_매체

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		if i응답값 == nil {
			break
		}

		값, ok := i응답값.(*xt.T0150_현물_당일_매매일지_응답)
		lib.F조건부_패닉(!ok, "TrT0150() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키_매매구분 = 값.M헤더.CTS_매매구분
		연속키_종목코드 = 값.M헤더.CTS_종목코드
		연속키_단가 = 값.M헤더.CTS_단가
		연속키_매체 = 값.M헤더.CTS_매체

		응답값_모음 = append(값.M반복값_모음, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키_매매구분) == "" &&
			lib.F2문자열_공백제거(연속키_종목코드) == "" &&
			lib.F2문자열_공백제거(연속키_단가) == "" &&
			lib.F2문자열_공백제거(연속키_매체) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT0151_현물_일자별_매매일지(계좌번호 string, 일자 time.Time) (응답값_모음 []*xt.T0151_현물_일자별_매매일지_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	계좌번호_모음 := lib.F확인(F계좌번호_모음()).([]string)

	존재함 := false
	for _, 계좌번호_값 := range 계좌번호_모음 {
		if 계좌번호 == 계좌번호_값 {
			존재함 = true
			break
		}
	}

	lib.F조건부_패닉(!존재함, "존재하지 않는 계좌번호 : '%v'", 계좌번호)

	var 연속키_매매구분, 연속키_종목코드, 연속키_단가, 연속키_매체 string

	응답값_모음 = make([]*xt.T0151_현물_일자별_매매일지_응답_반복값, 0)

	for {
		질의값 := new(xt.T0151_현물_일자별_매매일지_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물_일자별_매매일지_t0151)
		질의값.M일자 = 일자.Format("20060102")
		질의값.M계좌번호 = 계좌번호
		질의값.M연속키_매매구분 = 연속키_매매구분
		질의값.M연속키_종목코드 = 연속키_종목코드
		질의값.M연속키_단가 = 연속키_단가
		질의값.M연속키_매체 = 연속키_매체

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		if i응답값 == nil {
			break
		}

		값, ok := i응답값.(*xt.T0151_현물_일자별_매매일지_응답)
		lib.F조건부_패닉(!ok, "TrT0151() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키_매매구분 = 값.M헤더.CTS_매매구분
		연속키_종목코드 = 값.M헤더.CTS_종목코드
		연속키_단가 = 값.M헤더.CTS_단가
		연속키_매체 = 값.M헤더.CTS_매체

		응답값_모음 = append(값.M반복값_모음, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키_매매구분) == "" &&
			lib.F2문자열_공백제거(연속키_종목코드) == "" &&
			lib.F2문자열_공백제거(연속키_단가) == "" &&
			lib.F2문자열_공백제거(연속키_매체) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT0167_시각_조회() chan *xt.T0167_시각_조회_응답 {
	ch응답 := make(chan *xt.T0167_시각_조회_응답, 1)

	go func(ch응답 chan *xt.T0167_시각_조회_응답) {
		var 에러 error

		defer lib.S예외처리{M에러: &에러, M함수: func() {
			if ch응답 != nil {
				응답값 := new(xt.T0167_시각_조회_응답)
				응답값.M시각 = time.Time{}
				응답값.M에러 = 에러

				ch응답 <- 응답값
			}
		}}.S실행()

		질의값 := lib.S질의값_기본형{M구분: xt.TR조회, M코드: xt.TR시간_조회_t0167}
		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		시각값, ok := i응답값.(time.Time)
		lib.F조건부_패닉(!ok, "TrT0167 예상하지 못한 자료형 : '%T", i응답값)

		응답값 := new(xt.T0167_시각_조회_응답)
		응답값.M시각 = 시각값
		응답값.M에러 = nil

		ch응답 <- 응답값
	}(ch응답)

	return ch응답
}

func TrT0425_현물_체결_미체결_조회(계좌번호, 종목코드 string, 체결_구분 lib.T체결_구분,
	매도_매수_구분 lib.T매도_매수_구분) (응답값_모음 []*xt.T0425_현물_체결_미체결_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))

	응답값_모음 = make([]*xt.T0425_현물_체결_미체결_조회_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := new(xt.T0425_현물_체결_미체결_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR현물_체결_미체결_조회_t0425)
		질의값.M계좌번호 = 계좌번호
		질의값.M종목코드 = 종목코드
		질의값.M체결구분 = 체결_구분
		질의값.M매도_매수_구분 = 매도_매수_구분
		질의값.M정렬구분 = lib.P정렬_정순
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		if i응답값 == nil {
			break
		}

		값, ok := i응답값.(*xt.T0425_현물_체결_미체결_조회_응답)
		lib.F조건부_패닉(!ok, "TrT0425() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(값.M반복값_모음, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT0434_선물옵션_체결_미체결_조회(계좌번호, 종목코드 string, 체결구분 lib.T체결_구분, 정렬구분 lib.T정렬_구분) (
	응답값_모음 []*xt.T0434_선물옵션_체결_미체결_조회_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*xt.T0434_선물옵션_체결_미체결_조회_반복값, 0)
	연속키 := ""

	for {
		질의값 := new(xt.T0434_선물옵션_체결_미체결_조회_질의값)
		질의값.S질의값_단일_종목 = lib.New질의값_단일_종목()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR선물옵션_체결_미체결_조회_t0434
		질의값.M계좌번호 = 계좌번호
		질의값.M종목코드 = 종목코드
		질의값.M체결구분 = 체결구분
		질의값.M정렬구분 = lib.P정렬_정순
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		if i응답값 == nil {
			break
		}

		값, ok := i응답값.(*xt.T0434_선물옵션_체결_미체결_조회_응답)
		lib.F조건부_패닉(!ok, "TrT0434() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M연속키
		응답값_모음 = append(값.M반복값_모음, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1101_현물_호가_조회(종목코드 string) (응답값 *xt.T1101_현물_호가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR현물_호가_조회_t1101
	질의값.M종목코드 = 종목코드

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값, ok := i응답값.(*xt.T1101_현물_호가_조회_응답)
	lib.F조건부_패닉(!ok, "TrT1101() 예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값, nil
}

func TrT1102_현물_시세_조회(종목코드 string) (응답값 *xt.T1102_현물_시세_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR현물_시세_조회_t1102
	질의값.M종목코드 = 종목코드

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값, ok := i응답값.(*xt.T1102_현물_시세_조회_응답)
	lib.F조건부_패닉(!ok, "TrT1102() 예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값, nil
}

func TrT1305_기간별_주가_조회(종목코드 string, 일주월_구분 xt.T일주월_구분, 추가_옵션_모음 ...interface{}) (
	응답값_모음 []*xt.T1305_현물_기간별_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	var 일자 time.Time

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			일자 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	lib.F조건부_패닉(일주월_구분 != xt.P일주월_일 && 일주월_구분 != xt.P일주월_주 &&
		일주월_구분 != xt.P일주월_월, "예상하지 못한 일주월 구분값 : '%v'", 일주월_구분)

	연속키 := ""
	응답값_모음 = make([]*xt.T1305_현물_기간별_조회_응답_반복값, 0)

	defer func() { // 순서 거꾸로 뒤집기.
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*xt.T1305_현물_기간별_조회_응답_반복값, 수량)

		for i, 응답값 := range 응답값_모음_임시 {
			응답값_모음[수량-i-1] = 응답값
		}
	}()

	for {
		질의값 := xt.NewT1305_현물_기간별_조회_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_기간별_조회_t1305
		질의값.M종목코드 = 종목코드
		질의값.M일주월_구분 = 일주월_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		값, ok := i응답값.(*xt.T1305_현물_기간별_조회_응답)
		lib.F조건부_패닉(!ok, "TrT1305() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		lib.F조건부_패닉(값.M헤더.M수량 != int64(len(값.M반복값_모음.M배열)),
			"반복값 수량 불일치. '%v', '%v'", 값.M헤더.M수량, len(값.M반복값_모음.M배열))

		if !일자.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M일자.Equal(일자) || 응답값.M일자.Before(일자) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1310_현물_당일전일_분틱_조회(종목코드 string, 당일전일_구분 xt.T당일전일_구분, 분틱_구분 xt.T분틱_구분,
	종료시각 time.Time, 수량_옵션 ...int) (응답값_모음 []*xt.T1310_현물_전일당일분틱조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	if len(수량_옵션) == 1 {
		수량 = 수량_옵션[0]
	}

	lib.F조건부_패닉(당일전일_구분 != xt.P당일전일구분_당일 && 당일전일_구분 != xt.P당일전일구분_전일,
		"예상하지 못한 당일_전일 구분값 : '%v'", 당일전일_구분)

	lib.F조건부_패닉(분틱_구분 != xt.P분틱구분_분 && 분틱_구분 != xt.P분틱구분_틱,
		"예상하지 못한 분_틱 구분값 : '%v'", 분틱_구분)

	응답값_모음_역순 := make([]*xt.T1310_현물_전일당일분틱조회_응답_반복값, 0)
	연속키 := ""

	defer func() {
		일자 := lib.F조건부_시간(당일전일_구분 == xt.P당일전일구분_당일, F당일(), F전일())
		수량 = len(응답값_모음_역순)
		응답값_모음 = make([]*xt.T1310_현물_전일당일분틱조회_응답_반복값, len(응답값_모음_역순))

		// 종목코드, 당일/전일 설정. 시간 기준 정렬순서 변경.
		for i, 응답값 := range 응답값_모음_역순 {
			응답값.M종목코드 = 종목코드

			시각 := 응답값.M시각
			응답값.M시각 = time.Date(일자.Year(), 일자.Month(), 일자.Day(),
				시각.Hour(), 시각.Minute(), 시각.Second(), 시각.Nanosecond(), 시각.Location())

			응답값_모음[수량-1-i] = 응답값
		}
	}()

	for {
		질의값 := xt.NewT1310_현물_전일당일_분틱_조회_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_당일_전일_분틱_조회_t1310
		질의값.M종목코드 = 종목코드
		질의값.M당일전일구분 = 당일전일_구분
		질의값.M분틱구분 = 분틱_구분
		질의값.M종료시각 = 종료시각.Format("1504")
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T1310_현물_전일당일분틱조회_응답)
		lib.F조건부_패닉(!ok, "TrT1310() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음_역순 = append(응답값_모음_역순, 값.M반복값_모음.M배열...)

		if 수량 > 0 && len(응답값_모음_역순) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1404_관리종목_조회(시장_구분 lib.T시장구분, 관리_질의_구분 xt.T관리_질의_구분) (응답값_모음 []*xt.T1404_관리종목_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*xt.T1404_관리종목_조회_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := new(xt.T1404_관리종목_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR관리_불성실_투자유의_조회_t1404)
		질의값.M시장_구분 = 시장_구분
		질의값.M관리_질의_구분 = 관리_질의_구분
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T1404_관리종목_조회_응답)
		lib.F조건부_패닉(!ok, "TrT1404() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1405_투자경고_조회(시장_구분 lib.T시장구분, 투자경고_질의_구분 xt.T투자경고_질의_구분) (응답값_모음 []*xt.T1405_투자경고_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*xt.T1405_투자경고_조회_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := new(xt.T1405_투자경고_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR투자경고_매매정지_정리매매_조회_t1405)
		질의값.M시장_구분 = 시장_구분
		질의값.M투자경고_질의_구분 = 투자경고_질의_구분
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T1405_투자경고_조회_응답)
		lib.F조건부_패닉(!ok, "TrT1405() 예상하지 못한 자료형 : '%T'", i응답값)

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if 연속키 = 값.M헤더.M연속키; lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1902_ETF_시간별_추이(종목코드 string, 추가_옵션_모음 ...interface{}) (응답값_모음 []*xt.T1902_ETF시간별_추이_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	var 시각 time.Time

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			시각 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	응답값_모음 = make([]*xt.T1902_ETF시간별_추이_응답_반복값, 0)
	연속키 := ""

	defer func() { // 순서 거꾸로 뒤집고, 종목코드 정보 및 누락된 시각 데이터 추가.
		nil시각 := time.Time{}
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*xt.T1902_ETF시간별_추이_응답_반복값, 수량)

		for i, 응답값 := range 응답값_모음_임시 {
			if 응답값.M시각.Equal(nil시각) && i != 0 && !응답값_모음_임시[i-1].M시각.Equal(nil시각) {
				응답값.M시각 = 응답값_모음_임시[i-1].M시각.Add(-1 * lib.P10초)
			}

			응답값.M종목코드 = 종목코드
			응답값_모음[수량-i-1] = 응답값
		}

		for i, 응답값 := range 응답값_모음 {
			if 응답값.M시각.Equal(nil시각) && i != 0 && !응답값_모음_임시[i-1].M시각.Equal(nil시각) {
				응답값.M시각 = 응답값_모음[i-1].M시각.Add(lib.P10초)
			}
		}
	}()

	for {
		질의값 := lib.New질의값_단일종목_연속키()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR_ETF_시간별_추이_t1902
		질의값.M종목코드 = 종목코드
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T1902_ETF시간별_추이_응답)
		lib.F조건부_패닉(!ok, "TrT1902() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		if !시각.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M시각.Equal(시각) || 응답값.M시각.Before(시각) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

// HTS 3302 화면. t3320 은 참고자료로서 정보의 정확성이나 완전성은 보장하기는 어렵습니다. 숫자 엉망이다.
func TrT3320_F기업정보_요약(종목코드 string) (응답값 *xt.T3320_기업정보_요약_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR기업정보_요약_t3320
	질의값.M종목코드 = 종목코드

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
	// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
	lib.F대기(lib.P3초)

	응답값, ok := i응답값.(*xt.T3320_기업정보_요약_응답)
	lib.F조건부_패닉(!ok, "TrT3320() 예상하지 못한 자료형 : '%T'", i응답값)

	응답값.M종목코드 = 종목코드
	return 응답값, nil
}

// HTS 3303 화면
func TrT3341_재무_순위_종합(시장구분 lib.T시장구분, 재무순위_구분 xt.T재무순위_구분,
	추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T3341_재무순위_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	switch 시장구분 {
	case lib.P시장구분_전체,
		lib.P시장구분_코스피,
		lib.P시장구분_코스닥: // OK
	default:
		panic(lib.New에러("잘못된 시장구분값 : '%s' '%d'", 시장구분, 시장구분))
	}

	switch 재무순위_구분 {
	case xt.P재무순위_매출액증가율,
		xt.P재무순위_영업이익증가율,
		xt.P재무순위_세전계속이익증가율,
		xt.P재무순위_부채비율,
		xt.P재무순위_유보율,
		xt.P재무순위_EPS,
		xt.P재무순위_BPS,
		xt.P재무순위_ROE,
		xt.P재무순위_PER,
		xt.P재무순위_PBR,
		xt.P재무순위_PEG:
		// OK
	default:
		panic(lib.New에러("잘못된 재무순위 구분값 : '%s' '%s'", string(재무순위_구분), 재무순위_구분.String()))
	}

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T3341_재무순위_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := xt.NewT3341_재무순위_질의값()
		질의값.M시장구분 = 시장구분
		질의값.M재무순위_구분 = 재무순위_구분
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		값, ok := i응답값.(*xt.T3341_재무순위_응답)
		lib.F조건부_패닉(!ok, "TrT3341() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}
	}

	return 응답값_모음, nil
}

func TrT8407_현물_멀티_현재가_조회_전종목() (응답값_모음 []*xt.T8407_현물_멀티_현재가_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	종목코드_모음_전체 := F종목코드_모음_전체()
	종목코드_모음_나머지 := 종목코드_모음_전체
	응답값_모음 = make([]*xt.T8407_현물_멀티_현재가_조회_응답_반복값, 0)

	for {
		var 종목코드_모음 []string

		if len(종목코드_모음_나머지) == 0 {
			break
		} else if len(종목코드_모음_나머지) >= 50 {
			종목코드_모음 = 종목코드_모음_나머지[:50]
			종목코드_모음_나머지 = 종목코드_모음_나머지[50:]
		} else {
			종목코드_모음 = 종목코드_모음_나머지
			종목코드_모음_나머지 = nil
		}

		값_모음, 에러 := TrT8407_현물_멀티_현재가_조회(종목코드_모음)
		lib.F확인(에러)

		if len(값_모음) > 0 {
			응답값_모음 = append(응답값_모음, 값_모음...)
		}
	}

	return 응답값_모음, nil
}

func TrT8407_현물_멀티_현재가_조회(종목코드_모음 []string) (응답값_모음 []*xt.T8407_현물_멀티_현재가_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(len(종목코드_모음) > 50, "한 번에 50개 종목만 질의 가능합니다. : '%v'", len(종목코드_모음))

	for _, 종목코드 := range 종목코드_모음 {
		lib.F확인(F종목코드_검사(종목코드))
	}

	// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
	// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
	lib.F대기(lib.P1초)

	질의값 := xt.NewT8407_현물_멀티_현재가_조회_질의값(종목코드_모음)
	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값_모음, ok := i응답값.([]*xt.T8407_현물_멀티_현재가_조회_응답_반복값)
	lib.F조건부_패닉(!ok, "TrT8407() 예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값_모음, nil
}

func TrT8411_현물_차트_틱(종목코드 string, 시작일자, 종료일자 time.Time, 추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8411_현물_차트_틱_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8411_현물_차트_틱_응답_반복값, 0)
	연속일자 := ""
	연속시간 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	for {
		질의값 := xt.NewT8411_현물_차트_틱_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_틱_t8411
		질의값.M종목코드 = 종목코드
		질의값.M단위 = 1
		질의값.M요청건수 = 2000
		질의값.M조회영업일수 = 0
		질의값.M시작일자 = 시작일자.Format("20060102")
		질의값.M종료일자 = 종료일자.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M연속시간 = 연속시간
		질의값.M압축여부 = true

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T8411_현물_차트_틱_응답)
		lib.F조건부_패닉(!ok, "TrT8411() 예상하지 못한 자료형 : '%T'", i응답값)

		연속일자 = 값.M헤더.M연속일자
		연속시간 = 값.M헤더.M연속시간

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" || lib.F2문자열_공백제거(연속시간) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8412_현물_차트_분(종목코드 string, 시작일자, 종료일자 time.Time, 추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8412_현물_차트_분_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8412_현물_차트_분_응답_반복값, 0)
	연속일자 := ""
	연속시간 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	for {
		질의값 := xt.NewT8412_현물_차트_분_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_분_t8412
		질의값.M종목코드 = 종목코드
		질의값.M단위 = 0 // 0:30초, 1: 1분, 2: 2분, ....., n: n분
		질의값.M요청건수 = 2000
		질의값.M조회영업일수 = 0
		질의값.M시작일자 = 시작일자.Format("20060102")
		질의값.M종료일자 = 종료일자.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M연속시간 = 연속시간
		질의값.M압축여부 = true

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T8412_현물_차트_분_응답)
		lib.F조건부_패닉(!ok, "TrT8412() 예상하지 못한 자료형 : '%T', '%v'", i응답값, len(응답값_모음))

		연속일자 = 값.M헤더.M연속일자
		연속시간 = 값.M헤더.M연속시간

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" || lib.F2문자열_공백제거(연속시간) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8413_현물_차트_일주월(종목코드 string, 시작일자, 종료일자 time.Time, 주기구분 xt.T일주월_구분,
	추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8413_현물_차트_일주월_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8413_현물_차트_일주월_응답_반복값, 0)
	연속일자 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	for {
		질의값 := xt.NewT8413_현물_차트_일주월_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_일주월_t8413
		질의값.M종목코드 = 종목코드
		질의값.M주기구분 = 주기구분
		질의값.M요청건수 = 2000 // 최대 압축 2000, 비압축 500
		질의값.M시작일자 = 시작일자.Format("20060102")
		질의값.M종료일자 = 종료일자.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M압축여부 = true

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T8413_현물_차트_일주월_응답)
		lib.F조건부_패닉(!ok, "TrT8413() 예상하지 못한 자료형 : '%T'", i응답값)

		연속일자 = 값.M헤더.M연속일자

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

// HTS 1503 화면
func TrT8428_증시주변자금추이(시장_구분 lib.T시장구분, 추가_옵션_모음 ...interface{}) (응답값_모음 []*xt.T8428_증시주변_자금추이_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(
		시장_구분 != lib.P시장구분_코스피 && 시장_구분 != lib.P시장구분_코스닥,
		"예상하지 못한 시장 구분값 : '%v'", 시장_구분)

	var 수량 int
	var 일자 time.Time
	var 연속키 string

	응답값_모음 = make([]*xt.T8428_증시주변_자금추이_응답_반복값, 0)

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			일자 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	for {
		질의값 := xt.NewT8428_증시주변자금추이_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR증시_주변_자금_추이_t8428
		질의값.M시장구분 = 시장_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		값, ok := i응답값.(*xt.T8428_증시주변_자금추이_응답)
		lib.F조건부_패닉(!ok, "TrT8428() 예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		if !일자.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M일자.Equal(일자) || 응답값.M일자.Before(일자) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if len(lib.F정규식_검색(연속키, []string{"[0-9]*"})) < 8 {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8432_지수선물_마스터_조회(구분 string) (응답값_모음 []*xt.T8432_지수선물_마스터_조회_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	질의값 := lib.New질의값_문자열(xt.TR조회, xt.TR지수선물_마스터_조회_t8432, 구분)
	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	var ok bool
	응답값_모음, ok = i응답값.([]*xt.T8432_지수선물_마스터_조회_반복값)
	lib.F조건부_패닉(!ok, "TrT8432() 예상하지 못한 자료형 : '%T'", 응답값_모음)

	return 응답값_모음, nil
}

func TrT8436_주식종목_조회(시장_구분 lib.T시장구분) (응답값_모음 []*xt.T8436_현물_종목조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 시장구분_문자열 string

	switch 시장_구분 {
	case lib.P시장구분_전체:
		시장구분_문자열 = "0"
	case lib.P시장구분_코스피:
		시장구분_문자열 = "1"
	case lib.P시장구분_코스닥:
		시장구분_문자열 = "2"
	default:
		panic(lib.New에러("예상하지 못한 시장 구분값 : '%v'", 시장_구분))
	}

	질의값 := lib.New질의값_문자열(xt.TR조회, xt.TR현물_종목_조회_t8436, 시장구분_문자열)
	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	값, ok := i응답값.(*xt.T8436_현물_종목조회_응답_반복값_모음)
	lib.F조건부_패닉(!ok, "TrT8436() 예상하지 못한 자료형 : '%T'", i응답값)

	return 값.M배열, nil
}

func F질의(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 *lib.S바이트_변환_모음) {
	var 에러 error

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		값 = lib.New바이트_변환_모음_단순형(lib.MsgPack, 에러)
	}}.S실행()

	lib.F확인(F질의값_종목코드_검사(질의값))

	switch 질의값.TR구분() {
	case xt.TR조회, xt.TR주문:
		f전송_권한_획득(질의값.TR코드())

		defer f전송_시각_기록(질의값.TR코드())
	}

	소켓REQ := 소켓REQ_저장소.G소켓()
	defer 소켓REQ_저장소.S회수(소켓REQ)

	if len(옵션_모음) > 0 {
		소켓REQ.S옵션(옵션_모음...)
	}

	return 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 질의값)
}

func F질의_단일TR(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 interface{}, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 에러 }}.S실행()

	타임아웃 := lib.P1분

	for _, 옵션 := range 옵션_모음 {
		switch 변환값 := 옵션.(type) {
		case time.Duration:
			타임아웃 = 변환값
		}
	}

	i식별번호 := F질의(질의값, 옵션_모음...).G해석값_단순형(0)

	switch 변환값 := i식별번호.(type) {
	case int:
		break
	case error:
		lib.F에러_출력(변환값)
		return nil, 변환값
	default:
		panic(lib.New에러with출력("F질의_단일TR() 예상하지 못한 자료형.\n" +
			"Xing API에서 식별번호를 부여받고, 콜백을 통해서 응답이 있는 경우에만 사용할 것.\n" +
			"그렇지 않은 경우에는 F질의()를 사용할 것.\n'%T'\n'%v'\n"))
	}

	식별번호 := i식별번호.(int)
	ch회신 := 대기소_C32.S추가(식별번호, 질의값.TR코드())

	select {
	case 값 := <-ch회신:
		switch 변환값 := 값.(type) {
		case error:
			if strings.Contains(변환값.Error(), "주문이 접수 대기") ||
				strings.Contains(변환값.Error(), "원주문번호를 잘못 입력") {
				return nil, 변환값
			}

			println("*********************************************************")
			println(변환값.Error())
			lib.F문자열_출력("*********************************************************")

			return nil, 변환값
		default:
			return 값, nil
		}
	case <-time.After(타임아웃):
		return nil, lib.New에러("타임아웃. '%v' '%v'", 질의값.TR코드(), 식별번호)
	}
}

func F접속됨() (접속됨 bool, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 접속됨 = false }}.S실행()

	질의값 := lib.New질의값_기본형(xt.TR접속됨, "")
	접속됨 = lib.F확인(F질의(질의값, lib.P10초).G해석값(0)).(bool)

	return 접속됨, nil
}

func F계좌번호_모음() (응답값 []string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌번호_모음 = nil }}.S실행()

	if len(계좌번호_모음) != 0 {
		return 계좌번호_모음, nil
	}

	질의값 := lib.New질의값_기본형(xt.TR계좌번호_모음, "")

	계좌번호_모음 = make([]string, 0)
	lib.F확인(F질의(질의값, lib.P10초).G값(0, &계좌번호_모음))

	return 계좌번호_모음, nil
}

func F계좌_수량() (계좌_수량 int, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_수량 = 0 }}.S실행()

	회신_메시지 := F질의(lib.New질의값_기본형(xt.TR계좌_수량, ""))
	계좌_수량 = lib.F확인(회신_메시지.G해석값(0)).(int)
	lib.F조건부_패닉(계좌_수량 == 0, "계좌 수량 0.")

	return 계좌_수량, nil
}

func F계좌_번호(인덱스 int) (계좌_번호 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_번호 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_정수(xt.TR계좌번호_모음, "", 인덱스))
	계좌_번호 = lib.F확인(회신_메시지.G해석값(0)).([]string)[인덱스]

	return 계좌_번호, nil
}

func F계좌_이름(계좌_번호 string) (계좌_이름 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_이름 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(xt.TR계좌_이름, "", 계좌_번호))
	계좌_이름 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_이름, nil
}

func F계좌_상세명(계좌_번호 string) (계좌_상세명 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_상세명 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(xt.TR계좌_상세명, "", 계좌_번호))
	계좌_상세명 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_상세명, nil
}

func F계좌_별명(계좌_번호 string) (계좌_별명 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_별명 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(xt.TR계좌_별명, "", 계좌_번호))
	계좌_별명 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_별명, nil
}

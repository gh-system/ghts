/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	"strings"
	"testing"
	"time"
)

func TestI안전한_bool(테스트 *testing.T) {
	테스트.Parallel()

	안전한_bool := New안전한_bool(false)

	F테스트_거짓임(테스트, 안전한_bool.G값())
	F테스트_에러발생(테스트, 안전한_bool.S값(false))
	F테스트_거짓임(테스트, 안전한_bool.G값())

	F테스트_에러없음(테스트, 안전한_bool.S값(true))
	F테스트_참임(테스트, 안전한_bool.G값())
}

func TestI안전한_string(테스트 *testing.T) {
	테스트.Parallel()

	안전한_string := New안전한_string("테스트")

	F테스트_같음(테스트, 안전한_string.G값(), "테스트")
	안전한_string.S값("테스트 2")
	F테스트_같음(테스트, 안전한_string.G값(), "테스트 2")
}

func TestI안전한_zmq소켓(테스트 *testing.T) {
	F메모("TestI안전한_zmq소켓()")
}

func TestI메시지(테스트 *testing.T) {
	테스트.Parallel()

	메시지 := New메시지(P메시지_GET, "테스트")
	F테스트_같음(테스트, 메시지.G구분(), P메시지_GET)
	F테스트_같음(테스트, 메시지.G내용(0), "테스트")
	F테스트_같음(테스트, 메시지.G길이(), 1)
	F테스트_같음(테스트, len(메시지.G내용_전체()), 메시지.G길이())
	F테스트_같음(테스트, 메시지.G내용_전체()[0], 메시지.G내용(0))

	문자열 := 메시지.String()
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G구분()))
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G내용(0)))

	메시지 = New메시지(P메시지_OK, "테스트", 1)
	F테스트_같음(테스트, 메시지.G구분(), P메시지_OK)
	F테스트_같음(테스트, 메시지.G내용(0), "테스트")
	F테스트_같음(테스트, 메시지.G내용(1), "1")
	F테스트_같음(테스트, 메시지.G길이(), 2)
	F테스트_같음(테스트, len(메시지.G내용_전체()), 메시지.G길이())
	F테스트_같음(테스트, 메시지.G내용_전체()[0], 메시지.G내용(0))
	F테스트_같음(테스트, 메시지.G내용_전체()[1], 메시지.G내용(1))

	문자열 = 메시지.String()
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G구분()))
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G내용(0)))
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G내용(1)))
}

func TestI질의_I회신(테스트 *testing.T) {
	// New질의()
	질의 := New질의(P메시지_GET, "질의", 1)

	_, ok := 질의.(I질의)
	F테스트_참임(테스트, ok)

	F테스트_같음(테스트, 질의.G구분(), P메시지_GET)
	F테스트_같음(테스트, 질의.G길이(), 2)
	F테스트_같음(테스트, 질의.G내용(0), "질의")
	F테스트_같음(테스트, 질의.G내용(1), "1")

	// New질의_zmq메시지()
	질의 = New질의_zmq메시지([]string{P메시지_GET, "질의", "1"})

	_, ok = 질의.(I질의)
	F테스트_참임(테스트, ok)

	F테스트_같음(테스트, 질의.G구분(), P메시지_GET)
	F테스트_같음(테스트, 질의.G길이(), 2)
	F테스트_같음(테스트, 질의.G내용(0), "질의")
	F테스트_같음(테스트, 질의.G내용(1), "1")

	회신 := New회신(F에러_생성("테스트용 에러"))

	_, ok = 회신.(I회신)
	F테스트_참임(테스트, ok)
	F테스트_참임(테스트, strings.HasPrefix(회신.G에러().Error(), "테스트용 에러"))
	F테스트_같음(테스트, 회신.G구분(), P메시지_에러)
	F테스트_같음(테스트, 회신.G길이(), 0)

	회신 = New회신(nil, "회신", 10)

	F테스트_같음(테스트, 회신.G에러(), nil)
	F테스트_같음(테스트, 회신.G구분(), P메시지_OK)
	F테스트_같음(테스트, 회신.G길이(), 2)
	F테스트_같음(테스트, 회신.G내용(0), "회신")
	F테스트_같음(테스트, 회신.G내용(1), "10")

	// 실제로 'I질의'와 'I회신'을 주고받는 테스트
	ch초기화_대기 := make(chan bool)
	ch질의 := make(chan I질의)
	ch종료 := make(chan S비어있는_구조체)

	go testI질의_도우미_Go루틴(ch초기화_대기, ch질의, ch종료)
	<-ch초기화_대기

	회신 = New질의(P메시지_SET, "키1", "값1").G회신(ch질의, P타임아웃_Go)
	F테스트_에러없음(테스트, 회신.G에러())
	F테스트_같음(테스트, 회신.G길이(), 0)

	회신 = New질의(P메시지_SET, "키2", "값2-1", "값2-2").G회신(ch질의, P타임아웃_Go)
	F테스트_에러없음(테스트, 회신.G에러())
	F테스트_같음(테스트, 회신.G길이(), 0)

	회신 = New질의(P메시지_SET, "키3", "값3-1", "값3-2", "값3-3").G회신(ch질의, P타임아웃_Go)
	F테스트_에러없음(테스트, 회신.G에러())
	F테스트_같음(테스트, 회신.G길이(), 0)

	회신 = New질의(P메시지_GET, "키1").G회신(ch질의, P타임아웃_Go)
	F테스트_에러없음(테스트, 회신.G에러())
	F테스트_같음(테스트, 회신.G길이(), 2)
	F테스트_같음(테스트, len(회신.G내용_전체()), 회신.G길이())
	F테스트_같음(테스트, 회신.G내용(0), "키1")
	F테스트_같음(테스트, 회신.G내용(1), "값1")

	회신 = New질의(P메시지_GET, "키2").G회신(ch질의, P타임아웃_Go)
	F테스트_에러없음(테스트, 회신.G에러())
	F테스트_같음(테스트, 회신.G길이(), 3)
	F테스트_같음(테스트, len(회신.G내용_전체()), 회신.G길이())
	F테스트_같음(테스트, 회신.G내용(0), "키2")
	F테스트_같음(테스트, 회신.G내용(1), "값2-1")
	F테스트_같음(테스트, 회신.G내용(2), "값2-2")

	회신 = New질의(P메시지_GET, "키3").G회신(ch질의, P타임아웃_Go)
	F테스트_에러없음(테스트, 회신.G에러())
	F테스트_같음(테스트, 회신.G길이(), 4)
	F테스트_같음(테스트, len(회신.G내용_전체()), 회신.G길이())
	F테스트_같음(테스트, 회신.G내용(0), "키3")
	F테스트_같음(테스트, 회신.G내용(1), "값3-1")
	F테스트_같음(테스트, 회신.G내용(2), "값3-2")
	F테스트_같음(테스트, 회신.G내용(3), "값3-3")

	회신 = New질의(P메시지_DEL, "키2").G회신(ch질의, P타임아웃_Go)
	F테스트_에러없음(테스트, 회신.G에러())
	F테스트_같음(테스트, 회신.G길이(), 0)

	회신 = New질의(P메시지_GET, "키2").G회신(ch질의, P타임아웃_Go)
	F테스트_에러발생(테스트, 회신.G에러())
	F테스트_같음(테스트, 회신.G길이(), 0)

	ch종료 <- S비어있는_구조체{}
}

func testI질의_도우미_Go루틴(ch초기화 chan bool, ch질의 chan I질의, ch종료 chan S비어있는_구조체) {
	맵 := make(map[string][]string)
	ch초기화 <- true

	// 받은 문자열을 그대로 되돌려 줌.

	for {
		select {
		case 질의 := <-ch질의:
			switch 질의.G구분() {
			case P메시지_GET:
				문자열_모음, 존재함 := 맵[질의.G내용(0)]
				if !존재함 {
					질의.S회신(F에러_생성("존재하지 않는 값. %v", 질의.G내용(0)))
				} else {
					질의.S회신(nil, F문자열_모음2인터페이스_모음(문자열_모음)...)
				}
			case P메시지_SET:
				맵[질의.G내용(0)] = 질의.G내용_전체()
				질의.S회신(nil)
			case P메시지_DEL:
				delete(맵, 질의.G내용(0))
				질의.S회신(nil)
			default:
				panic("")
			}
		case <-ch종료:
			return
		}
	}
}

func TestI종목(테스트 *testing.T) {
	테스트.Parallel()

	종목 := New종목("코드", "이름")
	F테스트_같음(테스트, 종목.G코드(), "코드")
	F테스트_같음(테스트, 종목.G이름(), "이름")
}

func TestI통화(테스트 *testing.T) {
	테스트.Parallel()

	통화 := New통화(KRW, 100.01)
	F테스트_같음(테스트, 통화.G단위(), KRW)
	F테스트_같음(테스트, 통화.G실수값(), 100.01)
	F테스트_같음(테스트, 통화.G정밀값().Float(), 100.01)
	F테스트_같음(테스트, 통화.G문자열값(), "100.01")
	F테스트_같음(테스트, 통화.G문자열값_고정소숫점(1), "100.0")
	F테스트_같음(테스트, 통화.G문자열값_고정소숫점(2), "100.01")
	F테스트_같음(테스트, 통화.G문자열값_고정소숫점(3), "100.010")

	F테스트_같음(테스트, 통화.G비교(New통화(KRW, 100.01)), P같음)
	F테스트_같음(테스트, 통화.G비교(New통화(KRW, 100.02)), P큼)
	F테스트_같음(테스트, 통화.G비교(New통화(KRW, 100.00)), P작음)
	F테스트_같음(테스트, 통화.G비교(New통화(USD, 100.00)), P비교불가)

	F테스트_같음(테스트, New통화(KRW, 100.00).G부호(), P양수)
	F테스트_같음(테스트, New통화(KRW, -100.00).G부호(), P음수)
	F테스트_같음(테스트, New통화(KRW, 0.0).G부호(), P영)

	// 복사본에 변경을 가해도 원본값은 변경되지 않는 지 확인.
	통화 = New통화(KRW, 100.01)
	F테스트_같음(테스트, 통화.G복사본().G비교(New통화(KRW, 100.01)), P같음)
	F테스트_같음(테스트, 통화.G복사본().S금액(10.00).G비교(New통화(KRW, 10.00)), P같음)
	F테스트_같음(테스트, 통화.G비교(New통화(KRW, 100.01)), P같음)
	F테스트_같음(테스트, 통화.G복사본().G비교(New통화(KRW, 100.01)), P같음)

	통화 = New통화(KRW, 100.01)
	F테스트_거짓임(테스트, 통화.G변경불가())
	통화.S동결()
	F테스트_참임(테스트, 통화.G변경불가())
	F테스트_패닉발생(테스트, 통화.S더하기, 100.0)
	F테스트_패닉발생(테스트, 통화.S빼기, 100.0)
	F테스트_패닉발생(테스트, 통화.S곱하기, 100.0)
	F테스트_패닉발생(테스트, 통화.S나누기, 100.0)
	F테스트_패닉발생(테스트, 통화.S금액, 100.0)

	F테스트_같음(테스트, New통화(KRW, 100.00).S더하기(100).G비교(New통화(KRW, 200.00)), P같음)
	F테스트_같음(테스트, New통화(KRW, 100.00).S빼기(100).G비교(New통화(KRW, 0.00)), P같음)
	F테스트_같음(테스트, New통화(KRW, 100.00).S곱하기(100.0).G비교(New통화(KRW, 10000.00)), P같음)

	통화, 에러 := New통화(KRW, 100.00).S나누기(100.0)
	F테스트_에러없음(테스트, 에러)
	F테스트_같음(테스트, 통화.G비교(New통화(KRW, 1.00)), P같음)

	F테스트_같음(테스트, New통화(KRW, 100.00).String(), "KRW 100")

	F테스트_같음(테스트, New원화(100.00).G비교(New통화(KRW, 100.00)), P같음)
	F테스트_같음(테스트, New달러(100.00).G비교(New통화(USD, 100.00)), P같음)
	F테스트_같음(테스트, New유로(100.00).G비교(New통화(EUR, 100.00)), P같음)
	F테스트_같음(테스트, New위안(100.00).G비교(New통화(CNY, 100.00)), P같음)

	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_해제()

	통화, 에러 = New통화(KRW, 100.00).S나누기(0.0)
	F테스트_에러발생(테스트, 에러)
	F테스트_같음(테스트, 통화, nil)
}

func TestI가격정보(테스트 *testing.T) {
	테스트.Parallel()

	시점1 := time.Now()
	가격정보 := New가격정보("종목코드", New원화(100.00))
	시점2 := time.Now()

	F테스트_같음(테스트, 가격정보.G종목코드(), "종목코드")
	F테스트_같음(테스트, 가격정보.G가격().G단위(), KRW)
	F테스트_같음(테스트, 가격정보.G가격().G실수값(), 100.0)

	F테스트_참임(테스트, 가격정보.G시점().Equal(시점1) || 가격정보.G시점().After(시점1))
	F테스트_참임(테스트, 가격정보.G시점().Equal(시점2) || 가격정보.G시점().Before(시점2))
}

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

package common

import (
	"C"
	//"bytes"
	"reflect"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

func F에러_패닉(에러 error) {
	if 에러 != nil {
		panic(에러)
	}
}

func F2문자열(값 interface{}) string {
	switch 값.(type) {
	case time.Time:
		return 값.(time.Time).Format(P시간_형식)
	case float64:
		return strconv.FormatFloat(값.(float64), 'f', -1, 64)
	case *byte:
		return C.GoString((*C.char)(unsafe.Pointer(값.(*byte))))
	case []byte:
		return string(값.([]byte))
	default:
		return F포맷된_문자열("%v", 값)
	}
}

func F2문자열_모음(인터페이스_모음 []interface{}) []string {
	if 인터페이스_모음 == nil {
		return nil
	}

	문자열_모음 := make([]string, len(인터페이스_모음))

	for i := 0; i < len(인터페이스_모음); i++ {
		문자열_모음[i] = F2문자열(인터페이스_모음[i])
	}

	return 문자열_모음
}

func F2정수(문자열 string) (int, error) {
	return strconv.Atoi(문자열)
}

func F2정수64(문자열 string) (int64, error) {
	return strconv.ParseInt(문자열, 10, 64)
}

func F2정수64_바이트(바이트_모음 []byte) int64 {
	반환값, 에러 := strconv.ParseInt(F2문자열(바이트_모음), 10, 64)
	F에러_패닉(에러)
	
	return 반환값
}

func F2실수(문자열 string) (float64, error) {
	return strconv.ParseFloat(문자열, 64)
}

func F2실수_바이트(바이트_모음 []byte) float64 {
	반환값, 에러 := strconv.ParseFloat(F2문자열(바이트_모음), 64)
	F에러_패닉(에러)
	
	return 반환값
}

func F2시각(문자열 string) (time.Time, error) {
	return time.Parse(P시간_형식, 문자열)
}

func F2시각_바이트(바이트_모음 []byte, 포맷_문자열 string) time.Time {
	반환값, 에러 := time.Parse(포맷_문자열, F2문자열(바이트_모음))
	F에러_패닉(에러)
	
	return 반환값
}

func F2참거짓_바이트(바이트_모음 []byte, 조건 interface{}, 결과 bool) bool {
	switch 조건.(type) {
	case string:
		if string(바이트_모음) == 조건.(string) {
			return 결과
		} else {
			return !결과
		}
	case int:
		if len(바이트_모음) != 1 {
			에러 := F에러_생성("바이트_모음 길이가 %v임. 예상치 못한 경우.", len(바이트_모음))
			F에러_출력(에러); panic(에러)
		}
		
		if int(uint(바이트_모음[0])) == 조건.(int) {
			return 결과
		} else {
			return !결과
		}
	}
	
	에러 := F에러_생성("예상치 못한 경우.")
	F에러_출력(에러); panic(에러)
}

func F2인터페이스_모음(문자열_모음 []string) []interface{} {
	if 문자열_모음 == nil {
		return nil
	}

	인터페이스_모음 := make([]interface{}, len(문자열_모음))

	for i := 0; i < len(문자열_모음); i++ {
		인터페이스_모음[i] = 문자열_모음[i]
	}

	return 인터페이스_모음
}

func F바이트_모음_늘리기(바이트_모음 []byte, 길이 int) []byte {
	if len(바이트_모음) > 길이 {
		에러 := F에러_생성("지정된 길이가 더 짧음.")
		F에러_출력(에러)
		panic(에러)
	}
	
	반환값 := make([]byte, 길이)
	
	for i:=0 ; i < len(바이트_모음) ; i++ {
		반환값[i] = 바이트_모음[i]
	}
	
	return 반환값
}

func F타입_이름(i interface{}) string {
	return reflect.TypeOf(i).Name()
}

func F문자열_복사(문자열 string) string {
	return (문자열 + " ")[:len(문자열)]
}

// 이하 최대 스레드 수량 관련 함수

func F단일_스레드_모드() { runtime.GOMAXPROCS(1) }
func F멀티_스레드_모드() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func F단일_스레드_모드임() bool {
	if runtime.GOMAXPROCS(-1) == 1 {
		return true
	} else {
		return false
	}
}

func F멀티_스레드_모드임() bool { return !F단일_스레드_모드임() }

// 이하 종료 시 존재하는 모든 Go루틴 정리(혹은 종료) 관련 함수 모음
var ch공통_종료_채널 = make(chan S비어있는_구조체)

func F공통_종료_채널() chan S비어있는_구조체 {
	return ch공통_종료_채널
}

func F공통_종료_채널_재설정() {
	ch공통_종료_채널 = make(chan S비어있는_구조체)
}

func F등록된_Go루틴_종료() {
	close(ch공통_종료_채널)
}

func F_nil에러() error { return nil }

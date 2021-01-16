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

// #include "func_C.h"
import "C"
import (
	xing "github.com/ghts/ghts/experimental/xing/go"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/c"
	"github.com/ghts/ghts/xing/base"
	"gopkg.in/ini.v1"
	"strings"
	"syscall"

	"bytes"
	"os"
	"path/filepath"
	"time"
	"unsafe"
)

func f초기화_XingAPI() {
	if API_초기화_완료.G값() {
		return
	} else {
		defer API_초기화_완료.S값(true)
	}

	lib.F조건부_패닉(lib.F환경변수("GOARCH") != "386", "C32 모듈은 32비트 전용입니다.")

	// DLL파일이 있는 디렉토리로 이동. (빼먹으면 안 됨)
	원래_디렉토리, 에러 := os.Getwd()
	lib.F확인(에러)

	xing디렉토리, 에러 := XingAPI디렉토리()
	lib.F확인(에러)

	lib.F확인(os.Chdir(xing디렉토리))

	// XingAPI 초기화 ('반드시' DLL파일이 있는 디렉토리에서 실행해야 함.)
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	xing_api_dll = syscall.Handle(uintptr(C.XingApiDll()))
	//xing_api_dll, 에러 := syscall.LoadLibrary("xingAPI.dll")
	//lib.F확인(에러)

	// 원래 디렉토리로 이동
	lib.F확인(os.Chdir(원래_디렉토리))

	// Xing API 함수 포인터
	etkConnect, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Connect")
	lib.F확인(에러)

	etkIsConnected, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_IsConnected")
	lib.F확인(에러)

	etkLogin, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Login")
	lib.F확인(에러)

	etkLogout, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Logout")
	lib.F확인(에러)

	// syscall, cgo 방식 모두 에러 발생.
	//etkDisconnect, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Disconnect")
	//lib.F확인(에러)

	etkRequest, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Request")
	lib.F확인(에러)

	etkAdviseRealData, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_AdviseRealData")
	lib.F확인(에러)

	etkUnadviseRealData, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_UnadviseRealData")
	lib.F확인(에러)

	etkUnadviseWindow, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_UnadviseWindow")
	lib.F확인(에러)

	etkGetAccountListCount, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountListCount")
	lib.F확인(에러)

	etkGetAccountList, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountList")
	lib.F확인(에러)

	// syscall 방식은 에러 발생. cgo 방식은 정상 작동.
	//etkGetAccountName, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAccountName")
	//lib.F확인(에러)

	// syscall 방식은 에러 발생. cgo 방식은 정상 작동.
	//etkGetAccountDetailName, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAcctDetailName")
	//lib.F확인(에러)

	// syscall 방식은 에러 발생. cgo 방식은 정상 작동.
	//etkGetAccountNickName, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetAcctNickname")
	//lib.F확인(에러)

	etkGetServerName, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetServerName")
	lib.F확인(에러)

	etkGetLastError, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetLastError")
	lib.F확인(에러)

	etkGetErrorMessage, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetErrorMessage")
	lib.F확인(에러)

	etkGetTRCountPerSec, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountPerSec")
	lib.F확인(에러)

	etkGetTRCountBaseSec, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountBaseSec")
	lib.F확인(에러)

	etkGetTRCountLimit, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountLimit")
	lib.F확인(에러)

	etkGetTRCountRequest, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_GetTRCountRequest")
	lib.F확인(에러)

	etkReleaseRequestData, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_ReleaseRequestData")
	lib.F확인(에러)

	etkReleaseMessageData, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_ReleaseMessageData")
	lib.F확인(에러)

	etkDecompress, 에러 = syscall.GetProcAddress(xing_api_dll, "ETK_Decompress")
	lib.F확인(에러)
}

func F접속() error {
	if 접속됨, 에러 := f접속됨(); 에러 != nil {
		return 에러
	} else if 접속됨 {
		return nil // 이미 접속됨.
	}

	var 서버_이름 string
	var 포트_번호 int

	switch xt.F서버_구분() {
	case xt.P서버_실거래:
		서버_이름 = "hts.ebestsec.co.kr"
		포트_번호 = 20001
	case xt.P서버_모의투자:
		서버_이름 = "demo.ebestsec.co.kr"
		포트_번호 = 20001
	case xt.P서버_XingACE:
		서버_이름 = "127.0.0.1"
		포트_번호 = 0
	}

	c서버_이름 := c.F2C문자열(서버_이름)
	defer c.F메모리_해제(unsafe.Pointer(c서버_이름))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	마이너스_일 := -1

	참거짓, _, 에러_번호 := syscall.Syscall6(etkConnect, 6,
		메시지_윈도우,
		uintptr(unsafe.Pointer(c서버_이름)),
		uintptr(포트_번호),
		WM_USER,
		uintptr(unsafe.Pointer(&마이너스_일)),
		uintptr(unsafe.Pointer(&마이너스_일)))

	if 에러_번호 != 0 {
		return lib.New에러("F접속() 에러 발생.\n'%v'", 에러_번호)
	} else if 참거짓 == FALSE {
		return lib.New에러("F접속() 실패.")
	}

	return nil
}

func F접속됨(질의 *S질의) {
	접속됨, 에러 := f접속됨()

	질의.Ch응답 <- New응답(lib.F조건부_값(에러 == nil, 접속됨, 에러))
}

func f접속됨() (bool, error) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkIsConnected, 0, 0, 0, 0)

	switch 에러_번호 {
	case 0:
		return 참거짓 == TRUE || 참거짓 != FALSE, nil
	default:
		return false, lib.New에러("f접속됨() 에러 발생.\n'%v'", 에러_번호)
	}
}

func F로그인() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	if lib.F파일_없음(설정파일_경로) {
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString("Xing 설정화일 없음\n")
		버퍼.WriteString("%v가 존재하지 않습니다.\n")
		버퍼.WriteString("sample_config.ini를 참조하여 새로 생성하십시오.")

		return lib.New에러(버퍼.String(), 설정파일_경로)
	}

	설정파일_복사본_이름 := lib.F2문자열("config_%v.ini", lib.F지금().Format("20060102_150406"))
	설정파일_복사본_경로 := filepath.Join(설정파일_디렉토리, 설정파일_복사본_이름)
	lib.F확인(lib.F파일_복사(설정파일_경로, 설정파일_복사본_경로))
	defer lib.F파일_삭제(설정파일_복사본_경로)

	cfg파일 := lib.F확인(ini.Load(설정파일_복사본_경로)).(*ini.File)
	섹션 := lib.F확인(cfg파일.GetSection("XingAPI_LogIn_Info")).(*ini.Section)

	키_ID := lib.F확인(섹션.GetKey("ID")).(*ini.Key)
	c아이디 := c.F2C문자열(키_ID.String())
	defer c.F메모리_해제(unsafe.Pointer(c아이디))

	키_PWD := lib.F확인(섹션.GetKey("PWD")).(*ini.Key)
	c암호 := c.F2C문자열(키_PWD.String())
	defer c.F메모리_해제(unsafe.Pointer(c암호))

	키_CertPWD := lib.F확인(섹션.GetKey("CertPWD")).(*ini.Key)
	공인인증서_암호 := lib.F조건부_문자열(xt.F서버_구분() == xt.P서버_실거래, 키_CertPWD.String(), "")
	c공인인증서_암호 := c.F2C문자열(공인인증서_암호)
	defer c.F메모리_해제(unsafe.Pointer(c공인인증서_암호))

	키_AcctPWD := lib.F확인(섹션.GetKey("AcctPWD")).(*ini.Key)
	계좌_비밀번호 = lib.F조건부_문자열(xt.F서버_구분() == xt.P서버_실거래, 키_AcctPWD.String(), "0000")

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkLogin, 6,
		메시지_윈도우,
		uintptr(unsafe.Pointer(c아이디)),
		uintptr(unsafe.Pointer(c암호)),
		uintptr(unsafe.Pointer(c공인인증서_암호)),
		0,
		uintptr(FALSE))

	switch {
	case 에러_번호 != 0:
		return lib.New에러("F로그인() 에러 발생.\n'%v'", 에러_번호)
	case 참거짓 == FALSE:
		return lib.New에러with출력("F로그인() 실패.")
	}

	return nil
}

func F로그아웃() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkLogout, 1,
		메시지_윈도우,
		0, 0)

	if 에러_번호 != 0 {
		return lib.New에러with출력("로그아웃 에러 : '%v'", 에러_번호)
	}

	// ETK_Disconnect() 에러 발생. syscall, cgo 모두 마찬가지.
	//_, _, 에러_번호 = syscall.Syscall(etkDisconnect, 0, 0, 0,0)
	//
	//if 에러_번호 != 0 {
	//	return lib.New에러with출력("접속 해제 에러 : '%v'", 에러_번호)
	//}

	return nil
}

func F질의(TR코드 string, c데이터 unsafe.Pointer, 길이 int,
	연속_조회_여부 bool, 연속키 string, 타임아웃 time.Duration) (반환값 int, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	접속됨 := false

	for i := 0; i < 3; i++ {
		if 접속됨, 에러 = f접속됨(); 에러 != nil {
			return -1, 에러
		} else if !접속됨 {
			if 에러 := xing.C32_재시작(); 에러 == nil {
				break
			}
		}
	}

	if !접속됨 {
		return -1, 에러
	}

	cTR코드 := c.F2C문자열(TR코드)
	defer c.F메모리_해제(unsafe.Pointer(cTR코드))

	c연속_조회_키 := c.F2C문자열(연속키)
	defer c.F메모리_해제(unsafe.Pointer(c연속_조회_키))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	질의ID, _, 에러_번호 := syscall.Syscall9(etkRequest, 7,
		메시지_윈도우,
		uintptr(unsafe.Pointer(cTR코드)),
		uintptr(c데이터),
		uintptr(길이),
		uintptr(lib.F조건부_정수(연속_조회_여부, TRUE, FALSE)),
		uintptr(unsafe.Pointer(c연속_조회_키)),
		uintptr(타임아웃/time.Second),
		0, 0)

	if 에러_번호 != 0 {
		에러 := lib.New에러with출력("F질의() 에러 발생. 에러 코드 : '%v'", 에러_번호)

		if strings.Contains(에러.Error(), "Access is denied.") {
			lib.F체크포인트("C32 자체 종료.")
			f종료()
		}
	}

	return int(질의ID), nil
}

func F실시간_정보_구독(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	cTR코드 := c.F2C문자열(TR코드)
	defer c.F메모리_해제(unsafe.Pointer(cTR코드))

	c전체_종목코드 := c.F2C문자열(전체_종목코드)
	defer c.F메모리_해제(unsafe.Pointer(c전체_종목코드))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkAdviseRealData, 4,
		메시지_윈도우,
		uintptr(unsafe.Pointer(cTR코드)),
		uintptr(unsafe.Pointer(c전체_종목코드)),
		uintptr(단위_길이),
		0, 0)

	if 에러_번호 != 0 || 참거짓 == FALSE {
		return lib.New에러with출력("실시간 정보 구독 에러 : '%v', '%v'", TR코드, 전체_종목코드)
	}

	return nil
}

func F실시간_정보_해지(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	cTR코드 := c.F2C문자열(TR코드)
	defer c.F메모리_해제(unsafe.Pointer(cTR코드))

	c전체_종목코드 := c.F2C문자열(전체_종목코드)
	defer c.F메모리_해제(unsafe.Pointer(c전체_종목코드))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall6(etkUnadviseRealData, 4,
		메시지_윈도우,
		uintptr(unsafe.Pointer(cTR코드)),
		uintptr(unsafe.Pointer(c전체_종목코드)),
		uintptr(단위_길이),
		0, 0)

	if 에러_번호 != 0 || 참거짓 == FALSE {
		return lib.New에러with출력("실시간 정보 해지 에러 : '%v', '%v'", TR코드, 전체_종목코드)
	}

	return nil
}

func F실시간_정보_일괄_해지(질의 *S질의) {
	에러 := f실시간_정보_일괄_해지()

	질의.Ch응답 <- New응답(lib.F조건부_값(에러 == nil, lib.P신호_OK, New응답(에러)))
}

func f실시간_정보_일괄_해지() error {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkUnadviseWindow, 1,
		메시지_윈도우,
		0, 0)

	switch {
	case 에러_번호 != 0:
		return lib.New에러("F실시간_정보_일괄_해지() 에러 발생.\n'%v'", 에러_번호)
	case 참거짓 == FALSE:
		return lib.New에러("F실시간_정보_일괄_해지() 실패.")
	}

	return nil
}

func F계좌_수량(질의 *S질의) {
	계좌_수량, 에러 := f계좌_수량()

	질의.Ch응답 <- New응답(lib.F조건부_값(에러 == nil, 계좌_수량, 에러))
}

func f계좌_수량() (int, error) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	계좌_수량, _, 에러_번호 := syscall.Syscall(etkGetAccountListCount, 0, 0, 0, 0)

	if 에러_번호 != 0 {
		return 0, lib.New에러("f계좌_수량() 에러 발생.\n'%v'", 에러_번호)
	}

	return int(계좌_수량), nil
}

func f계좌_번호(인덱스 int) (string, error) {
	버퍼_초기값 := "            " // 12자리 공백문자열
	버퍼_길이 := len(버퍼_초기값)

	c버퍼 := c.F2C문자열(버퍼_초기값)
	defer c.F메모리_해제(unsafe.Pointer(c버퍼))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	참거짓, _, 에러_번호 := syscall.Syscall(etkGetAccountList, 3,
		uintptr(인덱스),
		uintptr(unsafe.Pointer(c버퍼)),
		uintptr(버퍼_길이))

	if 에러_번호 != 0 {
		return "", lib.New에러("f계좌_번호() 에러 발생.\n'%v'", 에러_번호)
	} else if 참거짓 == FALSE {
		return "", lib.New에러("f계좌_번호() 호출 결과 FALSE.")
	}

	return string(bytes.Trim(c.F2Go바이트_모음(unsafe.Pointer(c버퍼), 버퍼_길이), "\x00")), nil
}

func f계좌_리스트_설정() (에러 error) {
	if 수량, 에러 := f계좌_수량(); 에러 != nil {
		return 에러
	} else if 수량 != len(계좌번호_모음) {
		계좌번호_모음 = make([]string, 0)

		for i := 0; i < 수량; i++ {
			if 계좌번호, 에러 := f계좌_번호(i); 에러 != nil {
				return lib.New에러with출력(에러)
			} else {
				계좌번호_모음 = append(계좌번호_모음, 계좌번호)
			}
		}
	}

	return nil
}

func F계좌번호_모음(질의 *S질의) {
	if len(계좌번호_모음) == 0 {
		if 에러 := f계좌_리스트_설정(); 에러 != nil {
			질의.Ch응답 <- New응답(에러)
			return
		}
	}

	질의.Ch응답 <- New응답(계좌번호_모음)
}

func F계좌_이름(질의 *S질의) {
	defer lib.S예외처리{M함수: func() {
		질의.Ch응답 <- New응답(lib.New에러("F계좌_이름() 에러 발생."))
	}}.S실행()

	계좌_번호 := 질의.M값.(*lib.S질의값_문자열).M문자열
	c계좌번호 := c.F2C문자열(계좌_번호)
	defer c.F메모리_해제(unsafe.Pointer(c계좌번호))

	var 버퍼_배열 [41]byte // 버퍼 길이 41로 고정
	버퍼_길이 := len(버퍼_배열)

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	C.GetAccountName(
		unsafe.Pointer(c계좌번호),
		unsafe.Pointer(&버퍼_배열[0]),
		C.int(버퍼_길이))

	계좌_이름 := lib.F2문자열_EUC_KR_공백제거(버퍼_배열[:])
	질의.Ch응답 <- New응답(계좌_이름)

	// syscall 방식 호출은 에러 발생
	//_, _, 에러_번호 := syscall.Syscall(etkGetAccountName, 3,
	//	uintptr(unsafe.Pointer(c계좌번호)),
	//	uintptr(unsafe.Pointer(c버퍼)),
	//	uintptr(버퍼_길이))
	//
	//switch 에러_번호 {
	//case 0:
	//	질의.Ch회신값 <- lib.F2문자열_EUC_KR_공백제거(c.F2Go바이트_모음(unsafe.Pointer(c버퍼), 버퍼_길이))
	//default:
	//	질의.Ch에러 <- lib.New에러("F계좌_이름() 에러 발생.\n'%v'", 에러_번호)
	//}
}

func F계좌_상세명(질의 *S질의) {
	defer lib.S예외처리{M함수: func() {
		질의.Ch응답 <- New응답(lib.New에러("F계좌_상세명() 에러 발생."))
	}}.S실행()

	계좌_번호 := 질의.M값.(*lib.S질의값_문자열).M문자열
	c계좌번호 := c.F2C문자열(계좌_번호)
	defer c.F메모리_해제(unsafe.Pointer(c계좌번호))

	var 버퍼_배열 [41]byte // 버퍼 길이 41로 고정
	버퍼_길이 := len(버퍼_배열)

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	C.GetAcctDetailName(
		unsafe.Pointer(c계좌번호),
		unsafe.Pointer(&버퍼_배열[0]),
		C.int(버퍼_길이))

	계좌_상세명 := lib.F2문자열_EUC_KR_공백제거(버퍼_배열[:])

	lib.F체크포인트("계좌 상세명", 계좌_상세명)

	질의.Ch응답 <- New응답(계좌_상세명)

	// syscall 방식 호출은 에러 발생
	//_, _, 에러_번호 := syscall.Syscall(etkGetAccountDetailName, 3,
	//	uintptr(unsafe.Pointer(c계좌번호)),
	//	uintptr(unsafe.Pointer(&버퍼_배열[0])),
	//	uintptr(버퍼_길이))
	//
	//switch 에러_번호 {
	//case 0:
	//	질의.Ch회신값 <- lib.F2문자열_EUC_KR_공백제거(버퍼_배열)
	//default:
	//	질의.Ch에러 <- lib.New에러("F계좌_상세명() 에러 발생.\n'%v'", 에러_번호)
	//}
}

func F계좌_별명(질의 *S질의) {
	defer lib.S예외처리{M함수: func() {
		질의.Ch응답 <- New응답(lib.New에러("F계좌_별명() 에러 발생."))
	}}.S실행()

	계좌_번호 := 질의.M값.(*lib.S질의값_문자열).M문자열
	c계좌번호 := c.F2C문자열(계좌_번호)
	defer c.F메모리_해제(unsafe.Pointer(c계좌번호))

	var 버퍼_배열 [41]byte // 버퍼 길이 41로 고정
	버퍼_길이 := len(버퍼_배열)

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	C.GetAcctNickName(
		unsafe.Pointer(c계좌번호),
		unsafe.Pointer(&버퍼_배열[0]),
		C.int(버퍼_길이))

	계좌_별명 := lib.F2문자열_EUC_KR_공백제거(버퍼_배열[:])
	질의.Ch응답 <- New응답(계좌_별명)

	// syscall 방식 호출은 에러 발생
	//_, _, 에러_번호 := syscall.Syscall(etkGetAccountNickName, 3,
	//	uintptr(unsafe.Pointer(c계좌번호)),
	//	uintptr(unsafe.Pointer(c버퍼)),
	//	uintptr(버퍼_길이))
	//
	//switch 에러_번호 {
	//case 0:
	//	질의.Ch회신값 <- lib.F2문자열_EUC_KR_공백제거(c.F2Go바이트_모음(unsafe.Pointer(c버퍼), 버퍼_길이))
	//default:
	//	질의.Ch에러 <- lib.New에러("F계좌_별명() 에러 발생.\n'%v'", 에러_번호)
	//}
}

func F서버_이름(질의 *S질의) {
	버퍼 := "                                                   "
	c버퍼 := c.F2C문자열(버퍼)
	defer c.F메모리_해제(unsafe.Pointer(c버퍼))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkGetServerName, 1,
		uintptr(unsafe.Pointer(c버퍼)),
		0, 0)

	질의.Ch응답 <- New응답(lib.F조건부_값(에러_번호 == 0,
		lib.F2문자열_EUC_KR_공백제거(c.F2Go바이트_모음(unsafe.Pointer(c버퍼), len(버퍼))),
		lib.New에러("F서버_이름() 에러 발생.\n'%v'", 에러_번호)))
}

func F에러_코드(질의 *S질의) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	에러_코드, _, 에러_번호 := syscall.Syscall(etkGetLastError, 0, 0, 0, 0)

	질의.Ch응답 <- New응답(lib.F조건부_값(에러_번호 == 0, int(에러_코드), lib.New에러("F에러_코드() 에러 발생.\n'%v'", 에러_번호)))
}

func F에러_메시지(질의 *S질의) {
	에러_코드 := 질의.M값.(*lib.S질의값_정수).M정수값

	go버퍼 := new(bytes.Buffer)
	for i := 0; i < 512; i++ {
		go버퍼.WriteString(" ")
	}

	버퍼 := go버퍼.String()
	버퍼_길이 := len(버퍼)
	c버퍼 := c.F2C문자열(버퍼)
	defer c.F메모리_해제(unsafe.Pointer(c버퍼))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	에러_메시지_길이, _, 에러_번호 := syscall.Syscall(etkGetErrorMessage, 3,
		uintptr(에러_코드),
		uintptr(unsafe.Pointer(c버퍼)),
		uintptr(버퍼_길이))

	switch {
	case 에러_번호 != 0:
		질의.Ch응답 <- New응답(lib.New에러("F에러_메시지() 에러 발생.\n'%v'", 에러_번호))
	case 에러_메시지_길이 == 0:
		질의.Ch응답 <- New응답(lib.New에러("에러 메시지를 구할 수 없습니다."))
	default:
		질의.Ch응답 <- New응답(lib.F2문자열_EUC_KR_공백제거(c.F2Go바이트_모음(unsafe.Pointer(c버퍼), int(에러_메시지_길이))))
	}
}

func TR코드별_전송_제한(질의 *S질의) {
	TR코드_모음 := 질의.M값.(*lib.S질의값_문자열_모음).M문자열_모음

	정보_모음 := new(xt.TR코드별_전송_제한_정보_모음)
	정보_모음.M맵 = make(map[string]*xt.TR코드별_전송_제한_정보)

	for _, TR코드 := range TR코드_모음 {
		값 := new(xt.TR코드별_전송_제한_정보)
		값.TR코드 = TR코드
		값.M초당_전송_제한 = f초당_TR쿼터(TR코드)
		값.M초_베이스 = f초당_TR쿼터_역수(TR코드)
		값.M10분당_전송_제한 = f10분당_TR쿼터(TR코드)
		값.M10분간_전송한_수량 = f10분간_요청한_TR수량(TR코드)

		정보_모음.M맵[TR코드] = 값
	}

	질의.Ch응답 <- New응답(정보_모음)
}

func f초당_TR쿼터(TR코드 string) int {
	cTR코드 := c.F2C문자열(TR코드)
	defer c.F메모리_해제(unsafe.Pointer(cTR코드))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	초당_전송_가능_횟수, _, 에러_번호 := syscall.Syscall(etkGetTRCountPerSec, 1,
		uintptr(unsafe.Pointer(cTR코드)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("f초당_TR쿼터() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(초당_전송_가능_횟수)
}

func f초당_TR쿼터_역수(TR코드 string) int {
	cTR코드 := c.F2C문자열(TR코드)
	defer c.F메모리_해제(unsafe.Pointer(cTR코드))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	초당_전송_가능_횟수_역수, _, 에러_번호 := syscall.Syscall(etkGetTRCountBaseSec, 1,
		uintptr(unsafe.Pointer(cTR코드)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("f초당_TR쿼터_역수() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(초당_전송_가능_횟수_역수)
}

func f10분당_TR쿼터(TR코드 string) int {
	cTR코드 := c.F2C문자열(TR코드)
	defer c.F메모리_해제(unsafe.Pointer(cTR코드))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	십분당_TR쿼터, _, 에러_번호 := syscall.Syscall(etkGetTRCountLimit, 1,
		uintptr(unsafe.Pointer(cTR코드)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("f10분당_TR쿼터() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(십분당_TR쿼터)
}

func f10분간_요청한_TR수량(TR코드 string) int {
	cTR코드 := c.F2C문자열(TR코드)
	defer c.F메모리_해제(unsafe.Pointer(cTR코드))

	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	십분간_요청한_TR수량, _, 에러_번호 := syscall.Syscall(etkGetTRCountRequest, 1,
		uintptr(unsafe.Pointer(cTR코드)),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("f10분간_요청한_TR수량() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(십분간_요청한_TR수량)
}

func F데이터_해제(식별번호 int) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkReleaseRequestData, 1,
		uintptr(식별번호),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F데이터_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}
}

func F메시지_해제(메시지_포인터 unsafe.Pointer) {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	_, _, 에러_번호 := syscall.Syscall(etkReleaseMessageData, 1,
		uintptr(메시지_포인터),
		0, 0)

	if 에러_번호 != 0 {
		lib.New에러with출력("F메시지_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}
}

func F압축_해제(압축된_원본_데이터 unsafe.Pointer, 버퍼 *byte, 원본_데이터_길이 int32) int {
	api_호출_잠금.Lock()
	defer api_호출_잠금.Unlock()

	압축_해제된_데이터_길이, _, 에러_번호 := syscall.Syscall(etkDecompress, 3,
		uintptr(압축된_원본_데이터),
		uintptr(unsafe.Pointer(버퍼)),
		uintptr(원본_데이터_길이))

	if 에러_번호 != 0 {
		lib.New에러with출력("F압축_해제() 에러 발생. 에러 코드 : '%v'", 에러_번호)
	}

	return int(압축_해제된_데이터_길이)
}
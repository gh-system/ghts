GHTS
====

자동 주식 매매 시스템 구현에 유용한 라이브러리.  


- 현재 주식 관련 일부 TR이 구현된 상태.
- 실제 거래에 적용 중이지만 여전히 변동이 많은 알파 내지 베타 상태.

*********************************************************

사전준비물
- Go언어 : https://golang.org/dl/
- C 컴파일러 (MSYS2) : https://www.msys2.org/ 
- Git : https://git-scm.com/download/win
- 이베스트 Xing API : https://www.ebestsec.co.kr/

*********************************************************
MSYS2에서 C 컴파일러 설치 ('MSYS2 -> MSYS' 터미널에서 아래 명령을 실행.)

<pre><code>pacman -Syuu 
pacman -S base-devel
pacman -S mingw-w64-i686-toolchain
pacman -S mingw-w64-x86_64-toolchain</code></pre>

*********************************************************

GHTS 라이브러리 설치

<pre><code>go get -u github.com/ghts/ghts</code></pre>
 
*********************************************************    

Go언어
- 복수의 매매 전략을 간편하게 동시 운용.
- 윈도우와의 호환성은 좋지 않은 편.
- DLL 호출만 제대로 지원되며, OLE/OCX등에 대한 지원은 빈약함.

이베스트투자증권 Xing API
- DLL 호출과 OLE 호출을 모두 지원. (Go언어와 호환성 우수.)
- API 문서가 비교적 충실함.
- 이베스트 증권사 홈페이지의 질답 게시판에서 비교적 원활한 기술 지원 진행.

*********************************************************

증권사 API 호출 기능은 별도의 프로세스로 분리됨.

- 매매 전략은 64비트에서 실행해야 할 필요성 제기됨.
- 증권사 API는 32비트에서'만' 호출할 수 있음. (국내 모든 증권사 API 공통)
- 2개의 프로세스로 분리한 후 상호 연동하는 방식으로 작성.
- 자세한 기술적 사항은 https://ghts.tistory.com/48 참조

*********************************************************

디렉토리별 설명
- lib : 공용 기능.
- xing/go : Xing API를 간접적으로 호출하는 순수 Go언어 모듈. (32/64비트 모두 가능)
- xing/c32 : Xing API을 직접 호출하는 모듈. (32비트 전용)
- xing/base : Xing API 공용 자료형.

*********************************************************

API 사용 관련 파일  

- xing/go/func_TR.go : TR호출 함수.
- xing/go/tr_<TR코드>_test.go : 개별 TR코드 함수 사용법.
- xing/go/test_main_test.go : 

*********************************************************

GHTS 사용 전 초기화
 
- 문자열 상수 사용

<pre><code>import (
    xt "github.com/ghts/ghts/xing/base"
    xing "github.com/ghts/ghts/xing/go"
    
    <... 기타 의존성 라이브러리 ...>
)

func main() {
    // <> 자리에 계좌 정보 문자열 상수 설정.
    // 'xt.P서버_실거래' 대신 'xt.P서버_모의투자' 사용 가능.    
    xing.F초기화(xt.P서버_실거래, <로그인_ID>, <로그인_암호>, <인증서_암호>, <계좌_암호>)  
    defer xing.F종료()    
       
    <... 이하 Xing API 호출 코드 ...>
}
</code></pre>

*********************************************************

GO MODULE 지원 미비.

- GHTS는 GO 모듈 지원 관련 리팩터링이 일부 진행되었으나, 아직 충분한 테스트가 이루어지지 않았습니다.
- GOPATH모드로 사용하려면 GO111MODULE 환경변수를 auto 혹은 off로 설정하면 됩니다.
- 더 자세한 사항은 GO111MODULE, GOPATH, GO모듈에 대해서는 인터넷을 검색해 보십시오.

*********************************************************

TR이란
- 트랜잭션(TRansaction)의 약자. 
- 증권사 서버와 사용자 컴퓨터 간의 상호 작용을 의미.
- 사용자가 증권사 서버에게 시킬 작업을 의미하게 됨.


TR코드
- 증권사 서버가 수행할 작업(TR)을 나타내는 코드.

자주 사용하는 TR코드
- 현재가 : t1101/t8407
- 호가 : t1102
- 주문 : CSPAT00600/CSPAT00700/CSPAT00800 (정상/정정/취소)
- 계좌 평가액 : CSPAQ12200
- 보유 수량 : CSPAQ12300
- 주문 가능 금액 : CSPAQ22200
- 일일 가격 정보 : t1305/t8413
- 전체 종목 코드 : t8436
- 현재 서버 시각 : t0167

TR코드에 대한 더 자세한 내용은 이베스트 Xing API 패키지를 설치하면 함께 설치되는 DevCenter를 참고.

주의 : GHTS에는 아직 일부 TR만 구현되어 있습니다.

*********************************************************


악성코드로 오진되는 경우.

xing32 모듈은 API호출 결과를 네트워크를 통해서 전달하는 동작 방식으로 인해,<br>
'안랩 세이프 트랜잭션'(AST : Ahnlab Safe Transaction)에서<br>
 멀웨어(MDP.Connect.M1924)로 잘못 진단되며,
AST에서 '위협 행위 차단'을 해제해야 정상 작동합니다.<br>
윈도우 기본 백신인 '윈도우 디펜더'에서는 이런 문제가 발생하지 않습니다.

*********************************************************

<주의>
------
저작권자, 개발자, 개발에 참여한 기여자들은 이 소프트웨어에 대한 어떠한 보증도 하지 않습니다.  
이 소프트웨어를 사용하면서 발생하는 그 어떠한 손실 및 손상에 대해서 책임지지 않습니다.  
소스코드 파일에 별도의 언급이 없는 한, 모든 소스코드는 GNU LGPL V2.1 라이센스를 따릅니다.  
저작권에 대한 자세한 사항은 'LICENSE' 파일을 참고하십시오.  

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
	"testing"
)

func TestCSPAQ22200_현물계좌_예수금_주문가능금액(t *testing.T) {
	t.Parallel()

	계좌번호, 에러 := 현물_계좌번호()
	lib.F테스트_에러없음(t, 에러)

	값, 에러 := TrCSPAQ22200_현물계좌_예수금_주문가능금액(계좌번호)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 값.M계좌번호, 계좌번호)
	lib.F테스트_다름(t, 값.M지점명, "")
	lib.F테스트_다름(t, 값.M계좌명, "")
	lib.F테스트_참임(t, 값.M현금주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M대용주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M코스피금액 >= 0)
	lib.F테스트_참임(t, 값.M코스닥금액 >= 0)
	lib.F테스트_참임(t, 값.M신용담보주문금액 >= 0)
	lib.F테스트_참임(t, 값.M증거금률100퍼센트주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M증거금률35퍼센트주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M증거금률50퍼센트주문가능금액 >= 0)
	lib.F테스트_참임(t, 값.M신용주문가능금액 >= 0)
	//lib.F테스트_참임(t, 값.M예수금 >= 0, 값.M예수금)
	lib.F테스트_참임(t, 값.M대용금액 >= 0)
	lib.F테스트_참임(t, 값.M증거금_현금 >= 0)
	lib.F테스트_참임(t, 값.M증거금_대용 >= 0)
	lib.F테스트_참임(t, 값.D1_예수금 >= 0)
	lib.F테스트_참임(t, 값.D2_예수금 >= 0)
	lib.F테스트_참임(t, 값.M미수금액 >= 0)
	lib.F테스트_참임(t, 값.D1연체변제소요금액 >= 0)
	lib.F테스트_참임(t, 값.D2연체변제소요금액 >= 0)
	lib.F테스트_같음(t, 값.M융자금액, 0)
	lib.F테스트_같음(t, 값.M변경후담보비율, 0.0)
	lib.F테스트_같음(t, 값.M소요담보금액, 0)
	lib.F테스트_같음(t, 값.M담보부족금액, 0)
	lib.F테스트_같음(t, 값.M원담보합계금액, 0)
	lib.F테스트_같음(t, 값.M부담보합계금액, 0)
	lib.F테스트_같음(t, 값.M신용담보금현금, 0)
	lib.F테스트_같음(t, 값.M신용담보대용금액, 0)
	lib.F테스트_같음(t, 값.M신용설정보증금, 0)
	lib.F테스트_같음(t, 값.M신용담보재사용금액, 0)
	lib.F테스트_같음(t, 값.M처분제한금액, 0)
	lib.F테스트_참임(t, 값.M전일매도정산금액 >= 0)
	lib.F테스트_참임(t, 값.M전일매수정산금액 >= 0)
	lib.F테스트_참임(t, 값.M금일매도정산금액 >= 0)
	lib.F테스트_참임(t, 값.M금일매수정산금액 >= 0)
	lib.F테스트_같음(t, 값.M매도대금담보대출금액, 0)
}

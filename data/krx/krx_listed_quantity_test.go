/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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

package krx

import (
	"github.com/ghts/ghts/lib"
	"testing"
)

func TestOTP(t *testing.T) {
	OTP, 에러 := otp()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(OTP) > 0)
}

func TestCSV다운로드(t *testing.T) {
	CSV, 에러 := csv다운로드()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(CSV) > 0)
}

func TestF상장_주식_수량_맵(t *testing.T) {
	상장_주식_수량_맵, 에러 := F상장_주식_수량_맵()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(상장_주식_수량_맵) > 2000)
}
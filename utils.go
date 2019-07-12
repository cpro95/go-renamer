package main

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

// Semigraphics provides an easy way to access unicode characters for drawing.
//
// Named like the unicode characters, 'Semigraphics'-prefix used if unicode block
// isn't prefixed itself.
const (
	// Block: General Punctation U+2000-U+206F (http://unicode.org/charts/PDF/U2000.pdf)
	SemigraphicsHorizontalEllipsis rune = '\u2026' // …

	// Block: Box Drawing U+2500-U+257F (http://unicode.org/charts/PDF/U2500.pdf)
	BoxDrawingsLightHorizontal                    rune = '\u2500' // ─
	BoxDrawingsHeavyHorizontal                    rune = '\u2501' // ━
	BoxDrawingsLightVertical                      rune = '\u2502' // │
	BoxDrawingsHeavyVertical                      rune = '\u2503' // ┃
	BoxDrawingsLightTripleDashHorizontal          rune = '\u2504' // ┄
	BoxDrawingsHeavyTripleDashHorizontal          rune = '\u2505' // ┅
	BoxDrawingsLightTripleDashVertical            rune = '\u2506' // ┆
	BoxDrawingsHeavyTripleDashVertical            rune = '\u2507' // ┇
	BoxDrawingsLightQuadrupleDashHorizontal       rune = '\u2508' // ┈
	BoxDrawingsHeavyQuadrupleDashHorizontal       rune = '\u2509' // ┉
	BoxDrawingsLightQuadrupleDashVertical         rune = '\u250a' // ┊
	BoxDrawingsHeavyQuadrupleDashVertical         rune = '\u250b' // ┋
	BoxDrawingsLightDownAndRight                  rune = '\u250c' // ┌
	BoxDrawingsDownLighAndRightHeavy              rune = '\u250d' // ┍
	BoxDrawingsDownHeavyAndRightLight             rune = '\u250e' // ┎
	BoxDrawingsHeavyDownAndRight                  rune = '\u250f' // ┏
	BoxDrawingsLightDownAndLeft                   rune = '\u2510' // ┐
	BoxDrawingsDownLighAndLeftHeavy               rune = '\u2511' // ┑
	BoxDrawingsDownHeavyAndLeftLight              rune = '\u2512' // ┒
	BoxDrawingsHeavyDownAndLeft                   rune = '\u2513' // ┓
	BoxDrawingsLightUpAndRight                    rune = '\u2514' // └
	BoxDrawingsUpLightAndRightHeavy               rune = '\u2515' // ┕
	BoxDrawingsUpHeavyAndRightLight               rune = '\u2516' // ┖
	BoxDrawingsHeavyUpAndRight                    rune = '\u2517' // ┗
	BoxDrawingsLightUpAndLeft                     rune = '\u2518' // ┘
	BoxDrawingsUpLightAndLeftHeavy                rune = '\u2519' // ┙
	BoxDrawingsUpHeavyAndLeftLight                rune = '\u251a' // ┚
	BoxDrawingsHeavyUpAndLeft                     rune = '\u251b' // ┛
	BoxDrawingsLightVerticalAndRight              rune = '\u251c' // ├
	BoxDrawingsVerticalLightAndRightHeavy         rune = '\u251d' // ┝
	BoxDrawingsUpHeavyAndRightDownLight           rune = '\u251e' // ┞
	BoxDrawingsDownHeacyAndRightUpLight           rune = '\u251f' // ┟
	BoxDrawingsVerticalHeavyAndRightLight         rune = '\u2520' // ┠
	BoxDrawingsDownLightAnbdRightUpHeavy          rune = '\u2521' // ┡
	BoxDrawingsUpLightAndRightDownHeavy           rune = '\u2522' // ┢
	BoxDrawingsHeavyVerticalAndRight              rune = '\u2523' // ┣
	BoxDrawingsLightVerticalAndLeft               rune = '\u2524' // ┤
	BoxDrawingsVerticalLightAndLeftHeavy          rune = '\u2525' // ┥
	BoxDrawingsUpHeavyAndLeftDownLight            rune = '\u2526' // ┦
	BoxDrawingsDownHeavyAndLeftUpLight            rune = '\u2527' // ┧
	BoxDrawingsVerticalheavyAndLeftLight          rune = '\u2528' // ┨
	BoxDrawingsDownLightAndLeftUpHeavy            rune = '\u2529' // ┨
	BoxDrawingsUpLightAndLeftDownHeavy            rune = '\u252a' // ┪
	BoxDrawingsHeavyVerticalAndLeft               rune = '\u252b' // ┫
	BoxDrawingsLightDownAndHorizontal             rune = '\u252c' // ┬
	BoxDrawingsLeftHeavyAndRightDownLight         rune = '\u252d' // ┭
	BoxDrawingsRightHeavyAndLeftDownLight         rune = '\u252e' // ┮
	BoxDrawingsDownLightAndHorizontalHeavy        rune = '\u252f' // ┯
	BoxDrawingsDownHeavyAndHorizontalLight        rune = '\u2530' // ┰
	BoxDrawingsRightLightAndLeftDownHeavy         rune = '\u2531' // ┱
	BoxDrawingsLeftLightAndRightDownHeavy         rune = '\u2532' // ┲
	BoxDrawingsHeavyDownAndHorizontal             rune = '\u2533' // ┳
	BoxDrawingsLightUpAndHorizontal               rune = '\u2534' // ┴
	BoxDrawingsLeftHeavyAndRightUpLight           rune = '\u2535' // ┵
	BoxDrawingsRightHeavyAndLeftUpLight           rune = '\u2536' // ┶
	BoxDrawingsUpLightAndHorizontalHeavy          rune = '\u2537' // ┷
	BoxDrawingsUpHeavyAndHorizontalLight          rune = '\u2538' // ┸
	BoxDrawingsRightLightAndLeftUpHeavy           rune = '\u2539' // ┹
	BoxDrawingsLeftLightAndRightUpHeavy           rune = '\u253a' // ┺
	BoxDrawingsHeavyUpAndHorizontal               rune = '\u253b' // ┻
	BoxDrawingsLightVerticalAndHorizontal         rune = '\u253c' // ┼
	BoxDrawingsLeftHeavyAndRightVerticalLight     rune = '\u253d' // ┽
	BoxDrawingsRightHeavyAndLeftVerticalLight     rune = '\u253e' // ┾
	BoxDrawingsVerticalLightAndHorizontalHeavy    rune = '\u253f' // ┿
	BoxDrawingsUpHeavyAndDownHorizontalLight      rune = '\u2540' // ╀
	BoxDrawingsDownHeavyAndUpHorizontalLight      rune = '\u2541' // ╁
	BoxDrawingsVerticalHeavyAndHorizontalLight    rune = '\u2542' // ╂
	BoxDrawingsLeftUpHeavyAndRightDownLight       rune = '\u2543' // ╃
	BoxDrawingsRightUpHeavyAndLeftDownLight       rune = '\u2544' // ╄
	BoxDrawingsLeftDownHeavyAndRightUpLight       rune = '\u2545' // ╅
	BoxDrawingsRightDownHeavyAndLeftUpLight       rune = '\u2546' // ╆
	BoxDrawingsDownLightAndUpHorizontalHeavy      rune = '\u2547' // ╇
	BoxDrawingsUpLightAndDownHorizontalHeavy      rune = '\u2548' // ╈
	BoxDrawingsRightLightAndLeftVerticalHeavy     rune = '\u2549' // ╉
	BoxDrawingsLeftLightAndRightVerticalHeavy     rune = '\u254a' // ╊
	BoxDrawingsHeavyVerticalAndHorizontal         rune = '\u254b' // ╋
	BoxDrawingsLightDoubleDashHorizontal          rune = '\u254c' // ╌
	BoxDrawingsHeavyDoubleDashHorizontal          rune = '\u254d' // ╍
	BoxDrawingsLightDoubleDashVertical            rune = '\u254e' // ╎
	BoxDrawingsHeavyDoubleDashVertical            rune = '\u254f' // ╏
	BoxDrawingsDoubleHorizontal                   rune = '\u2550' // ═
	BoxDrawingsDoubleVertical                     rune = '\u2551' // ║
	BoxDrawingsDownSingleAndRightDouble           rune = '\u2552' // ╒
	BoxDrawingsDownDoubleAndRightSingle           rune = '\u2553' // ╓
	BoxDrawingsDoubleDownAndRight                 rune = '\u2554' // ╔
	BoxDrawingsDownSingleAndLeftDouble            rune = '\u2555' // ╕
	BoxDrawingsDownDoubleAndLeftSingle            rune = '\u2556' // ╖
	BoxDrawingsDoubleDownAndLeft                  rune = '\u2557' // ╗
	BoxDrawingsUpSingleAndRightDouble             rune = '\u2558' // ╘
	BoxDrawingsUpDoubleAndRightSingle             rune = '\u2559' // ╙
	BoxDrawingsDoubleUpAndRight                   rune = '\u255a' // ╚
	BoxDrawingsUpSingleAndLeftDouble              rune = '\u255b' // ╛
	BoxDrawingsUpDobuleAndLeftSingle              rune = '\u255c' // ╜
	BoxDrawingsDoubleUpAndLeft                    rune = '\u255d' // ╝
	BoxDrawingsVerticalSingleAndRightDouble       rune = '\u255e' // ╞
	BoxDrawingsVerticalDoubleAndRightSingle       rune = '\u255f' // ╟
	BoxDrawingsDoubleVerticalAndRight             rune = '\u2560' // ╠
	BoxDrawingsVerticalSingleAndLeftDouble        rune = '\u2561' // ╡
	BoxDrawingsVerticalDoubleAndLeftSingle        rune = '\u2562' // ╢
	BoxDrawingsDoubleVerticalAndLeft              rune = '\u2563' // ╣
	BoxDrawingsDownSingleAndHorizontalDouble      rune = '\u2564' // ╤
	BoxDrawingsDownDoubleAndHorizontalSingle      rune = '\u2565' // ╥
	BoxDrawingsDoubleDownAndHorizontal            rune = '\u2566' // ╦
	BoxDrawingsUpSingleAndHorizontalDouble        rune = '\u2567' // ╧
	BoxDrawingsUpDoubleAndHorizontalSingle        rune = '\u2568' // ╨
	BoxDrawingsDoubleUpAndHorizontal              rune = '\u2569' // ╩
	BoxDrawingsVerticalSingleAndHorizontalDouble  rune = '\u256a' // ╪
	BoxDrawingsVerticalDoubleAndHorizontalSingle  rune = '\u256b' // ╫
	BoxDrawingsDoubleVerticalAndHorizontal        rune = '\u256c' // ╬
	BoxDrawingsLightArcDownAndRight               rune = '\u256d' // ╭
	BoxDrawingsLightArcDownAndLeft                rune = '\u256e' // ╮
	BoxDrawingsLightArcUpAndLeft                  rune = '\u256f' // ╯
	BoxDrawingsLightArcUpAndRight                 rune = '\u2570' // ╰
	BoxDrawingsLightDiagonalUpperRightToLowerLeft rune = '\u2571' // ╱
	BoxDrawingsLightDiagonalUpperLeftToLowerRight rune = '\u2572' // ╲
	BoxDrawingsLightDiagonalCross                 rune = '\u2573' // ╳
	BoxDrawingsLightLeft                          rune = '\u2574' // ╴
	BoxDrawingsLightUp                            rune = '\u2575' // ╵
	BoxDrawingsLightRight                         rune = '\u2576' // ╶
	BoxDrawingsLightDown                          rune = '\u2577' // ╷
	BoxDrawingsHeavyLeft                          rune = '\u2578' // ╸
	BoxDrawingsHeavyUp                            rune = '\u2579' // ╹
	BoxDrawingsHeavyRight                         rune = '\u257a' // ╺
	BoxDrawingsHeavyDown                          rune = '\u257b' // ╻
	BoxDrawingsLightLeftAndHeavyRight             rune = '\u257c' // ╼
	BoxDrawingsLightUpAndHeavyDown                rune = '\u257d' // ╽
	BoxDrawingsHeavyLeftAndLightRight             rune = '\u257e' // ╾
	BoxDrawingsHeavyUpAndLightDown                rune = '\u257f' // ╿
)

// print Strings to Screen
func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

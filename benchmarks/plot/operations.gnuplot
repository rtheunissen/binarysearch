load "benchmarks/plot/colors.gnuplot"

AVLBottomUp                = 1
AVLJoinBased               = 2
AVLWeakTopDown             = 3
AVLWeakBottomUp            = 4
AVLWeakJoinBased           = 5
AVLRelaxedTopDown          = 6
AVLRelaxedBottomUp         = 7
RedBlackRelaxedBottomUp    = 8
RedBlackRelaxedTopDown     = 9
LBSTBottomUp               = 10
LBSTTopDown                = 11
LBSTJoinBased              = 12
LBSTRelaxed                = 13
TreapTopDown               = 14
TreapJoinBased             = 15
TreapFingerTree            = 16
Randomized                 = 17
Zip                        = 18
Splay                      = 19
Conc                       = 20
RedBlackBottomUp           = 21


set style line AVLBottomUp                  dashtype 1 ps 2 lw 2 pt   1 pn 2 lc rgb COLOR_BLACK
set style line AVLJoinBased                 dashtype 1 ps 2 lw 2 pt   2 pn 2 lc rgb COLOR_CYAN
set style line AVLWeakBottomUp              dashtype 1 ps 2 lw 2 pt   4 pn 2 lc rgb COLOR_BLUE
set style line AVLWeakTopDown               dashtype 1 ps 2 lw 2 pt   3 pn 2 lc rgb COLOR_RED
set style line AVLWeakJoinBased             dashtype 1 ps 2 lw 2 pt   5 pn 2 lc rgb COLOR_CYAN
set style line AVLRelaxedTopDown            dashtype 1 ps 2 lw 2 pt   6 pn 2 lc rgb COLOR_BLUE
set style line AVLRelaxedBottomUp           dashtype 1 ps 2 lw 2 pt   7 pn 2 lc rgb COLOR_YELLOW
set style line RedBlackBottomUp             dashtype 1 ps 2 lw 2 pt  12 pn 2 lc rgb COLOR_RED
set style line RedBlackRelaxedBottomUp      dashtype 1 ps 2 lw 2 pt   8 pn 2 lc rgb COLOR_GREEN
set style line RedBlackRelaxedTopDown       dashtype 1 ps 2 lw 2 pt   9 pn 2 lc rgb COLOR_CYAN
set style line LBSTBottomUp                 dashtype 1 ps 2 lw 2 pt  10 pn 2 lc rgb COLOR_GREEN
set style line LBSTTopDown                  dashtype 1 ps 2 lw 2 pt  11 pn 2 lc rgb COLOR_BLUE
set style line LBSTJoinBased                dashtype 1 ps 2 lw 2 pt  12 pn 2 lc rgb COLOR_YELLOW
set style line LBSTRelaxed                  dashtype 1 ps 2 lw 2 pt  13 pn 2 lc rgb COLOR_RED
set style line TreapTopDown                 dashtype 1 ps 2 lw 2 pt  14 pn 2 lc rgb COLOR_RED
set style line TreapJoinBased               dashtype 1 ps 2 lw 2 pt  15 pn 2 lc rgb COLOR_CYAN
set style line TreapFingerTree              dashtype 1 ps 2 lw 2 pt  16 pn 2 lc rgb COLOR_BLUE
set style line Randomized                   dashtype 1 ps 2 lw 2 pt  17 pn 2 lc rgb COLOR_YELLOW
set style line Zip                          dashtype 1 ps 2 lw 2 pt  18 pn 2 lc rgb COLOR_GREEN
set style line Splay                        dashtype 1 ps 2 lw 2 pt  19 pn 2 lc rgb COLOR_PINK
set style line Conc                         dashtype 1 ps 2 lw 2 pt  20 pn 2 lc rgb COLOR_PURPLE


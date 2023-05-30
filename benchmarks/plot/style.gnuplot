
##################################################################
#
#           GRAPH
#
##################################################################

set terminal svg size 640,420 dynamic butt

set samples 1000

set border 6 front lc '#000000' lt 1 lw 0

set style line 1000 dashtype 1 lw 1 pt 1 ps 0 lc "#EEEEEE"

#set grid y2tics back linestyle 1000
set grid xtics back linestyle 1000

#set lmargin 1
set tmargin 7
#set rmargin 7
#set bmargin 5

set title font "Arial,20"
set title offset 0,4

set xlabel  font "Arial,16" offset 0,-1
set ylabel  font "Arial,16"

unset y2label
unset ylabel

set key reverse Left
set key horizontal outside center top
set key at screen 0.5, screen 0.85
set key width 0
set key height 0
set key spacing 1.1
set key samplen 4
set key font "monospace,16"

set autoscale xfix

unset xtics
unset ytics

unset mxtics
unset mytics

unset my2tics
unset mx2tics

unset x2tics
unset y2tics

set xtics font "monospace,16"
set xrange [0.5:*]

unset y2tics
set y2tics autofreq font "monospace,16"
set y2range [*:*]

set datafile missing "NaN"


##################################################################
#
#           COLORS
#
##################################################################

COLOR_RED       = "#F44336"
COLOR_YELLOW    = "#FFAB00"
COLOR_GREEN     = "#4CAF50"
COLOR_CYAN      = "#29C3F6"
COLOR_BLUE      = "#304FFE"
COLOR_PURPLE    = "#D336C8"
COLOR_BLACK     = "#000000"
COLOR_PINK      = "#EE82EE"

##################################################################
#
#           TREES
#
##################################################################

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

set style line AVLBottomUp                  dashtype 1 ps 1 lw 1 pt   1 pn 2 lc rgb COLOR_BLACK
set style line AVLJoinBased                 dashtype 1 ps 1 lw 1 pt   2 pn 2 lc rgb COLOR_CYAN

set style line AVLWeakBottomUp              dashtype 1 ps 1 lw 1 pt   4 pn 2 lc rgb COLOR_BLUE
set style line AVLWeakTopDown               dashtype 1 ps 1 lw 1 pt   3 pn 2 lc rgb COLOR_RED
set style line AVLWeakJoinBased             dashtype 1 ps 1 lw 1 pt   5 pn 2 lc rgb COLOR_CYAN

set style line AVLRelaxedTopDown            dashtype 1 ps 1 lw 1 pt   6 pn 2 lc rgb COLOR_BLUE
set style line AVLRelaxedBottomUp           dashtype 1 ps 1 lw 1 pt   7 pn 2 lc rgb COLOR_YELLOW
set style line RedBlackRelaxedBottomUp      dashtype 1 ps 1 lw 1 pt   8 pn 2 lc rgb COLOR_GREEN
set style line RedBlackRelaxedTopDown       dashtype 1 ps 1 lw 1 pt   9 pn 2 lc rgb COLOR_RED

set style line LBSTBottomUp                 dashtype 1 ps 1 lw 1 pt  10 pn 2 lc rgb COLOR_GREEN
set style line LBSTTopDown                  dashtype 1 ps 1 lw 1 pt  11 pn 2 lc rgb COLOR_BLUE
set style line LBSTJoinBased                dashtype 1 ps 1 lw 1 pt  12 pn 2 lc rgb COLOR_YELLOW
set style line LBSTRelaxed                  dashtype 1 ps 1 lw 1 pt  13 pn 2 lc rgb COLOR_RED

set style line TreapTopDown                 dashtype 1 ps 1 lw 1 pt  14 pn 2 lc rgb COLOR_RED
set style line TreapJoinBased               dashtype 1 ps 1 lw 1 pt  15 pn 2 lc rgb COLOR_CYAN
set style line TreapFingerTree              dashtype 1 ps 1 lw 1 pt  16 pn 2 lc rgb COLOR_BLUE
set style line Randomized                   dashtype 1 ps 1 lw 1 pt  17 pn 2 lc rgb COLOR_YELLOW
set style line Zip                          dashtype 1 ps 1 lw 1 pt  18 pn 2 lc rgb COLOR_GREEN

set style line Splay                        dashtype 1 ps 1 lw 1 pt  19 pn 2 lc rgb COLOR_PINK
set style line Conc                         dashtype 1 ps 1 lw 1 pt  20 pn 2 lc rgb COLOR_PURPLE


##################################################################
#
#           BALANCERS
#
##################################################################

Median     = 101
Weight     = 102
Height     = 103
DSW        = 104

set style line Median  dashtype 1  ps 2 lw 2 pt  10 pn 2 lc rgb COLOR_PURPLE
set style line Height  dashtype 4  ps 2 lw 2 pt   4 pn 2 lc rgb COLOR_GREEN
set style line Weight  dashtype 1  ps 2 lw 2 pt   9 pn 2 lc rgb COLOR_BLACK
set style line DSW     dashtype 1  ps 2 lw 2 pt   6 pn 2 lc rgb COLOR_YELLOW


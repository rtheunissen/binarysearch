load "docs/benchmarks/plot/colors.gnuplot"

##################################################################
#
#           STRATEGIES
#
##################################################################

AVLBottomUp                = 1
AVLTopDown                 = 2
AVLWeakTopDown             = 3
AVLWeakBottomUp            = 4
AVLRelaxedTopDown          = 5
AVLRelaxedBottomUp         = 6
RedBlackBottomUp           = 7
RedBlackTopDown            = 8
RedBlackRelaxedBottomUp    = 9
RedBlackRelaxedTopDown     = 10
LBSTBottomUp               = 11
LBSTTopDown                = 12
LBSTRelaxed                = 13
WBSTBottomUp               = 14
WBSTTopDown                = 15
WBSTRelaxed                = 16
TreapTopDown               = 17
TreapFingerTree            = 18
Randomized                 = 19
Zip                        = 20
Splay                      = 21
Conc                       = 22


COLOR_RED       = "#F44336"
COLOR_YELLOW    = "#FFAB00"
COLOR_GREEN     = "#00C853"
COLOR_CYAN      = "#29C3F6"
COLOR_BLUE      = "#3B75EA"
COLOR_PURPLE    = "#D336C8"
COLOR_BLACK     = "#000000"
COLOR_PINK      = "#EE82EE"
COLOR_BROWN     = "#9F650D"
COLOR_ORANGE    = "#EC6F1A"


set style line AVLBottomUp                  dashtype 1 ps 1 lw 1 pt   1 pn 2 lc rgb COLOR_BLACK
set style line AVLTopDown                   dashtype 3 ps 1 lw 1 pt   2 pn 2 lc rgb COLOR_CYAN
set style line AVLWeakBottomUp              dashtype 1 ps 1 lw 1 pt   4 pn 2 lc rgb COLOR_BLUE
set style line AVLWeakTopDown               dashtype 5 ps 1 lw 1 pt   3 pn 2 lc rgb COLOR_GREEN
set style line AVLRelaxedTopDown            dashtype 5 ps 1 lw 1 pt   6 pn 2 lc rgb COLOR_BLUE
set style line AVLRelaxedBottomUp           dashtype 1 ps 1 lw 1 pt   7 pn 2 lc rgb COLOR_YELLOW
set style line RedBlackBottomUp             dashtype 1 ps 1 lw 1 pt  12 pn 2 lc rgb COLOR_RED
set style line RedBlackTopDown              dashtype 3 ps 1 lw 1 pt   5 pn 2 lc rgb COLOR_BLUE
set style line RedBlackRelaxedBottomUp      dashtype 1 ps 1 lw 1 pt   8 pn 2 lc rgb COLOR_GREEN
set style line RedBlackRelaxedTopDown       dashtype 5 ps 1 lw 1 pt   9 pn 2 lc rgb COLOR_CYAN

set style line LBSTBottomUp                 dashtype 1 ps 1 lw 1 pt   6 pn 2 lc rgb COLOR_BLUE
set style line LBSTTopDown                  dashtype 5 ps 1 lw 1 pt   7 pn 2 lc rgb COLOR_RED
set style line LBSTRelaxed                  dashtype 5 ps 1 lw 1 pt  13 pn 2 lc rgb COLOR_PINK

set style line WBSTBottomUp                 dashtype 1 ps 1 lw 1 pt   8 pn 2 lc rgb COLOR_BLACK
set style line WBSTTopDown                  dashtype 4 ps 1 lw 1 pt   9 pn 2 lc rgb COLOR_CYAN
set style line WBSTRelaxed                  dashtype 4 ps 1 lw 1 pt   9 pn 2 lc rgb COLOR_YELLOW

set style line TreapTopDown                 dashtype 5 ps 1 lw 1 pt  14 pn 2 lc rgb COLOR_RED
set style line TreapFingerTree              dashtype 1 ps 1 lw 1 pt  16 pn 2 lc rgb COLOR_BLUE
set style line Randomized                   dashtype 1 ps 1 lw 1 pt  17 pn 2 lc rgb COLOR_YELLOW
set style line Zip                          dashtype 1 ps 1 lw 1 pt  18 pn 2 lc rgb COLOR_GREEN
set style line Splay                        dashtype 1 ps 1 lw 1 pt  19 pn 2 lc rgb COLOR_PINK
set style line Conc                         dashtype 3 ps 1 lw 1 pt  20 pn 2 lc rgb COLOR_PURPLE


##################################################################
#
#           GROUPS
#
##################################################################

AVL                    = "AVLBottomUp AVLTopDown"
AVLRedBlack            = "AVLBottomUp AVLTopDown RedBlackBottomUp RedBlackTopDown"
AVLWeak                = "AVLBottomUp AVLTopDown AVLWeakBottomUp AVLWeakTopDown"
AVLRelaxed             = "AVLBottomUp AVLTopDown AVLRelaxedBottomUp AVLRelaxedTopDown"
RedBlackRelaxed        = "RedBlackBottomUp RedBlackTopDown RedBlackRelaxedBottomUp RedBlackRelaxedTopDown"
RankBalanced           = "AVLBottomUp AVLWeakTopDown AVLRelaxedTopDown RedBlackRelaxedTopDown"

HeightBalanced        = "AVLBottomUp RedBlackBottomUp AVLWeakBottomUp AVLRelaxedBottomUp Conc"

WeightBalanced         = "LBSTBottomUp LBSTTopDown WBSTBottomUp WBSTTopDown"
WeightBalancedRelaxed  = "LBSTBottomUp LBSTRelaxed WBSTBottomUp WBSTRelaxed"

Probabilistic          = "TreapTopDown TreapFingerTree Randomized Zip"
Combination            = "AVLRelaxedBottomUp AVLWeakBottomUp LBSTBottomUp TreapTopDown"
CombinationSplay       = "AVLRelaxedBottomUp AVLWeakBottomUp LBSTBottomUp TreapTopDown Splay"
SizeOnly               = "LBSTBottomUp LBSTRelaxed Randomized Splay"

GROUPS = "AVLWeakRedBlack AVLRelaxed RedBlack"
GROUPS = "RankBalanced AVLRedBlack AVLWeak AVLRelaxed RedBlackRelaxed WeightBalanced WeightBalancedRelaxed"

OPERATIONS = "Insert InsertPersistent InsertDelete InsertDeletePersistent InsertDeleteCycles InsertDeleteCyclesPersistent InsertDeleteSearch InsertDeleteSearchPersistent"
OPERATIONS = "Insert InsertDeleteCycles"


DISTRIBUTIONS = "Uniform"

do for [GROUP in GROUPS] {

    do for [OPERATION in OPERATIONS] {

        ##################################################################
        #
        #           Allocations
        #
        ##################################################################

        MEASUREMENT = 'Allocations'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold Allocations } / log_2(Size)"

        DATA = sprintf("docs/benchmarks/data/operations/measurements/%s", OPERATION)

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Allocations')/column('Step')/log2(column('Size'))"

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           MAXIMUM PATH LENGTH
        #
        ##################################################################

        MEASUREMENT = 'MaximumPathLength'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold Maximum Path Length } / log_2(Size)"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('MaximumPathLength')/log2(column('Size'))"

        DATA = sprintf("docs/benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           AVERAGE PATH LENGTH
        #
        ##################################################################

        MEASUREMENT = 'AveragePathLength'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold AveragePathLength } / log_2(Size)"

        DATA = sprintf("docs/benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('AveragePathLength')/log2(column('Size'))"

        SMOOTH = "sbezier"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           ROTATIONS
        #
        ##################################################################

        MEASUREMENT = "Rotations"

        set xlabel "{Operations × 10^5}"
        set ylabel "{/:Bold Rotations} / Operation"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Rotations')/column('Step')"

        DATA = sprintf("docs/benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           DURATION
        #
        ##################################################################

        MEASUREMENT = 'Duration'

        set xlabel "Operations / 10^6"
        set ylabel "{/:Bold Duration } in milliseconds"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Duration')/1000000" # nano / 10^6 is 1k

        set format y2 "%.2f"

        DATA = sprintf("docs/benchmarks/data/operations/benchmarks/%s", OPERATION)

        SMOOTH = "sbezier"
        load "docs/benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "docs/benchmarks/plot/operations/lines.gnuplot"
    }
}